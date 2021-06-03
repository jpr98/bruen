package quads

import (
	"fmt"
	"log"
	"strings"

	"github.com/jpr98/bruen/constants"
	"github.com/jpr98/bruen/memory"
	"github.com/jpr98/bruen/semantic"
	"github.com/jpr98/bruen/utils"
)

// Manager contains all the logic for generating code (Quad) for the virtual machine.
// This struct contains stacks to know the context in which the parser is parsing a
// rule and uses this information to create virtual machine instructions.
type Manager struct {
	operands   ElementStack      // Operand stack (i.e. Elements)
	operators  QuadActionStack   // Operator stack (i.e. Actions)
	quads      []Quad            // Generated Quads
	jumpStack  []int             // Contains the instruction pointer for Quads with pending updates
	scopeStack utils.StringStack // Keeps track of the current scope

	functionTable semantic.FunctionTable
	classTable    semantic.ClassTable

	globalName      string // Name of the program
	currentFunction string // Name of the function we are currently parsing, no need for stack bc no nested functions

	currentFunctionCall      utils.StringStack // Keeps track of the current function call
	currentFunctionCallClass utils.StringStack // Keeps track of the current instance (for method calls or attribute access)
	paramCounter             utils.IntStack    // Keeps track of the number of parameters parsed for current function call
	avail                    int               // Counter used to give debug names to temporary addresses
}

// GetQuads returns the generated code after successful compilation
func (m Manager) GetQuads() []Quad {
	return m.quads
}

// NewManager creates a Manager given a FunctionTable and a ClassTable
func NewManager(functionTable semantic.FunctionTable, classTable semantic.ClassTable) Manager {
	return Manager{
		operands:      NewElementStack(),
		operators:     NewQuadActionStack(),
		quads:         make([]Quad, 0),
		scopeStack:    utils.StringStack{},
		functionTable: functionTable,
		classTable:    classTable,
		avail:         0,
	}
}

// PushOp adds an operator to the operators stack
func (m *Manager) PushOp(op string) {
	opCode, err := stringToOp(op)
	if err != nil {
		log.Fatalln(err)
		return
	}

	m.operators.Push(opCode)
}

// PushConstantOperand adds a previously defined constant as an operand to the operand stack
func (m *Manager) PushConstantOperand(operand string) {
	op, exists := semantic.ConstantsTable[operand]
	if !exists {
		log.Fatalf("Error: (PushConstantOperand) %s is not a known constant", operand)
	}
	element := NewElement(op.Dir, operand, op.TypeOf, "")
	m.operands.Push(element)
}

// getOperandData validates that a variable is declared, checks if it is an attribute
// access and returns its VariableAttributes
//
// Access to attributes:
//
// a.b = explicit = a is the instance
//
// b = implicit = we assume self is the instance
func (m *Manager) getOperandData(operand string) *semantic.VariableAttributes {
	data := &semantic.VariableAttributes{TypeOf: constants.ERR}

	// Explicit access to attribute
	if strings.Contains(operand, ".") {
		strElements := strings.Split(operand, ".")
		if len(strElements) != 2 {
			log.Fatalf("Error: (getOperandData) unexpected instance access syntax\n")
		}

		instance, attribute := strElements[0], strElements[1]
		instanceData := &semantic.VariableAttributes{TypeOf: constants.ERR}
		if attr, exists := m.getCurrentFunctionTable()[m.currentFunction].Vars[instance]; exists {
			instanceData = attr
		} else {
			if attr, exists := m.functionTable[m.globalName].Vars[instance]; exists {
				instanceData = attr
			}
		}

		// Validate the required instance has been declared
		if instanceData.TypeOf == constants.ERR {
			log.Fatalf(
				"Error: (PushOperand) undeclared variable %s\n",
				instance,
			)
			return nil
		}

		if attr, exists := m.classTable[instanceData.Class].Attributes[attribute]; exists {
			data = attr
			data.FromSelf = true
			data.SelfDir = instanceData.Dir

			if data.IsPrivate && m.scopeStack.Top() != instanceData.Class {
				log.Fatalf("Error: attribute %s is private\n", attribute)
				return nil
			}
		}

		// Validate that attribute attributes have been found
		if data.TypeOf == constants.ERR {
			log.Fatalf(
				"Error: (PushOperand) undeclared attribute %s in class %s\n",
				attribute,
				instanceData.Class,
			)
			return nil
		}
		return data
	}

	if attr, exists := m.getCurrentFunctionTable()[m.currentFunction].Vars[operand]; exists {
		data = attr
	} else if m.scopeStack.Top() != m.globalName {
		// Implicit access to attribute
		if attr, exists := m.classTable[m.scopeStack.Top()].Attributes[operand]; exists {
			data = attr
			data.FromSelf = true
			data.SelfDir = m.classTable[m.scopeStack.Top()].Methods[m.currentFunction].Vars["self"].Dir
		}
	} else {
		if attr, exists := m.functionTable[m.globalName].Vars[operand]; exists {
			data = attr
		}
	}

	// Validate that variable attributes have been found
	if data.TypeOf == constants.ERR {
		log.Fatalf(
			"Error: (PushOperand) undeclared variable %s\n",
			operand,
		)
		return nil
	}
	return data
}

