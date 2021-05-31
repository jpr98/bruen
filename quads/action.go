package quads

// QuadAction represents an action in Quadruples
type QuadAction int

const (
	ADD QuadAction = iota
	SUB
	MUL
	DIV
	GT
	LT
	EQ
	NEQ
	AND
	OR
	LPAREN
	READ
	WRITE
	WRITENEWLINE
	ASSIGN
	VER
	CALCDIR
	GOTOF
	GOTO
	ENDFUNC
	ERA
	PARAM
	GOSUB
	RETURN
	INIT
	INITRETURN
)

func (a QuadAction) String() string {
	return [...]string{"+", "-", "*", "/", ">", "<", "==", "!=", "&&", "||", "(", "READ", "WRITE", "WRITENEWLINE", "=", "VER", "CALCDIR", "GOTOF", "GOTO", "ENDFUNC", "ERA", "PARAM", "GOSUB", "RETURN", "INIT", "INITRETURN"}[a]
}

type QuadActionStack struct {
	stack []QuadAction
}

// NewQuadActionStack creates a new stack
func NewQuadActionStack() QuadActionStack {
	return QuadActionStack{stack: make([]QuadAction, 0)}
}

// Push adds an action to the stack
func (qas *QuadActionStack) Push(e QuadAction) {
	qas.stack = append(qas.stack, e)
}

// Pop removes and returns the next action in the stack
func (qas *QuadActionStack) Pop() QuadAction {
	e := qas.stack[qas.Size()-1]
	qas.stack = qas.stack[:qas.Size()-1]
	return e
}

// Top returns the next action in the stack
func (qas *QuadActionStack) Top() QuadAction {
	return qas.stack[qas.Size()-1]
}

// Size returns the length of the stack
func (qas *QuadActionStack) Size() int {
	return len(qas.stack)
}
