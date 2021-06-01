package semantic

import (
	"encoding/gob"

	"github.com/jpr98/compis/memory"
)

type ClassTable map[string]*ClassTableContent

type ClassTableContent struct {
	Attributes  map[string]*VariableAttributes
	Methods     FunctionTable
	VarsSize    [4]int
	ObjSize     []memory.MemObjInfo
	Objects     []string
	ObjectCount int
}

func init() {
	gob.Register(ClassTable{})
}