// PushOperand adds an operand (Element) to the operands stack after getting it's attributes
func (m *Manager) PushOperand(operand string) {
	operandData := m.getOperandData(operand)
	operandName := operandData.Name
	// If the operand is an attribute from an instance we set tag in it's ID
	if operandData.FromSelf {
		operandName = fmt.Sprintf("self_%d_%s", operandData.SelfDir, operandData.Name)
	}
	element := NewElement(operandData.Dir, operandName, operandData.TypeOf, operandData.Class)
	m.operands.Push(element)
}

// GenerateQuad creates a quad for an expression given that the top of the operators
// belongs to a given set of Actions
//
// isForLoop is used to retain the iterator variable for next iterations of the loop
func (m *Manager) GenerateQuad(validOps []int, isForLoop bool) {
	// 1: Validate that top operator is in validOps
	if arrayContainsElement(m.operators.Top(), validOps) {
		// 2: Get right and left operands (in that order)
		rOperand := m.operands.Pop()
		var lOperand Element
		if isForLoop {
			lOperand = m.operands.Top()
		} else {
			lOperand = m.operands.Pop()
		}
		// 4: Remove operator from stack
		op := m.operators.Pop()

		// 5: Get operation result type given the operands and operator
		resultType := semantic.Cube.ValidateBinaryOperation(lOperand.Type(), rOperand.Type(), int(op))
		// 6.a: Validate that result type is not an error
		if resultType == constants.ERR {
			log.Fatalf(
				"Error: (GenerateQuad) type mismatch, operator (%s) can't be done between %s and %s\n",
				op,
				lOperand,
				rOperand,
			)
		}
		// 6.b: If result type is constants.TYPECLASS validate that both operands' classes are the same
		if resultType == constants.TYPECLASS {
			if lOperand.ClassName() != rOperand.ClassName() {
				log.Fatalf(
					"Error: (GenerateQuad) type mismatch, operator (%s) can't be done between %s and %s\n",
					op,
					lOperand,
					rOperand,
				)
			}
		}
		// 7: Get memory address for result
		dir, err := memory.Manager.GetNextAddr(resultType, memory.Temp)
		if err != nil {
			log.Fatalf("Error: (GenerateQuad) %s\n", err)
		}
		result := NewElement(dir, m.getNextAvail(), resultType, lOperand.ClassName())

		// 8: Create quad
		q := Quad{op, lOperand, rOperand, result}

		// 9: Push quad to list of quads
		m.quads = append(m.quads, q)

		// 10: Add result to operands stack
		m.operands.Push(result)
	}
}

