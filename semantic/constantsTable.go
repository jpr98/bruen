package semantic

import (
	"encoding/gob"

	"github.com/jpr98/compis/constants"
)

type constantsTableContent struct {
	TypeOf constants.Type
	Dir    int
}

func init() {
	gob.Register(constantsTableContent{})
}

var ConstantsTable map[string]*constantsTableContent = make(map[string]*constantsTableContent)
