package semantic

import (
	"encoding/gob"

	"github.com/jpr98/bruen/constants"
	"github.com/jpr98/bruen/memory"
)

// FunctionTable is a map of function IDs to FunctionTableContent. Its purpose is
// to serve as a registry of all global functions defined in a program. This map
// happens to also have information about the global scope.
type FunctionTable map[string]*FunctionTableContent

// VariableAttributes stores all attributes of variables
type VariableAttributes struct {
	Name       string         // Name of variable
	TypeOf     constants.Type // Type of variable
	Dir        int            // Virtual memory address assigned to variable
	ArrayOrMat int            // Identifies if variable is nothing, array, matrix (0 = nothing, 1 = array, 2 = mat)
	Dim        [2]int         // Dim holds the dimensions of an array or a matrix, don't use if ArrayOrMat is 0
	Class      string         // The class type of variable, only use if TypeOf is constants.TYPECLASS
	FromSelf   bool           // Identifies if variable is really a class attribute being accessed
	SelfDir    int            // Holds the address of the instance where variable lives, don't if FromSelf is false
	IsPrivate  bool           // Access modifier of variable
}

// NewVariableAttributes creates a new VariableAttributes instance with given attributes and default values
func NewVariableAttributes(id string, typeOf constants.Type, dir int) *VariableAttributes {
	return &VariableAttributes{
		Name:       id,
		TypeOf:     typeOf,
		Dir:        dir,
		ArrayOrMat: 0,
		Dim:        [2]int{},
		Class:      "",
		FromSelf:   false,
		IsPrivate:  false,
	}
}

// FunctionTableContent contains all relevant data of a function. It only exists
// inside of a FunctionTable.
type FunctionTableContent struct {
	TypeOf      constants.Type                 // Return type of function
	Dir         int                            // Instruction pointer of function start
	ReturnDir   int                            // Virtual memory address for temporal return value holding
	Vars        map[string]*VariableAttributes // Map of local variables
	Params      []int                          // List of expected types for parameters
	Scope       string                         // Scope in which a function exists, mainly global or a class
	VarsSize    [4]int                         // Number of variables per type, used to assign memory
	TempSize    [4]int                         // Number of temporary variables, used to assign memory
	ObjSize     []memory.MemObjInfo            // Number of class variables, used to assign memory
	Objects     []string                       // Class names of all class variables
	ObjectCount int                            // Expected length of Objects
	IsPrivate   bool                           // Access modifier of function
}

func init() {
	gob.Register(FunctionTable{})
}
