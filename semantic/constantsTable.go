package semantic

import (
	"encoding/gob"

	"github.com/jpr98/compis/constants"
)

type constantsTableContent struct {
	TypeOf constants.Type // Type of constant
	Dir    int            // Virtual memory address assigned to constant
}

func init() {
	gob.Register(constantsTableContent{})
}

// ConstantsTable holds all constants in a program.
var ConstantsTable map[string]*constantsTableContent = make(map[string]*constantsTableContent)
