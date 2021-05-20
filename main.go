package main

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/quads"
	"github.com/jpr98/compis/semantic"
)

func main() {
	testReadObjFile()
	// if len(os.Args) < 2 {
	// 	log.Fatal("No filename provided")
	// }

	// // Open file stream
	// filename := os.Args[1]
	// is, err := antlr.NewFileStream(filename)
	// if err != nil {
	// 	log.Fatalf("Error: %s", err.Error())
	// }

	// // Creates the lexer
	// lexer := parser.NewProyectoLexer(is)

	// // Read all tokens
	// // for {
	// // 	t := lexer.NextToken()
	// // 	if t.GetTokenType() == antlr.TokenEOF {
	// // 		break
	// // 	}
	// // 	fmt.Printf("%s (%q)\n",
	// // 		lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	// // }
	// stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// // Creates the parser
	// p := parser.NewProyectoParser(stream)

	// listener := semantic.NewListener()
	// antlr.ParseTreeWalkerDefault.Walk(&listener, p.Program())

	// fmt.Println("\n---------")
	// stream.Seek(0)
	// p = parser.NewProyectoParser(stream)

	// var quadListener quads.QuadGenListener = quads.NewListener(listener.GetFunctionTable())
	// antlr.ParseTreeWalkerDefault.Walk(&quadListener, p.Program())
	// quads := quadListener.GetManager().GetQuads()
	// for i, q := range quads {
	// 	fmt.Printf("%d. %s\n", i, q)
	// }
	// fmt.Println("\n---------")
	// //debugFT(listener.GetFunctionTable())

	// f, err := os.Create("out.obj")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// enc := gob.NewEncoder(f)
	// enc.Encode(quads)
	// enc.Encode(listener.GetFunctionTable())
}

func testReadObjFile() {
	f, err := os.Open("out.obj")
	if err != nil {
		panic(err)
	}

	var m []quads.Quad
	dec := gob.NewDecoder(f)
	err = dec.Decode(&m)
	if err != nil {
		panic(err)
	}
	for i, q := range m {
		fmt.Printf("%d. %s\n", i, q)
	}

	var ft semantic.FunctionTable
	err = dec.Decode(&ft)
	if err != nil {
		panic(err)
	}
	debugFT(ft)
}

func testSC() {
	sc := semantic.NewCube(nil)
	fmt.Println(sc.ValidateBinaryOperation(constants.TYPEBOOL, constants.TYPEBOOL, constants.OPAND))
}

func debugFT(ft semantic.FunctionTable) {
	for fname, function := range ft {
		fmt.Printf("\n *Function: %s Returns: %s Scope: %s\n", fname, function.TypeOf, function.Scope)
		fmt.Println("VarTable:")
		for id, variable := range function.Vars {
			fmt.Printf("%s: %s \n", id, variable.TypeOf)
		}
		fmt.Printf("Dir: %d\n", function.Dir)
	}
}
