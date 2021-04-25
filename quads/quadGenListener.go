package quads

import (
	"github.com/jpr98/compis/constants"
	"github.com/jpr98/compis/parser"
)

type QuadGenListener struct {
	*parser.BaseProyectoListener
	m *Manager
}

func NewListener() QuadGenListener {
	m := NewManager()
	return QuadGenListener{m: &m}
}

func (l *QuadGenListener) EnterG_exp(c *parser.G_expContext) {
	if c.Relop() != nil {
		l.m.PushOp(c.Relop().GetText())
	}
}

func (l *QuadGenListener) ExitG_exp(c *parser.G_expContext) {
	l.m.GenerateQuad([]int{
		constants.OPGT,
		constants.OPLT,
		constants.OPEQ,
		constants.OPNEQ,
	})
}

func (l *QuadGenListener) EnterM_exp2(c *parser.M_exp2Context) {
	if c.ADD() != nil {
		l.m.PushOp(c.ADD().GetText())
	} else if c.SUB() != nil {
		l.m.PushOp(c.SUB().GetText())
	}
}
func (l *QuadGenListener) ExitM_exp2(c *parser.M_exp2Context) {
	l.m.GenerateQuad([]int{
		constants.OPPLUS,
		constants.OPMINUS,
	})
}

func (l *QuadGenListener) EnterTerm2(c *parser.Term2Context) {
	if c.MUL() != nil {
		l.m.PushOp(c.MUL().GetText())
		return
	}

	if c.DIV() != nil {
		l.m.PushOp(c.DIV().GetText())
		return
	}
}

func (l *QuadGenListener) ExitTerm2(c *parser.Term2Context) {
	l.m.GenerateQuad([]int{
		constants.OPMULT,
		constants.OPDIV,
	})
}

func (l *QuadGenListener) EnterFactor(c *parser.FactorContext) {
	if c.Vars() != nil {
		l.m.PushOperand(c.Vars().GetText())
	}
}
