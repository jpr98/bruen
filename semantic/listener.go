package semantic

import (
	"log"
	"strconv"

	"github.com/jpr98/bruen/constants"
	"github.com/jpr98/bruen/memory"
	"github.com/jpr98/bruen/parser"
	"github.com/jpr98/bruen/utils"
)

// Listener contains all logic to analyze a program on a semantic level. The main
// function of this pass over the source code is to construct and populate our
// ClassTable and FunctionTable. It manages definitions, virtual memory address
// assignation, access modifiers, etc. Almost all methods are called when the
// parser visits rule nodes and emits events.
type Listener struct {
	*parser.BaseBruenListener
	scopeStack          utils.StringStack // Keeps track of the current scope
	currentFunction     string
	ProgramName         string
	functionTable       FunctionTable
	unassignedVariables []string

	classTable            ClassTable
	isInClassAttributes   bool
	isInClassMethods      bool
	pendingAccessModifier string
}

// NewListener creates an instance of Listener and initialices all of its
// stacks, lists and maps
func NewListener() Listener {
	var listener Listener
	listener.scopeStack = utils.StringStack{}
	listener.currentFunction = ""
	listener.ProgramName = ""
	listener.functionTable = make(map[string]*FunctionTableContent)
	listener.classTable = make(map[string]*ClassTableContent)
	listener.unassignedVariables = make([]string, 0)
	return listener
}

// GetFunctionTable returns the populated FunctionTable
func (l *Listener) GetFunctionTable() FunctionTable {
	return l.functionTable
}

// GetClassTable returns the populated ClassTable
func (l *Listener) GetClassTable() ClassTable {
	return l.classTable
}

// handleFunctionRedefinition validates that a function is not previously defined
func (l *Listener) handleFunctionRedefinition(scope string) {
	if l.isInClassMethods {
		if _, exists := l.classTable[l.scopeStack.Top()].Methods[l.currentFunction]; exists {
			log.Fatal(l.currentFunction + " is already defined in scope " + l.scopeStack.Top())
		}
		if _, exists := l.classTable[l.scopeStack.Top()].Attributes[l.currentFunction]; exists {
			log.Fatal(l.currentFunction + " is already defined in scope " + l.scopeStack.Top())
		}
	} else {
		if function, exists := l.functionTable[l.currentFunction]; exists {
			if function.Scope == scope {
				log.Fatal(l.currentFunction + " is already defined in scope " + l.scopeStack.Top())
			}
		}
	}
}

// getCurrentFunctionTable returns the FunctionTable of the current scope
func (l *Listener) getCurrentFunctionTable() FunctionTable {
	if l.scopeStack.Top() == l.ProgramName {
		return l.functionTable
	}
	return l.classTable[l.scopeStack.Top()].Methods
}

// addToFunctionTable adds a function to the current FunctionTable
func (l *Listener) addToFunctionTable(typeOf constants.Type) {
	functionTable := l.getCurrentFunctionTable()
	l.handleFunctionRedefinition(l.scopeStack.Top())
	functionTable[l.currentFunction] = &FunctionTableContent{}
	functionTable[l.currentFunction].Vars = make(map[string]*VariableAttributes)
	functionTable[l.currentFunction].TypeOf = typeOf
	functionTable[l.currentFunction].Scope = l.scopeStack.Top()
}

// EnterProgram adds the global scope to the function table and to the scope stack
func (l *Listener) EnterProgram(c *parser.ProgramContext) {
	if c.ID().GetText() == "self" {
		log.Fatalln("Error: self is a reserved keyword")
	}
	l.scopeStack.Push(c.ID().GetText())
	l.currentFunction = c.ID().GetText()
	l.ProgramName = c.ID().GetText()
	l.addToFunctionTable(constants.VOID)

	// Setting a constant 1 int since almost all programs will use it
	dir, err := memory.Manager.GetNextAddr(constants.TYPEINT, memory.Constant)
	if err != nil {
		log.Fatalf("Error: (EnterProgram) setting constant 1, %s", err)
	}
	ConstantsTable["1"] = &constantsTableContent{constants.TYPEINT, dir}
}

