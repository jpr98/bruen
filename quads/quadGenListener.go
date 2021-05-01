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

func (l QuadGenListener) GetManager() Manager {
	return *l.m
}

func (l *QuadGenListener) EnterAssignation(c *parser.AssignationContext) {
	l.m.PushOperand(c.ID().GetText())
	l.m.PushOp(c.ASSIGN().GetText())
}

func (l *QuadGenListener) ExitAssignation(c *parser.AssignationContext) {
	l.m.GenerateAssignationQuad(false)
}

func (l *QuadGenListener) EnterG_exp2(c *parser.G_exp2Context) {
	if c.Relop() != nil {
		l.m.PushOp(c.Relop().GetText())
	}
}

func (l *QuadGenListener) ExitG_exp2(c *parser.G_exp2Context) {
	l.m.GenerateQuad([]int{
		constants.OPGT,
		constants.OPLT,
		constants.OPEQ,
		constants.OPNEQ,
	}, false)
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
	}, false)
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
	}, false)
}

func (l *QuadGenListener) EnterFactor(c *parser.FactorContext) {
	if c.Vars() != nil {
		l.m.PushOperand(c.Vars().GetText())
	} else if c.VarCte() != nil {
		l.m.PushOperand(c.VarCte().GetText())
	}
}

func (l *QuadGenListener) ExitConditional(c *parser.ConditionalContext) {
	// Neuralgic point 2
	l.m.UpdateGoto()
}

func (l *QuadGenListener) ExitConditional2(c *parser.Conditional2Context) {
	// Neuralgic point 1
	l.m.AddGotoF()
}

func (l *QuadGenListener) EnterConditional4(c *parser.Conditional4Context) {
	// Neuralgic point 3
	l.m.AddAndUpdateGoto()
}

func (l *QuadGenListener) ExitWhileLoop(c *parser.WhileLoopContext) {
	l.m.AddAndUpdateWhileGotos()
}

func (l *QuadGenListener) EnterWhileLoop2(c *parser.WhileLoop2Context) {
	l.m.SaveJumpPosition()
}

func (l *QuadGenListener) ExitWhileLoop2(c *parser.WhileLoop2Context) {
	l.m.AddGotoF()
}

func (l *QuadGenListener) ExitForLoop(c *parser.ForLoopContext) {
	// TODO: change this string to constant(1) address
	l.m.PushOperand("1")
	l.m.PushOp("+")
	l.m.GenerateQuad([]int{constants.OPPLUS}, true)

	l.m.PushOp("=")
	l.m.GenerateAssignationQuad(true)

	l.m.AddAndUpdateWhileGotos()
}

func (l *QuadGenListener) EnterForLoop2(c *parser.ForLoop2Context) {
	l.m.PushOperand(c.ID().GetText())
	l.m.PushOp(c.ASSIGN().GetText())
}
func (l *QuadGenListener) ExitForLoop2(c *parser.ForLoop2Context) {
	l.m.GenerateAssignationQuad(true)
	l.m.SaveJumpPosition()
}

func (l *QuadGenListener) ExitForLoop3(c *parser.ForLoop3Context) {
	l.m.PushOp("<")
	l.m.GenerateQuad([]int{constants.OPLT}, true)
	l.m.AddGotoF()
}
