package virtualMachine

import (
	"fmt"
	"strconv"

	"github.com/jpr98/bruen/constants"
	"github.com/jpr98/bruen/memory"
	"github.com/jpr98/bruen/semantic"
)

// Memory is the interface representing a memory block in the virtual machine
type Memory interface {
	// Get returns the value stored in a given address
	Get(addr int) interface{}
	// Set assigns a value to a given address
	Set(value interface{}, addr int) error
}

// mem is the memory block on which the Memory interfacee makes operations
type mem struct {
	vars    memblock // Local atomic variables
	temps   memblock // Temporary atomic variables
	objects []Memory // Memory blocks for constants.TYPECLASS variables
}

// memblock is a mapping from type to an array of stores values
type memblock map[constants.Type][]interface{}

func (m mem) String() string {
	return fmt.Sprintf("Local: %v, Temp: %v, Objects: %v", m.vars, m.temps, m.objects)
}

// NewMemory creates a new Memory with the given sizes
func NewMemory(varSizes, tempSizes [4]int, objInfo []memory.MemObjInfo) Memory {
	m := mem{}
	m.vars = make(memblock)
	m.temps = make(memblock)
	m.objects = make([]Memory, len(objInfo))
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

	for index, obj := range objInfo {
		m.objects[index] = NewMemory(obj.VarSize, [4]int{}, obj.ObjSize)
	}
	return &m
}

func (m *mem) Get(addr int) interface{} {
	vAddr := memory.ConvertAddr(addr)
	typeOf := memory.TypeForAddr(addr)

	if typeOf == constants.TYPECLASS {
		return m.objects[vAddr]
	}
	if memory.IsTempAddr(addr) {
		return m.temps[typeOf][vAddr]
	}
	return m.vars[typeOf][vAddr]
}

func (m *mem) Set(value interface{}, addr int) error {
	vAddr := memory.ConvertAddr(addr)
	typeOf := memory.TypeForAddr(addr)

	switch typeOf {
	case constants.TYPEINT, constants.TYPEFLOAT:
		numberValue, ok := value.(float64)
		if !ok {
			return fmt.Errorf("Error: (Set) couldn't cast %v to float64", value)
		}
		if memory.IsTempAddr(addr) {
			m.temps[typeOf][vAddr] = numberValue
		} else {
			m.vars[typeOf][vAddr] = numberValue
		}
		return nil
	case constants.TYPECHAR:
		runeValue, ok := value.(rune)
		if !ok {
			return fmt.Errorf("Error: (Set) couldn't cast %v to rune", value)
		}
		if memory.IsTempAddr(addr) {
			m.temps[typeOf][vAddr] = runeValue
		} else {
			m.vars[typeOf][vAddr] = runeValue
		}
		return nil
	case constants.TYPEBOOL:
		boolValue, ok := value.(bool)
		if !ok {
			return fmt.Errorf("Error: (Set) couldn't cast %v to bool", value)
		}
		if memory.IsTempAddr(addr) {
			m.temps[typeOf][vAddr] = boolValue
		} else {
			m.vars[typeOf][vAddr] = boolValue
		}
		return nil
	case constants.TYPECLASS:
		memValue, ok := value.(Memory)
		if !ok {
			return fmt.Errorf("Error: (Set) couldn't cast %v to Memory", value)
		}
		m.objects[vAddr] = memValue
		return nil
	default:
		return fmt.Errorf("Error: (Set) unexpected type to set in memory block")
	}
}

// MakeConstantMemory creates the Memory for constants taken from the global variable semantic.ConstantsTable
func MakeConstantMemory() Memory {
	m := mem{}
	m.vars = make(memblock)
	var size [4]int
	for _, content := range semantic.ConstantsTable {
		size[content.TypeOf]++
	}
	m.vars[constants.TYPEINT] = make([]interface{}, size[constants.TYPEINT])
	m.vars[constants.TYPEFLOAT] = make([]interface{}, size[constants.TYPEFLOAT])
	m.vars[constants.TYPECHAR] = make([]interface{}, size[constants.TYPECHAR])
	m.vars[constants.TYPEBOOL] = make([]interface{}, size[constants.TYPEBOOL])

	for cte, content := range semantic.ConstantsTable {
		vAddr := memory.ConvertAddr(content.Dir)
		switch content.TypeOf {
		case constants.TYPEINT:
			if intValue, err := strconv.ParseFloat(cte, 64); err == nil {
				m.vars[constants.TYPEINT][vAddr] = intValue
			}
		case constants.TYPEFLOAT:
			if floatValue, err := strconv.ParseFloat(cte, 64); err == nil {
				m.vars[constants.TYPEFLOAT][vAddr] = floatValue
			}
		case constants.TYPECHAR:
			runes := []rune(cte)
			if len(runes) > 0 {
				m.vars[constants.TYPECHAR][vAddr] = runes[0]
			}
		case constants.TYPEBOOL:
			if boolValue, err := strconv.ParseBool(cte); err == nil {
				m.vars[constants.TYPEBOOL][vAddr] = boolValue
			}
		}
	}
	return &m
}
