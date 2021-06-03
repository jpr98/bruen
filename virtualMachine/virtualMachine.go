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

// VirtualMachine is in charge of executing a given set of instructions. It contains
// and manages the memory needed to execute a program.
type VirtualMachine struct {
	globalMemBlock   Memory       // The global memory block of the program
	memBlocks        MemoryStack  // The stack of functions memory block
	constantMemBlock Memory       // The constants memory block
	pointer          int          // The instruction pointer of current Quad
	pointerStack     PointerStack // Keeps track of pending instruction pointers

	currentSelf Memory // The memory block for an instance, acting as self in a method

	quads         []quads.Quad           // The given program instructions
	programName   string                 // The name of the program (also known as globalName)
	functionTable semantic.FunctionTable // The FunctionTable of the compiled program
	classTable    semantic.ClassTable    // The ClassTable of the compiled program
}

// NewVM creates a virtual machine for the given program, function and class tables and instruction set
func NewVM(programName string, functionTable semantic.FunctionTable, classTable semantic.ClassTable, quads []quads.Quad) VirtualMachine {
	gmb := NewMemory(functionTable[programName].VarsSize, functionTable[programName].TempSize, functionTable[programName].ObjSize)
	return VirtualMachine{
		globalMemBlock:   gmb,
		constantMemBlock: MakeConstantMemory(),
		currentSelf:      nil,
		quads:            quads,
		programName:      programName,
		functionTable:    functionTable,
		classTable:       classTable,
	}
}

// Run iterates over the program's instructions and handles each quad.Action correspondingly
func (vm *VirtualMachine) Run() {
	fmb := MemoryStack{} // keeps Memory of functions not yet activated
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
			vm.handleAssign(quad)
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
			if quad.Left.ClassName() == "" {
				ftc := vm.functionTable[quad.Left.ID()]
				fmb.Push(NewMemory(ftc.VarsSize, ftc.TempSize, ftc.ObjSize))
				if quad.Left.ID() == "main" {
					vm.memBlocks.Push(fmb.Pop())
				}
			} else {
				methodInfo := vm.classTable[quad.Left.ClassName()].Methods[quad.Left.ID()]
				fmb.Push(NewMemory(methodInfo.VarsSize, methodInfo.TempSize, methodInfo.ObjSize))
			}

			vm.pointer++

		case quads.PARAM: // TODO : revisar con arreglos
			memblock := vm.getMemBlockForAddr(quad.Left.GetAddr())
			value := memblock.Get(quad.Left.GetAddr())

			err := fmb.Top().Set(value, quad.Right.GetAddr())
			if err != nil {
				log.Fatalf("Error: (Run) quads.PARAM %s", err)
			}
			vm.pointer++

		case quads.INSTANCE:
			vm.handleInstance(quad)
			vm.pointer++

		case quads.GOSUB:
			vm.memBlocks.Push(fmb.Pop())
			vm.pointerStack.Push(vm.pointer + 1)
			if quad.Left.ClassName() != "" {
				vm.pointer = vm.classTable[quad.Left.ClassName()].Methods[quad.Left.ID()].Dir
			} else {
				vm.pointer = vm.functionTable[quad.Left.ID()].Dir
			}

		case quads.ENDFUNC:
			vm.memBlocks.Pop()
			vm.currentSelf = nil
			vm.pointer = vm.pointerStack.Pop()

		case quads.RETURN:
			if quad.Result != nil {
				value := vm.getValueForElement(quad.Result)
				memblock := vm.getMemBlockForElement(quad.Left)
				memblock.Set(value, quad.Left.GetAddr())
			}
			vm.memBlocks.Pop()
			vm.pointer = vm.pointerStack.Pop()

		default:
			vm.pointer++
		}
	}
}

