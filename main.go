package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/jpr98/bruen/memory"
	"github.com/jpr98/bruen/parser"
	"github.com/jpr98/bruen/quads"
	"github.com/jpr98/bruen/semantic"
	"github.com/jpr98/bruen/virtualMachine"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("No args")
	}

	switch os.Args[2] {
	case "b": // build
		compile()
	case "e": // execute
		execute()
	case "r": // run
		compile()
		execute()
	default:
		fmt.Println("What? Unknown command")
	}
}

func compile() {
	if len(os.Args) < 2 {
		log.Fatal("No filename provided")
	}

	// Open file stream
	filename := os.Args[1]
	is, err := antlr.NewFileStream(filename)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	// Creates the lexer
	lexer := parser.NewBruenLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Creates the parser
	p := parser.NewBruenParser(stream)

	p.RemoveErrorListeners()
	errorListener := semantic.NewErrorListener()
	p.AddErrorListener(&errorListener)

	listener := semantic.NewListener()
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Program())

	if errorListener.FoundError {
		os.Exit(1)
	}

	stream.Seek(0)
	p = parser.NewBruenParser(stream)

	var quadListener quads.QuadGenListener = quads.NewListener(listener.GetFunctionTable(), listener.GetClassTable())
	antlr.ParseTreeWalkerDefault.Walk(&quadListener, p.Program())
	quads := quadListener.GetManager().GetQuads()

	if len(os.Args) > 3 {
		if os.Args[3] == "-d" {
			for i, q := range quads {
				fmt.Printf("%d. %s\n", i, q)
			}
		}
	}

	f, err := os.Create("out.obj")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	enc.Encode(listener.ProgramName)
	enc.Encode(quads)
	enc.Encode(listener.GetFunctionTable())
	enc.Encode(listener.GetClassTable())
	enc.Encode(semantic.ConstantsTable)
}

func execute() {
	f, err := os.Open("out.obj")
	if err != nil {
		panic(err)
	}

	dec := gob.NewDecoder(f)

	var programName string
	err = dec.Decode(&programName)
	if err != nil {
		panic(err)
	}

	var m []quads.Quad
	err = dec.Decode(&m)
	if err != nil {
		panic(err)
	}
	// fmt.Println("\n---------")
	// for i, q := range m {
	// 	fmt.Printf("%d. %s\n", i, q)
	// }
	// fmt.Println("\n---------")

	var ft semantic.FunctionTable
	err = dec.Decode(&ft)
	if err != nil {
		panic(err)
	}
	//debugFT(ft)

	var classTable semantic.ClassTable
	err = dec.Decode(&classTable)
	if err != nil {
		panic(err)
	}

	err = dec.Decode(&semantic.ConstantsTable)
	if err != nil {
		panic(err)
	}

	vm := virtualMachine.NewVM(programName, ft, classTable, m)
	vm.Run()
}

func debugFT(ft semantic.FunctionTable) {
	for fname, function := range ft {
		fmt.Printf("\n *Function: %s Returns: %s Scope: %s, VarSize: %v, TempSize: %v\n", fname, function.TypeOf, function.Scope, function.VarsSize, function.TempSize)
		fmt.Println("VarTable:")
		for id, variable := range function.Vars {
			fmt.Printf("%s: %s - %d\n", id, variable.TypeOf, variable.Dir)
		}
		fmt.Printf("Dir: %d\n", function.Dir)
		debugObjSize(function.ObjSize)
	}
}

func debugCT(ct semantic.ClassTable) {
	for class, classContent := range ct {
		fmt.Println(class)
		for attr, attrContent := range classContent.Attributes {
			fmt.Printf("\t%s: %s\n", attr, attrContent.TypeOf)
		}
		for funct, functContent := range classContent.Methods {
			fmt.Printf("Func %s: %s\n", funct, functContent.TypeOf)
			for v, vattr := range functContent.Vars {
				fmt.Printf("\t%s : %s\n", v, vattr.TypeOf)
			}
			debugObjSize(functContent.ObjSize)
		}
	}
}

func debugObjSize(objsize []memory.MemObjInfo) {
	for _, moi := range objsize {
		fmt.Println(moi.VarSize)
		debugObjSize(moi.ObjSize)
	}
}
