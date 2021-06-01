package quads

import (
	"encoding/gob"
	"fmt"
	"log"
	"strings"

	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/memory"
	"github.com/jpr98/compis/semantic"
	"github.com/jpr98/compis/utils"
)

type Quad struct {
	Action QuadAction // operator
	Left   Element    // operand
	Right  Element    // operand
	Result Element
}

func init() {
	gob.Register(Quad{})
	gob.Register(quadElement{})
}

func (q Quad) String() string {
	return fmt.Sprintf(
		"Quad {%s %s %s %s}",
		q.Action,
		q.Left,
		q.Right,
		q.Result,
	)
}

type Manager struct {
	operands   ElementStack    // stack operands
	operators  QuadActionStack // stack operators
	quads      []Quad
	jumpStack  []int
	scopeStack utils.StringStack

	functionTable semantic.FunctionTable
	classTable    semantic.ClassTable

	globalName      string
	currentFunction string

	currentFunctionCall      string
	currentFunctionCallClass string
	paramCounter             int
	avail                    int
}

func (m Manager) GetQuads() []Quad {
	return m.quads
}

func NewManager(functionTable semantic.FunctionTable, classTable semantic.ClassTable, ss utils.StringStack) Manager {
	return Manager{
		operands:      NewElementStack(),
		operators:     NewQuadActionStack(),
		quads:         make([]Quad, 0),
		scopeStack:    ss,
		functionTable: functionTable,
		classTable:    classTable,
		avail:         0,
	}
}

func (m *Manager) PushOp(op string) {
	opCode, err := stringToOp(op)
	if err != nil {
		log.Fatalln(err)
		return
	}

	m.operators.Push(opCode)
}

// PushConstantOperand sets an operand with a defined (hardcoded) type
func (m *Manager) PushConstantOperand(operand string) {
	op, exists := semantic.ConstantsTable[operand]
	if !exists {
		log.Fatalf("Error: (PushConstantOperand) %s is not a known constant", operand)
	}
	element := NewElement(op.Dir, operand, op.TypeOf, "")
	m.operands.Push(element)
}

// Acceso a atributos
// a.b = explicito = a es el objecto
// b = implicito = self es el objecto
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

	if data.TypeOf == constants.ERR {
		log.Fatalf(
			"Error: (PushOperand) undeclared variable %s\n",
			operand,
		)
		return nil
	}
	return data
}

func (m *Manager) PushOperand(operand string) {
	operandData := m.getOperandData(operand)
	operandName := operandData.Name
	if operandData.FromSelf {
		operandName = fmt.Sprintf("self_%d_%s", operandData.SelfDir, operandData.Name)
	}
	element := NewElement(operandData.Dir, operandName, operandData.TypeOf, operandData.Class)
	m.operands.Push(element)
}

func (m *Manager) GenerateQuad(validOps []int, isForLoop bool) {
	// 1: validar que top op este en validOps - done
	if arrayContainsElement(m.operators.Top(), validOps) {
		// 2: obtener Right operand - done
		rOperand := m.operands.Pop()
		// 3: obtener Left operand - done
		var lOperand Element
		if isForLoop {
			lOperand = m.operands.Top()
		} else {
			lOperand = m.operands.Pop()
		}
		// 4: sacar op de stack - done
		op := m.operators.Pop()

		// 5: obtener tipo de cubo semantico - done
		resultType := semantic.Cube.ValidateBinaryOperation(lOperand.Type(), rOperand.Type(), int(op))
		// 6: validar tipo no sea error (manejar error en caso de haber) - done
		if resultType == constants.ERR {
			log.Fatalf(
				"Error: (GenerateQuad) type mismatch, operator (%s) can't be done between %s and %s\n",
				op,
				lOperand,
				rOperand,
			)
		}
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
		// 7: pedir espacio para resultado - done
		dir, err := memory.Manager.GetNextAddr(resultType, memory.Temp)
		if err != nil {
			log.Fatalf("Error: (GenerateQuad) %s\n", err)
		}
		result := NewElement(dir, m.getNextAvail(), resultType, lOperand.ClassName())

		// 8: generar quad - done
		q := Quad{op, lOperand, rOperand, result}

		// 9: meter quad a lista - done
		m.quads = append(m.quads, q)

		// 10: meter result a stack operandos - done
		m.operands.Push(result)
	}
}

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
				"Error: (GenerateQuad) type mismatch, operator (%s) can't be done between %s and %s\n",
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

