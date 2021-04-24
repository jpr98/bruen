package semantic

import (
	"fmt"
)

type Quad struct {
	action QuadAction  // operator
	left   QuadElement // operand
	right  QuadElement // operand
	result QuadElement
}

type QuadrupleManager struct {
	operands  QuadElementStack // stack operands
	operators QuadActionStack  // stack operators
	quads     []Quad
}

func NewQuadrupleManager() QuadrupleManager {
	return QuadrupleManager{
		operands:  NewQuadElementStack(),
		operators: NewQuadActionStack(),
		quads:     make([]Quad, 0),
	}
}

func (m *QuadrupleManager) PushOp(op string) {
	opCode, err := stringToOp(op)
	if err != nil {
		return
	}

	m.operators.Push(opCode)
}

func (m *QuadrupleManager) PushOperand(operand string) {
	// TODO: obtener tipo de tabla de variables
	element := QuadElement{QuadValue(operand), 0}
	m.operands.Push(element)
	fmt.Println("Adding to stack")
}

func (m *QuadrupleManager) GenerateQuad(validOps []int) {
	// 1: validar que top op este en validOps
	if arrayContainsElement(m.operators.Top(), validOps) {
		// 2: obtener right operand
		rOperand := m.operands.Pop()
		// 3: obtener left operand
		lOperand := m.operands.Pop()
		// 4: sacar op de stack
		op := m.operators.Top()
		//fmt.Println("found ", op)
		fmt.Printf("obtuvimos %s %d %s \n", rOperand.value, op, lOperand.value)

		// 5: obtener tipo de cubo semantico
		// 6: validar tipo no sea error (manejar error en caso de haber)
		// 7: pedir espacio para resultado
		// 8: generar quad
		// 9: meter quad a lista
		// 10: meter result a stack operandos
		return
	}
	fmt.Println("no hubo match")
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
	case "+":
		return 0, nil
	case "-":
		return 1, nil
	case ">":
		return 4, nil
	case "<":
		return 5, nil
	case "==":
		return 6, nil
	case "!=":
		return 7, nil
	default:
		// TODO: log error
		return -1, fmt.Errorf("no op code for given string")
	}
}
