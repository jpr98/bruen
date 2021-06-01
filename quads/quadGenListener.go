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
	m *Manager

	enteredVarInit  bool
	isArgumentParam bool
}

func NewListener(functionTable semantic.FunctionTable, classTable semantic.ClassTable) QuadGenListener {
	ss := utils.StringStack{}
	m := NewManager(functionTable, classTable, ss)
	return QuadGenListener{m: &m}
}

func (l QuadGenListener) GetManager() Manager {
	return *l.m
}

func (l *QuadGenListener) EnterAssignation(c *parser.AssignationContext) {
	l.m.PushOp(c.ASSIGN().GetText())
}

func (l *QuadGenListener) ExitAssignation(c *parser.AssignationContext) {
	l.m.GenerateAssignationQuad(false)
}

func (l *QuadGenListener) EnterVarsDec(c *parser.VarsDecContext) {
	l.m.PushOperand(c.ID().GetText())
}

func (l *QuadGenListener) EnterVarsTypeInit(c *parser.VarsTypeInitContext) {
	if c.ID() != nil {
		if _, exists := l.m.classTable[c.ID().GetText()]; !exists {
			log.Fatalf("Error: Undefined type %s", c.ID().GetText())
		}
	}
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

func (l *QuadGenListener) EnterExp2(c *parser.Exp2Context) {
	if c.OR() != nil {
		l.m.PushOp(c.OR().GetText())
	}
}

func (l *QuadGenListener) ExitExp2(c *parser.Exp2Context) {
	l.m.GenerateQuad([]int{
		constants.OPOR,
	}, false)
}

func (l *QuadGenListener) EnterT_exp2(c *parser.T_exp2Context) {
	if c.AND() != nil {
		l.m.PushOp(c.AND().GetText())
	}
}

func (l *QuadGenListener) ExitT_exp2(c *parser.T_exp2Context) {
	l.m.GenerateQuad([]int{
		constants.OPAND,
	}, false)
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
	if c.VarCte() != nil {
		l.m.PushConstantOperand(c.VarCte().GetText())
	}
}

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

func (l *QuadGenListener) EnterVars(c *parser.VarsContext) {
	if c.ID(1) != nil {
		l.m.PushOperand(c.GetText())
	} else {
		l.m.PushOperand(c.ID(0).GetText())
	}
}

func (l *QuadGenListener) EnterVarArray(c *parser.VarArrayContext) {
	if c.ID(1) != nil {
		operand := fmt.Sprintf("%s.%s", c.ID(0).GetText(), c.ID(1).GetText())
		l.m.PushOperand(operand)
	} else {
		l.m.PushOperand(c.ID(0).GetText())
	}
}

func (l *QuadGenListener) ExitVarArray(c *parser.VarArrayContext) {
	if c.ID(1) != nil {
		operand := fmt.Sprintf("%s.%s", c.ID(0).GetText(), c.ID(1).GetText())
		l.m.AddVerifyQuad(operand, true)
	} else {
		l.m.AddVerifyQuad(c.ID(0).GetText(), true)
	}
	l.m.AddArrayAccessQuad()
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
	l.m.scopeStack.Push(c.ID().GetText())
	l.m.currentFunction = l.m.scopeStack.Top()
	l.m.globalName = l.m.scopeStack.Top()
}

func (l *QuadGenListener) EnterProgram2(c *parser.Program2Context) {
	l.m.SaveJumpPosition()
	l.m.AddSimpleGOTO()
}

func (l *QuadGenListener) EnterFunctions(c *parser.FunctionsContext) {
	l.m.currentFunction = c.ID().GetText()
	l.m.getCurrentFunctionTable()[l.m.currentFunction].Dir = len(l.m.GetQuads())
}

func (l *QuadGenListener) ExitFunctions(c *parser.FunctionsContext) {
	if len(l.m.quads) > 0 {
		if c.TypeRule() != nil && l.m.quads[len(l.m.quads)-1].Action != RETURN {
			log.Fatalf("Error: (ExitFunctions) missing return statement in non-void function\n")
		}
		if c.VOID() != nil && l.m.quads[len(l.m.quads)-1].Action == RETURN && l.m.quads[len(l.m.quads)-1].Result != nil {
			log.Fatalf("Error: (ExitFunctions) void function shouldn't return a value")
		}
	}

	l.m.AddEndFuncQuad()
	tempSize, objSize := memory.Manager.ResetTempCounter()
	l.m.getCurrentFunctionTable()[l.m.currentFunction].TempSize = tempSize
	l.m.getCurrentFunctionTable()[l.m.currentFunction].ObjectCount += objSize
	l.m.setObjectSize()
}

func (l *QuadGenListener) ExitReturnRule(c *parser.ReturnRuleContext) {
	l.m.AddReturnQuad()
}

func (l *QuadGenListener) ExitRead2(c *parser.Read2Context) {
	l.m.AddReadQuad()
}

func (l *QuadGenListener) ExitWrite(c *parser.WriteContext) {
	l.m.AddNewLineWriteQuad()
}

func (l *QuadGenListener) EnterMain(c *parser.MainContext) {
	l.m.currentFunction = c.MAIN().GetText()
	l.m.UpdateGoto()
	l.m.AddEraQuad(c.MAIN().GetText(), "")
}

func (l *QuadGenListener) ExitMain(c *parser.MainContext) {
	tempSize, objSize := memory.Manager.ResetTempCounter()
	l.m.functionTable[l.m.currentFunction].TempSize = tempSize
	l.m.functionTable[l.m.currentFunction].ObjectCount += objSize
	l.m.setObjectSize()
	l.m.currentFunction = l.m.scopeStack.Top()
}

func (l *QuadGenListener) ExitProgram(c *parser.ProgramContext) {
	varsSize, objSize := memory.Manager.GetGlobalSize()
	l.m.functionTable[l.m.globalName].VarsSize = varsSize
	l.m.functionTable[l.m.globalName].ObjectCount = objSize
	l.m.setObjectSize()
	l.m.scopeStack.Pop()
	fmt.Println(l.m.operands)
}

func (l *QuadGenListener) EnterFunctionCall(c *parser.FunctionCallContext) {
	exists := false
	if _, ok := l.m.functionTable[c.ID().GetText()]; ok {
		exists = true
	} else {
		if _, ok := l.m.classTable[c.ID().GetText()]; ok {
			exists = true
		}
	}
	if !exists {
		log.Fatalf("Error: (EnterFunctionCall) undeclared function %s", c.ID().GetText())
	}
	l.m.AddEraQuad(c.ID().GetText(), "")
	l.isArgumentParam = true
}

func (l *QuadGenListener) ExitFunctionCall(c *parser.FunctionCallContext) {
	l.m.AddGoSubQuad()
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

// Classes
func (l *QuadGenListener) EnterClassDef(c *parser.ClassDefContext) {
	l.m.scopeStack.Push(c.ID(0).GetText())
	l.m.currentFunction = l.m.scopeStack.Top()
}

func (l *QuadGenListener) ExitClassDef(c *parser.ClassDefContext) {
	l.m.scopeStack.Pop()
	l.m.currentFunction = l.m.scopeStack.Top()
}

func (l *QuadGenListener) EnterAttributesDec(c *parser.AttributesDecContext) {
	if c.ID(1) != nil {
		if _, exists := l.m.classTable[c.ID(1).GetText()]; !exists {
			log.Fatalf("Error: Undefined type %s", c.ID(1).GetText())
		}
	}
}

func (l *QuadGenListener) EnterClassInit(c *parser.ClassInitContext) {
	l.m.currentFunction = c.INIT().GetText()
	l.m.classTable[l.m.scopeStack.Top()].Methods[c.INIT().GetText()].Dir = len(l.m.GetQuads())
}

func (l *QuadGenListener) ExitClassInit(c *parser.ClassInitContext) {
	l.m.AddInitReturnQuad()
	l.m.AddEndFuncQuad()
	tempSize, objSize := memory.Manager.ResetTempCounter()
	l.m.classTable[l.m.scopeStack.Top()].Methods[c.INIT().GetText()].TempSize = tempSize
	l.m.classTable[l.m.scopeStack.Top()].Methods[c.INIT().GetText()].ObjectCount += objSize
	l.m.setObjectSize()
}

func (l *QuadGenListener) EnterInitCall(c *parser.InitCallContext) {
	l.m.AddEraQuad("init", c.ID().GetText())
}

func (l *QuadGenListener) ExitInitCall(c *parser.InitCallContext) {
	l.m.AddClassGoSubQuad(c.ID().GetText())
}
