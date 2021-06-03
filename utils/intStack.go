package utils

// IntStack is a stack of ints
type IntStack struct {
	stack []int
}

// Top returns the top int of the stack
func (is *IntStack) Top() int {
	return is.stack[len(is.stack)-1]
}

// Push adds a new int to the stack
func (is *IntStack) Push(n int) {
	is.stack = append(is.stack, n)
}

// Pop removes and returns the top int of the stack
func (is *IntStack) Pop() int {
	result := is.Top()
	is.stack = is.stack[:len(is.stack)-1]
	return result
}
