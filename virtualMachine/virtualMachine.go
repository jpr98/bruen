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
			memblock := vm.getMemBlockForAddr(quad.Left.GetAddr())
			value := memblock.Get(quad.Left.GetAddr())
			memblock = vm.getMemBlockForAddr(quad.Result.GetAddr())
			err := memblock.Set(value, quad.Result.GetAddr())
			if err != nil {
				log.Fatalf("Error: (Run) quads.ASSIGN %s", err)
			}
			vm.pointer++

		case quads.GOTO:
			vm.pointer = quad.Result.GetAddr()

		case quads.GOTOF:
			memblock := vm.getMemBlockForAddr(quad.Left.GetAddr())
			value, ok := memblock.Get(quad.Left.GetAddr()).(bool)
			if !ok {
				log.Fatalf(
					"Error: (Run) quads.GOTOF failed to cast %v to bool",
					memblock.Get(quad.Left.GetAddr()),
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

		case quads.PARAM:
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
	memblock := vm.getMemBlockForAddr(quad.Left.GetAddr())
	left, ok := memblock.Get(quad.Left.GetAddr()).(float64)
	if !ok {
		log.Fatalf(
			"Error: (handleArithmeticOp) 1 couldn't cast %v to float64",
			memblock.Get(quad.Left.GetAddr()),
		)
	}

	memblock = vm.getMemBlockForAddr(quad.Right.GetAddr())
	right, ok := memblock.Get(quad.Right.GetAddr()).(float64)
	if !ok {
		log.Fatalf(
			"Error: (handleArithmeticOp) 2 couldn't cast %v to float64",
			memblock.Get(quad.Right.GetAddr()),
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

	memblock = vm.getMemBlockForAddr(quad.Result.GetAddr())
	err := memblock.Set(res, quad.Result.GetAddr())
	if err != nil {
		log.Fatalf("Error (handleArithmeticOp) %s", err)
	}
}

func (vm *VirtualMachine) handleRelOp(quad quads.Quad) {
	var res bool
	memblock := vm.getMemBlockForAddr(quad.Left.GetAddr())
	lTemp := memblock.Get(quad.Left.GetAddr())
	switch lTemp.(type) {
	case rune:
		left := lTemp.(rune)
		memblock = vm.getMemBlockForAddr(quad.Right.GetAddr())
		right, ok := memblock.Get(quad.Right.GetAddr()).(rune)
		if !ok {
			log.Fatalf(
				"Error: (handleRelOp) couldn't cast %v to rune",
				memblock.Get(quad.Right.GetAddr()),
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
		memblock = vm.getMemBlockForAddr(quad.Right.GetAddr())
		right, ok := memblock.Get(quad.Right.GetAddr()).(float64)
		if !ok {
			log.Fatalf(
				"Error: (handleRelOp) couldn't cast %v to float64",
				memblock.Get(quad.Right.GetAddr()),
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

	memblock = vm.getMemBlockForAddr(quad.Result.GetAddr())
	err := memblock.Set(res, quad.Result.GetAddr())
	if err != nil {
		log.Fatalf("Error: (handleRelOp) %s", err)
	}
}

func (vm *VirtualMachine) handleLogicOp(quad quads.Quad) {
	memblock := vm.getMemBlockForAddr(quad.Left.GetAddr())
	left, ok := memblock.Get(quad.Left.GetAddr()).(bool)
	if !ok {
		log.Fatalf(
			"Error: (handleLogicOp) couldn't cast %v to bool",
			memblock.Get(quad.Left.GetAddr()),
		)
	}
	memblock = vm.getMemBlockForAddr(quad.Right.GetAddr())
	right, ok := memblock.Get(quad.Right.GetAddr()).(bool)
	if !ok {
		log.Fatalf(
			"Error: (handleLogicOp) couldn't cast %v to bool",
			memblock.Get(quad.Right.GetAddr()),
		)
	}

	var res bool
	switch quad.Action {
	case quads.AND:
		res = left && right
	case quads.OR:
		res = left || right
	}
	memblock = vm.getMemBlockForAddr(quad.Result.GetAddr())
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

	switch quad.Result.Type() {
	case constants.TYPEINT:
		floatVal, err := strconv.ParseFloat(str, 64)
		if err != nil {
			log.Println("Warning: read to int expects int")
		}

		err = memblock.Set(float64(int(floatVal)), quad.Result.GetAddr())
		if err != nil {
			log.Fatalf("Error: (handleRead) %s", err)
		}

	case constants.TYPEFLOAT:
		floatVal, err := strconv.ParseFloat(str, 64)
		if err != nil {
			//TODO: handle err
		}
		err = memblock.Set(floatVal, quad.Result.GetAddr())
		if err != nil {
			//TODO: handle err
		}

	case constants.TYPECHAR:
		if len(str) > 0 {
			runeVal := str[0]
			err := memblock.Set(runeVal, quad.Result.GetAddr())
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
		err = memblock.Set(boolVal, quad.Result.GetAddr())
		if err != nil {
			//TODO: handle err
		}
	}
}

func (vm *VirtualMachine) handleWrite(quad quads.Quad) {
	memblock := vm.getMemBlockForAddr(quad.Result.GetAddr())
	value := memblock.Get(quad.Result.GetAddr())
	switch value.(type) {
	case float64:
		fmt.Print(value.(float64))
	case rune:
		fmt.Print(strconv.QuoteRune(value.(rune)))
	case bool:
		fmt.Println(value.(bool))
	default:
		log.Fatalln("Accesing uninitialized variable")
	}
}

/*

	LPAREN
	ASSIGN

	ENDFUNC
	ERA
	PARAM
	GOSUB
*/
