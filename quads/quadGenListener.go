package quads

import (
	"fmt"
	"log"

	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/memory"
	"github.com/jpr98/compis/parser"
	"github.com/jpr98/compis/semantic"
)

// QuadGenListener contains logic to analyze a program on a semantic level. The main
// function of this pass over the source code is to generate code for our virtual
// machine. It manages expressions, conditional statements, loops, function and method
// calls, etc. Almost all methods are called when the parser visits rule nodes and
// emits events.
type QuadGenListener struct {
	*parser.BaseProyectoListener
	m *Manager

	enteredVarInit bool
}

// NewListener creates a new QuadGenListener with a Manager
func NewListener(functionTable semantic.FunctionTable, classTable semantic.ClassTable) QuadGenListener {
	m := NewManager(functionTable, classTable)
	return QuadGenListener{m: &m}
}

// GetManager returns the Manager
func (l QuadGenListener) GetManager() Manager {
	return *l.m
}

// EnterAssignation adds the assign operator to the operator stack
func (l *QuadGenListener) EnterAssignation(c *parser.AssignationContext) {
	l.m.PushOp(c.ASSIGN().GetText())
}

// ExitAssignation generates the assignation quad
func (l *QuadGenListener) ExitAssignation(c *parser.AssignationContext) {
	l.m.GenerateAssignationQuad(false)
}

// EnterVarsDec adds an operand to the operand stack
func (l *QuadGenListener) EnterVarsDec(c *parser.VarsDecContext) {
	l.m.PushOperand(c.ID().GetText())
}

// EnterVarsTypeInit validates that the Class used to declare a variable exists
func (l *QuadGenListener) EnterVarsTypeInit(c *parser.VarsTypeInitContext) {
	if c.ID() != nil {
		if _, exists := l.m.classTable[c.ID().GetText()]; !exists {
			log.Fatalf("Error: Undefined type %s", c.ID().GetText())
		}
	}
}

// ExitVarsTypeInit removes an operand from the stack if it wasn't initialized in declaration
func (l *QuadGenListener) ExitVarsTypeInit(c *parser.VarsTypeInitContext) {
	defer func() { l.enteredVarInit = false }()
	if !l.enteredVarInit {
		l.m.operands.Pop()
	}
}

// EnterVarsTypeInit2 adds an assignation quad for variable initialization on declaration
func (l *QuadGenListener) EnterVarsTypeInit2(c *parser.VarsTypeInit2Context) {
	l.enteredVarInit = true
	l.m.PushOp(c.ASSIGN().GetText())
}

// ExitVarsTypeInit2 generates the assignation quad for variable initialization on declaration
func (l *QuadGenListener) ExitVarsTypeInit2(c *parser.VarsTypeInit2Context) {
	l.m.GenerateAssignationQuad(false)
}

// EnterExp2 adds OR operator to stack if present
func (l *QuadGenListener) EnterExp2(c *parser.Exp2Context) {
	if c.OR() != nil {
		l.m.PushOp(c.OR().GetText())
	}
}

// ExitExp2 generates quad for OR expression
func (l *QuadGenListener) ExitExp2(c *parser.Exp2Context) {
	l.m.GenerateQuad([]int{
		constants.OPOR,
	}, false)
}

// EnterT_exp2 adds AND operator to stack if present
func (l *QuadGenListener) EnterT_exp2(c *parser.T_exp2Context) {
	if c.AND() != nil {
		l.m.PushOp(c.AND().GetText())
	}
}

// ExitT_exp2 generates quad for AND expression
func (l *QuadGenListener) ExitT_exp2(c *parser.T_exp2Context) {
	l.m.GenerateQuad([]int{
		constants.OPAND,
	}, false)
}

// EnterG_exp2 adds RELOP operator to stack if present
func (l *QuadGenListener) EnterG_exp2(c *parser.G_exp2Context) {
	if c.Relop() != nil {
		l.m.PushOp(c.Relop().GetText())
	}
}

// ExitG_exp2 generates quad for RELOP expression
func (l *QuadGenListener) ExitG_exp2(c *parser.G_exp2Context) {
	l.m.GenerateQuad([]int{
		constants.OPGT,
		constants.OPLT,
		constants.OPEQ,
		constants.OPNEQ,
	}, false)
}

// EnterM_exp2 adds ADD or SUB operator to stack if present
func (l *QuadGenListener) EnterM_exp2(c *parser.M_exp2Context) {
	if c.ADD() != nil {
		l.m.PushOp(c.ADD().GetText())
	} else if c.SUB() != nil {
		l.m.PushOp(c.SUB().GetText())
	}
}

