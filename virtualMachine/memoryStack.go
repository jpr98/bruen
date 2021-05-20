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
