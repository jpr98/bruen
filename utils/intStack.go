package utils

type IntStack struct {
	stack []int
}

func (is *IntStack) Top() int {
	return is.stack[len(is.stack)-1]
}

func (is *IntStack) Push(n int) {
	is.stack = append(is.stack, n)
}

func (is *IntStack) Pop() int {
	result := is.Top()
	is.stack = is.stack[:len(is.stack)-1]
	return result
}