// ExitM_exp2 generates quad for ADD or SUB expression
func (l *QuadGenListener) ExitM_exp2(c *parser.M_exp2Context) {
	l.m.GenerateQuad([]int{
		constants.OPPLUS,
		constants.OPMINUS,
	}, false)
}

// EnterTerm2 adds MUL or DIV operator to stack if present
func (l *QuadGenListener) EnterTerm2(c *parser.Term2Context) {
	if c.MUL() != nil {
		l.m.PushOp(c.MUL().GetText())
		return
	} else if c.DIV() != nil {
		l.m.PushOp(c.DIV().GetText())
		return
	}
}

// ExitTerm2 generates quad for MUL or DIV expression
func (l *QuadGenListener) ExitTerm2(c *parser.Term2Context) {
	l.m.GenerateQuad([]int{
		constants.OPMULT,
		constants.OPDIV,
	}, false)
}

// EnterFactor adds constant operand to stack if present
func (l *QuadGenListener) EnterFactor(c *parser.FactorContext) {
	if c.VarCte() != nil {
		l.m.PushConstantOperand(c.VarCte().GetText())
	}
}

// EnterFactor2 adds LPAREN false bottom to operator stack if present
func (l *QuadGenListener) EnterFactor2(c *parser.Factor2Context) {
	if c.LPAREN() != nil {
		l.m.PushOp(c.LPAREN().GetText())
	}
}

// ExitFactor2 removes false bottom
func (l *QuadGenListener) ExitFactor2(c *parser.Factor2Context) {
	// Removes false bottom
	l.m.operators.Pop()
}

// EnterVars adds an operand to the stack, if it's an instance access it passes both ids
func (l *QuadGenListener) EnterVars(c *parser.VarsContext) {
	if c.ID(1) != nil {
		l.m.PushOperand(c.GetText())
	} else {
		l.m.PushOperand(c.ID(0).GetText())
	}
}

// EnterVarArray adds an array operand to the stack, if it's an instance access it passes both ids
func (l *QuadGenListener) EnterVarArray(c *parser.VarArrayContext) {
	if c.ID(1) != nil {
		operand := fmt.Sprintf("%s.%s", c.ID(0).GetText(), c.ID(1).GetText())
		l.m.PushOperand(operand)
	} else {
		l.m.PushOperand(c.ID(0).GetText())
	}
}

// ExitVarArray generates the Verify and ArrayAccess quads needed for array operands,
// if its an instance access it passes both ids
func (l *QuadGenListener) ExitVarArray(c *parser.VarArrayContext) {
	if c.ID(1) != nil {
		operand := fmt.Sprintf("%s.%s", c.ID(0).GetText(), c.ID(1).GetText())
		l.m.AddVerifyQuad(operand, true)
	} else {
		l.m.AddVerifyQuad(c.ID(0).GetText(), true)
	}
	l.m.AddArrayAccessQuad()
}

// ExitConditional updates the pending goto quad
func (l *QuadGenListener) ExitConditional(c *parser.ConditionalContext) {
	// Neuralgic point 2
	l.m.UpdateGoto()
}

// ExitConditional2 generates a GOTOF quad
func (l *QuadGenListener) ExitConditional2(c *parser.Conditional2Context) {
	// Neuralgic point 1
	l.m.AddGotoF()
}

// EnterConditional4 generates a GOTO quad and updates pending quad
func (l *QuadGenListener) EnterConditional4(c *parser.Conditional4Context) {
	// Neuralgic point 3
	l.m.AddAndUpdateGoto()
}

// ExitWhileLoop generates while GOTO quad and updates two pending quads
func (l *QuadGenListener) ExitWhileLoop(c *parser.WhileLoopContext) {
	l.m.AddAndUpdateWhileGotos()
}

// EnterWhileLoop2 saves the current instruction pointer in the jumpstack
func (l *QuadGenListener) EnterWhileLoop2(c *parser.WhileLoop2Context) {
	l.m.SaveJumpPosition()
}

// ExitWhileLoop2 generates a GOTOF quad
func (l *QuadGenListener) ExitWhileLoop2(c *parser.WhileLoop2Context) {
	l.m.AddGotoF()
}

// ExitForLoop generates the quads for incrementing the loop iterator, adds a GOTO
// quad and updates two pending quads. Finally we delete the info of the iterator
// from the current funciton's variables table.
func (l *QuadGenListener) ExitForLoop(c *parser.ForLoopContext) {
	l.m.PushConstantOperand("1")
	l.m.PushOp("+")
	l.m.GenerateQuad([]int{constants.OPPLUS}, true)

	l.m.PushOp("=")
	l.m.GenerateAssignationQuad(true)

	l.m.AddAndUpdateWhileGotos()

	operand := l.m.operands.Pop()
	delete(l.m.getCurrentFunctionTable()[l.m.currentFunction].Vars, operand.ID())
}

