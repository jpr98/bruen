package virtualMachine

import (
	"github.com/jpr98/compis/quads"
	"github.com/jpr98/compis/semantic"
)

type VirtualMachine struct {
	globalMemBlock Memory
	memBlocks      MemoryStack

	quads []quads.Quad
}

func NewVM(functionTable semantic.FunctionTable, quads []quads.Quad) VirtualMachine {
	gmb := NewMemory([4]int{0, 0, 0, 0}, [4]int{0, 0, 0, 0}) // TODO: NewMemory() de functionTable sacar tamaño para el bloque global
	return VirtualMachine{globalMemBlock: gmb, quads: quads}
}
