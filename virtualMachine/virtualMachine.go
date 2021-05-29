package virtualMachine

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/memory"
	"github.com/jpr98/compis/quads"
	"github.com/jpr98/compis/semantic"
)

type VirtualMachine struct {
	globalMemBlock   Memory
	memBlocks        MemoryStack
	constantMemBlock Memory
	pointer          int // the index of current Quad
	pointerStack     PointerStack

	quads         []quads.Quad
	programName   string
	functionTable semantic.FunctionTable
}

func NewVM(programName string, functionTable semantic.FunctionTable, quads []quads.Quad) VirtualMachine {
	gmb := NewMemory(functionTable[programName].VarsSize, functionTable[programName].TempSize)
	return VirtualMachine{
		globalMemBlock:   gmb,
		constantMemBlock: MakeConstantMemory(),
		quads:            quads,
		programName:      programName,
		functionTable:    functionTable,
	}
}

func (vm *VirtualMachine) Run() {
	var fmb Memory
	for vm.pointer < len(vm.quads) {
		quad := vm.quads[vm.pointer]
		switch quad.Action {
		case quads.ADD, quads.SUB, quads.MUL, quads.DIV:
			vm.handleArithmeticOp(quad)
			vm.pointer++

		case quads.GT, quads.LT, quads.EQ, quads.NEQ:
			vm.handleRelOp(quad)
			vm.pointer++

		case quads.AND, quads.OR:
			vm.handleLogicOp(quad)
			vm.pointer++

		case quads.READ:
			vm.handleRead(quad)
			vm.pointer++

		case quads.WRITE:
			vm.handleWrite(quad)
			vm.pointer++

		case quads.WRITENEWLINE:
			fmt.Println("")
			vm.pointer++

		case quads.ASSIGN:
			value := vm.getValueForElement(quad.Left)
			memblock := vm.getMemBlockForAddr(quad.Result.GetAddr())
			if strings.Contains(quad.Result.ID(), "ptr_") {
				addr, ok := memblock.Get(quad.Result.GetAddr()).(float64)
				if !ok {
					log.Fatalf("Error: (RUN) quads.ASSIGN couldn't cast %v to float64",
						memblock.Get(quad.Result.GetAddr()),
					)
				}
				memblock = vm.getMemBlockForAddr(int(addr))
				fmt.Println("here: ", value)
				err := memblock.Set(value, int(addr))
				if err != nil {
					log.Fatalf("Error: (Run) quads.ASSIGN %s", err)
				}
			} else {
				fmt.Println("here1: ", value)
				err := memblock.Set(value, quad.Result.GetAddr())
				if err != nil {
					log.Fatalf("Error: (Run) quads.ASSIGN %s", err)
				}
			}
			vm.pointer++

		case quads.VER:
			indexValue, ok := vm.getValueForElement(quad.Left).(float64)
			if !ok {
				log.Fatalf("Error: (Run) quads.VER couldn't cast index to float64")
			}
			// Result.GetAddr contains the dimension of the array
			if int(indexValue) >= quad.Result.GetAddr() || indexValue < 0 {
				log.Fatalf("Error: Index out of range %s[%d]", quad.Result.ID(), int(indexValue))
			}
			vm.pointer++

		case quads.CALCDIR:
			vm.handleCalcDir(quad)
			vm.pointer++

		case quads.GOTO:
			vm.pointer = quad.Result.GetAddr()

		case quads.GOTOF:
			value, ok := vm.getValueForElement(quad.Left).(bool)
			if !ok {
				log.Fatalf(
					"Error: (Run) quads.GOTOF failed to cast %v to bool",
					vm.getValueForElement(quad.Left),
				)
			}
			if !value {
				vm.pointer = quad.Result.GetAddr()
			} else {
				vm.pointer++
			}

		case quads.ERA:
			fmb = NewMemory(vm.functionTable[quad.Left.ID()].VarsSize, vm.functionTable[quad.Left.ID()].TempSize)
			if quad.Left.ID() == "main" {
				vm.memBlocks.Push(fmb)
			}
			vm.pointer++

		case quads.PARAM: // TODO : revisar con arreglos
			memblock := vm.getMemBlockForAddr(quad.Left.GetAddr())
			value := memblock.Get(quad.Left.GetAddr())

			err := fmb.Set(value, quad.Right.GetAddr())
			if err != nil {
				log.Fatalf("Error: (Run) quads.PARAM %s", err)
			}
			vm.pointer++

		case quads.GOSUB:
			vm.memBlocks.Push(fmb)
			vm.pointerStack.Push(vm.pointer + 1)
			vm.pointer = vm.functionTable[quad.Left.ID()].Dir

		case quads.ENDFUNC:
			vm.memBlocks.Pop()
			vm.pointer = vm.pointerStack.Pop()

		case quads.RETURN:
			if quad.Result != nil {
				memblock := vm.getMemBlockForAddr(quad.Result.GetAddr())
				value := memblock.Get(quad.Result.GetAddr())
				memblock = vm.getMemBlockForAddr(quad.Left.GetAddr())
				memblock.Set(value, quad.Left.GetAddr())
			}
			vm.memBlocks.Pop()
			vm.pointer = vm.pointerStack.Pop()

		default:
			vm.pointer++
		}
	}
}