// GenerateAssignationQuad creates an assignation quad between the last two operands
// in the stack. The last operand is the value to assign and the second to last is
// where to assign the value.
//
// If retainresult is true we push back the element where the value was just assigned.
// This is mainly used to retain the for loop iterator for next iterations.
func (m *Manager) GenerateAssignationQuad(retainresult bool) {
	op := m.operators.Pop()
	operand := m.operands.Pop() // Value to assign
	result := m.operands.Pop()  // Place to assing value

	resultType := semantic.Cube.ValidateBinaryOperation(result.Type(), operand.Type(), constants.OPASSIGN)
	if resultType == constants.ERR {
		log.Fatalf(
			"Error: (GenerateAssignationQuad) type mismatch, assignation to %s can't be %s\n",
			result,
			operand,
		)
	}
	if resultType == constants.TYPECLASS {
		if operand.ClassName() != result.ClassName() {
			log.Fatalf(
				"Error: (GenerateAssignationQuad) type mismatch, operator (%s) can't be done between %s and %s\n",
				op,
				operand,
				result,
			)
		}
	}

	if retainresult {
		m.operands.Push(result)
	}
	q := Quad{op, operand, nil, result}
	m.quads = append(m.quads, q)
}

// AddVerifyQuad creates a quad which validates that the top operand is a valid
// index to access an array.
//
// isArray is set to true if the id represents an array and false if it is a matrix
// (currently unsupported).
func (m *Manager) AddVerifyQuad(id string, isArray bool) {
	index := m.operands.Top()
	if index.Type() != constants.TYPEINT {
		log.Fatalf("Error: (AddVerifyQuad) index to %s should be an int, got %s", id, index.Type())
	}

	if isArray {
		array := m.getOperandData(id)
		maxIndexVal := NewElement(array.Dim[0], array.Name, constants.TYPEINT, "")
		q := Quad{VER, index, nil, maxIndexVal}
		m.quads = append(m.quads, q)
	}
}

// AddArrayAccessQuad generates a quad that makes the calculation for the actual
// array address given the index found in the top of operands stack.
func (m *Manager) AddArrayAccessQuad() {
	index := m.operands.Pop()
	array := m.operands.Pop()

	dir, err := memory.Manager.GetNextAddr(constants.TYPEINT, memory.Temp)
	if err != nil {
		log.Fatalf("Error: (GenerateQuad) %s\n", err)
	}
	var result Element
	// validate if array is an attribute from an instance, set self dir if so
	if strings.Contains(array.ID(), "self_") {
		strElements := strings.Split(array.ID(), "_")
		ptrID := fmt.Sprintf("self_%s_ptr_%s", strElements[1], m.getNextAvail())
		result = NewElement(dir, ptrID, array.Type(), "")
	} else {
		result = NewElement(dir, "ptr_"+m.getNextAvail(), array.Type(), "")
	}

	q := Quad{CALCDIR, array, index, result}
	m.quads = append(m.quads, q)

	m.operands.Push(result)
}

// AddGotoF generates a GOTOF quad given the last operand in the stack and saves the
// instruction pointer to the jumpstack.
func (m *Manager) AddGotoF() {
	m.jumpStack = append(m.jumpStack, len(m.quads))
	operand := m.operands.Pop()
	if operand.Type() != constants.TYPEBOOL {
		log.Fatalf(
			"Error: (AddGotoF) if statement needs bool, received %s\n",
			operand,
		)
	}
	q := Quad{GOTOF, operand, nil, nil}
	m.quads = append(m.quads, q)
}

// AddAndUpdateGoto generates a GOTO quad and updates the quad that the top of
// the jumpstack points to. This is done in a specific order so that the new
// quad doesn't affect the Result address of the past quad.
func (m *Manager) AddAndUpdateGoto() {
	// Get pos of quad to update
	pos := m.jumpStack[len(m.jumpStack)-1]
	m.jumpStack = m.jumpStack[:len(m.jumpStack)-1]

	// Add goto quad
	m.jumpStack = append(m.jumpStack, len(m.quads))
	q := Quad{GOTO, nil, nil, nil}
	m.quads = append(m.quads, q)

	// Update quad
	m.quads[pos].Result = NewElement(len(m.quads), "", constants.ADDR, "")
}

