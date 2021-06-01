package semantic

import (
	"log"
	"strconv"

	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/memory"
	"github.com/jpr98/compis/parser"
	"github.com/jpr98/compis/utils"
)

// Our Listener that extends proyecto_listener interface generated by antlr
type MyListener struct {
	*parser.BaseProyectoListener
	scopeStack          utils.StringStack
	currentFunction     string
	ProgramName         string
	functionTable       FunctionTable
	unassignedVariables []string

	classTable            ClassTable
	isInClassAttributes   bool
	isInClassMethods      bool
	pendingAccessModifier string
}

func NewListener() MyListener {
	var listener MyListener
	listener.scopeStack = utils.StringStack{}
	listener.currentFunction = ""
	listener.ProgramName = ""
	listener.functionTable = make(map[string]*FunctionTableContent)
	listener.classTable = make(map[string]*ClassTableContent)
	listener.unassignedVariables = make([]string, 0)
	return listener
}

func (l *MyListener) GetFunctionTable() map[string]*FunctionTableContent {
	return l.functionTable
}

func (l *MyListener) GetClassTable() ClassTable {
	return l.classTable
}

// Validates uniqueness of the function ID defined by the developer
func (l *MyListener) handleFunctionRedefinition(scope string) {
	if l.isInClassMethods {
		if _, exists := l.classTable[l.scopeStack.Top()].Methods[l.currentFunction]; exists {
			log.Fatal(l.currentFunction + " is already Defined in scope " + l.scopeStack.Top())
		}
		if _, exists := l.classTable[l.scopeStack.Top()].Attributes[l.currentFunction]; exists {
			log.Fatal(l.currentFunction + " is already Defined in scope " + l.scopeStack.Top())
		}
	} else {
		if function, exists := l.functionTable[l.currentFunction]; exists {
			if function.Scope == scope {
				log.Fatal(l.currentFunction + " is already Defined in scope " + l.scopeStack.Top())
			}
		}
	}
}

func (l *MyListener) getCurrentFunctionTable() FunctionTable {
	if l.scopeStack.Top() == l.ProgramName {
		return l.functionTable
	}
	return l.classTable[l.scopeStack.Top()].Methods
}

// Adds function ID to the functionTable
func (l *MyListener) addToFunctionTable(typeOf constants.Type) {
	functionTable := l.getCurrentFunctionTable()
	l.handleFunctionRedefinition(l.scopeStack.Top())
	functionTable[l.currentFunction] = &FunctionTableContent{}
	functionTable[l.currentFunction].Vars = make(map[string]*VariableAttributes)
	functionTable[l.currentFunction].TypeOf = typeOf
	functionTable[l.currentFunction].Scope = l.scopeStack.Top()
}

func (l *MyListener) EnterProgram(c *parser.ProgramContext) {
	if c.ID().GetText() == "self" {
		log.Fatalln("Error: self is a reserved keyword")
	}
	l.scopeStack.Push(c.ID().GetText())
	l.currentFunction = c.ID().GetText()
	l.ProgramName = c.ID().GetText()
	l.addToFunctionTable(constants.VOID)

	// Setting a constant 1 int
	dir, err := memory.Manager.GetNextAddr(constants.TYPEINT, memory.Constant)
	if err != nil {
		log.Fatalf("Error: (EnterProgram) setting constant 1, %s", err)
	}
	ConstantsTable["1"] = &constantsTableContent{constants.TYPEINT, dir}
}

func (l *MyListener) EnterClassDef(c *parser.ClassDefContext) {
	if c.ID().GetText() == "self" {
		log.Fatalln("Error: self is a reserved keyword")
	}
	l.scopeStack.Push(c.ID().GetText())
	l.classTable[c.ID().GetText()] = &ClassTableContent{}
	l.classTable[c.ID().GetText()].Methods = make(FunctionTable)
	l.classTable[c.ID().GetText()].Attributes = make(map[string]*VariableAttributes)
}

// After defining the classes, we need to continue addressing variables to the program function.
func (l *MyListener) ExitClassDef(c *parser.ClassDefContext) {
	l.scopeStack.Pop()
	l.isInClassMethods = false
	l.currentFunction = l.ProgramName
}

