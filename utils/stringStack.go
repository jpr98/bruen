package utils

// StringStack is a stack of strings
type StringStack struct {
	stack []string
}

// Top returns the top string of the stack
func (ss *StringStack) Top() string {
	return ss.stack[len(ss.stack)-1]
}

// Push adds a string to the stack
func (ss *StringStack) Push(str string) {
	ss.stack = append(ss.stack, str)
}

// Removes the top string of the stack
func (ss *StringStack) Pop() {
	ss.stack = ss.stack[:len(ss.stack)-1]
}