// EnterClassDef adds a class to ClassTable and adds the class to scopeStack
func (l *Listener) EnterClassDef(c *parser.ClassDefContext) {
	if c.ID().GetText() == "self" {
		log.Fatalln("Error: self is a reserved keyword")
	}
	l.scopeStack.Push(c.ID().GetText())
	l.classTable[c.ID().GetText()] = &ClassTableContent{}
	l.classTable[c.ID().GetText()].Methods = make(FunctionTable)
	l.classTable[c.ID().GetText()].Attributes = make(map[string]*VariableAttributes)
}

// ExitClassDef removes class from scope stack
func (l *Listener) ExitClassDef(c *parser.ClassDefContext) {
	l.scopeStack.Pop()
	l.isInClassMethods = false
	// After defining the classes, we need to continue addressing variables to the program function
	l.currentFunction = l.ProgramName
}

// EnterClassInit adds the class initializer as a method to its ClassTableContent and
// gets a virtual memory address for the return value
func (l *Listener) EnterClassInit(c *parser.ClassInitContext) {
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

// ExitClassInit adds a self variable to the method (this will be the newly created instance)
// and sets the local variables memory size to the class init
func (l *Listener) ExitClassInit(c *parser.ClassInitContext) {
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

// EnterClassAttributes raises a flag indicating that we are parsing
// class attributes instead of normal variables
func (l *Listener) EnterClassAttributes(c *parser.ClassAttributesContext) {
	l.isInClassAttributes = true
}

// ExitClassAttributes sets the attributes memory size for the class and raises a flag
// indicating that we are parsing methods instead of global functions
func (l *Listener) ExitClassAttributes(c *parser.ClassAttributesContext) {
	varSize, objSize := memory.Manager.ResetLocalCounter()
	l.classTable[l.scopeStack.Top()].VarsSize = varSize
	l.classTable[l.scopeStack.Top()].ObjectCount = objSize
	// Sets objects size according to attributes which are objects (currently not supported)
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

// EnterClassMethod sets a pending access modifier for later use (defaults to public)
func (l *Listener) EnterClassMethod(c *parser.ClassMethodContext) {
	if c.PRIVATE() != nil {
		l.pendingAccessModifier = c.PRIVATE().GetText()
		return
	}
	l.pendingAccessModifier = "public"
}

// EnterFunctions adds a new function to the current FunctionTable and assigns
// a memory address for the function return value if it's not void. We also set
// the access modifier to the function and discard the previous value
func (l *Listener) EnterFunctions(c *parser.FunctionsContext) {
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

// ExitFunctions adds self as a variable for class methods only and sets the memory
// size needed to execute the function
func (l *Listener) ExitFunctions(c *parser.FunctionsContext) {
	if l.scopeStack.Top() != l.ProgramName {
		selfVariable := NewVariableAttributes("self", constants.TYPECLASS, -1)
		selfVariable.Class = l.scopeStack.Top()
		l.classTable[l.scopeStack.Top()].Methods[l.currentFunction].Vars["self"] = selfVariable
	}
	varSize, objSize := memory.Manager.ResetLocalCounter()
	l.getCurrentFunctionTable()[l.currentFunction].VarsSize = varSize
	l.getCurrentFunctionTable()[l.currentFunction].ObjectCount = objSize
}

// EnterMain sets main as the current function and adds it to the FunctionTable
func (l *Listener) EnterMain(c *parser.MainContext) {
	l.currentFunction = c.MAIN().GetText()
	l.addToFunctionTable(constants.VOID)
}

// EnterAttributesDeclaration sets a pending access modifier for later use (defaults to public)
func (l *Listener) EnterAttributesDeclaration(c *parser.AttributesDeclarationContext) {
	if c.PRIVATE() != nil {
		l.pendingAccessModifier = c.PRIVATE().GetText()
		return
	}
	l.pendingAccessModifier = "public" // hardcoded because it can be nil but public is default value
}

// EnterAttributesDec adds a new attribute to the current class and discards previous
// access modifier value
func (l *Listener) EnterAttributesDec(c *parser.AttributesDecContext) {
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

// countInObjSize adds a class name to the list of objects which a function or class
// needs to be able to hold in memory
func (l *Listener) countInObjSize(className string) {
	if l.isInClassAttributes {
		l.classTable[l.scopeStack.Top()].Objects = append(l.classTable[l.scopeStack.Top()].Objects, className)
	} else {
		l.getCurrentFunctionTable()[l.currentFunction].Objects = append(l.getCurrentFunctionTable()[l.currentFunction].Objects, className)
	}
}

// EnterVarsDec holds a declared variable id for later use (this is done because the
// grammar rule has two parts, we use an array for possiblity of declaring multiple
// variables of same type in one line although it's not currently supported)
func (l *Listener) EnterVarsDec(c *parser.VarsDecContext) {
	id := c.ID().GetText()
	if id == "self" {
		log.Fatalln("Error: self is a reserved keyword")
	}
	l.unassignedVariables = append(l.unassignedVariables, id)
}

// EnterVarsTypeInit gets the pending variables (right now it is always only one) and
// adds it to its corresponding FunctionTable
func (l *Listener) EnterVarsTypeInit(c *parser.VarsTypeInitContext) {
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

// EnterVarsDecArray adds an array variable to its corresponding FunctionTable, it
// also assigns memory addresses for the solicited array dimensions
func (l *Listener) EnterVarsDecArray(c *parser.VarsDecArrayContext) {
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

// EnterVarsDecMat should add a matrix to it's corresponding FunctionTable (currently unsupported)
func (l *Listener) EnterVarsDecMat(c *parser.VarsDecMatContext) {
	id := c.ID().GetText()
	l.unassignedVariables = append(l.unassignedVariables, id)
}

// EnterParameter adds a parameter to the local variables of a function, assigns
// it a virtual memory address and updates the expected parameter list of a function
func (l *Listener) EnterParameter(c *parser.ParameterContext) {
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

// EnterVarCte adds constants to the ConstantsTable
func (l *Listener) EnterVarCte(c *parser.VarCteContext) {
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

// EnterForLoop2 adds the for loop iterator as a local variable to the current function
func (l *Listener) EnterForLoop2(c *parser.ForLoop2Context) {
	id := c.ID().GetText()
	if id == "self" {
		log.Fatalln("Error: self is a reserved keyword")
	}

	dir, err := memory.Manager.GetNextAddr(constants.TYPEINT, memory.Local)
	if err != nil {
		log.Fatalf("Error: (AddForLoopIterator) %s\n", err)
	}

	currVariable := NewVariableAttributes(id, constants.TYPEINT, dir)
	l.getCurrentFunctionTable()[l.currentFunction].Vars[id] = currVariable
}

// ExitMain sets the memory size needed for main
func (l *Listener) ExitMain(c *parser.MainContext) {
	varSize, objSize := memory.Manager.ResetLocalCounter()
	l.functionTable[l.currentFunction].VarsSize = varSize
	l.functionTable[l.currentFunction].ObjectCount = objSize
}

// ExitProgram removes the global scope from the scope stack
func (l *Listener) ExitProgram(c *parser.ProgramContext) {
	l.scopeStack.Pop()
}

// validateVariableInScope validates that a variable hasn't been previously defined
func (l *Listener) validateVariableInScope(id string) {
	if l.isInClassAttributes {
		if _, exists := l.classTable[l.scopeStack.Top()].Attributes[id]; exists {
			log.Fatal(id + "is already defined")
		}
	} else {
		if _, exists := l.getCurrentFunctionTable()[l.currentFunction].Vars[id]; exists {
			log.Fatal(id + " is already defined")
		}
	}
}