// EnterForLoop2 adds the iterator to the operands stack and adds an assign operator
func (l *QuadGenListener) EnterForLoop2(c *parser.ForLoop2Context) {
	l.m.AddForLoopIterator(c.ID().GetText())
	l.m.PushOp(c.ASSIGN().GetText())
}

// ExitForLoop2 generates the initial value of the for loop assignation quad and adds
// the instruction pointer to the jumpstack
func (l *QuadGenListener) ExitForLoop2(c *parser.ForLoop2Context) {
	l.m.GenerateAssignationQuad(true)
	l.m.SaveJumpPosition()
}

// ExitForLoop3 adds LT operator to stack and generates a quad to make the comparison with
// the iterator and a GOTOF quad with the comparison result
func (l *QuadGenListener) ExitForLoop3(c *parser.ForLoop3Context) {
	l.m.PushOp("<")
	l.m.GenerateQuad([]int{constants.OPLT}, true)
	l.m.AddGotoF()
}

// EnterProgram adds the global scope to the stack and sets global name and current function
func (l *QuadGenListener) EnterProgram(c *parser.ProgramContext) {
	l.m.scopeStack.Push(c.ID().GetText())
	l.m.currentFunction = l.m.scopeStack.Top()
	l.m.globalName = l.m.scopeStack.Top()
}

// EnterProgram2 saves instruction pointer to jump stack and adds a simple goto, which is for
// main. We do this after global variable declaration in case of any of them being initialized
// at declaration.
func (l *QuadGenListener) EnterProgram2(c *parser.Program2Context) {
	l.m.SaveJumpPosition()
	l.m.AddSimpleGOTO()
}

// EnterFunctions sets the instruction pointer address of the current function
func (l *QuadGenListener) EnterFunctions(c *parser.FunctionsContext) {
	l.m.currentFunction = c.ID().GetText()
	l.m.getCurrentFunctionTable()[l.m.currentFunction].Dir = len(l.m.GetQuads())
}

// ExitFunctions validates that the current function has a return statement, generates an
// Endfunc quad and sets the memory size for the function.
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

// ExitReturnRule generates a return quad
func (l *QuadGenListener) ExitReturnRule(c *parser.ReturnRuleContext) {
	l.m.AddReturnQuad()
}

// ExitRead2 generates a read quad
func (l *QuadGenListener) ExitRead2(c *parser.Read2Context) {
	l.m.AddReadQuad()
}

// ExitWrite generates a write new line quad, this helps to print in different lines
// each call to write() but same line contents of each call.
func (l *QuadGenListener) ExitWrite(c *parser.WriteContext) {
	l.m.AddNewLineWriteQuad()
}

// EnterMain updates pending goto with main function instruction pointer (Dir) and
// generates main ERA quad
func (l *QuadGenListener) EnterMain(c *parser.MainContext) {
	l.m.currentFunction = c.MAIN().GetText()
	l.m.UpdateGoto()
	l.m.AddEraQuad(c.MAIN().GetText(), "")
}

// ExitMain sets the memory size for main
func (l *QuadGenListener) ExitMain(c *parser.MainContext) {
	tempSize, objSize := memory.Manager.ResetTempCounter()
	l.m.functionTable[l.m.currentFunction].TempSize = tempSize
	l.m.functionTable[l.m.currentFunction].ObjectCount += objSize
	l.m.setObjectSize()
	l.m.currentFunction = l.m.scopeStack.Top()
}

// ExitProgram sets the global memory size
func (l *QuadGenListener) ExitProgram(c *parser.ProgramContext) {
	varsSize, objSize := memory.Manager.GetGlobalSize()
	l.m.functionTable[l.m.globalName].VarsSize = varsSize
	l.m.functionTable[l.m.globalName].ObjectCount = objSize
	l.m.setObjectSize()
	l.m.scopeStack.Pop()
}

// EnterFunctionCall validates that a function exists and adds it's era quad, we also
// check if an explicit method call is being made (a method being called inside of its
// own class, so it's syntax is indistinguishable from a normal funciton).
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
		if isMethodCall, className := l.m.CheckImplicitMethodCall(c.ID().GetText()); isMethodCall {
			l.m.AddEraQuad(c.ID().GetText(), className)
			return
		}
		log.Fatalf("Error: (EnterFunctionCall) undeclared function %s", c.ID().GetText())
	}
	l.m.AddEraQuad(c.ID().GetText(), "")
}

