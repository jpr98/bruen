package semantic

import "github.com/jpr98/compis/constants"

type constantsTableContent struct {
	TypeOf constants.Type
	Dir    int
}

var ConstantsTable map[string]*constantsTableContent = make(map[string]*constantsTableContent)