// UpdateGoto updates the quad that the top of the jumpstack points to.
func (m *Manager) UpdateGoto() {
	pos := m.jumpStack[len(m.jumpStack)-1]
	m.jumpStack = m.jumpStack[:len(m.jumpStack)-1]

	m.quads[pos].Result = NewElement(len(m.quads), "", constants.ADDR, "")
}

// SaveJumpPosition adds the current instruction pointer to the jumpstack.
func (m *Manager) SaveJumpPosition() {
	m.jumpStack = append(m.jumpStack, len(m.quads))
}

// AddAndUpdateWhileGotos generates a GOTO quad for a while loop and updates the GOTOF
// result address.
func (m *Manager) AddAndUpdateWhileGotos() {
	posF := m.jumpStack[len(m.jumpStack)-1]
	m.jumpStack = m.jumpStack[:len(m.jumpStack)-1]

	posR := m.jumpStack[len(m.jumpStack)-1]
	m.jumpStack = m.jumpStack[:len(m.jumpStack)-1]

	q := Quad{GOTO, nil, nil, NewElement(posR, "", constants.ADDR, "")}
	m.quads = append(m.quads, q)

	m.quads[posF].Result = NewElement(len(m.quads), "", constants.ADDR, "")
}

// AddSimpleGOTO adds an empty GOTO quad.
func (m *Manager) AddSimpleGOTO() {
	q := Quad{GOTO, nil, nil, nil}
	m.quads = append(m.quads, q)
}

// AddForLoopIterator pushes the for loop iterator for the operands stack.
func (m *Manager) AddForLoopIterator(id string) {
	attr, exists := m.getCurrentFunctionTable()[m.currentFunction].Vars[id]
	if !exists {
		log.Fatalf("Error: (AddForLoopIterator) undeclared variable %s", id)
	}

	e := NewElement(attr.Dir, id, attr.TypeOf, "")
	m.operands.Push(e)
}

// AddReadQuad creates a read quad with the top operand.
func (m *Manager) AddReadQuad() {
	operand := m.operands.Pop()
	q := Quad{READ, nil, nil, operand}
	m.quads = append(m.quads, q)
}

// AddWriteQuad generates a write quad with the top operand.
func (m *Manager) AddWriteQuad() {
	operand := m.operands.Pop()
	q := Quad{WRITE, nil, nil, operand}
	m.quads = append(m.quads, q)
}

// AddNewLineWriteQuad generates an empty WRITENEWLINE quad.
func (m *Manager) AddNewLineWriteQuad() {
	q := Quad{WRITENEWLINE, nil, nil, nil}
	m.quads = append(m.quads, q)
}

// AddEndFuncQuad generates an empty ENDFUNC quad.
func (m *Manager) AddEndFuncQuad() {
	q := Quad{ENDFUNC, nil, nil, nil}
	m.quads = append(m.quads, q)
}

// getCurrentFunctionTable returns the FunctionTable of the current sope.
func (m *Manager) getCurrentFunctionTable() semantic.FunctionTable {
	if m.scopeStack.Top() == m.globalName {
		return m.functionTable
	}
	return m.classTable[m.scopeStack.Top()].Methods
}

// AddReturnQuad generates a Return quad. If the currentFunction return type is Void
// the quad is empty. If the currentFunction has a return type other than Void it validates
// the the return type and the value type match, then it sets the return address and return
// value as elements in the quad.
func (m *Manager) AddReturnQuad() {
	if m.getCurrentFunctionTable()[m.currentFunction].TypeOf == constants.VOID {
		q := Quad{RETURN, nil, nil, nil}
		m.quads = append(m.quads, q)
		return
	}

	returnValue := m.operands.Pop()
	functionType := m.getCurrentFunctionTable()[m.currentFunction].TypeOf
	if returnValue.Type() != functionType {
		log.Fatalf("Error: (AddReturnQuad) %s expects to return type %s", m.currentFunction, functionType)
	}

	returnDir := m.getCurrentFunctionTable()[m.currentFunction].ReturnDir
	funcReturnElement := NewElement(returnDir, "", constants.ADDR, "")
	q := Quad{RETURN, funcReturnElement, nil, returnValue}
	m.quads = append(m.quads, q)
}

