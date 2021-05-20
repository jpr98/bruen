package quads

import (
	"fmt"

	"github.com/jpr98/compis/constants"
)

// Element represents an element in a Quadruple
type Element interface {
	GetAddr() int
	Type() constants.Type
	String() string
}

type quadElement struct {
	Addr   int
	Id     string
	TypeOf constants.Type
}

// NewElement creates a new quad.Element
func NewElement(addr int, id string, typeOf constants.Type) Element {
	return quadElement{Addr: addr, Id: id, TypeOf: typeOf}
}

func (e quadElement) GetAddr() int {
	return e.Addr
}

func (e quadElement) Type() constants.Type {
	return e.TypeOf
}

func (e quadElement) String() string {
	return fmt.Sprintf("%s (type: %s, addr: %d)", e.Id, e.Type(), e.GetAddr())
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
