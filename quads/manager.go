package quads

import (
	"encoding/gob"
	"fmt"
	"log"

	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/memory"
	"github.com/jpr98/compis/semantic"
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
	operands  ElementStack    // stack operands
	operators QuadActionStack // stack operators
	quads     []Quad
	jumpStack []int

	functionTable semantic.FunctionTable

	currentFunctionCall string
	paramCounter        int
	avail               int
}

func (m Manager) GetQuads() []Quad {
	return m.quads
}

func NewManager(functionTable semantic.FunctionTable) Manager {
	return Manager{
		operands:      NewElementStack(),
		operators:     NewQuadActionStack(),
		quads:         make([]Quad, 0),
		functionTable: functionTable,
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
	element := NewElement(op.Dir, operand, op.TypeOf)
	m.operands.Push(element)
}

func (m *Manager) getOperandData(operand, currentFunction, globalName string) *semantic.VariableAttributes {
	data := &semantic.VariableAttributes{TypeOf: constants.ERR}
	if attr, exists := m.functionTable[currentFunction].Vars[operand]; exists {
		data = attr
	} else {
		if attr, exists := m.functionTable[globalName].Vars[operand]; exists {
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

func (m *Manager) PushOperand(operand, currentFunction, globalName string) {
	operandData := m.getOperandData(operand, currentFunction, globalName)
	element := NewElement(operandData.Dir, operand, operandData.TypeOf)
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

		// 7: pedir espacio para resultado - done
		dir, err := memory.Manager.GetNextAddr(resultType, memory.Temp)
		if err != nil {
			log.Fatalf("Error: (GenerateQuad) %s\n", err)
		}
		result := NewElement(dir, m.getNextAvail(), resultType)

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
	operand := m.operands.Pop()
	result := m.operands.Pop()

	resultType := semantic.Cube.ValidateBinaryOperation(result.Type(), operand.Type(), constants.OPASSIGN)
	if resultType == constants.ERR {
		log.Fatalf(
			"Error: (GenerateAssignationQuad) type mismatch, assignation to %s can't be %s\n",
			result,
			operand,
		)
	}

	if retainresult {
		m.operands.Push(result)
	}
	q := Quad{op, operand, nil, result}
	m.quads = append(m.quads, q)
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
	m.quads[pos].Result = NewElement(len(m.quads), "", constants.ADDR)
}

func (m *Manager) UpdateGoto() {
	pos := m.jumpStack[len(m.jumpStack)-1]
	m.jumpStack = m.jumpStack[:len(m.jumpStack)-1]

	m.quads[pos].Result = NewElement(len(m.quads), "", constants.ADDR)
}

func (m *Manager) SaveJumpPosition() {
	m.jumpStack = append(m.jumpStack, len(m.quads))
}

func (m *Manager) AddAndUpdateWhileGotos() {
	posF := m.jumpStack[len(m.jumpStack)-1]
	m.jumpStack = m.jumpStack[:len(m.jumpStack)-1]

	posR := m.jumpStack[len(m.jumpStack)-1]
	m.jumpStack = m.jumpStack[:len(m.jumpStack)-1]

	q := Quad{GOTO, nil, nil, NewElement(posR, "", constants.ADDR)}
	m.quads = append(m.quads, q)

	m.quads[posF].Result = NewElement(len(m.quads), "", constants.ADDR)
}

func (m *Manager) AddSimpleGOTO() {
	q := Quad{GOTO, nil, nil, nil}
	m.quads = append(m.quads, q)
}

func (m *Manager) AddForLoopIterator(id string) {
	dir, err := memory.Manager.GetNextAddr(constants.TYPEINT, memory.Local)
	if err != nil {
		log.Fatalf("Error: (AddForLoopIterator) %s\n", err)
	}

	e := NewElement(dir, id, constants.TYPEINT)
	m.operands.Push(e)
}

func (m *Manager) AddReadQuad(operand, functionName, globalName string) {
	operandData := m.getOperandData(operand, functionName, globalName)
	element := NewElement(operandData.Dir, operandData.Name, operandData.TypeOf)
	q := Quad{READ, nil, nil, element}
	m.quads = append(m.quads, q)
}

func (m *Manager) AddEndFuncQuad() {
	q := Quad{ENDFUNC, nil, nil, nil}
	m.quads = append(m.quads, q)
}

func (m *Manager) AddEraQuad(name string) {
	n := NewElement(0, name, constants.TYPEINT)
	q := Quad{ERA, n, nil, nil}
	m.quads = append(m.quads, q)

	m.currentFunctionCall = name
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

	argNum := NewElement(m.functionTable[m.currentFunctionCall].Params[m.paramCounter], fmt.Sprintf("%d", m.paramCounter), expectedType)
	q := Quad{PARAM, arg, argNum, nil}
	m.quads = append(m.quads, q)
}

func (m *Manager) AddGoSubQuad(name string) {
	if m.paramCounter < len(m.functionTable[m.currentFunctionCall].Params) {
		log.Fatalf(
			"Error: (AddGoSubQuad) function %s has too few arguments",
			m.currentFunctionCall,
		)
	}

	dir := m.functionTable[m.currentFunctionCall].Dir
	n := NewElement(0, name, constants.TYPEINT)
	dirElement := NewElement(dir, "", constants.ADDR)
	q := Quad{GOSUB, n, nil, dirElement}
	m.quads = append(m.quads, q)
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
	case ASSIGN.String():
		return ASSIGN, nil
	case LPAREN.String():
		return LPAREN, nil
	default:
		return -1, fmt.Errorf("no op code for given string")
	}
}
