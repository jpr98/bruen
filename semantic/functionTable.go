package semantic

type FunctionTable map[string]*FunctionTableContent

// Stores all attributes of the variables
type VariableAttributes struct {
	Name   string
	TypeOf string
}

type FunctionTableContent struct {
	TypeOf  string
	Dir     int
	EraSize string
	Vars    map[string]*VariableAttributes
	// Params constants.TYPE[]
	Scope string
}

func (ft FunctionTable) GetVarType(varName string) string {
	for _, v := range ft {
		if attr, ok := v.Vars[varName]; ok {
			return attr.TypeOf
		}
	}
	return "error"
}