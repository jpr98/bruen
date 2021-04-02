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

// EnterVars is called when production vars is entered.
func (s *BaseProyectoListener) EnterVars(ctx *VarsContext) {}

// ExitVars is called when production vars is exited.
func (s *BaseProyectoListener) ExitVars(ctx *VarsContext) {}

// EnterVarsTypeInit is called when production varsTypeInit is entered.
func (s *BaseProyectoListener) EnterVarsTypeInit(ctx *VarsTypeInitContext) {}

// ExitVarsTypeInit is called when production varsTypeInit is exited.
func (s *BaseProyectoListener) ExitVarsTypeInit(ctx *VarsTypeInitContext) {}

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

// EnterRead is called when production read is entered.
func (s *BaseProyectoListener) EnterRead(ctx *ReadContext) {}

// ExitRead is called when production read is exited.
func (s *BaseProyectoListener) ExitRead(ctx *ReadContext) {}

// EnterIds is called when production ids is entered.
func (s *BaseProyectoListener) EnterIds(ctx *IdsContext) {}

// ExitIds is called when production ids is exited.
func (s *BaseProyectoListener) ExitIds(ctx *IdsContext) {}

// EnterId is called when production id is entered.
func (s *BaseProyectoListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseProyectoListener) ExitId(ctx *IdContext) {}

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

// EnterExpression is called when production expression is entered.
func (s *BaseProyectoListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseProyectoListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseProyectoListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseProyectoListener) ExitExp(ctx *ExpContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseProyectoListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseProyectoListener) ExitTerm(ctx *TermContext) {}

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
