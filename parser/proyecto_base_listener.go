// Code generated from Proyecto.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Proyecto

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseProyectoListener is a complete listener for a parse tree produced by ProyectoParser.
type BaseProyectoListener struct{}

var _ ProyectoListener = &BaseProyectoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProyectoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProyectoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProyectoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProyectoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseProyectoListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseProyectoListener) ExitProgram(ctx *ProgramContext) {}

// EnterClassDef is called when production classDef is entered.
func (s *BaseProyectoListener) EnterClassDef(ctx *ClassDefContext) {}

// ExitClassDef is called when production classDef is exited.
func (s *BaseProyectoListener) ExitClassDef(ctx *ClassDefContext) {}

// EnterClassBlock is called when production classBlock is entered.
func (s *BaseProyectoListener) EnterClassBlock(ctx *ClassBlockContext) {}

// ExitClassBlock is called when production classBlock is exited.
func (s *BaseProyectoListener) ExitClassBlock(ctx *ClassBlockContext) {}

// EnterVarsDec is called when production varsDec is entered.
func (s *BaseProyectoListener) EnterVarsDec(ctx *VarsDecContext) {}

// ExitVarsDec is called when production varsDec is exited.
func (s *BaseProyectoListener) ExitVarsDec(ctx *VarsDecContext) {}

// EnterVarsTypeInit is called when production varsTypeInit is entered.
func (s *BaseProyectoListener) EnterVarsTypeInit(ctx *VarsTypeInitContext) {}

// ExitVarsTypeInit is called when production varsTypeInit is exited.
func (s *BaseProyectoListener) ExitVarsTypeInit(ctx *VarsTypeInitContext) {}

// EnterVars is called when production vars is entered.
func (s *BaseProyectoListener) EnterVars(ctx *VarsContext) {}

// ExitVars is called when production vars is exited.
func (s *BaseProyectoListener) ExitVars(ctx *VarsContext) {}

// EnterFunctions is called when production functions is entered.
func (s *BaseProyectoListener) EnterFunctions(ctx *FunctionsContext) {}

// ExitFunctions is called when production functions is exited.
func (s *BaseProyectoListener) ExitFunctions(ctx *FunctionsContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseProyectoListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseProyectoListener) ExitParameters(ctx *ParametersContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseProyectoListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseProyectoListener) ExitParameter(ctx *ParameterContext) {}

// EnterFunctionBlock is called when production functionBlock is entered.
func (s *BaseProyectoListener) EnterFunctionBlock(ctx *FunctionBlockContext) {}

// ExitFunctionBlock is called when production functionBlock is exited.
func (s *BaseProyectoListener) ExitFunctionBlock(ctx *FunctionBlockContext) {}

// EnterReturnRule is called when production returnRule is entered.
func (s *BaseProyectoListener) EnterReturnRule(ctx *ReturnRuleContext) {}

// ExitReturnRule is called when production returnRule is exited.
func (s *BaseProyectoListener) ExitReturnRule(ctx *ReturnRuleContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseProyectoListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseProyectoListener) ExitBlock(ctx *BlockContext) {}

// EnterStatutes is called when production statutes is entered.
func (s *BaseProyectoListener) EnterStatutes(ctx *StatutesContext) {}

// ExitStatutes is called when production statutes is exited.
func (s *BaseProyectoListener) ExitStatutes(ctx *StatutesContext) {}

// EnterAssignation is called when production assignation is entered.
func (s *BaseProyectoListener) EnterAssignation(ctx *AssignationContext) {}

// ExitAssignation is called when production assignation is exited.
func (s *BaseProyectoListener) ExitAssignation(ctx *AssignationContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseProyectoListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseProyectoListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseProyectoListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseProyectoListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseProyectoListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseProyectoListener) ExitArgument(ctx *ArgumentContext) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *BaseProyectoListener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BaseProyectoListener) ExitMethodCall(ctx *MethodCallContext) {}

// EnterCall is called when production call is entered.
func (s *BaseProyectoListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseProyectoListener) ExitCall(ctx *CallContext) {}

// EnterRead is called when production read is entered.
func (s *BaseProyectoListener) EnterRead(ctx *ReadContext) {}

// ExitRead is called when production read is exited.
func (s *BaseProyectoListener) ExitRead(ctx *ReadContext) {}

// EnterWrite is called when production write is entered.
func (s *BaseProyectoListener) EnterWrite(ctx *WriteContext) {}

// ExitWrite is called when production write is exited.
func (s *BaseProyectoListener) ExitWrite(ctx *WriteContext) {}

// EnterConditional is called when production conditional is entered.
func (s *BaseProyectoListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BaseProyectoListener) ExitConditional(ctx *ConditionalContext) {}

// EnterForLoop is called when production forLoop is entered.
func (s *BaseProyectoListener) EnterForLoop(ctx *ForLoopContext) {}

// ExitForLoop is called when production forLoop is exited.
func (s *BaseProyectoListener) ExitForLoop(ctx *ForLoopContext) {}

// EnterWhileLoop is called when production whileLoop is entered.
func (s *BaseProyectoListener) EnterWhileLoop(ctx *WhileLoopContext) {}

