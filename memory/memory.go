package memory

import (
	"fmt"

	"github.com/jpr98/compis/constants"
)

const (
	GLOBAL_INT   = 1000
	GLOBAL_FLOAT = 2000
	GLOBAL_CHAR  = 3000
	GLOBAL_BOOL  = 4000

	LOCAL_INT   = 5000
	LOCAL_FLOAT = 6000
	LOCAL_CHAR  = 7000
	LOCAL_BOOL  = 8000

	TEMP_INT   = 9000
	TEMP_FLOAT = 10000
	TEMP_CHAR  = 11000
	TEMP_BOOL  = 12000

	CONST_INT   = 13000
	CONST_FLOAT = 14000
	CONST_CHAR  = 15000
	CONST_BOOL  = 16000
)

type Context string

const (
	Global   Context = "global"
	Local    Context = "local"
	Temp     Context = "temp"
	Constant Context = "constant"
)

func TypeForAddr(addr int) constants.Type {
	n := addr / 1000
	switch n {
	case 1, 5, 9, 13:
		return constants.TYPEINT
	case 2, 6, 10, 14:
		return constants.TYPEFLOAT
	case 3, 7, 11, 15:
		return constants.TYPECHAR
	case 4, 8, 12, 16:
		return constants.TYPEBOOL
	default:
		return constants.ERR
	}
}

var Manager *manager = newManager()

type manager struct {
	gInt   int
	gFloat int
	gChar  int
	gBool  int

	lInt   int
	lFloat int
	lChar  int
	lBool  int

	tInt   int
	tFloat int
	tChar  int
	tBool  int

	cInt   int
	cFloat int
	cChar  int
	cBool  int
}

func newManager() *manager {
	return &manager{
		gInt:   GLOBAL_INT,
		gFloat: GLOBAL_FLOAT,
		gChar:  GLOBAL_CHAR,
		gBool:  GLOBAL_BOOL,

		lInt:   LOCAL_INT,
		lFloat: LOCAL_FLOAT,
		lChar:  LOCAL_CHAR,
		lBool:  LOCAL_BOOL,

		tInt:   TEMP_INT,
		tFloat: TEMP_FLOAT,
		tChar:  TEMP_CHAR,
		tBool:  TEMP_BOOL,

		cInt:   CONST_INT,
		cFloat: CONST_FLOAT,
		cChar:  CONST_CHAR,
		cBool:  CONST_BOOL,
	}
}

func (m *manager) GetNextAddr(typeOf constants.Type, context Context) (int, error) {
	switch context {
	case Global:
		return m.getGlobalForType(typeOf), nil
	case Local:
		return m.getLocalForType(typeOf), nil
	case Temp:
		return m.getTempForType(typeOf), nil
	case Constant:
		return m.getConstForType(typeOf), nil
	}
	return 0, fmt.Errorf("(GetNextAddr) invalid context %s", context)
}

// ConvertAddr returns the memory reference minus the offset to get the real memory address
func ConvertAddr(addr int) int {
	if addr < GLOBAL_FLOAT {
		return addr - GLOBAL_INT
	} else if addr < GLOBAL_CHAR {
		return addr - GLOBAL_FLOAT
	} else if addr < GLOBAL_BOOL {
		return addr - GLOBAL_CHAR
	} else if addr < LOCAL_INT {
		return addr - GLOBAL_BOOL
	} else if addr < LOCAL_FLOAT {
		return addr - LOCAL_INT
	} else if addr < LOCAL_CHAR {
		return addr - LOCAL_FLOAT
	} else if addr < LOCAL_BOOL {
		return addr - LOCAL_CHAR
	} else if addr < TEMP_INT {
		return addr - LOCAL_BOOL
	} else if addr < TEMP_FLOAT {
		return addr - TEMP_INT
	} else if addr < TEMP_CHAR {
		return addr - TEMP_FLOAT
	} else if addr < TEMP_BOOL {
		return addr - TEMP_CHAR
	} else if addr < CONST_INT {
		return addr - TEMP_BOOL
	} else if addr < CONST_FLOAT {
		return addr - CONST_INT
	} else if addr < CONST_CHAR {
		return addr - CONST_FLOAT
	} else if addr < CONST_BOOL {
		return addr - CONST_CHAR
	} else {
		return addr - CONST_BOOL
	}
}

func (m *manager) getGlobalForType(typeOf constants.Type) int {
	var result int
	switch typeOf {
	case constants.TYPEINT:
		result = m.gInt
		m.gInt++
	case constants.TYPEFLOAT:
		result = m.gFloat
		m.gFloat++
	case constants.TYPECHAR:
		result = m.gChar
		m.gChar++
	case constants.TYPEBOOL:
		result = m.gBool
		m.gBool++
	}
	return result
}

func (m *manager) getLocalForType(typeOf constants.Type) int {
	var result int
	switch typeOf {
	case constants.TYPEINT:
		result = m.lInt
		m.lInt++
	case constants.TYPEFLOAT:
		result = m.lFloat
		m.lFloat++
	case constants.TYPECHAR:
		result = m.lChar
		m.lChar++
	case constants.TYPEBOOL:
		result = m.lBool
		m.lBool++
	}
	return result
}

func (m *manager) getTempForType(typeOf constants.Type) int {
	var result int
	switch typeOf {
	case constants.TYPEINT:
		result = m.tInt
		m.tInt++
	case constants.TYPEFLOAT:
		result = m.tFloat
		m.tFloat++
	case constants.TYPECHAR:
		result = m.tChar
		m.tChar++
	case constants.TYPEBOOL:
		result = m.tBool
		m.tBool++
	}
	return result
}

func (m *manager) getConstForType(typeOf constants.Type) int {
	var result int
	switch typeOf {
	case constants.TYPEINT:
		result = m.cInt
		m.cInt++
	case constants.TYPEFLOAT:
		result = m.cFloat
		m.cFloat++
	case constants.TYPECHAR:
		result = m.cChar
		m.cChar++
	case constants.TYPEBOOL:
		result = m.cBool
		m.cBool++
	}
	return result
}

func (m *manager) ResetLocalCounter() {
	m.lInt = LOCAL_INT
	m.lFloat = LOCAL_FLOAT
	m.lChar = LOCAL_CHAR
	m.lBool = LOCAL_BOOL
}

func (m *manager) ResetTempCounter() {
	m.tInt = TEMP_INT
	m.tFloat = TEMP_FLOAT
	m.tChar = TEMP_CHAR
	m.tBool = TEMP_BOOL
}
