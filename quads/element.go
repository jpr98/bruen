package quads

import (
	"fmt"

	"github.com/jpr98/compis/constants"
)

type value interface {
	//Value()
}

// Element represents an element in a Quadruple
type Element interface {
	Value() value
	Type() constants.Type
	String() string
}

type quadElement struct {
	value  value
	typeOf constants.Type
}

// NewElement creates a new quad.Element
func NewElement(val value, typeOf constants.Type) Element {
	return quadElement{value: val, typeOf: typeOf}
}

func (e quadElement) Value() value {
	return e.value
}

func (e quadElement) Type() constants.Type {
	return e.typeOf
}

func (e quadElement) String() string {
	return fmt.Sprintf("%s (type: %s)", e.Value(), e.Type())
}

// ElementStack is a stack of quad.Element
type ElementStack interface {
	Push(Element)
	Pop() Element
	Top() Element
	Size() int
}

type elementStack struct {
	stack []Element
}

// NewElementStack creates a new stack
func NewElementStack() ElementStack {
	return &elementStack{stack: make([]Element, 0)}
}

// Push adds an element to the stack
func (es *elementStack) Push(e Element) {
	es.stack = append(es.stack, e)
}

// Pop removes and returns the next element in the stack
func (es *elementStack) Pop() Element {
	e := es.Top()
	es.stack = es.stack[:es.Size()-1]
	return e
}

// Top returns the next element in the stack
func (es *elementStack) Top() Element {
	return es.stack[es.Size()-1]
}

// Size returns the length of the stack
func (es *elementStack) Size() int {
	return len(es.stack)
}
