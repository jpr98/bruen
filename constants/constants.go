package constants

type Type int

const (
	TYPEINT   Type = 0 // 0
	TYPEFLOAT Type = 1 // 1
	TYPECHAR  Type = 2 // 2
	TYPEBOOL  Type = 3 // 3

	ERR Type = 4 // 4
)

func (t Type) String() string {
	switch t {
	case TYPEINT:
		return "int"
	case TYPEFLOAT:
		return "float"
	case TYPECHAR:
		return "char"
	case TYPEBOOL:
		return "bool"
	case ERR:
		return "Type error"
	default:
		return "Unkown type"
	}
}

const (
	OPPLUS = iota
	OPMINUS
	OPDIV
	OPMULT
	OPGT
	OPLT
	OPEQ
	OPNEQ
	OPAND
	OPOR
)
