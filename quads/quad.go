package quads

import (
	"encoding/gob"
	"fmt"
)

// Quad is the structure of the instruction code generated by the compiler
type Quad struct {
	Action QuadAction // Operation to realize
	Left   Element    // Left operand
	Right  Element    // Right operand
	Result Element    // Result operand
}

func init() {
	gob.Register(Quad{})
	gob.Register(quadElement{})
}

func (q Quad) String() string {
	return fmt.Sprintf(
		"Quad {%s %s %s %s}",
		q.Action,
		q.Left,
		q.Right,
		q.Result,
	)
}
