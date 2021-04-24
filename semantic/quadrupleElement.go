package semantic

import (
	"fmt"

	"github.com/jpr98/compis/constants"
)

type QuadValue interface {
	//Value()
}

type QuadElement struct {
	value  QuadValue
	typeOf constants.Type
}

type QuadElementStack struct {
	stack []QuadElement
}

// NewQuadElementStack creates a new stack
func NewQuadElementStack() QuadElementStack {
	return QuadElementStack{stack: make([]QuadElement, 0)}
}

// Push adds an element to the stack
func (qes *QuadElementStack) Push(e QuadElement) {
	qes.stack = append(qes.stack, e)
	fmt.Println("adding to element stack")
}

// Pop removes and returns the next element in the stack
func (qes *QuadElementStack) Pop() QuadElement {
	e := qes.Top()
	qes.stack = qes.stack[:qes.Size()-1]
	return e
}

// Top returns the next element in the stack
func (qes *QuadElementStack) Top() QuadElement {
	return qes.stack[qes.Size()-1]
}

// Size returns the length of the stack
func (qes *QuadElementStack) Size() int {
	return len(qes.stack)
}
