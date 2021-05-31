package semantic

import (
	"encoding/gob"

	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/memory"
)

type FunctionTable map[string]*FunctionTableContent

// Stores all attributes of the variables
type VariableAttributes struct {
	Name       string
	TypeOf     constants.Type
	Dir        int
	ArrayOrMat int // 0 = nothing, 1 = array, 2 = mat
	Dim        [2]int
	Class      string
	FromSelf   bool
	SelfDir    int
}

func NewVariableAttributes(id string, typeOf constants.Type, dir int) *VariableAttributes {
	return &VariableAttributes{
		Name:       id,
		TypeOf:     typeOf,
		Dir:        dir,
		ArrayOrMat: 0,
		Dim:        [2]int{},
		Class:      "",
		FromSelf:   false,
	}
}

type FunctionTableContent struct {
	TypeOf      constants.Type
	Dir         int
	ReturnDir   int
	Vars        map[string]*VariableAttributes
	Params      []int
	Scope       string
	VarsSize    [4]int
	TempSize    [4]int
	ObjSize     []memory.MemObjInfo
	Objects     []string
	ObjectCount int
}

func init() {
	gob.Register(FunctionTable{})
}

func (ft FunctionTable) GetVarType(varName string) constants.Type {
	for _, v := range ft {
		if attr, ok := v.Vars[varName]; ok {
			return attr.TypeOf
		}
	}
	return constants.ERR
}