func (l *MyListener) EnterClassInit(c *parser.ClassInitContext) {
	l.currentFunction = c.INIT().GetText()
	ftc := &FunctionTableContent{}
	ftc.TypeOf = constants.TYPECLASS
	returnDir, err := memory.Manager.GetNextAddr(constants.TYPECLASS, memory.Global)
	if err != nil {
		log.Fatalf("Error: (EnterClassInit) %s", err)
	}
	l.functionTable[l.ProgramName].Objects = append(l.functionTable[l.ProgramName].Objects, l.scopeStack.Top())

	ftc.ReturnDir = returnDir
	l.classTable[l.scopeStack.Top()].Methods[l.currentFunction] = ftc
	l.classTable[l.scopeStack.Top()].Methods[l.currentFunction].Vars = make(map[string]*VariableAttributes)
}

func (l *MyListener) ExitClassInit(c *parser.ClassInitContext) {
	selfDir, err := memory.Manager.GetNextAddr(constants.TYPECLASS, memory.Local)
	if err != nil {
		log.Fatalf("Error: (ExitClassInit) %s", err)
	}
	selfVariable := NewVariableAttributes("self", constants.TYPECLASS, selfDir)
	selfVariable.Class = l.scopeStack.Top()
	l.classTable[l.scopeStack.Top()].Methods[l.currentFunction].Vars["self"] = selfVariable
	l.countInObjSize(l.scopeStack.Top())

	varSize, objSize := memory.Manager.ResetLocalCounter()
	l.classTable[l.scopeStack.Top()].Methods[l.currentFunction].VarsSize = varSize
	l.classTable[l.scopeStack.Top()].Methods[l.currentFunction].ObjectCount = objSize
	l.currentFunction = l.scopeStack.Top()
}

func (l *MyListener) EnterClassAttributes(c *parser.ClassAttributesContext) {
	l.isInClassAttributes = true
}

func (l *MyListener) ExitClassAttributes(c *parser.ClassAttributesContext) {
	varSize, objSize := memory.Manager.ResetLocalCounter()
	l.classTable[l.scopeStack.Top()].VarsSize = varSize
	l.classTable[l.scopeStack.Top()].ObjectCount = objSize
	// Sets objects size according to attributes which are objects
	if len(l.classTable[l.scopeStack.Top()].Objects) == objSize {
		for _, objName := range l.classTable[l.scopeStack.Top()].Objects {
			for _, objInfo := range l.classTable[objName].ObjSize {
				l.classTable[l.scopeStack.Top()].ObjSize = append(l.classTable[l.scopeStack.Top()].ObjSize, objInfo)
			}
		}
	}
	l.classTable[l.scopeStack.Top()].Objects = nil
	l.classTable[l.scopeStack.Top()].ObjectCount = 0
	l.isInClassAttributes = false
	l.isInClassMethods = true
}

func (l *MyListener) EnterClassMethod(c *parser.ClassMethodContext) {
	if c.PRIVATE() != nil {
		l.pendingAccessModifier = c.PRIVATE().GetText()
		return
	}
	l.pendingAccessModifier = "public"
}

func (l *MyListener) EnterFunctions(c *parser.FunctionsContext) {
	if c.ID().GetText() == "self" {
		log.Fatalln("Error: self is a reserved keyword")
	}
	l.currentFunction = c.ID().GetText()
	var typeOf constants.Type
	var returnDir int
	var err error
	if c.TypeRule() != nil {
		typeOf = constants.StringToType(c.TypeRule().GetText())
		returnDir, err = memory.Manager.GetNextAddr(typeOf, memory.Global)
		if err != nil {
			log.Fatalf("Error: (EnterFunctions) %s", err)
		}
	} else if c.VOID() != nil {
		typeOf = constants.VOID
	}

	l.addToFunctionTable(typeOf)
	l.getCurrentFunctionTable()[l.currentFunction].ReturnDir = returnDir
	l.getCurrentFunctionTable()[l.currentFunction].IsPrivate = l.pendingAccessModifier == "private"
	l.pendingAccessModifier = ""
}

func (l *MyListener) ExitFunctions(c *parser.FunctionsContext) {
	if l.scopeStack.Top() != l.ProgramName {
		selfVariable := NewVariableAttributes("self", constants.TYPECLASS, -1)
		selfVariable.Class = l.scopeStack.Top()
		l.classTable[l.scopeStack.Top()].Methods[l.currentFunction].Vars["self"] = selfVariable
	}
	varSize, objSize := memory.Manager.ResetLocalCounter()
	l.getCurrentFunctionTable()[l.currentFunction].VarsSize = varSize
	l.getCurrentFunctionTable()[l.currentFunction].ObjectCount = objSize
}

