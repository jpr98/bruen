package constants

type Type int

const (
	TYPEINT   Type = 0 // 0
	TYPEFLOAT Type = 1 // 1
	TYPECHAR  Type = 2 // 2
	TYPEBOOL  Type = 3 // 3

	TYPECLASS Type = 4

	VOID Type = 5
	ERR  Type = 6 // 5
	ADDR Type = 7
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
	case VOID:
		return "void"
	case ERR:
		return "Type error"
	case ADDR:
		return "Address"
	case TYPECLASS:
		return "Class"
	default:
		return "Unkown type"
	}
}

func StringToType(str string) Type {
	switch str {
	case TYPEINT.String():
		return TYPEINT
	case TYPEFLOAT.String():
		return TYPEFLOAT
	case TYPECHAR.String():
		return TYPECHAR
	case TYPEBOOL.String():
		return TYPEBOOL
	case ADDR.String():
		return ADDR
	case VOID.String():
		return VOID
	default:
		return ERR
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
	OPASSIGN
)
