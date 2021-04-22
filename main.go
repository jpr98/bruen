package main

import (
	"fmt"
	"log"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/parser"
)

func main() {
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

	// Read all tokens
	// for {
	// 	t := lexer.NextToken()
	// 	if t.GetTokenType() == antlr.TokenEOF {
	// 		break
	// 	}
	// 	fmt.Printf("%s (%q)\n",
	// 		lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	// }
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Creates the parser
	p := parser.NewProyectoParser(stream)

	var listener MyListener = NewListener()
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Program())

	testSC()
}

func testSC() {
	sc := NewSemanticCube()
	fmt.Println(sc.cube[constants.TYPEINT][constants.TYPEFLOAT][constants.OPPLUS])
}