func (l *MyListener) EnterMain(c *parser.MainContext) {
	l.currentFunction = c.MAIN().GetText()
	l.addToFunctionTable(constants.VOID)
}

func (l *MyListener) EnterAttributesDeclaration(c *parser.AttributesDeclarationContext) {
	if c.PRIVATE() != nil {
		l.pendingAccessModifier = c.PRIVATE().GetText()
		return
	}
	l.pendingAccessModifier = "public" // hardcoded because it can be nil but public is default value
}

func (l *MyListener) EnterAttributesDec(c *parser.AttributesDecContext) {
	id := c.ID().GetText()
	if id == "self" {
		log.Fatalln("Error: self is a reserved keyword")
	}

	t := constants.StringToType(c.TypeRule().GetText())
	if t == constants.ERR {
		log.Fatalf("Error: (EnterVarsTypeInit) unkown type from %s", c.TypeRule().GetText())
	}

	var memCtx memory.Context
	if l.scopeStack.Top() == l.ProgramName && l.currentFunction == l.ProgramName {
		memCtx = memory.Global
	} else {
		memCtx = memory.Local
	}

	dir, err := memory.Manager.GetNextAddr(t, memCtx)
	if err != nil {
		log.Fatalf("Error: (EnterVarsTypeInit) %s\n", err)
	}

	currVariable := NewVariableAttributes(id, t, dir)
	currVariable.IsPrivate = l.pendingAccessModifier == "private"
	l.pendingAccessModifier = ""
	l.classTable[l.scopeStack.Top()].Attributes[id] = currVariable
}

func (l *MyListener) countInObjSize(className string) {
	if l.isInClassAttributes {
		l.classTable[l.scopeStack.Top()].Objects = append(l.classTable[l.scopeStack.Top()].Objects, className)
	} else {
		l.getCurrentFunctionTable()[l.currentFunction].Objects = append(l.getCurrentFunctionTable()[l.currentFunction].Objects, className)
	}
}

func (l *MyListener) EnterVarsDec(c *parser.VarsDecContext) {
	id := c.ID().GetText()
	if id == "self" {
		log.Fatalln("Error: self is a reserved keyword")
	}
	l.unassignedVariables = append(l.unassignedVariables, id)
}

func (l *MyListener) EnterVarsTypeInit(c *parser.VarsTypeInitContext) {
	for _, id := range l.unassignedVariables {
		l.validateVariableInScope(id)

		var t constants.Type
		if c.TypeRule() != nil {
			t = constants.StringToType(c.TypeRule().GetText())
			if t == constants.ERR {
				log.Fatalf("Error: (EnterVarsTypeInit) unkown type from %s", c.TypeRule().GetText())
			}

		} else if c.ID() != nil {
			t = constants.TYPECLASS
		}

		var memCtx memory.Context
		if l.scopeStack.Top() == l.ProgramName && l.currentFunction == l.ProgramName {
			memCtx = memory.Global
		} else {
			memCtx = memory.Local
		}

		dir, err := memory.Manager.GetNextAddr(t, memCtx)
		if err != nil {
			log.Fatalf("Error: (EnterVarsTypeInit) %s\n", err)
		}

		currVariable := NewVariableAttributes(id, t, dir)
		if t == constants.TYPECLASS {
			currVariable.Class = c.ID().GetText()
			l.countInObjSize(c.ID().GetText())
		}

		if l.isInClassAttributes {
			currVariable.IsPrivate = l.pendingAccessModifier == "private"
			l.pendingAccessModifier = ""
			l.classTable[l.scopeStack.Top()].Attributes[id] = currVariable
		} else {
			l.getCurrentFunctionTable()[l.currentFunction].Vars[id] = currVariable
		}
	}
	l.unassignedVariables = nil
}