// getMemBlockForElement returns the memory block for a given element's address. If the
// element has the `self_` tag it gets the memory block of the self instance indicated by
// the addr in the tag. If the tag is -1 we use the currentSelf.
func (vm *VirtualMachine) getMemBlockForElement(elem quads.Element) Memory {
	if strings.Contains(elem.ID(), "self_") {
		strElements := strings.Split(elem.ID(), "_")
		if len(strElements) < 2 {
			log.Fatalf("Error: (getMemBlockForElement) unexpected object attribute id format")
		}
		objInstanceAddr, err := strconv.Atoi(strElements[1])
		if err != nil {
			log.Fatalf("Error: (getMemBlockForElement) couldn't cast objInstanceAddr to int")
		}
		if objInstanceAddr == -1 {
			return vm.currentSelf
		}
		memblock := vm.getMemBlockForAddr(objInstanceAddr)
		objInstance, ok := memblock.Get(objInstanceAddr).(Memory)
		if !ok {
			log.Fatalf("Error: (getMemBlockForElement) couldn't cast %v to Memory", memblock.Get(objInstanceAddr))
		}
		return objInstance
	}
	return vm.getMemBlockForAddr(elem.GetAddr())
}

// getMemBlockForAddr returns the memory block for a given address
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

// handleInstance sets the self instance for the given quad, if the self dir is -1
// we continue using the currentSelf instance
func (vm *VirtualMachine) handleInstance(quad quads.Quad) {
	if strings.Contains(quad.Left.ID(), "self_") {
		strElements := strings.Split(quad.Left.ID(), "_")
		if len(strElements) < 2 {
			log.Fatalf("Error: (handleInstance) unexpected element id structure")
		}
		objInstanceAddr, err := strconv.Atoi(strElements[1])
		if err != nil {
			log.Fatalf("Error: (handleInstance) couldn't cast objInstanceAddr to int")
		}
		if objInstanceAddr == -1 {
			return
		}
	}

	memblock := vm.getMemBlockForAddr(quad.Left.GetAddr())
	var ok bool
	vm.currentSelf, ok = memblock.Get(quad.Left.GetAddr()).(Memory)
	if !ok {
		log.Fatalf(
			"Error: (Run) quads.INSTANCE couldn't cast %v to Memory",
			memblock.Get(quad.Left.GetAddr()),
		)
	}
}

// handleArithmeticOp casts the operands to float, makes the appropiate arithemtic
// operation and sets the result in the result memory address. If the result type
// is constants.TYPEINT we round the result.
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
			log.Fatalf("Error: (handleArithmeticOp) zero division")
		}
		res = left / right
	}

	if quad.Result.Type() == constants.TYPEINT {
		res = math.Round(res)
	}

	memblock := vm.getMemBlockForElement(quad.Result)
	err := memblock.Set(res, quad.Result.GetAddr())
	if err != nil {
		log.Fatalf("Error (handleArithmeticOp) %s", err)
	}
}

// handleRelOp performs a relational operation between two values, it gets the type
// of the left operand and switches over its type to also cast the right operand to
// the same type, then it performs the operation and sets the result value.
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

	memblock := vm.getMemBlockForElement(quad.Result)
	err := memblock.Set(res, quad.Result.GetAddr())
	if err != nil {
		log.Fatalf("Error: (handleRelOp) %s", err)
	}
}

// handleLogicOp performs a logic operation between to boolean values
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
	memblock := vm.getMemBlockForElement(quad.Result)
	err := memblock.Set(res, quad.Result.GetAddr())
	if err != nil {
		log.Fatalf("Error: (handleLogicOp) %s", err)
	}
}

