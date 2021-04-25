package semantic

import (
	"log"

	"github.com/jpr98/compis/constants"
)

// SCube contains the validity table for type operations
type SCube interface {
	ValidateBinaryOperation(constants.Type, constants.Type, int) constants.Type
}

// Cube is an instance of the semantic cube
var Cube SCube = NewCube(nil)

type cube struct {
	cube   [][][]constants.Type
	logger *log.Logger
}

// NewCube creates a new Cube
func NewCube(logger *log.Logger) SCube {
	sc := cube{logger: logger}
	sc.cube = make([][][]constants.Type, 4)
	for i := 0; i < 4; i++ {
		sc.cube[i] = make([][]constants.Type, 4)
		for j := 0; j < 4; j++ {
			sc.cube[i][j] = make([]constants.Type, 10)
		}
	}

	// Int
	sc.cube[constants.TYPEINT][constants.TYPEINT][constants.OPPLUS] = constants.TYPEINT
	sc.cube[constants.TYPEINT][constants.TYPEINT][constants.OPMINUS] = constants.TYPEINT
	sc.cube[constants.TYPEINT][constants.TYPEINT][constants.OPDIV] = constants.TYPEINT
	sc.cube[constants.TYPEINT][constants.TYPEINT][constants.OPMULT] = constants.TYPEINT
	sc.cube[constants.TYPEINT][constants.TYPEINT][constants.OPGT] = constants.TYPEBOOL
	sc.cube[constants.TYPEINT][constants.TYPEINT][constants.OPLT] = constants.TYPEBOOL
	sc.cube[constants.TYPEINT][constants.TYPEINT][constants.OPEQ] = constants.TYPEBOOL
	sc.cube[constants.TYPEINT][constants.TYPEINT][constants.OPNEQ] = constants.TYPEBOOL
	sc.cube[constants.TYPEINT][constants.TYPEINT][constants.OPAND] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPEINT][constants.OPOR] = constants.ERR

	sc.cube[constants.TYPEINT][constants.TYPEFLOAT][constants.OPPLUS] = constants.TYPEFLOAT
	sc.cube[constants.TYPEINT][constants.TYPEFLOAT][constants.OPMINUS] = constants.TYPEFLOAT
	sc.cube[constants.TYPEINT][constants.TYPEFLOAT][constants.OPDIV] = constants.TYPEFLOAT
	sc.cube[constants.TYPEINT][constants.TYPEFLOAT][constants.OPMULT] = constants.TYPEFLOAT
	sc.cube[constants.TYPEINT][constants.TYPEFLOAT][constants.OPGT] = constants.TYPEBOOL
	sc.cube[constants.TYPEINT][constants.TYPEFLOAT][constants.OPLT] = constants.TYPEBOOL
	sc.cube[constants.TYPEINT][constants.TYPEFLOAT][constants.OPEQ] = constants.TYPEBOOL
	sc.cube[constants.TYPEINT][constants.TYPEFLOAT][constants.OPNEQ] = constants.TYPEBOOL
	sc.cube[constants.TYPEINT][constants.TYPEFLOAT][constants.OPAND] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPEFLOAT][constants.OPOR] = constants.ERR

	sc.cube[constants.TYPEINT][constants.TYPECHAR][constants.OPPLUS] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPECHAR][constants.OPMINUS] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPECHAR][constants.OPDIV] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPECHAR][constants.OPMULT] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPECHAR][constants.OPGT] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPECHAR][constants.OPLT] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPECHAR][constants.OPEQ] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPECHAR][constants.OPNEQ] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPECHAR][constants.OPAND] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPECHAR][constants.OPOR] = constants.ERR

	sc.cube[constants.TYPEINT][constants.TYPEBOOL][constants.OPPLUS] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPEBOOL][constants.OPMINUS] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPEBOOL][constants.OPDIV] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPEBOOL][constants.OPMULT] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPEBOOL][constants.OPGT] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPEBOOL][constants.OPLT] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPEBOOL][constants.OPEQ] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPEBOOL][constants.OPNEQ] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPEBOOL][constants.OPAND] = constants.ERR
	sc.cube[constants.TYPEINT][constants.TYPEBOOL][constants.OPOR] = constants.ERR

	// Float
	sc.cube[constants.TYPEFLOAT][constants.TYPEINT][constants.OPPLUS] = constants.TYPEFLOAT
	sc.cube[constants.TYPEFLOAT][constants.TYPEINT][constants.OPMINUS] = constants.TYPEFLOAT
	sc.cube[constants.TYPEFLOAT][constants.TYPEINT][constants.OPDIV] = constants.TYPEFLOAT
	sc.cube[constants.TYPEFLOAT][constants.TYPEINT][constants.OPMULT] = constants.TYPEFLOAT
	sc.cube[constants.TYPEFLOAT][constants.TYPEINT][constants.OPGT] = constants.TYPEBOOL
	sc.cube[constants.TYPEFLOAT][constants.TYPEINT][constants.OPLT] = constants.TYPEBOOL
	sc.cube[constants.TYPEFLOAT][constants.TYPEINT][constants.OPEQ] = constants.TYPEBOOL
	sc.cube[constants.TYPEFLOAT][constants.TYPEINT][constants.OPNEQ] = constants.TYPEBOOL
	sc.cube[constants.TYPEFLOAT][constants.TYPEINT][constants.OPAND] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPEINT][constants.OPOR] = constants.ERR

	sc.cube[constants.TYPEFLOAT][constants.TYPEFLOAT][constants.OPPLUS] = constants.TYPEFLOAT
	sc.cube[constants.TYPEFLOAT][constants.TYPEFLOAT][constants.OPMINUS] = constants.TYPEFLOAT
	sc.cube[constants.TYPEFLOAT][constants.TYPEFLOAT][constants.OPDIV] = constants.TYPEFLOAT
	sc.cube[constants.TYPEFLOAT][constants.TYPEFLOAT][constants.OPMULT] = constants.TYPEFLOAT
	sc.cube[constants.TYPEFLOAT][constants.TYPEFLOAT][constants.OPGT] = constants.TYPEBOOL
	sc.cube[constants.TYPEFLOAT][constants.TYPEFLOAT][constants.OPLT] = constants.TYPEBOOL
	sc.cube[constants.TYPEFLOAT][constants.TYPEFLOAT][constants.OPEQ] = constants.TYPEBOOL
	sc.cube[constants.TYPEFLOAT][constants.TYPEFLOAT][constants.OPNEQ] = constants.TYPEBOOL
	sc.cube[constants.TYPEFLOAT][constants.TYPEFLOAT][constants.OPAND] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPEFLOAT][constants.OPOR] = constants.ERR

	sc.cube[constants.TYPEFLOAT][constants.TYPECHAR][constants.OPPLUS] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPECHAR][constants.OPMINUS] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPECHAR][constants.OPDIV] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPECHAR][constants.OPMULT] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPECHAR][constants.OPGT] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPECHAR][constants.OPLT] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPECHAR][constants.OPEQ] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPECHAR][constants.OPNEQ] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPECHAR][constants.OPAND] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPECHAR][constants.OPOR] = constants.ERR

	sc.cube[constants.TYPEFLOAT][constants.TYPEBOOL][constants.OPPLUS] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPEBOOL][constants.OPMINUS] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPEBOOL][constants.OPDIV] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPEBOOL][constants.OPMULT] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPEBOOL][constants.OPGT] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPEBOOL][constants.OPLT] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPEBOOL][constants.OPEQ] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPEBOOL][constants.OPNEQ] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPEBOOL][constants.OPAND] = constants.ERR
	sc.cube[constants.TYPEFLOAT][constants.TYPEBOOL][constants.OPOR] = constants.ERR

	// Char
	sc.cube[constants.TYPECHAR][constants.TYPEINT][constants.OPPLUS] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEINT][constants.OPMINUS] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEINT][constants.OPDIV] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEINT][constants.OPMULT] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEINT][constants.OPGT] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEINT][constants.OPLT] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEINT][constants.OPEQ] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEINT][constants.OPNEQ] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEINT][constants.OPAND] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEINT][constants.OPOR] = constants.ERR

	sc.cube[constants.TYPECHAR][constants.TYPEFLOAT][constants.OPPLUS] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEFLOAT][constants.OPMINUS] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEFLOAT][constants.OPDIV] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEFLOAT][constants.OPMULT] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEFLOAT][constants.OPGT] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEFLOAT][constants.OPLT] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEFLOAT][constants.OPEQ] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEFLOAT][constants.OPNEQ] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEFLOAT][constants.OPAND] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEFLOAT][constants.OPOR] = constants.ERR

	sc.cube[constants.TYPECHAR][constants.TYPECHAR][constants.OPPLUS] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPECHAR][constants.OPMINUS] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPECHAR][constants.OPDIV] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPECHAR][constants.OPMULT] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPECHAR][constants.OPGT] = constants.TYPEBOOL
	sc.cube[constants.TYPECHAR][constants.TYPECHAR][constants.OPLT] = constants.TYPEBOOL
	sc.cube[constants.TYPECHAR][constants.TYPECHAR][constants.OPEQ] = constants.TYPEBOOL
	sc.cube[constants.TYPECHAR][constants.TYPECHAR][constants.OPNEQ] = constants.TYPEBOOL
	sc.cube[constants.TYPECHAR][constants.TYPECHAR][constants.OPAND] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPECHAR][constants.OPOR] = constants.ERR

	sc.cube[constants.TYPECHAR][constants.TYPEBOOL][constants.OPPLUS] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEBOOL][constants.OPMINUS] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEBOOL][constants.OPDIV] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEBOOL][constants.OPMULT] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEBOOL][constants.OPGT] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEBOOL][constants.OPLT] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEBOOL][constants.OPEQ] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEBOOL][constants.OPNEQ] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEBOOL][constants.OPAND] = constants.ERR
	sc.cube[constants.TYPECHAR][constants.TYPEBOOL][constants.OPOR] = constants.ERR

	// Bool
	sc.cube[constants.TYPEBOOL][constants.TYPEINT][constants.OPPLUS] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEINT][constants.OPMINUS] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEINT][constants.OPDIV] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEINT][constants.OPMULT] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEINT][constants.OPGT] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEINT][constants.OPLT] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEINT][constants.OPEQ] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEINT][constants.OPNEQ] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEINT][constants.OPAND] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEINT][constants.OPOR] = constants.ERR

	sc.cube[constants.TYPEBOOL][constants.TYPEFLOAT][constants.OPPLUS] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEFLOAT][constants.OPMINUS] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEFLOAT][constants.OPDIV] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEFLOAT][constants.OPMULT] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEFLOAT][constants.OPGT] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEFLOAT][constants.OPLT] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEFLOAT][constants.OPEQ] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEFLOAT][constants.OPNEQ] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEFLOAT][constants.OPAND] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEFLOAT][constants.OPOR] = constants.ERR

	sc.cube[constants.TYPEBOOL][constants.TYPECHAR][constants.OPPLUS] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPECHAR][constants.OPMINUS] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPECHAR][constants.OPDIV] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPECHAR][constants.OPMULT] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPECHAR][constants.OPGT] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPECHAR][constants.OPLT] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPECHAR][constants.OPEQ] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPECHAR][constants.OPNEQ] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPECHAR][constants.OPAND] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPECHAR][constants.OPOR] = constants.ERR

	sc.cube[constants.TYPEBOOL][constants.TYPEBOOL][constants.OPPLUS] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEBOOL][constants.OPMINUS] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEBOOL][constants.OPDIV] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEBOOL][constants.OPMULT] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEBOOL][constants.OPGT] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEBOOL][constants.OPLT] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEBOOL][constants.OPEQ] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEBOOL][constants.OPNEQ] = constants.ERR
	sc.cube[constants.TYPEBOOL][constants.TYPEBOOL][constants.OPAND] = constants.TYPEBOOL
	sc.cube[constants.TYPEBOOL][constants.TYPEBOOL][constants.OPOR] = constants.TYPEBOOL

	return sc
}

// ValidateBinaryOperation gets the type of a given pair of types applying a given operator
func (sc cube) ValidateBinaryOperation(lType, rType constants.Type, op int) constants.Type {
	if int(lType) > len(sc.cube) || lType < 0 {
		if sc.logger != nil {
			sc.logger.Printf("Error: (ValidateBinaryOperation) accesing binary semantic cube, lType out of bounds: lType value (%d)", lType)
		}
		return constants.ERR
	}
	if int(rType) > len(sc.cube[lType]) || rType < 0 {
		if sc.logger != nil {
			sc.logger.Printf("Error: (ValidateBinaryOperation) accesing binary semantic cube, rType out of bounds: rType value (%d)", rType)
		}
		return constants.ERR
	}
	if int(op) > len(sc.cube[lType][rType]) || op < 0 {
		if sc.logger != nil {
			sc.logger.Printf("Error: (ValidateBinaryOperation) accesing binary semantic cube, op out of bounds: op value (%d)", op)
		}
		return constants.ERR
	}
	return sc.cube[lType][rType][op]
}