func (l *MyListener) EnterVarsDecArray(c *parser.VarsDecArrayContext) {
	id := c.ID().GetText()
	if id == "self" {
		log.Fatalln("Error: self is a reserved keyword")
	}
	l.validateVariableInScope(id)

	dim, err := strconv.Atoi(c.INT().GetText())
	if err != nil {
		log.Fatalf("Error: (EnterVarsDecArray) dimension can't be parsed to int")
	}

	if dim < 1 {
		log.Fatalf("Error: Array %s needs to be at least of size 1", id)
	}

	var currVariable *VariableAttributes
	var t constants.Type
	var memCtx memory.Context
	if l.scopeStack.Top() == l.ProgramName && l.currentFunction == l.ProgramName {
		memCtx = memory.Global
	} else {
		memCtx = memory.Local
	}

	if c.TypeRule() != nil {
		t = constants.StringToType(c.TypeRule().GetText())
		if t == constants.ERR {
			log.Fatalf("Error: (EnterVarsTypeInit) unkown type from %s", c.TypeRule().GetText())
		}

		dir, err := memory.Manager.GetNextAddr(t, memCtx)
		if err != nil {
			log.Fatalf("Error: (EnterVarsTypeInit) %s\n", err)
		}

		currVariable = NewVariableAttributes(id, t, dir)

	} /*else if c.ID(1) != nil { Arreglos de objetos
		// TODO: Asegurarnos que la clase existe
		if _, exists := l.classTable[c.ID(1).GetText()]; !exists {
			log.Fatalf("Error: Undefined type %s", c.ID(0).GetText())
		}
		currVariable = NewVariableAttributes(id, constants.TYPECLASS, 20000)
		currVariable.Class = c.ID(1).GetText()
	}*/

	for i := 1; i < dim; i++ {
		_, err := memory.Manager.GetNextAddr(t, memCtx)
		if err != nil {
			log.Fatalf("Error: (EnterVarsDecArray) %s", err)
		}
	}

	currVariable.ArrayOrMat = 1
	currVariable.Dim[0] = dim
	if l.isInClassAttributes {
		currVariable.IsPrivate = l.pendingAccessModifier == "private"
		l.pendingAccessModifier = ""
		l.classTable[l.scopeStack.Top()].Attributes[id] = currVariable
	} else {
		l.getCurrentFunctionTable()[l.currentFunction].Vars[id] = currVariable
	}
}

func (l *MyListener) EnterVarsDecMat(c *parser.VarsDecMatContext) {
	id := c.ID().GetText()
	l.unassignedVariables = append(l.unassignedVariables, id)
}

func (l *MyListener) EnterParameter(c *parser.ParameterContext) {
	id := c.ID().GetText()
	if id == "self" {
		log.Fatalln("Error: self is a reserved keyword")
	}
	t := constants.StringToType(c.TypeRule().GetText())
	if t == constants.ERR {
		log.Fatalf("Error: (EnterParameter) unkown type from %s", c.TypeRule().GetText())
	}

	dir, err := memory.Manager.GetNextAddr(t, memory.Local)
	if err != nil {
		log.Fatalf("Error: (EnterParameter) %s\n", err)
	}

	// TODO: Manejar arreglos en ultimo atributo
	currVariable := NewVariableAttributes(id, t, dir)
	l.getCurrentFunctionTable()[l.currentFunction].Vars[id] = currVariable
	l.getCurrentFunctionTable()[l.currentFunction].Params = append(l.getCurrentFunctionTable()[l.currentFunction].Params, dir)
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
	if id == "self" {
		log.Fatalln("Error: self is a reserved keyword")
	}
	currVariable := NewVariableAttributes(id, constants.TYPEINT, memory.LOCAL_INT)
	l.getCurrentFunctionTable()[l.currentFunction].Vars[id] = currVariable
}

func (l *MyListener) ExitMain(c *parser.MainContext) {
	varSize, objSize := memory.Manager.ResetLocalCounter()
	l.functionTable[l.currentFunction].VarsSize = varSize
	l.functionTable[l.currentFunction].ObjectCount = objSize
}

func (l *MyListener) ExitProgram(c *parser.ProgramContext) {
	l.scopeStack.Pop()
}

func (l *MyListener) validateVariableInScope(id string) {
	if l.isInClassAttributes {
		if _, exists := l.classTable[l.scopeStack.Top()].Attributes[id]; exists {
			log.Fatal(id + "is already Defined")
		}
	} else {
		if _, exists := l.getCurrentFunctionTable()[l.currentFunction].Vars[id]; exists {
			log.Fatal(id + " is already Defined")
		}
	}
}
