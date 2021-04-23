package main

type QuadAction int

type QuadActionStack struct {
	stack []QuadAction
}

// NewQuadActionStack creates a new stack
func NewQuadActionStack() QuadActionStack {
	return QuadActionStack{stack: make([]QuadAction, 0)}
}

// Push adds an action to the stack
func (qas *QuadActionStack) Push(e QuadAction) {
	qas.stack = append(qas.stack, e)
}

// Pop removes and returns the next action in the stack
func (qas *QuadActionStack) Pop() QuadAction {
	e := qas.stack[qas.Size()-1]
	qas.stack = qas.stack[:qas.Size()-1]
	return e
}

// Top returns the next action in the stack
func (qas *QuadActionStack) Top() QuadAction {
	return qas.stack[qas.Size()-1]
}

// Size returns the length of the stack
func (qas *QuadActionStack) Size() int {
	return len(qas.stack)
}
