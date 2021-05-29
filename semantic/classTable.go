package semantic

type ClassTable map[string]*ClassTableContent

type ClassTableContent struct {
	Attributes map[string]*VariableAttributes
	Methods    FunctionTable
	VarsSize   [4]int
	ObjSize    int
}