// ExitFunctionCall generates the GOSUB quad of a function, same as EnterFunctionCall we check
// if the function call is really an implicit method call
func (l *QuadGenListener) ExitFunctionCall(c *parser.FunctionCallContext) {
	if isMethodCall, className := l.m.CheckImplicitMethodCall(c.ID().GetText()); isMethodCall {
		l.m.AddClassGoSubQuad(className, "self")
		return
	}
	l.m.AddGoSubQuad()
}

// EnterArguments2 generates Param quads with recently parsed argument
func (l *QuadGenListener) EnterArguments2(c *parser.Arguments2Context) {
	if l.m.currentFunctionCallClass.Top() != "" { // This is how we know in this point we are calling a method
		l.m.AddClassParamQuad()
		aux := l.m.paramCounter.Pop()
		l.m.paramCounter.Push(aux + 1)
		return
	}
	l.m.AddParamQuad()
	aux := l.m.paramCounter.Pop()
	l.m.paramCounter.Push(aux + 1)
}

// EnterW_arguments2 generates a write quad
func (l *QuadGenListener) EnterW_arguments2(c *parser.W_arguments2Context) {
	l.m.AddWriteQuad()
}

// EnterClassDef sets the class as the current scope
func (l *QuadGenListener) EnterClassDef(c *parser.ClassDefContext) {
	l.m.scopeStack.Push(c.ID().GetText())
	l.m.currentFunction = l.m.scopeStack.Top()
}

// ExitClassDef removes the class from the current scope and sets global as the current function
func (l *QuadGenListener) ExitClassDef(c *parser.ClassDefContext) {
	l.m.scopeStack.Pop()
	l.m.currentFunction = l.m.scopeStack.Top()
}

// EnterClassInit sets the instruction pointer addres to the init function of the current class
func (l *QuadGenListener) EnterClassInit(c *parser.ClassInitContext) {
	l.m.currentFunction = c.INIT().GetText()
	l.m.classTable[l.m.scopeStack.Top()].Methods[c.INIT().GetText()].Dir = len(l.m.GetQuads())
}

// ExitClassInit generates the init return and gosub quads and sets the memory size for init
func (l *QuadGenListener) ExitClassInit(c *parser.ClassInitContext) {
	l.m.AddInitReturnQuad()
	l.m.AddEndFuncQuad()
	tempSize, objSize := memory.Manager.ResetTempCounter()
	l.m.classTable[l.m.scopeStack.Top()].Methods[c.INIT().GetText()].TempSize = tempSize
	l.m.classTable[l.m.scopeStack.Top()].Methods[c.INIT().GetText()].ObjectCount += objSize
	l.m.setObjectSize()
}

// EnterInitCall generates Era quad for init
func (l *QuadGenListener) EnterInitCall(c *parser.InitCallContext) {
	l.m.AddEraQuad("init", c.ID().GetText())
}

// ExitInitCall generates GOSUB quad for init
func (l *QuadGenListener) ExitInitCall(c *parser.InitCallContext) {
	l.m.AddClassGoSubQuad(c.ID().GetText(), "")
}

// EnterMethodCall validates that the instance of a method is declared, then that the method exits
// in the class of the instance and finally generates an ERA quad for the method call
func (l *QuadGenListener) EnterMethodCall(c *parser.MethodCallContext) {
	className := ""
	if attr, exists := l.m.getCurrentFunctionTable()[l.m.currentFunction].Vars[c.ID(0).GetText()]; exists {
		className = attr.Class
	} else {
		if attr, exists := l.m.functionTable[l.m.globalName].Vars[c.ID(0).GetText()]; exists {
			className = attr.Class
		}
	}

	if className == "" {
		log.Fatalf("Error: Undeclared variable %s", c.ID(0).GetText())
	}

	methodData, exists := l.m.classTable[className].Methods[c.ID(1).GetText()]
	if !exists {
		log.Fatalf("Error: Class %s doesn't have a method %s", className, c.ID(1).GetText())
	}
	if methodData.IsPrivate && l.m.scopeStack.Top() != className {
		log.Fatalf("Error: method %s is private", c.ID(1).GetText())
	}

	l.m.AddEraQuad(c.ID(1).GetText(), className)
}

// ExitMethodCall validates that the instance of a method is declared and then generates
// the GOSUB quad for the method call
func (l *QuadGenListener) ExitMethodCall(c *parser.MethodCallContext) {
	className := ""
	if attr, exists := l.m.getCurrentFunctionTable()[l.m.currentFunction].Vars[c.ID(0).GetText()]; exists {
		className = attr.Class
	} else {
		if attr, exists := l.m.functionTable[l.m.globalName].Vars[c.ID(0).GetText()]; exists {
			className = attr.Class
		}
	}

	if className == "" {
		log.Fatalf("Error: Undeclared variable %s", c.ID(0).GetText())
	}

	l.m.AddClassGoSubQuad(className, c.ID(0).GetText())
}
