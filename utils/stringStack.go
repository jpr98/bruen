package utils

type StringStack struct {
	stack []string
}

func (ss *StringStack) Top() string {
	return ss.stack[len(ss.stack)-1]
}

func (ss *StringStack) Push(scope string) {
	ss.stack = append(ss.stack, scope)
}

func (ss *StringStack) Pop() {
	ss.stack = ss.stack[:len(ss.stack)-1]
}
