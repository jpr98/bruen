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

	// EnterVarsDec is called when entering the varsDec production.
	EnterVarsDec(c *VarsDecContext)

	// EnterVarsTypeInit is called when entering the varsTypeInit production.
	EnterVarsTypeInit(c *VarsTypeInitContext)

	// EnterVars is called when entering the vars production.
	EnterVars(c *VarsContext)

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

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterRead is called when entering the read production.
	EnterRead(c *ReadContext)

	// EnterWrite is called when entering the write production.
	EnterWrite(c *WriteContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterForLoop is called when entering the forLoop production.
	EnterForLoop(c *ForLoopContext)

	// EnterWhileLoop is called when entering the whileLoop production.
	EnterWhileLoop(c *WhileLoopContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterExp2 is called when entering the exp2 production.
	EnterExp2(c *Exp2Context)

	// EnterT_exp is called when entering the t_exp production.
	EnterT_exp(c *T_expContext)

	// EnterT_exp2 is called when entering the t_exp2 production.
	EnterT_exp2(c *T_exp2Context)

	// EnterG_exp is called when entering the g_exp production.
	EnterG_exp(c *G_expContext)

	// EnterM_exp is called when entering the m_exp production.
	EnterM_exp(c *M_expContext)

	// EnterM_exp2 is called when entering the m_exp2 production.
	EnterM_exp2(c *M_exp2Context)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterTerm2 is called when entering the term2 production.
	EnterTerm2(c *Term2Context)

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

	// EnterCte_b is called when entering the cte_b production.
	EnterCte_b(c *Cte_bContext)

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

	// ExitVarsDec is called when exiting the varsDec production.
	ExitVarsDec(c *VarsDecContext)

	// ExitVarsTypeInit is called when exiting the varsTypeInit production.
	ExitVarsTypeInit(c *VarsTypeInitContext)

	// ExitVars is called when exiting the vars production.
	ExitVars(c *VarsContext)

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

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitRead is called when exiting the read production.
	ExitRead(c *ReadContext)

	// ExitWrite is called when exiting the write production.
	ExitWrite(c *WriteContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitForLoop is called when exiting the forLoop production.
	ExitForLoop(c *ForLoopContext)

	// ExitWhileLoop is called when exiting the whileLoop production.
	ExitWhileLoop(c *WhileLoopContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitExp2 is called when exiting the exp2 production.
	ExitExp2(c *Exp2Context)

	// ExitT_exp is called when exiting the t_exp production.
	ExitT_exp(c *T_expContext)

	// ExitT_exp2 is called when exiting the t_exp2 production.
	ExitT_exp2(c *T_exp2Context)

	// ExitG_exp is called when exiting the g_exp production.
	ExitG_exp(c *G_expContext)

	// ExitM_exp is called when exiting the m_exp production.
	ExitM_exp(c *M_expContext)

	// ExitM_exp2 is called when exiting the m_exp2 production.
	ExitM_exp2(c *M_exp2Context)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitTerm2 is called when exiting the term2 production.
	ExitTerm2(c *Term2Context)

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

	// ExitCte_b is called when exiting the cte_b production.
	ExitCte_b(c *Cte_bContext)

	// ExitCte_s is called when exiting the cte_s production.
	ExitCte_s(c *Cte_sContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitTypeRule is called when exiting the typeRule production.
	ExitTypeRule(c *TypeRuleContext)

	// ExitRelop is called when exiting the relop production.
	ExitRelop(c *RelopContext)
}
