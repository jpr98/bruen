package constants

// Type represents a type in the compiler
type Type int

const (
	TYPEINT   Type = 0 // An integer
	TYPEFLOAT Type = 1 // A float
	TYPECHAR  Type = 2 // A character
	TYPEBOOL  Type = 3 // A boolean

	TYPECLASS Type = 4 // A class, not specific

	VOID Type = 5 // Void return type
	ERR  Type = 6 // Error in a type
	ADDR Type = 7 // An address in the compiler
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

// StringToType takes a string and tries to match it to a defined type
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