// AddInitReturnQuad generates a return quad for a class init. Class inits always
// return the self local variable, which holds the newly created instance of the class.
func (m *Manager) AddInitReturnQuad() {
	selfAtrr := m.getCurrentFunctionTable()[m.currentFunction].Vars["self"]
	returnValue := NewElement(selfAtrr.Dir, selfAtrr.Name, selfAtrr.TypeOf, selfAtrr.Class)

	returnDir := m.getCurrentFunctionTable()[m.currentFunction].ReturnDir
	initReturnElement := NewElement(returnDir, "", constants.ADDR, "")
	q := Quad{RETURN, initReturnElement, nil, returnValue}
	m.quads = append(m.quads, q)
}

// AddEraQuadgenerates an ERA quad for the given function name. If the function happens
// to be a method, then className won't be empty. We also push the current function call
// name and class to their stacks and push a new param counter.
func (m *Manager) AddEraQuad(name, className string) {
	n := NewElement(0, name, constants.ADDR, className)
	q := Quad{ERA, n, nil, nil}
	m.quads = append(m.quads, q)

	m.currentFunctionCall.Push(name)
	m.currentFunctionCallClass.Push(className)
	m.paramCounter.Push(0)
}

// AddParamQuad validates that the current function doesn't have extra arguments, checks
// that the argument matches the expected argument type in it's argument position and then
// creates a PARAM quad. This is only for global functions!
func (m *Manager) AddParamQuad() {
	if m.paramCounter.Top() >= len(m.functionTable[m.currentFunctionCall.Top()].Params) {
		log.Fatalf(
			"Error: (AddGoSubQuad) function %s has too many arguments %d",
			m.currentFunctionCall,
			m.paramCounter,
		)
	}

	expectedType := memory.TypeForAddr(m.functionTable[m.currentFunctionCall.Top()].Params[m.paramCounter.Top()])
	arg := m.operands.Pop()
	t := semantic.Cube.ValidateBinaryOperation(expectedType, arg.Type(), int(constants.OPASSIGN))
	if t == constants.ERR {
		log.Fatalf(
			"Error: (AddParamQuad) type mismatch, parameter %s must be of type %s",
			arg,
			expectedType,
		)
	}

	argNum := NewElement(m.functionTable[m.currentFunctionCall.Top()].Params[m.paramCounter.Top()], fmt.Sprintf("%d", m.paramCounter), expectedType, "")
	q := Quad{PARAM, arg, argNum, nil}
	m.quads = append(m.quads, q)
}

// AddClassParamQuad validates that the current method doesn't have extra arguments, checks
// that the argument matches the expected argument type in it's argument position and then
// creates a PARAM quad. This is only for class methods!
func (m *Manager) AddClassParamQuad() {
	if m.paramCounter.Top() >= len(m.classTable[m.currentFunctionCallClass.Top()].Methods[m.currentFunctionCall.Top()].Params) {
		log.Fatalf(
			"Error: (AddGoSubQuad) function %s has too many arguments",
			m.currentFunctionCall,
		)
	}

	expectedType := memory.TypeForAddr(m.classTable[m.currentFunctionCallClass.Top()].Methods[m.currentFunctionCall.Top()].Params[m.paramCounter.Top()])
	arg := m.operands.Pop()
	t := semantic.Cube.ValidateBinaryOperation(expectedType, arg.Type(), int(constants.OPASSIGN))
	if t == constants.ERR {
		log.Fatalf(
			"Error: (AddParamQuad) type mismatch, parameter %s must be of type %s",
			arg,
			expectedType,
		)
	}

	argNum := NewElement(m.classTable[m.currentFunctionCallClass.Top()].Methods[m.currentFunctionCall.Top()].Params[m.paramCounter.Top()], fmt.Sprintf("%d", m.paramCounter), expectedType, "")
	q := Quad{PARAM, arg, argNum, nil}
	m.quads = append(m.quads, q)
}

