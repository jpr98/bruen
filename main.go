package main

import (
	"fmt"
	"log"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/parser"
	"github.com/jpr98/compis/quads"
	"github.com/jpr98/compis/semantic"
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

	// var listener MyListener = NewListener()
	// antlr.ParseTreeWalkerDefault.Walk(&listener, p.Program())
	var listener quads.QuadGenListener = quads.NewListener()
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Program())
}

func testSC() {
	sc := semantic.NewCube(nil)
	fmt.Println(sc.ValidateBinaryOperation(constants.TYPEBOOL, constants.TYPEBOOL, constants.OPAND))
}
