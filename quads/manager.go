package quads

import (
	"fmt"
	"log"

	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/semantic"
)

type Quad struct {
	action QuadAction // operator
	left   Element    // operand
	right  Element    // operand
	result Element
}

func (q Quad) String() string {
	return fmt.Sprintf(
		"Quad {%s %s %s %s}",
		q.action,
		q.left,
		q.right,
		q.result,
	)
}

type Manager struct {
	operands  ElementStack    // stack operands
	operators QuadActionStack // stack operators
	quads     []Quad

	avail int
}

func NewManager() Manager {
	return Manager{
		operands:  NewElementStack(),
		operators: NewQuadActionStack(),
		quads:     make([]Quad, 0),
		avail:     0,
	}
}

func (m *Manager) PushOp(op string) {
	opCode, err := stringToOp(op)
	if err != nil {
		return
	}

	m.operators.Push(opCode)
}

func (m *Manager) PushOperand(operand string) {
	// TODO: obtener tipo de tabla de variables
	element := NewElement(operand, 0)
	m.operands.Push(element)
}

func (m *Manager) GenerateQuad(validOps []int) {
	// 1: validar que top op este en validOps - done
	if arrayContainsElement(m.operators.Top(), validOps) {
		// 2: obtener right operand - done
		rOperand := m.operands.Pop()
		// 3: obtener left operand - done
		lOperand := m.operands.Pop()
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
		result := NewElement(m.getNextAvail(), resultType)
		// 8: generar quad - done
		q := Quad{op, lOperand, rOperand, result}
		fmt.Println(q)
		// 9: meter quad a lista - done
		m.quads = append(m.quads, q)
		// 10: meter result a stack operandos - done
		m.operands.Push(result)
		fmt.Println("Operands: ", m.operands)
		fmt.Println("Operators: ", m.operators)
	}
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
	case GT.String():
		return GT, nil
	case LT.String():
		return LT, nil
	case EQ.String():
		return EQ, nil
	case NEQ.String():
		return NEQ, nil
	default:
		return -1, fmt.Errorf("no op code for given string")
	}
}
