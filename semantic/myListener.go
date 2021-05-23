package semantic

import (
	"fmt"
	"log"

	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/memory"
	"github.com/jpr98/compis/parser"
	"github.com/jpr98/compis/utils"
)

// Our Listener that extends proyecto_listener interface generated by antlr
type MyListener struct {
	*parser.BaseProyectoListener
	scopeStack      utils.StringStack
	currentFunction string
	ProgramName     string
	//TODO: store type of function (create FunctionAttributes?)
	functionTable       FunctionTable
	unassignedVariables []string
}

func NewListener() MyListener {
	var listener MyListener
	listener.scopeStack = utils.StringStack{}
	listener.currentFunction = ""
	listener.ProgramName = ""
	listener.functionTable = make(map[string]*FunctionTableContent)
	listener.unassignedVariables = make([]string, 0)
	return listener
}

func (l *MyListener) GetFunctionTable() map[string]*FunctionTableContent {
	return l.functionTable
}

// Validates uniqueness of the function ID defined by the developer
func (l *MyListener) handleFunctionRedefinition(scope string) {
	if function, exists := l.functionTable[l.currentFunction]; exists {
		if function.Scope == scope {
			log.Fatal(l.currentFunction + " is already Defined in scope " + l.scopeStack.Top())
		}
	}
}

// Adds function ID to the functionTable
func (l *MyListener) addToFunctionTable(typeOf string, scope string) {
	l.handleFunctionRedefinition(scope)
	l.functionTable[l.currentFunction] = &FunctionTableContent{}
	l.functionTable[l.currentFunction].Vars = make(map[string]*VariableAttributes)
	l.functionTable[l.currentFunction].TypeOf = typeOf
	l.functionTable[l.currentFunction].Scope = scope
}

func (l *MyListener) EnterProgram(c *parser.ProgramContext) {
	l.currentFunction = c.ID().GetText()
	l.ProgramName = c.ID().GetText()
	l.addToFunctionTable("void", "")
	l.scopeStack.Push(c.ID().GetText())

	// Setting a constant 1 int
	dir, err := memory.Manager.GetNextAddr(constants.TYPEINT, memory.Constant)
	if err != nil {
		log.Fatalf("Error: (EnterProgram) setting constant 1, %s", err)
	}
	ConstantsTable["1"] = &constantsTableContent{constants.TYPEINT, dir}
}

func (l *MyListener) EnterClassDef(c *parser.ClassDefContext) {
	l.currentFunction = c.ID(0).GetText()
	l.addToFunctionTable("void", l.scopeStack.Top())
	l.scopeStack.Push(c.ID(0).GetText())
}

// After defining the classes, we need to continue addressing variables to the program function.
func (l *MyListener) ExitClassDef(c *parser.ClassDefContext) {
	l.scopeStack.Pop()
	l.currentFunction = l.scopeStack.Top()
}

func (l *MyListener) EnterFunctions(c *parser.FunctionsContext) {
	l.currentFunction = c.ID().GetText()
	var typeString string
	if c.TypeRule() != nil {
		typeString = c.TypeRule().GetText()
	} else if c.VOID() != nil {
		typeString = c.VOID().GetText()
	}

	l.addToFunctionTable(typeString, l.scopeStack.Top())
}

func (l *MyListener) ExitFunctions(c *parser.FunctionsContext) {
	l.functionTable[l.currentFunction].VarsSize = memory.Manager.ResetLocalCounter()
}

func (l *MyListener) EnterMain(c *parser.MainContext) {
	l.currentFunction = c.MAIN().GetText()
	l.addToFunctionTable("void", l.scopeStack.Top())
}

func (l *MyListener) EnterVarsDec(c *parser.VarsDecContext) {
	id := c.ID().GetText()
	l.unassignedVariables = append(l.unassignedVariables, id)
}

func (l *MyListener) EnterParameter(c *parser.ParameterContext) {
	id := c.ID().GetText()
	t := constants.StringToType(c.TypeRule().GetText())
	if t == constants.ERR {
		log.Fatalf("Error: (EnterParameter) unkown type from %s", c.TypeRule().GetText())
	}

	dir, err := memory.Manager.GetNextAddr(t, memory.Local)
	if err != nil {
		log.Fatalf("Error: (EnterParameter) %s\n", err)
	}

	currVariable := VariableAttributes{id, t, dir}
	l.functionTable[l.currentFunction].Vars[id] = &currVariable
	l.functionTable[l.currentFunction].Params = append(l.functionTable[l.currentFunction].Params, t)
}