func (vm *VirtualMachine) getMemBlockForAddr(addr int) Memory {
	context := memory.ContextForAddr(addr)
	var memBlock Memory
	switch context {
	case memory.Global:
		memBlock = vm.globalMemBlock
	case memory.Local, memory.Temp:
		memBlock = vm.memBlocks.Top()
	case memory.Constant:
		memBlock = vm.constantMemBlock
	case memory.Invalid:
		log.Fatalf("Error: (getMemBlockForAddr) invalid address %d", addr)
	}
	return memBlock
}

func (vm *VirtualMachine) handleArithmeticOp(quad quads.Quad) {
	left, ok := vm.getValueForElement(quad.Left).(float64)
	if !ok {
		log.Fatalf(
			"Error: (handleArithmeticOp) 1 couldn't cast %v to float64",
			vm.getValueForElement(quad.Left),
		)
	}

	right, ok := vm.getValueForElement(quad.Right).(float64)
	if !ok {
		log.Fatalf(
			"Error: (handleArithmeticOp) 2 couldn't cast %v to float64",
			vm.getValueForElement(quad.Right),
		)
	}

	var res float64
	switch quad.Action {
	case quads.ADD:
		res = left + right
	case quads.SUB:
		res = left - right
	case quads.MUL:
		res = left * right
	case quads.DIV:
		if right == 0 {
			log.Fatalf("Error: (handleAdd) zero division")
		}
		res = left / right
	}

	if quad.Result.Type() == constants.TYPEINT {
		res = math.Round(res)
	}

	memblock := vm.getMemBlockForAddr(quad.Result.GetAddr())
	err := memblock.Set(res, quad.Result.GetAddr())
	if err != nil {
		log.Fatalf("Error (handleArithmeticOp) %s", err)
	}
}

func (vm *VirtualMachine) handleRelOp(quad quads.Quad) {
	var res bool
	lTemp := vm.getValueForElement(quad.Left)
	switch lTemp.(type) {
	case rune:
		left := lTemp.(rune)
		right, ok := vm.getValueForElement(quad.Right).(rune)
		if !ok {
			log.Fatalf(
				"Error: (handleRelOp) couldn't cast %v to rune",
				vm.getValueForElement(quad.Right),
			)
		}

		switch quad.Action {
		case quads.GT:
			res = left > right
		case quads.LT:
			res = left < right
		case quads.EQ:
			res = left == right
		case quads.NEQ:
			res = left != right
		}
	case float64:
		left := lTemp.(float64)
		right, ok := vm.getValueForElement(quad.Right).(float64)
		if !ok {
			log.Fatalf(
				"Error: (handleRelOp) couldn't cast %v to float64",
				vm.getValueForElement(quad.Right),
			)
		}

		switch quad.Action {
		case quads.GT:
			res = left > right
		case quads.LT:
			res = left < right
		case quads.EQ:
			res = left == right
		case quads.NEQ:
			res = left != right
		}
	}

	memblock := vm.getMemBlockForAddr(quad.Result.GetAddr())
	err := memblock.Set(res, quad.Result.GetAddr())
	if err != nil {
		log.Fatalf("Error: (handleRelOp) %s", err)
	}
}

