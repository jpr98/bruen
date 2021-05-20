package virtualMachine

import (
	"log"

	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/memory"
)

type Memory interface {
	GetInt(vAddr int, isTemp bool) int
	GetFloat(vAddr int, isTemp bool) float32
	GetChar(vAddr int, isTemp bool) rune
	GetBool(vAddr int, isTemp bool) bool
}

type mem struct {
	vars  memblock
	temps memblock
}

type memblock map[constants.Type][]interface{}

func NewMemory(varSizes, tempSizes [4]int) Memory {
	m := mem{}
	// Variables block
	if varSizes[constants.TYPEINT] > 0 {
		m.vars[constants.TYPEINT] = make([]interface{}, varSizes[constants.TYPEINT])
	}
	if varSizes[constants.TYPEFLOAT] > 0 {
		m.vars[constants.TYPEFLOAT] = make([]interface{}, varSizes[constants.TYPEFLOAT])
	}
	if varSizes[constants.TYPECHAR] > 0 {
		m.vars[constants.TYPECHAR] = make([]interface{}, varSizes[constants.TYPECHAR])
	}
	if varSizes[constants.TYPEBOOL] > 0 {
		m.vars[constants.TYPEBOOL] = make([]interface{}, varSizes[constants.TYPEBOOL])
	}

	// Temporaries block
	if tempSizes[constants.TYPEINT] > 0 {
		m.temps[constants.TYPEINT] = make([]interface{}, tempSizes[constants.TYPEINT])
	}
	if tempSizes[constants.TYPEFLOAT] > 0 {
		m.temps[constants.TYPEFLOAT] = make([]interface{}, tempSizes[constants.TYPEFLOAT])
	}
	if tempSizes[constants.TYPECHAR] > 0 {
		m.temps[constants.TYPECHAR] = make([]interface{}, tempSizes[constants.TYPECHAR])
	}
	if tempSizes[constants.TYPEBOOL] > 0 {
		m.temps[constants.TYPEBOOL] = make([]interface{}, tempSizes[constants.TYPEBOOL])
	}
	return &m
}

func (m mem) GetInt(vAddr int, isTemp bool) int {
	addr := memory.ConvertAddr(vAddr)
	var value interface{}
	if isTemp {
		value = m.temps[constants.TYPEINT][addr]
	} else {
		value = m.vars[constants.TYPEINT][addr]
	}
	intValue, ok := value.(int)
	if !ok {
		log.Fatalf(
			"Error: (GetInt) failed to type-cast %v into an int",
			value,
		)
	}
	return intValue
}

func (m mem) GetFloat(vAddr int, isTemp bool) float32 {
	addr := memory.ConvertAddr(vAddr)
	var value interface{}
	if isTemp {
		value = m.temps[constants.TYPEFLOAT][addr]
	} else {
		value = m.vars[constants.TYPEFLOAT][addr]
	}
	floatValue, ok := value.(float32)
	if !ok {
		log.Fatalf(
			"Error: (GetFloat) failed to type-cast %v into a float",
			value,
		)
	}
	return floatValue
}

func (m mem) GetChar(vAddr int, isTemp bool) rune {
	addr := memory.ConvertAddr(vAddr)
	var value interface{}
	if isTemp {
		value = m.temps[constants.TYPECHAR][addr]
	} else {
		value = m.vars[constants.TYPECHAR][addr]
	}
	charValue, ok := value.(rune)
	if !ok {
		log.Fatalf(
			"Error: (GetChar) failed to type-cast %v into a char",
			value,
		)
	}
	return charValue
}

func (m mem) GetBool(vAddr int, isTemp bool) bool {
	addr := memory.ConvertAddr(vAddr)
	var value interface{}
	if isTemp {
		value = m.temps[constants.TYPECHAR][addr]
	} else {
		value = m.vars[constants.TYPECHAR][addr]
	}
	boolValue, ok := value.(bool)
	if !ok {
		log.Fatalf(
			"Error: (GetBool) failed to type-cast %v into a bool",
			value,
		)
	}
	return boolValue
}

/*
map[int][]int

x : float = 10.0

func test {
	a, b, c : int = 1, 2, 3
	ch : char
	a = b + x
	test()
}

global.getFloat(1001) -> 10.0

global = [TYPEINT : [],
		TYPEFLOAT : [10.0],
		TYPEBOOL : [],
		TYPECHAR : []]

stack [
memtest = [TYPEINT : [ 1, 2, 3],
		TYPEBOOL : [],
		TYPECHAR : []]

memtest2 = {var: [TYPEINT : [ 1, 2, 3],
					TYPEBOOL : [],
					TYPECHAR : []],

			temp: [TYPEINT : [ 1, 2, 3],
					TYPEBOOL : [],
					TYPECHAR : []]
			}

memtest3 = [TYPEINT : [ 1, 2, 3],
		TYPEBOOL : [],
		TYPECHAR : []]
]

		// 1002

*/