// AddGoSubQuad validates that the function has enough arguments and generates a GOSUB quad.
// Then if the function return type is different then Void, we generate assignation quad to
// handle the return value held in the global return address for the function. Finally we
// remove the function call, class and param counter from their respective stacks. This is only
// for global functions!
func (m *Manager) AddGoSubQuad() {
	if m.paramCounter.Top() < len(m.functionTable[m.currentFunctionCall.Top()].Params) {
		log.Fatalf(
			"Error: (AddGoSubQuad) function %s has too few arguments",
			m.currentFunctionCall.Top(),
		)
	}

	dir := m.functionTable[m.currentFunctionCall.Top()].Dir
	n := NewElement(0, m.currentFunctionCall.Top(), constants.TYPEINT, "")
	dirElement := NewElement(dir, "", constants.ADDR, "")
	q := Quad{GOSUB, n, nil, dirElement}
	m.quads = append(m.quads, q)

	if m.functionTable[m.currentFunctionCall.Top()].TypeOf != constants.VOID {
		resultType := m.functionTable[m.currentFunctionCall.Top()].TypeOf
		if resultType == constants.ERR {
			log.Fatalf(
				"Error: (AddGoSubQuad) error in return type %s",
				m.functionTable[m.currentFunctionCall.Top()].TypeOf,
			)
		}

		dir, err := memory.Manager.GetNextAddr(resultType, memory.Temp)
		if err != nil {
			log.Fatalf("Error: (AddGoSubQuad) %s\n", err)
		}
		result := NewElement(dir, m.getNextAvail(), resultType, "")
		m.operands.Push(result)

		returnDir := m.functionTable[m.currentFunctionCall.Top()].ReturnDir
		funcReturnElement := NewElement(returnDir, "", constants.ADDR, "")
		q := Quad{ASSIGN, funcReturnElement, nil, result}
		m.quads = append(m.quads, q)
	}
	m.currentFunctionCall.Pop()
	m.currentFunctionCallClass.Pop()
	m.paramCounter.Pop()
}

// AddGoSubQuad validates that the method has enough arguments and generates a GOSUB quad.
// Then if the method return type is different then Void, we generate assignation quad to
// handle the return value held in the global return address for the method. Finally we
// remove the function call, class and param counter from their respective stacks. This is only
// for class methods!
func (m *Manager) AddClassGoSubQuad(className, instance string) {
	if m.paramCounter.Top() < len(m.classTable[className].Methods[m.currentFunctionCall.Top()].Params) {
		log.Fatalf(
			"Error: (AddClassGoSubQuad) function %s has too few arguments",
			m.currentFunctionCall.Top(),
		)
	}

	if m.currentFunctionCall.Top() != "init" {
		selfDir := m.getCurrentFunctionTable()[m.currentFunction].Vars[instance].Dir
		m.classTable[className].Methods[m.currentFunctionCall.Top()].Vars["self"].Dir = selfDir
		instanceName := fmt.Sprintf("self_%d_%s", selfDir, instance)
		instanceElement := NewElement(selfDir, instanceName, constants.ADDR, className)
		q := Quad{INSTANCE, instanceElement, nil, nil}
		m.quads = append(m.quads, q)
	}

	dir := m.classTable[className].Methods[m.currentFunctionCall.Top()].Dir
	n := NewElement(0, m.currentFunctionCall.Top(), constants.TYPEINT, className)
	dirElement := NewElement(dir, "", constants.ADDR, "")
	q := Quad{GOSUB, n, nil, dirElement}
	m.quads = append(m.quads, q)

	if m.classTable[className].Methods[m.currentFunctionCall.Top()].TypeOf != constants.VOID {
		resultType := m.classTable[className].Methods[m.currentFunctionCall.Top()].TypeOf
		if resultType == constants.ERR {
			log.Fatalf(
				"Error: (AddClassGoSubQuad) error in return type %s",
				m.classTable[className].Methods[m.currentFunctionCall.Top()].TypeOf,
			)
		}

		dir, err := memory.Manager.GetNextAddr(resultType, memory.Temp)
		if err != nil {
			log.Fatalf("Error: (AddClassGoSubQuad) %s\n", err)
		}

		var resultClass string
		if m.currentFunctionCall.Top() == "init" {
			resultClass = className
			m.getCurrentFunctionTable()[m.currentFunction].Objects = append(m.getCurrentFunctionTable()[m.currentFunction].Objects, className)
		}
		result := NewElement(dir, m.getNextAvail(), resultType, resultClass)
		m.operands.Push(result)

		returnDir := m.classTable[className].Methods[m.currentFunctionCall.Top()].ReturnDir
		funcReturnElement := NewElement(returnDir, "", constants.ADDR, "")
		q := Quad{ASSIGN, funcReturnElement, nil, result}
		m.quads = append(m.quads, q)
	}
	m.currentFunctionCall.Pop()
	m.currentFunctionCallClass.Pop()
	m.paramCounter.Pop()
}