func (vm *VirtualMachine) handleLogicOp(quad quads.Quad) {
	left, ok := vm.getValueForElement(quad.Left).(bool)
	if !ok {
		log.Fatalf(
			"Error: (handleLogicOp) couldn't cast %v to bool",
			vm.getValueForElement(quad.Left),
		)
	}
	right, ok := vm.getValueForElement(quad.Right).(bool)
	if !ok {
		log.Fatalf(
			"Error: (handleLogicOp) couldn't cast %v to bool",
			vm.getValueForElement(quad.Right),
		)
	}

	var res bool
	switch quad.Action {
	case quads.AND:
		res = left && right
	case quads.OR:
		res = left || right
	}
	memblock := vm.getMemBlockForAddr(quad.Result.GetAddr())
	err := memblock.Set(res, quad.Result.GetAddr())
	if err != nil {
		log.Fatalf("Error: (handleLogicOp) %s", err)
	}
}

func (vm *VirtualMachine) handleRead(quad quads.Quad) {
	reader := bufio.NewReader(os.Stdin)
	bytes, _ := reader.ReadBytes('\n')
	str := strings.TrimSpace(string(bytes))

	memblock := vm.getMemBlockForAddr(quad.Result.GetAddr())
	addr := quad.Result.GetAddr()
	if strings.Contains(quad.Result.ID(), "ptr_") {
		addrFloat, ok := memblock.Get(addr).(float64)
		if !ok {
			log.Fatalf("Error: (handleRead) couldn't cast %v to float64",
				memblock.Get(int(addr)))
		}
		addr = int(addrFloat)
		memblock = vm.getMemBlockForAddr(addr)
	}

	switch quad.Result.Type() {
	case constants.TYPEINT:
		floatVal, err := strconv.ParseFloat(str, 64)
		if err != nil {
			log.Println("Warning: read to int expects int")
		}

		err = memblock.Set(float64(int(floatVal)), addr)
		if err != nil {
			log.Fatalf("Error: (handleRead) %s", err)
		}

	case constants.TYPEFLOAT:
		floatVal, err := strconv.ParseFloat(str, 64)
		if err != nil {
			//TODO: handle err
		}
		err = memblock.Set(floatVal, addr)
		if err != nil {
			//TODO: handle err
		}

	case constants.TYPECHAR:
		if len(str) > 0 {
			runeVal := str[0]
			err := memblock.Set(runeVal, addr)
			if err != nil {
				//TODO: handle err
			}
		} else {
			//TODO: handle err
		}

	case constants.TYPEBOOL:
		boolVal, err := strconv.ParseBool(str)
		if err != nil {
			fmt.Println("Warning:")
		}
		err = memblock.Set(boolVal, addr)
		if err != nil {
			//TODO: handle err
		}
	}
}

func (vm *VirtualMachine) handleWrite(quad quads.Quad) {
	value := vm.getValueForElement(quad.Result)
	switch value.(type) {
	case float64:
		fmt.Print(value.(float64))
	case rune:
		fmt.Print(strconv.QuoteRune(value.(rune)))
	case bool:
		fmt.Print(value.(bool))
	default:
		log.Fatalf("Accesing uninitialized variable %s", quad.Result.ID())
	}
}

func (vm *VirtualMachine) handleCalcDir(quad quads.Quad) {
	memblock := vm.getMemBlockForAddr(quad.Right.GetAddr())
	index, ok := memblock.Get(quad.Right.GetAddr()).(float64)
	if !ok {
		log.Fatalf("Error: (handleCalcDir) couldn't cast index to float64")
	}
	realAddr := float64(quad.Left.GetAddr()) + index
	memblock = vm.getMemBlockForAddr(quad.Result.GetAddr())
	err := memblock.Set(realAddr, quad.Result.GetAddr())
	if err != nil {
		log.Fatalf("Error: (handleCalcDir) %s", err)
	}
}

func (vm *VirtualMachine) getValueForElement(e quads.Element) interface{} {
	memblock := vm.getMemBlockForAddr(e.GetAddr())
	if strings.Contains(e.ID(), "ptr_") {
		ptrAddr, ok := memblock.Get(e.GetAddr()).(float64)
		if !ok {
			log.Fatalf("Error: (pointerValue) couldn't cast index to float64")
		}

		memblock = vm.getMemBlockForAddr(int(ptrAddr))
		realValue := memblock.Get(int(ptrAddr))
		return realValue
	}
	return memblock.Get(e.GetAddr())
}
