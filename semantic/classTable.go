package semantic

import (
	"encoding/gob"

	"github.com/jpr98/bruen/memory"
)

// ClassTable is a map of class IDs to ClassTableContent. Its purpose is to serve
// as the regisrty of all classes defined in a program. ClassTable is mainly used
// in compilation.
type ClassTable map[string]*ClassTableContent

// ClassTableContent contains all relevant data of a class. It only exists inside
// of a ClassTable map.
type ClassTableContent struct {
	Attributes  map[string]*VariableAttributes // Attributes of a given class
	Methods     FunctionTable                  // Methods of a given class (incl. init)
	VarsSize    [4]int                         // Number of variable attributes per type, used to assign memory
	ObjSize     []memory.MemObjInfo            // Number of class attributes, used to assign memory
	Objects     []string                       // Class names of all class attributes
	ObjectCount int                            // Expected length of Objects
}

func init() {
	gob.Register(ClassTable{})
}
