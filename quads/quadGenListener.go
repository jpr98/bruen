package quads

import (
	"fmt"
	"log"

	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/memory"
	"github.com/jpr98/compis/parser"
	"github.com/jpr98/compis/semantic"
	"github.com/jpr98/compis/utils"
)

type QuadGenListener struct {
	*parser.BaseProyectoListener
	m               *Manager
	scopeStack      utils.StringStack
	currentFunction string
	globalName      string

	enteredVarInit  bool
	isArgumentParam bool
}

func NewListener(functionTable semantic.FunctionTable) QuadGenListener {
	m := NewManager(functionTable)
	ss := utils.StringStack{}
	return QuadGenListener{m: &m, scopeStack: ss}
}

func (l QuadGenListener) GetManager() Manager {
	return *l.m
}

func (l *QuadGenListener) EnterAssignation(c *parser.AssignationContext) {
	l.m.PushOperand(c.ID().GetText(), l.currentFunction, l.globalName)
	l.m.PushOp(c.ASSIGN().GetText())
}

func (l *QuadGenListener) ExitAssignation(c *parser.AssignationContext) {
	l.m.GenerateAssignationQuad(false)
}

func (l *QuadGenListener) EnterVarsDec(c *parser.VarsDecContext) {
	l.m.PushOperand(c.ID().GetText(), l.currentFunction, l.globalName)
}

func (l *QuadGenListener) ExitVarsTypeInit(c *parser.VarsTypeInitContext) {
	defer func() { l.enteredVarInit = false }()
	if !l.enteredVarInit {
		l.m.operands.Pop()
	}
}

func (l *QuadGenListener) EnterVarsTypeInit2(c *parser.VarsTypeInit2Context) {
	l.enteredVarInit = true
	l.m.PushOp(c.ASSIGN().GetText())
}

func (l *QuadGenListener) ExitVarsTypeInit2(c *parser.VarsTypeInit2Context) {
	l.m.GenerateAssignationQuad(false)
}

func (l *QuadGenListener) EnterG_exp2(c *parser.G_exp2Context) {
	if c.Relop() != nil {
		l.m.PushOp(c.Relop().GetText())
	}
}

func (l *QuadGenListener) ExitG_exp2(c *parser.G_exp2Context) {
	l.m.GenerateQuad([]int{
		constants.OPGT,
		constants.OPLT,
		constants.OPEQ,
		constants.OPNEQ,
	}, false)
}

func (l *QuadGenListener) EnterM_exp2(c *parser.M_exp2Context) {
	if c.ADD() != nil {
		l.m.PushOp(c.ADD().GetText())
	} else if c.SUB() != nil {
		l.m.PushOp(c.SUB().GetText())
	}
}
func (l *QuadGenListener) ExitM_exp2(c *parser.M_exp2Context) {
	l.m.GenerateQuad([]int{
		constants.OPPLUS,
		constants.OPMINUS,
	}, false)
}

func (l *QuadGenListener) EnterTerm2(c *parser.Term2Context) {
	if c.MUL() != nil {
		l.m.PushOp(c.MUL().GetText())
		return
	}

	if c.DIV() != nil {
		l.m.PushOp(c.DIV().GetText())
		return
	}
}

func (l *QuadGenListener) ExitTerm2(c *parser.Term2Context) {
	l.m.GenerateQuad([]int{
		constants.OPMULT,
		constants.OPDIV,
	}, false)
}

func (l *QuadGenListener) EnterFactor(c *parser.FactorContext) {
	if c.Vars() != nil {
		l.m.PushOperand(c.Vars().GetText(), l.currentFunction, l.globalName)
	} else if c.VarCte() != nil {
		l.m.PushConstantOperand(c.VarCte().GetText())
	}
}

// func (l *QuadGenListener) EnterVarCte(c *parser.VarCteContext) {
// 	if c.Cte_i() != nil {
// 		l.m.PushConstantOperand(c.Cte_i().GetText())
// 	} else if c.Cte_f() != nil {
// 		l.m.PushConstantOperand(c.Cte_f().GetText())
// 	} else if c.Cte_c() != nil {
// 		l.m.PushConstantOperand(c.Cte_c().GetText())
// 	} else if c.Cte_b() != nil {
// 		l.m.PushConstantOperand(c.Cte_b().GetText())
// 	}
// }

func (l *QuadGenListener) EnterFactor2(c *parser.Factor2Context) {
	// Adds false bottom
	if c.LPAREN() != nil {
		l.m.PushOp(c.LPAREN().GetText())
	}
}

func (l *QuadGenListener) ExitFactor2(c *parser.Factor2Context) {
	// Removes false bottom
	l.m.operators.Pop()
}

func (l *QuadGenListener) ExitConditional(c *parser.ConditionalContext) {
	// Neuralgic point 2
	l.m.UpdateGoto()
}

