// Code generated from Proyecto.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Proyecto

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ProyectoListener is a complete listener for a parse tree produced by ProyectoParser.
type ProyectoListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterClassDef is called when entering the classDef production.
	EnterClassDef(c *ClassDefContext)

	// EnterClassBlock is called when entering the classBlock production.
	EnterClassBlock(c *ClassBlockContext)

	// EnterVars is called when entering the vars production.
	EnterVars(c *VarsContext)

	// EnterVarsTypeInit is called when entering the varsTypeInit production.
	EnterVarsTypeInit(c *VarsTypeInitContext)

	// EnterFunctions is called when entering the functions production.
	EnterFunctions(c *FunctionsContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterFunctionBlock is called when entering the functionBlock production.
	EnterFunctionBlock(c *FunctionBlockContext)

	// EnterReturnRule is called when entering the returnRule production.
	EnterReturnRule(c *ReturnRuleContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatutes is called when entering the statutes production.
	EnterStatutes(c *StatutesContext)

	// EnterAssignation is called when entering the assignation production.
	EnterAssignation(c *AssignationContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterMethodCall is called when entering the methodCall production.
	EnterMethodCall(c *MethodCallContext)

	// EnterRead is called when entering the read production.
	EnterRead(c *ReadContext)

	// EnterIds is called when entering the ids production.
	EnterIds(c *IdsContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterWrite is called when entering the write production.
	EnterWrite(c *WriteContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterForLoop is called when entering the forLoop production.
	EnterForLoop(c *ForLoopContext)

	// EnterWhileLoop is called when entering the whileLoop production.
	EnterWhileLoop(c *WhileLoopContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterVarCte is called when entering the varCte production.
	EnterVarCte(c *VarCteContext)

	// EnterCte_i is called when entering the cte_i production.
	EnterCte_i(c *Cte_iContext)

	// EnterCte_f is called when entering the cte_f production.
	EnterCte_f(c *Cte_fContext)

	// EnterCte_c is called when entering the cte_c production.
	EnterCte_c(c *Cte_cContext)

	// EnterCte_s is called when entering the cte_s production.
	EnterCte_s(c *Cte_sContext)

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// EnterTypeRule is called when entering the typeRule production.
	EnterTypeRule(c *TypeRuleContext)

	// EnterRelop is called when entering the relop production.
	EnterRelop(c *RelopContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitClassDef is called when exiting the classDef production.
	ExitClassDef(c *ClassDefContext)

	// ExitClassBlock is called when exiting the classBlock production.
	ExitClassBlock(c *ClassBlockContext)

	// ExitVars is called when exiting the vars production.
	ExitVars(c *VarsContext)

	// ExitVarsTypeInit is called when exiting the varsTypeInit production.
	ExitVarsTypeInit(c *VarsTypeInitContext)

	// ExitFunctions is called when exiting the functions production.
	ExitFunctions(c *FunctionsContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitFunctionBlock is called when exiting the functionBlock production.
	ExitFunctionBlock(c *FunctionBlockContext)

	// ExitReturnRule is called when exiting the returnRule production.
	ExitReturnRule(c *ReturnRuleContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatutes is called when exiting the statutes production.
	ExitStatutes(c *StatutesContext)

	// ExitAssignation is called when exiting the assignation production.
	ExitAssignation(c *AssignationContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitMethodCall is called when exiting the methodCall production.
	ExitMethodCall(c *MethodCallContext)

	// ExitRead is called when exiting the read production.
	ExitRead(c *ReadContext)

	// ExitIds is called when exiting the ids production.
	ExitIds(c *IdsContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitWrite is called when exiting the write production.
	ExitWrite(c *WriteContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitForLoop is called when exiting the forLoop production.
	ExitForLoop(c *ForLoopContext)

	// ExitWhileLoop is called when exiting the whileLoop production.
	ExitWhileLoop(c *WhileLoopContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitVarCte is called when exiting the varCte production.
	ExitVarCte(c *VarCteContext)

	// ExitCte_i is called when exiting the cte_i production.
	ExitCte_i(c *Cte_iContext)

	// ExitCte_f is called when exiting the cte_f production.
	ExitCte_f(c *Cte_fContext)

	// ExitCte_c is called when exiting the cte_c production.
	ExitCte_c(c *Cte_cContext)

	// ExitCte_s is called when exiting the cte_s production.
	ExitCte_s(c *Cte_sContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitTypeRule is called when exiting the typeRule production.
	ExitTypeRule(c *TypeRuleContext)

	// ExitRelop is called when exiting the relop production.
	ExitRelop(c *RelopContext)
}
