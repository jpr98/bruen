package semantic

import "github.com/jpr98/compis/constants"

type FunctionTable map[string]*FunctionTableContent

// Stores all attributes of the variables
type VariableAttributes struct {
	Name   string
	TypeOf constants.Type
	Dir    int
}

type FunctionTableContent struct {
	TypeOf  string
	Dir     int
	EraSize string
	Vars    map[string]*VariableAttributes
	Params  []constants.Type
	Scope   string
}

func (ft FunctionTable) GetVarType(varName string) constants.Type {
	for _, v := range ft {
		if attr, ok := v.Vars[varName]; ok {
			return attr.TypeOf
		}
	}
	return constants.ERR
}
