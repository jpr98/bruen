package quads

import (
	"fmt"

	"github.com/jpr98/compis/constants"
)

// Element represents an element in a Quadruple
type Element interface {
	// GetAddr returns the memory address of the element
	GetAddr() int
	// Type returns the type of the element
	Type() constants.Type
	// String is the string representation of the element
	String() string
	// ID returns the id of the element, we also add tags to this field
	//
	// Tags:
	// 	`self_`: element belongs to an instance, the instance address is found between the 1st and 2nd `_`
	// 	`ptr_`: element value is a pointer to an address
	ID() string
	// ClassName returns the class of the element, is empty string if Type() is not constants.TYPECLASS
	ClassName() string
}

type quadElement struct {
	Addr      int
	Id        string
	TypeOf    constants.Type
	IsPointer bool
	Class     string
}

// NewElement creates a new quad.Element
func NewElement(addr int, id string, typeOf constants.Type, class string) Element {
	return quadElement{Addr: addr, Id: id, TypeOf: typeOf, IsPointer: false, Class: class}
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

func (e quadElement) ID() string {
	return e.Id
}

func (e quadElement) ClassName() string {
	return e.Class
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
