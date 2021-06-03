package virtualMachine

// MemoryStack is a stack of Memory
type MemoryStack struct {
	stack []Memory
}

// Top returns the top element of the stack
func (ms *MemoryStack) Top() Memory {
	return ms.stack[len(ms.stack)-1]
}

// Push adds an element to the stack
func (ms *MemoryStack) Push(mem Memory) {
	ms.stack = append(ms.stack, mem)
}

// Pop removes and return the top element of the stack
func (ms *MemoryStack) Pop() Memory {
	memory := ms.Top()
	ms.stack = ms.stack[:len(ms.stack)-1]
	return memory
}

// Empty returns true if the stack is empty
func (ms *MemoryStack) Empty() bool {
	return len(ms.stack) == 0
}

// PointerStack is a stack of pointers
type PointerStack struct {
	stack []int
}

// Top returns the top element of the stack
func (ps *PointerStack) Top() int {
	return ps.stack[len(ps.stack)-1]
}

// Push adds an element to the stack
func (ps *PointerStack) Push(pointer int) {
	ps.stack = append(ps.stack, pointer)
}

// Pop removes and return the top element of the stack
func (ps *PointerStack) Pop() int {
	pointer := ps.Top()
	ps.stack = ps.stack[:len(ps.stack)-1]
	return pointer
}

// Empty returns true if the stack is empty
func (ps *PointerStack) Empty() bool {
	return len(ps.stack) == 0
}
