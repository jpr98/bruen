package main

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