func (m *Manager) AddArrayAccessQuad() {
	index := m.operands.Pop()
	array := m.operands.Pop()

	dir, err := memory.Manager.GetNextAddr(constants.TYPEINT, memory.Temp)
	if err != nil {
		log.Fatalf("Error: (GenerateQuad) %s\n", err)
	}
	var result Element
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

func (m *Manager) UpdateGoto() {
	pos := m.jumpStack[len(m.jumpStack)-1]
	m.jumpStack = m.jumpStack[:len(m.jumpStack)-1]

	m.quads[pos].Result = NewElement(len(m.quads), "", constants.ADDR, "")
}

func (m *Manager) SaveJumpPosition() {
	m.jumpStack = append(m.jumpStack, len(m.quads))
}

func (m *Manager) AddAndUpdateWhileGotos() {
	posF := m.jumpStack[len(m.jumpStack)-1]
	m.jumpStack = m.jumpStack[:len(m.jumpStack)-1]

	posR := m.jumpStack[len(m.jumpStack)-1]
	m.jumpStack = m.jumpStack[:len(m.jumpStack)-1]

	q := Quad{GOTO, nil, nil, NewElement(posR, "", constants.ADDR, "")}
	m.quads = append(m.quads, q)

	m.quads[posF].Result = NewElement(len(m.quads), "", constants.ADDR, "")
}

func (m *Manager) AddSimpleGOTO() {
	q := Quad{GOTO, nil, nil, nil}
	m.quads = append(m.quads, q)
}

func (m *Manager) AddForLoopIterator(id string) {
	attr, exists := m.getCurrentFunctionTable()[m.currentFunction].Vars[id]
	if !exists {
		log.Fatalf("Error: (AddForLoopIterator) undeclared variable %s", id)
	}

	e := NewElement(attr.Dir, id, attr.TypeOf, "")
	m.operands.Push(e)
}

func (m *Manager) AddReadQuad() {
	operand := m.operands.Pop()
	q := Quad{READ, nil, nil, operand}
	m.quads = append(m.quads, q)
}

func (m *Manager) AddWriteQuad() {
	operand := m.operands.Pop()
	q := Quad{WRITE, nil, nil, operand}
	m.quads = append(m.quads, q)
}

func (m *Manager) AddNewLineWriteQuad() {
	q := Quad{WRITENEWLINE, nil, nil, nil}
	m.quads = append(m.quads, q)
}

func (m *Manager) AddEndFuncQuad() {
	q := Quad{ENDFUNC, nil, nil, nil}
	m.quads = append(m.quads, q)
}

func (m *Manager) getCurrentFunctionTable() semantic.FunctionTable {
	if m.scopeStack.Top() == m.globalName {
		return m.functionTable
	}
	return m.classTable[m.scopeStack.Top()].Methods
}

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

func (m *Manager) AddInitReturnQuad() {
	selfAtrr := m.getCurrentFunctionTable()[m.currentFunction].Vars["self"]
	returnValue := NewElement(selfAtrr.Dir, selfAtrr.Name, selfAtrr.TypeOf, selfAtrr.Class)

	returnDir := m.getCurrentFunctionTable()[m.currentFunction].ReturnDir
	initReturnElement := NewElement(returnDir, "", constants.ADDR, "")
	q := Quad{RETURN, initReturnElement, nil, returnValue}
	m.quads = append(m.quads, q)
}

func (m *Manager) AddEraQuad(name, className string) {
	n := NewElement(0, name, constants.ADDR, className)
	q := Quad{ERA, n, nil, nil}
	m.quads = append(m.quads, q)

	m.currentFunctionCall = name
	m.currentFunctionCallClass = className
	m.paramCounter = 0
}

func (m *Manager) AddParamQuad() {
	if m.paramCounter >= len(m.functionTable[m.currentFunctionCall].Params) {
		log.Fatalf(
			"Error: (AddGoSubQuad) function %s has too many arguments",
			m.currentFunctionCall,
		)
	}

	expectedType := memory.TypeForAddr(m.functionTable[m.currentFunctionCall].Params[m.paramCounter])
	arg := m.operands.Pop()
	t := semantic.Cube.ValidateBinaryOperation(expectedType, arg.Type(), int(constants.OPASSIGN))
	if t == constants.ERR {
		log.Fatalf(
			"Error: (AddParamQuad) type mismatch, parameter %s must be of type %s",
			arg,
			expectedType,
		)
	}

	argNum := NewElement(m.functionTable[m.currentFunctionCall].Params[m.paramCounter], fmt.Sprintf("%d", m.paramCounter), expectedType, "")
	q := Quad{PARAM, arg, argNum, nil}
	m.quads = append(m.quads, q)
}

func (m *Manager) AddClassParamQuad() {
	if m.paramCounter >= len(m.classTable[m.currentFunctionCallClass].Methods[m.currentFunctionCall].Params) {
		log.Fatalf(
			"Error: (AddGoSubQuad) function %s has too many arguments",
			m.currentFunctionCall,
		)
	}

	expectedType := memory.TypeForAddr(m.classTable[m.currentFunctionCallClass].Methods[m.currentFunctionCall].Params[m.paramCounter])
	arg := m.operands.Pop()
	t := semantic.Cube.ValidateBinaryOperation(expectedType, arg.Type(), int(constants.OPASSIGN))
	if t == constants.ERR {
		log.Fatalf(
			"Error: (AddParamQuad) type mismatch, parameter %s must be of type %s",
			arg,
			expectedType,
		)
	}

	argNum := NewElement(m.classTable[m.currentFunctionCallClass].Methods[m.currentFunctionCall].Params[m.paramCounter], fmt.Sprintf("%d", m.paramCounter), expectedType, "")
	q := Quad{PARAM, arg, argNum, nil}
	m.quads = append(m.quads, q)
}

func (m *Manager) AddGoSubQuad() {
	if m.paramCounter < len(m.functionTable[m.currentFunctionCall].Params) {
		log.Fatalf(
			"Error: (AddGoSubQuad) function %s has too few arguments",
			m.currentFunctionCall,
		)
	}

	dir := m.functionTable[m.currentFunctionCall].Dir
	n := NewElement(0, m.currentFunctionCall, constants.TYPEINT, "")
	dirElement := NewElement(dir, "", constants.ADDR, "")
	q := Quad{GOSUB, n, nil, dirElement}
	m.quads = append(m.quads, q)

	if m.functionTable[m.currentFunctionCall].TypeOf != constants.VOID {
		resultType := m.functionTable[m.currentFunctionCall].TypeOf
		if resultType == constants.ERR {
			log.Fatalf(
				"Error: (AddGoSubQuad) error in return type %s",
				m.functionTable[m.currentFunctionCall].TypeOf,
			)
		}

		dir, err := memory.Manager.GetNextAddr(resultType, memory.Temp)
		if err != nil {
			log.Fatalf("Error: (AddGoSubQuad) %s\n", err)
		}
		result := NewElement(dir, m.getNextAvail(), resultType, "")
		m.operands.Push(result)

		returnDir := m.functionTable[m.currentFunctionCall].ReturnDir
		funcReturnElement := NewElement(returnDir, "", constants.ADDR, "")
		q := Quad{ASSIGN, funcReturnElement, nil, result}
		m.quads = append(m.quads, q)
	}
}

func (m *Manager) AddClassGoSubQuad(className, instance string) {
	if m.paramCounter < len(m.classTable[className].Methods[m.currentFunctionCall].Params) {
		log.Fatalf(
			"Error: (AddClassGoSubQuad) function %s has too few arguments",
			m.currentFunctionCall,
		)
	}

	if m.currentFunctionCall != "init" {
		selfDir := m.getCurrentFunctionTable()[m.currentFunction].Vars[instance].Dir
		m.classTable[className].Methods[m.currentFunctionCall].Vars["self"].Dir = selfDir
		instanceName := fmt.Sprintf("self_%d_%s", selfDir, instance)
		instanceElement := NewElement(selfDir, instanceName, constants.ADDR, className)
		q := Quad{INSTANCE, instanceElement, nil, nil}
		m.quads = append(m.quads, q)
	}

	dir := m.classTable[className].Methods[m.currentFunctionCall].Dir
	n := NewElement(0, m.currentFunctionCall, constants.TYPEINT, className)
	dirElement := NewElement(dir, "", constants.ADDR, "")
	q := Quad{GOSUB, n, nil, dirElement}
	m.quads = append(m.quads, q)

	if m.classTable[className].Methods[m.currentFunctionCall].TypeOf != constants.VOID {
		resultType := m.classTable[className].Methods[m.currentFunctionCall].TypeOf
		if resultType == constants.ERR {
			log.Fatalf(
				"Error: (AddClassGoSubQuad) error in return type %s",
				m.classTable[className].Methods[m.currentFunctionCall].TypeOf,
			)
		}

		dir, err := memory.Manager.GetNextAddr(resultType, memory.Temp)
		if err != nil {
			log.Fatalf("Error: (AddClassGoSubQuad) %s\n", err)
		}

		var resultClass string
		if m.currentFunctionCall == "init" {
			resultClass = className
			m.getCurrentFunctionTable()[m.currentFunction].Objects = append(m.getCurrentFunctionTable()[m.currentFunction].Objects, className)
		}
		result := NewElement(dir, m.getNextAvail(), resultType, resultClass)
		m.operands.Push(result)

		returnDir := m.classTable[className].Methods[m.currentFunctionCall].ReturnDir
		funcReturnElement := NewElement(returnDir, "", constants.ADDR, "")
		q := Quad{ASSIGN, funcReturnElement, nil, result}
		m.quads = append(m.quads, q)
	}
	m.currentFunctionCallClass = ""
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

func (m *Manager) getNextAvail() string {
	defer func() { m.avail++ }()
	return fmt.Sprintf("t%d", m.avail)
}

func arrayContainsElement(element QuadAction, array []int) bool {
	for _, value := range array {
		if int(element) == value {
			return true
		}
	}
	return false
}

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
