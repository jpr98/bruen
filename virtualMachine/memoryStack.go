package virtualMachine

type MemoryStack struct {
	stack []Memory
}

func (ms *MemoryStack) Top() Memory {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MemoryStack) Push(mem Memory) {
	ms.stack = append(ms.stack, mem)
}

func (ms *MemoryStack) Pop() {
	ms.stack = ms.stack[:len(ms.stack)-1]
}

func (ms *MemoryStack) Empty() bool {
	return len(ms.stack) == 0
}

type PointerStack struct {
	stack []int
}

func (ps *PointerStack) Top() int {
	return ps.stack[len(ps.stack)-1]
}

func (ps *PointerStack) Push(pointer int) {
	ps.stack = append(ps.stack, pointer)
}

func (ps *PointerStack) Pop() int {
	pointer := ps.Top()
	ps.stack = ps.stack[:len(ps.stack)-1]
	return pointer
}

func (ps *PointerStack) Empty() bool {
	return len(ps.stack) == 0
}
