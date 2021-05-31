package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/parser"
	"github.com/jpr98/compis/quads"
	"github.com/jpr98/compis/semantic"
	"github.com/jpr98/compis/virtualMachine"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("No args")
	}

	if os.Args[2] == "b" {
		compile()
	} else {
		execute()
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
	lexer := parser.NewProyectoLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Creates the parser
	p := parser.NewProyectoParser(stream)

	listener := semantic.NewListener()
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Program())
	ct := listener.GetClassTable()
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
		}
	}

	fmt.Println("\n---------")
	stream.Seek(0)
	p = parser.NewProyectoParser(stream)

	var quadListener quads.QuadGenListener = quads.NewListener(listener.GetFunctionTable(), ct)
	antlr.ParseTreeWalkerDefault.Walk(&quadListener, p.Program())
	quads := quadListener.GetManager().GetQuads()
	for i, q := range quads {
		fmt.Printf("%d. %s\n", i, q)
	}
	fmt.Println("\n---------")
	debugFT(listener.GetFunctionTable())

	f, err := os.Create("out.obj")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	enc.Encode(listener.ProgramName)
	enc.Encode(quads)
	enc.Encode(listener.GetFunctionTable())
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
	fmt.Println("\n---------")
	for i, q := range m {
		fmt.Printf("%d. %s\n", i, q)
	}
	fmt.Println("\n---------")

	var ft semantic.FunctionTable
	err = dec.Decode(&ft)
	if err != nil {
		panic(err)
	}
	//debugFT(ft)

	err = dec.Decode(&semantic.ConstantsTable)
	if err != nil {
		panic(err)
	}

	vm := virtualMachine.NewVM(programName, ft, m)
	vm.Run()
}

func testSC() {
	sc := semantic.NewCube(nil)
	fmt.Println(sc.ValidateBinaryOperation(constants.TYPEBOOL, constants.TYPEBOOL, constants.OPAND))
}

func debugFT(ft semantic.FunctionTable) {
	for fname, function := range ft {
		fmt.Printf("\n *Function: %s Returns: %s Scope: %s, VarSize: %v, TempSize: %v\n", fname, function.TypeOf, function.Scope, function.VarsSize, function.TempSize)
		fmt.Println("VarTable:")
		for id, variable := range function.Vars {
			fmt.Printf("%s: %s \n", id, variable.TypeOf)
		}
		fmt.Printf("Dir: %d\n", function.Dir)
	}
}