// ExitWhileLoop is called when production whileLoop is exited.
func (s *BaseProyectoListener) ExitWhileLoop(ctx *WhileLoopContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseProyectoListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseProyectoListener) ExitExp(ctx *ExpContext) {}

// EnterExp2 is called when production exp2 is entered.
func (s *BaseProyectoListener) EnterExp2(ctx *Exp2Context) {}

// ExitExp2 is called when production exp2 is exited.
func (s *BaseProyectoListener) ExitExp2(ctx *Exp2Context) {}

// EnterT_exp is called when production t_exp is entered.
func (s *BaseProyectoListener) EnterT_exp(ctx *T_expContext) {}

// ExitT_exp is called when production t_exp is exited.
func (s *BaseProyectoListener) ExitT_exp(ctx *T_expContext) {}

// EnterT_exp2 is called when production t_exp2 is entered.
func (s *BaseProyectoListener) EnterT_exp2(ctx *T_exp2Context) {}

// ExitT_exp2 is called when production t_exp2 is exited.
func (s *BaseProyectoListener) ExitT_exp2(ctx *T_exp2Context) {}

// EnterG_exp is called when production g_exp is entered.
func (s *BaseProyectoListener) EnterG_exp(ctx *G_expContext) {}

// ExitG_exp is called when production g_exp is exited.
func (s *BaseProyectoListener) ExitG_exp(ctx *G_expContext) {}

// EnterG_exp2 is called when production g_exp2 is entered.
func (s *BaseProyectoListener) EnterG_exp2(ctx *G_exp2Context) {}

// ExitG_exp2 is called when production g_exp2 is exited.
func (s *BaseProyectoListener) ExitG_exp2(ctx *G_exp2Context) {}

// EnterM_exp is called when production m_exp is entered.
func (s *BaseProyectoListener) EnterM_exp(ctx *M_expContext) {}

// ExitM_exp is called when production m_exp is exited.
func (s *BaseProyectoListener) ExitM_exp(ctx *M_expContext) {}

// EnterM_exp2 is called when production m_exp2 is entered.
func (s *BaseProyectoListener) EnterM_exp2(ctx *M_exp2Context) {}

// ExitM_exp2 is called when production m_exp2 is exited.
func (s *BaseProyectoListener) ExitM_exp2(ctx *M_exp2Context) {}

// EnterTerm is called when production term is entered.
func (s *BaseProyectoListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseProyectoListener) ExitTerm(ctx *TermContext) {}

// EnterTerm2 is called when production term2 is entered.
func (s *BaseProyectoListener) EnterTerm2(ctx *Term2Context) {}

// ExitTerm2 is called when production term2 is exited.
func (s *BaseProyectoListener) ExitTerm2(ctx *Term2Context) {}

// EnterFactor is called when production factor is entered.
func (s *BaseProyectoListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseProyectoListener) ExitFactor(ctx *FactorContext) {}

// EnterVarCte is called when production varCte is entered.
func (s *BaseProyectoListener) EnterVarCte(ctx *VarCteContext) {}

// ExitVarCte is called when production varCte is exited.
func (s *BaseProyectoListener) ExitVarCte(ctx *VarCteContext) {}

// EnterCte_i is called when production cte_i is entered.
func (s *BaseProyectoListener) EnterCte_i(ctx *Cte_iContext) {}

// ExitCte_i is called when production cte_i is exited.
func (s *BaseProyectoListener) ExitCte_i(ctx *Cte_iContext) {}

// EnterCte_f is called when production cte_f is entered.
func (s *BaseProyectoListener) EnterCte_f(ctx *Cte_fContext) {}

// ExitCte_f is called when production cte_f is exited.
func (s *BaseProyectoListener) ExitCte_f(ctx *Cte_fContext) {}

// EnterCte_c is called when production cte_c is entered.
func (s *BaseProyectoListener) EnterCte_c(ctx *Cte_cContext) {}

// ExitCte_c is called when production cte_c is exited.
func (s *BaseProyectoListener) ExitCte_c(ctx *Cte_cContext) {}

// EnterCte_b is called when production cte_b is entered.
func (s *BaseProyectoListener) EnterCte_b(ctx *Cte_bContext) {}

// ExitCte_b is called when production cte_b is exited.
func (s *BaseProyectoListener) ExitCte_b(ctx *Cte_bContext) {}

// EnterCte_s is called when production cte_s is entered.
func (s *BaseProyectoListener) EnterCte_s(ctx *Cte_sContext) {}

// ExitCte_s is called when production cte_s is exited.
func (s *BaseProyectoListener) ExitCte_s(ctx *Cte_sContext) {}

// EnterMain is called when production main is entered.
func (s *BaseProyectoListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BaseProyectoListener) ExitMain(ctx *MainContext) {}

// EnterTypeRule is called when production typeRule is entered.
func (s *BaseProyectoListener) EnterTypeRule(ctx *TypeRuleContext) {}

// ExitTypeRule is called when production typeRule is exited.
func (s *BaseProyectoListener) ExitTypeRule(ctx *TypeRuleContext) {}

// EnterRelop is called when production relop is entered.
func (s *BaseProyectoListener) EnterRelop(ctx *RelopContext) {}

// ExitRelop is called when production relop is exited.
func (s *BaseProyectoListener) ExitRelop(ctx *RelopContext) {}
