package semantic

import (
	"fmt"
	"os"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// ErrorListener is an implementation of antlr.ErrorListener. It is used to report
// syntax errors.
type ErrorListener struct {
	*antlr.DefaultErrorListener
	FoundError bool // FoundError is raised if an error is found
}

// NewErrorListener creates a new ErrorListener with a lowered error flag
func NewErrorListener() ErrorListener {
	return ErrorListener{FoundError: false}
}

// SyntaxError is called when the parser encounters an error in the source code,
// FoundError is set to true and an error message is printed to Stderr
func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.FoundError = true
	fmt.Fprintln(os.Stderr, "line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg)
}
