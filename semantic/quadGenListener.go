package semantic

import (
	"fmt"

	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/parser"
)

type QuadGenListener struct {
	*parser.BaseProyectoListener
	qm *QuadrupleManager
}

func NewListener() QuadGenListener {
	qm := NewQuadrupleManager()
	return QuadGenListener{qm: &qm}
}

func (l *QuadGenListener) EnterG_exp(c *parser.G_expContext) {
	if c.Relop() != nil {
		l.qm.PushOp(c.Relop().GetText())
	}
}

func (l *QuadGenListener) ExitG_exp(c *parser.G_expContext) {
	l.qm.GenerateQuad([]int{
		constants.OPGT,
		constants.OPLT,
		constants.OPEQ,
		constants.OPNEQ,
	})
}

func (l *QuadGenListener) EnterM_exp2(c *parser.M_exp2Context) {
	if c.ADD() != nil {
		l.qm.PushOp(c.ADD().GetText())
	} else if c.SUB() != nil {
		l.qm.PushOp(c.SUB().GetText())
	}
}
func (l *QuadGenListener) ExitM_exp2(c *parser.M_exp2Context) {
	l.qm.GenerateQuad([]int{
		constants.OPPLUS,
		constants.OPMINUS,
	})
}

func (l *QuadGenListener) EnterTerm2(c *parser.Term2Context) {
	if c.MUL() != nil {
		l.qm.PushOp(c.MUL().GetText())
		return
	}

	if c.DIV() != nil {
		l.qm.PushOp(c.DIV().GetText())
		return
	}
}

func (l *QuadGenListener) ExitTerm2(c *parser.Term2Context) {
	l.qm.GenerateQuad([]int{
		constants.OPMULT,
		constants.OPDIV,
	})
}

func (l *QuadGenListener) EnterFactor(c *parser.FactorContext) {
	fmt.Println("enter factor")
	if c.Vars() != nil {
		l.qm.PushOperand(c.Vars().GetText())
	}
}