func (l *MyListener) EnterVarsTypeInit(c *parser.VarsTypeInitContext) {
	for _, id := range l.unassignedVariables {
		l.validateVariableInScope(id)

		if c.TypeRule() != nil {
			t := constants.StringToType(c.TypeRule().GetText())
			if t == constants.ERR {
				log.Fatalf("Error: (EnterVarsTypeInit) unkown type from %s", c.TypeRule().GetText())
			}

			var memCtx memory.Context
			if l.currentFunction == l.ProgramName {
				memCtx = memory.Global
			} else {
				memCtx = memory.Local
			}
			dir, err := memory.Manager.GetNextAddr(t, memCtx)
			if err != nil {
				log.Fatalf("Error: (EnterVarsTypeInit) %s\n", err)
			}

			currVariable := VariableAttributes{id, t, dir}
			l.functionTable[l.currentFunction].Vars[id] = &currVariable

		} else if c.ID() != nil {
			// TODO: Asegurarnos que la clase existe
			currVariable := VariableAttributes{id, constants.TYPECLASS, 20000}
			l.functionTable[l.currentFunction].Vars[id] = &currVariable
		}
	}
	l.unassignedVariables = nil
}

func (l *MyListener) EnterVarCte(c *parser.VarCteContext) {
	if c.Cte_i() != nil {
		if _, exists := ConstantsTable[c.Cte_i().GetText()]; !exists {
			dir, err := memory.Manager.GetNextAddr(constants.TYPEINT, memory.Constant)
			if err != nil {
				log.Fatalf("Error: (EnterVarCte) %s\n", err)
			}
			ConstantsTable[c.Cte_i().GetText()] = &constantsTableContent{constants.TYPEINT, dir}
		}
	} else if c.Cte_f() != nil {
		if _, exists := ConstantsTable[c.Cte_f().GetText()]; !exists {
			dir, err := memory.Manager.GetNextAddr(constants.TYPEFLOAT, memory.Constant)
			if err != nil {
				log.Fatalf("Error: (EnterVarCte) %s\n", err)
			}
			ConstantsTable[c.Cte_f().GetText()] = &constantsTableContent{constants.TYPEFLOAT, dir}
		}
	} else if c.Cte_c() != nil {
		if _, exists := ConstantsTable[c.Cte_c().GetText()]; !exists {
			dir, err := memory.Manager.GetNextAddr(constants.TYPECHAR, memory.Constant)
			if err != nil {
				log.Fatalf("Error: (EnterVarCte) %s\n", err)
			}
			ConstantsTable[c.Cte_c().GetText()] = &constantsTableContent{constants.TYPECHAR, dir}
		}
	} else if c.Cte_b() != nil {
		if _, exists := ConstantsTable[c.Cte_b().GetText()]; !exists {
			dir, err := memory.Manager.GetNextAddr(constants.TYPEBOOL, memory.Constant)
			if err != nil {
				log.Fatalf("Error: (EnterVarCte) %s\n", err)
			}
			ConstantsTable[c.Cte_b().GetText()] = &constantsTableContent{constants.TYPEBOOL, dir}
		}
	}
}

func (l *MyListener) EnterForLoop2(c *parser.ForLoop2Context) {
	id := c.ID().GetText()
	currVariable := VariableAttributes{id, constants.TYPEINT, memory.LOCAL_INT}
	l.functionTable[l.currentFunction].Vars[id] = &currVariable
}

func (l *MyListener) ExitProgram(c *parser.ProgramContext) {
	l.scopeStack.Pop()
	//l.printFunctions()
}

//Debugger Helpers
func (l *MyListener) printFunctions() {
	for fname, function := range l.functionTable {
		fmt.Printf("\n *Function: %s Returns: %s Scope: %s\n", fname, function.TypeOf, function.Scope)
		for id, variable := range function.Vars {
			fmt.Printf("%s: %s \n", id, variable.TypeOf)
		}
	}
}

func (l *MyListener) validateVariableInScope(id string) {
	if _, exists := l.functionTable[l.currentFunction].Vars[id]; exists {
		log.Fatal(id + " is already Defined")
	}
}
