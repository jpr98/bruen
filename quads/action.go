package quads

// QuadAction represents an action in Quadruples
type QuadAction int

const (
	ADD          QuadAction = iota // Addition
	SUB                            // Substraction
	MUL                            // Multiplication
	DIV                            // Division
	GT                             // Greater than
	LT                             // Less then
	EQ                             // Equal
	NEQ                            // Not equal
	AND                            // And
	OR                             // Or
	LPAREN                         // Left parentheses, used as false bottom in code generation
	READ                           // Read from stdin
	WRITE                          // Write to stdout
	WRITENEWLINE                   // Write new line to stdout
	ASSIGN                         // Assignation
	VER                            // Verify index in address
	CALCDIR                        // Calculate indexed array address
	GOTOF                          // Goto on false
	GOTO                           // Goto
	ENDFUNC                        // Function ends
	ERA                            // Activation record expansion
	PARAM                          // Parameter
	GOSUB                          // Go to function
	RETURN                         // Return
	INSTANCE                       // Capture instance of object
)

func (a QuadAction) String() string {
	return [...]string{"+", "-", "*", "/", ">", "<", "==", "!=", "&&", "||", "(", "READ", "WRITE", "WRITENEWLINE", "=", "VER", "CALCDIR", "GOTOF", "GOTO", "ENDFUNC", "ERA", "PARAM", "GOSUB", "RETURN", "INSTANCE"}[a]
}

// QuadActionStack holds the parsed actions
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