// handleRead creates a stdin reader and reads until a newline is found, gets the memory
// block needed to store the value, tries to cast the value to the expected type and saves
// the result.
func (vm *VirtualMachine) handleRead(quad quads.Quad) {
	reader := bufio.NewReader(os.Stdin)
	bytes, _ := reader.ReadBytes('\n')
	str := strings.TrimSpace(string(bytes))

	var memblock Memory
	var addr int
	if strings.Contains(quad.Result.ID(), "ptr_") {
		memblock = vm.getMemBlockForAddr(quad.Result.GetAddr())
		addrFloat, ok := memblock.Get(quad.Result.GetAddr()).(float64)
		if !ok {
			log.Fatalf("Error: (handleRead) couldn't cast %v to float64",
				memblock.Get(int(addr)))
		}

		addr = int(addrFloat)
		auxElement := quads.NewElement(addr, quad.Result.ID(), quad.Result.Type(), "")
		memblock = vm.getMemBlockForElement(auxElement)
	} else {
		memblock = vm.getMemBlockForElement(quad.Result)
		addr = quad.Result.GetAddr()
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
			log.Println("Warning: read to float expects float")
		}
		err = memblock.Set(floatVal, addr)
		if err != nil {
			log.Fatalf("Error: (handleRead) %s", err)
		}

	case constants.TYPECHAR:
		if len(str) > 0 {
			runeVal := str[0]
			err := memblock.Set(runeVal, addr)
			if err != nil {
				log.Fatalf("Error: (handleRead) %s", err)
			}
		} else {
			log.Println("Warning: read to char expects a single char")
		}

	case constants.TYPEBOOL:
		boolVal, err := strconv.ParseBool(str)
		if err != nil {
			fmt.Println("Warning: read to bool expects bool")
		}
		err = memblock.Set(boolVal, addr)
		if err != nil {
			log.Fatalf("Error: (handleRead) %s", err)
		}
	}
}

// handleWrite prints to stdout
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

// handleCalcDir calculates the array direction for a given index
func (vm *VirtualMachine) handleCalcDir(quad quads.Quad) {
	memblock := vm.getMemBlockForElement(quad.Right)
	index, ok := memblock.Get(quad.Right.GetAddr()).(float64)
	if !ok {
		log.Fatalf("Error: (handleCalcDir) couldn't cast index to float64")
	}
	realAddr := float64(quad.Left.GetAddr()) + index
	memblock = vm.getMemBlockForAddr(quad.Result.GetAddr()) // Here we don't call getMemBlockForElement since a pointer is always a temporary value inside a function, never an attribute, thus we don't want to look for an instance
	err := memblock.Set(realAddr, quad.Result.GetAddr())
	if err != nil {
		log.Fatalf("Error: (handleCalcDir) %s", err)
	}
}

// handleAssign performs an assignation instruction, if the result is an array
// we first get the real address to the index of the array, then we set the value.
func (vm *VirtualMachine) handleAssign(quad quads.Quad) {
	value := vm.getValueForElement(quad.Left)

	if strings.Contains(quad.Result.ID(), "ptr_") {
		memblock := vm.getMemBlockForAddr(quad.Result.GetAddr())
		addr, ok := memblock.Get(quad.Result.GetAddr()).(float64)
		if !ok {
			log.Fatalf("Error: (handleAssign) couldn't cast %v to float64",
				memblock.Get(quad.Result.GetAddr()),
			)
		}

		auxElement := quads.NewElement(int(addr), quad.Result.ID(), quad.Result.Type(), "")
		memblock = vm.getMemBlockForElement(auxElement)
		err := memblock.Set(value, int(addr))
		if err != nil {
			log.Fatalf("Error: (handleAssign) 1 %s", err)
		}
	} else {
		memblock := vm.getMemBlockForElement(quad.Result)
		err := memblock.Set(value, quad.Result.GetAddr())
		if err != nil {
			log.Fatalf("Error: (handleAssign) 2 %s", err)
		}
	}
}

// getValueForElement returns the value for a given element memory address, it takes
// into consideration ptr_ and self_ tags.
func (vm *VirtualMachine) getValueForElement(e quads.Element) interface{} {
	if strings.Contains(e.ID(), "ptr_") {
		memblock := vm.getMemBlockForAddr(e.GetAddr())
		ptrAddr, ok := memblock.Get(e.GetAddr()).(float64)
		if !ok {
			log.Fatalf("Error: (getValueForElement) couldn't cast index to float64")
		}

		auxElement := quads.NewElement(int(ptrAddr), e.ID(), e.Type(), "")
		memblock = vm.getMemBlockForElement(auxElement)
		realValue := memblock.Get(int(ptrAddr))
		return realValue
	}
	memblock := vm.getMemBlockForElement(e)
	return memblock.Get(e.GetAddr())
}