func (l *QuadGenListener) ExitConditional2(c *parser.Conditional2Context) {
	// Neuralgic point 1
	l.m.AddGotoF()
}

func (l *QuadGenListener) EnterConditional4(c *parser.Conditional4Context) {
	// Neuralgic point 3
	l.m.AddAndUpdateGoto()
}

func (l *QuadGenListener) ExitWhileLoop(c *parser.WhileLoopContext) {
	l.m.AddAndUpdateWhileGotos()
}

func (l *QuadGenListener) EnterWhileLoop2(c *parser.WhileLoop2Context) {
	l.m.SaveJumpPosition()
}

func (l *QuadGenListener) ExitWhileLoop2(c *parser.WhileLoop2Context) {
	l.m.AddGotoF()
}

func (l *QuadGenListener) ExitForLoop(c *parser.ForLoopContext) {
	// TODO: change this string to constant(1) address
	l.m.PushConstantOperand("1")
	l.m.PushOp("+")
	l.m.GenerateQuad([]int{constants.OPPLUS}, true)

	l.m.PushOp("=")
	l.m.GenerateAssignationQuad(true)

	l.m.AddAndUpdateWhileGotos()
}

func (l *QuadGenListener) EnterForLoop2(c *parser.ForLoop2Context) {
	l.m.AddForLoopIterator(c.ID().GetText())
	l.m.PushOp(c.ASSIGN().GetText())
}
func (l *QuadGenListener) ExitForLoop2(c *parser.ForLoop2Context) {
	l.m.GenerateAssignationQuad(true)
	l.m.SaveJumpPosition()
}

func (l *QuadGenListener) ExitForLoop3(c *parser.ForLoop3Context) {
	l.m.PushOp("<")
	l.m.GenerateQuad([]int{constants.OPLT}, true)
	l.m.AddGotoF()
}

func (l *QuadGenListener) EnterProgram(c *parser.ProgramContext) {
	l.scopeStack.Push(c.ID().GetText())
	l.currentFunction = l.scopeStack.Top()
	l.globalName = l.scopeStack.Top()
	l.m.SaveJumpPosition()
	l.m.AddSimpleGOTO()
}

func (l *QuadGenListener) EnterClassDef(c *parser.ClassDefContext) {
	l.scopeStack.Push(c.ID(0).GetText())
	l.currentFunction = l.scopeStack.Top()
}

func (l *QuadGenListener) ExitClassDef(c *parser.ClassDefContext) {
	l.scopeStack.Pop()
	l.currentFunction = l.scopeStack.Top()
}

func (l *QuadGenListener) EnterFunctions(c *parser.FunctionsContext) {
	l.currentFunction = c.ID().GetText()
	l.m.functionTable[l.currentFunction].Dir = len(l.m.GetQuads())
}

func (l *QuadGenListener) ExitFunctions(c *parser.FunctionsContext) {
	l.m.AddEndFuncQuad()
	l.m.functionTable[l.currentFunction].Vars = nil
	l.m.functionTable[l.currentFunction].EraSize = "0i1f2c3b"
	l.m.functionTable[l.currentFunction].TempSize = memory.Manager.ResetTempCounter()
}

func (l *QuadGenListener) EnterRead2(c *parser.Read2Context) {
	l.m.AddReadQuad(c.Vars().GetText(), l.currentFunction, l.globalName)
}

func (l *QuadGenListener) EnterMain(c *parser.MainContext) {
	l.currentFunction = c.MAIN().GetText()
	l.m.UpdateGoto()
	l.m.AddEraQuad(c.MAIN().GetText())
}

func (l *QuadGenListener) ExitMain(c *parser.MainContext) {
	l.m.functionTable[l.currentFunction].TempSize = memory.Manager.ResetTempCounter()
}

func (l *QuadGenListener) ExitProgram(c *parser.ProgramContext) {
	l.scopeStack.Pop()
	l.m.functionTable[l.globalName].VarsSize = memory.Manager.GetGlobalSize()
	fmt.Println(l.m.operands)
}

func (l *QuadGenListener) EnterFunctionCall(c *parser.FunctionCallContext) {
	_, exists := l.m.functionTable[c.ID().GetText()]
	if !exists {
		log.Fatalf("Error: (EnterFunctionCall) undeclared function %s", c.ID().GetText())
	}
	l.m.AddEraQuad(c.ID().GetText())
	l.isArgumentParam = true
}

func (l *QuadGenListener) ExitFunctionCall(c *parser.FunctionCallContext) {
	l.m.AddGoSubQuad(c.ID().GetText())
	l.isArgumentParam = false
}

func (l *QuadGenListener) EnterArguments2(c *parser.Arguments2Context) {
	if !l.isArgumentParam {
		l.m.AddWriteQuad()
		return
	}
	l.m.AddParamQuad()
	l.m.paramCounter++
}
