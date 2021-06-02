package semantic

import (
	"fmt"
	"os"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
	FoundError bool
}

func NewErrorListener() ErrorListener {
	return ErrorListener{FoundError: false}
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.FoundError = true
	fmt.Fprintln(os.Stderr, "line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg)
}
