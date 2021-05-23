package virtualMachine

import (
	"log"
	"math"

	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/quads"
	"github.com/jpr98/compis/semantic"
)

type VirtualMachine struct {
	globalMemBlock   Memory
	memBlocks        MemoryStack
	constantMemBlock Memory
	pointer          int // the index of current Quad

	quads       []quads.Quad
	programName string
}

func NewVM(programName string, functionTable semantic.FunctionTable, quads []quads.Quad) VirtualMachine {
	gmb := NewMemory(functionTable[programName].VarsSize, functionTable[programName].TempSize)
	return VirtualMachine{globalMemBlock: gmb, constantMemBlock: MakeConstantMemory(), quads: quads, programName: programName}
}

func (vm *VirtualMachine) Run() {
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

		case quads.ASSIGN:
			value := vm.globalMemBlock.Get(quad.Left.GetAddr())
			err := vm.globalMemBlock.Set(value, quad.Result.GetAddr())
			if err != nil {
				log.Fatalf("Error: (Run) quads.ASSIGN %s", err)
			}
			log.Println("Assigning ", value)
			vm.pointer++

		case quads.GOTO:
			vm.pointer = quad.Result.GetAddr()

		case quads.GOTOF:
			value, ok := vm.globalMemBlock.Get(quad.Left.GetAddr()).(bool)
			if !ok {
				log.Fatalf(
					"Error: (Run) quads.GOTOF failed to cast %v to bool",
					vm.globalMemBlock.Get(quad.Left.GetAddr()),
				)
			}
			if !value {
				vm.pointer = quad.Result.GetAddr()
			} else {
				vm.pointer++
			}

		default:
			vm.pointer++
		}
	}
}

func (vm *VirtualMachine) handleArithmeticOp(quad quads.Quad) {
	left, ok := vm.globalMemBlock.Get(quad.Left.GetAddr()).(float64)
	if !ok {
		log.Fatalf(
			"Error: (handleAdd) couldn't cast %v to float64",
			vm.globalMemBlock.Get(quad.Left.GetAddr()),
		)
	}

	right, ok := vm.globalMemBlock.Get(quad.Right.GetAddr()).(float64)
	if !ok {
		log.Fatalf(
			"Error: (handleAdd) couldn't cast %v to float64",
			vm.globalMemBlock.Get(quad.Right.GetAddr()),
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
	err := vm.globalMemBlock.Set(res, quad.Result.GetAddr())
	if err != nil {
		log.Fatalf("Error (handleArithmeticOp) %s", err)
	}
}

func (vm *VirtualMachine) handleRelOp(quad quads.Quad) {
	var res bool
	lTemp := vm.globalMemBlock.Get(quad.Left.GetAddr())
	switch lTemp.(type) {
	case rune:
		left := lTemp.(rune)
		right, ok := vm.globalMemBlock.Get(quad.Right.GetAddr()).(rune)
		if !ok {
			log.Fatalf(
				"Error: (handleRelOp) couldn't cast %v to rune",
				vm.globalMemBlock.Get(quad.Right.GetAddr()),
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
		right, ok := vm.globalMemBlock.Get(quad.Right.GetAddr()).(float64)
		if !ok {
			log.Fatalf(
				"Error: (handleRelOp) couldn't cast %v to float64",
				vm.globalMemBlock.Get(quad.Right.GetAddr()),
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

	err := vm.globalMemBlock.Set(res, quad.Result.GetAddr())
	if err != nil {
		log.Fatalf("Error: (handleRelOp) %s", err)
	}
}

func (vm *VirtualMachine) handleLogicOp(quad quads.Quad) {
	left, ok := vm.globalMemBlock.Get(quad.Left.GetAddr()).(bool)
	if !ok {
		log.Fatalf(
			"Error: (handleLogicOp) couldn't cast %v to bool",
			vm.globalMemBlock.Get(quad.Right.GetAddr()),
		)
	}
	right, ok := vm.globalMemBlock.Get(quad.Right.GetAddr()).(bool)
	if !ok {
		log.Fatalf(
			"Error: (handleLogicOp) couldn't cast %v to bool",
			vm.globalMemBlock.Get(quad.Right.GetAddr()),
		)
	}

	var res bool
	switch quad.Action {
	case quads.AND:
		res = left && right
	case quads.OR:
		res = left || right
	}
	err := vm.globalMemBlock.Set(res, quad.Result.GetAddr())
	if err != nil {
		log.Fatalf("Error: (handleLogicOp) %s", err)
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