// CheckImplicitMethodCall decides if a function call should be thought of as a method
// this decision is made by checking if the current scope is a class and by checking if the
// current class contains a method with the given name. If the function is indeed a method call
// this function returns the class it belongs to.
func (m *Manager) CheckImplicitMethodCall(name string) (bool, string) {
	if m.scopeStack.Top() == m.globalName {
		return false, ""
	}
	_, exists := m.classTable[m.scopeStack.Top()].Methods[name]
	return exists, m.scopeStack.Top()
}

// getNextAvail returns a string with format t# that represents a temporary address (only used
// for debugging purposes)
func (m *Manager) getNextAvail() string {
	defer func() { m.avail++ }()
	return fmt.Sprintf("t%d", m.avail)
}

// arrayContainsElement validates that an array contains a specific action
func arrayContainsElement(element QuadAction, array []int) bool {
	for _, value := range array {
		if int(element) == value {
			return true
		}
	}
	return false
}

// setObjectSize sets the needed memory size to handle objects inside a function. We do
// this by validating that the length of the class names array (Objects) matches the expected
// object count in a function, then we append the MemObjInfo from the given class to the
// ObjSize array of the current function. Finally we delete the Objects array and counter.
func (m *Manager) setObjectSize() {
	if len(m.getCurrentFunctionTable()[m.currentFunction].Objects) == m.getCurrentFunctionTable()[m.currentFunction].ObjectCount {
		for _, objName := range m.getCurrentFunctionTable()[m.currentFunction].Objects {
			obj := memory.MemObjInfo{VarSize: m.classTable[objName].VarsSize, ObjSize: m.classTable[objName].ObjSize}
			m.getCurrentFunctionTable()[m.currentFunction].ObjSize = append(m.getCurrentFunctionTable()[m.currentFunction].ObjSize, obj)
		}
	}
	m.getCurrentFunctionTable()[m.currentFunction].Objects = nil
	m.getCurrentFunctionTable()[m.currentFunction].ObjectCount = 0
}

func stringToOp(text string) (QuadAction, error) {
	switch text {
	case ADD.String():
		return ADD, nil
	case SUB.String():
		return SUB, nil
	case MUL.String():
		return MUL, nil
	case DIV.String():
		return DIV, nil
	case GT.String():
		return GT, nil
	case LT.String():
		return LT, nil
	case EQ.String():
		return EQ, nil
	case NEQ.String():
		return NEQ, nil
	case OR.String():
		return OR, nil
	case AND.String():
		return AND, nil
	case ASSIGN.String():
		return ASSIGN, nil
	case LPAREN.String():
		return LPAREN, nil
	default:
		return -1, fmt.Errorf("no op code for given string")
	}
}
