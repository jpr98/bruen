// Code generated from Bruen.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Bruen

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBruenListener is a complete listener for a parse tree produced by BruenParser.
type BaseBruenListener struct{}

var _ BruenListener = &BaseBruenListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBruenListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBruenListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBruenListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBruenListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseBruenListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseBruenListener) ExitProgram(ctx *ProgramContext) {}

// EnterProgram2 is called when production program2 is entered.
func (s *BaseBruenListener) EnterProgram2(ctx *Program2Context) {}

// ExitProgram2 is called when production program2 is exited.
func (s *BaseBruenListener) ExitProgram2(ctx *Program2Context) {}

// EnterClassDef is called when production classDef is entered.
func (s *BaseBruenListener) EnterClassDef(ctx *ClassDefContext) {}

// ExitClassDef is called when production classDef is exited.
func (s *BaseBruenListener) ExitClassDef(ctx *ClassDefContext) {}

// EnterClassBlock is called when production classBlock is entered.
func (s *BaseBruenListener) EnterClassBlock(ctx *ClassBlockContext) {}

// ExitClassBlock is called when production classBlock is exited.
func (s *BaseBruenListener) ExitClassBlock(ctx *ClassBlockContext) {}

// EnterClassAttributes is called when production classAttributes is entered.
func (s *BaseBruenListener) EnterClassAttributes(ctx *ClassAttributesContext) {}

// ExitClassAttributes is called when production classAttributes is exited.
func (s *BaseBruenListener) ExitClassAttributes(ctx *ClassAttributesContext) {}

// EnterClassMethod is called when production classMethod is entered.
func (s *BaseBruenListener) EnterClassMethod(ctx *ClassMethodContext) {}

// ExitClassMethod is called when production classMethod is exited.
func (s *BaseBruenListener) ExitClassMethod(ctx *ClassMethodContext) {}

// EnterClassInit is called when production classInit is entered.
func (s *BaseBruenListener) EnterClassInit(ctx *ClassInitContext) {}

// ExitClassInit is called when production classInit is exited.
func (s *BaseBruenListener) ExitClassInit(ctx *ClassInitContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseBruenListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseBruenListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterVarsDec is called when production varsDec is entered.
func (s *BaseBruenListener) EnterVarsDec(ctx *VarsDecContext) {}

// ExitVarsDec is called when production varsDec is exited.
func (s *BaseBruenListener) ExitVarsDec(ctx *VarsDecContext) {}

// EnterVarsDecArray is called when production varsDecArray is entered.
func (s *BaseBruenListener) EnterVarsDecArray(ctx *VarsDecArrayContext) {}

// ExitVarsDecArray is called when production varsDecArray is exited.
func (s *BaseBruenListener) ExitVarsDecArray(ctx *VarsDecArrayContext) {}

// EnterVarsDecMat is called when production varsDecMat is entered.
func (s *BaseBruenListener) EnterVarsDecMat(ctx *VarsDecMatContext) {}

// ExitVarsDecMat is called when production varsDecMat is exited.
func (s *BaseBruenListener) ExitVarsDecMat(ctx *VarsDecMatContext) {}

// EnterAttributesDeclaration is called when production attributesDeclaration is entered.
func (s *BaseBruenListener) EnterAttributesDeclaration(ctx *AttributesDeclarationContext) {}

// ExitAttributesDeclaration is called when production attributesDeclaration is exited.
func (s *BaseBruenListener) ExitAttributesDeclaration(ctx *AttributesDeclarationContext) {}

// EnterAttributesDec is called when production attributesDec is entered.
func (s *BaseBruenListener) EnterAttributesDec(ctx *AttributesDecContext) {}

// ExitAttributesDec is called when production attributesDec is exited.
func (s *BaseBruenListener) ExitAttributesDec(ctx *AttributesDecContext) {}

// EnterVarsTypeInit is called when production varsTypeInit is entered.
func (s *BaseBruenListener) EnterVarsTypeInit(ctx *VarsTypeInitContext) {}

// ExitVarsTypeInit is called when production varsTypeInit is exited.
func (s *BaseBruenListener) ExitVarsTypeInit(ctx *VarsTypeInitContext) {}

// EnterVarsTypeInit2 is called when production varsTypeInit2 is entered.
func (s *BaseBruenListener) EnterVarsTypeInit2(ctx *VarsTypeInit2Context) {}

// ExitVarsTypeInit2 is called when production varsTypeInit2 is exited.
func (s *BaseBruenListener) ExitVarsTypeInit2(ctx *VarsTypeInit2Context) {}

// EnterVars is called when production vars is entered.
func (s *BaseBruenListener) EnterVars(ctx *VarsContext) {}

// ExitVars is called when production vars is exited.
func (s *BaseBruenListener) ExitVars(ctx *VarsContext) {}

// EnterVarArray is called when production varArray is entered.
func (s *BaseBruenListener) EnterVarArray(ctx *VarArrayContext) {}

// ExitVarArray is called when production varArray is exited.
func (s *BaseBruenListener) ExitVarArray(ctx *VarArrayContext) {}

// EnterVarMat is called when production varMat is entered.
func (s *BaseBruenListener) EnterVarMat(ctx *VarMatContext) {}

// ExitVarMat is called when production varMat is exited.
func (s *BaseBruenListener) ExitVarMat(ctx *VarMatContext) {}

// EnterFunctions is called when production functions is entered.
func (s *BaseBruenListener) EnterFunctions(ctx *FunctionsContext) {}

// ExitFunctions is called when production functions is exited.
func (s *BaseBruenListener) ExitFunctions(ctx *FunctionsContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseBruenListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseBruenListener) ExitParameters(ctx *ParametersContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseBruenListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseBruenListener) ExitParameter(ctx *ParameterContext) {}

// EnterFunctionBlock is called when production functionBlock is entered.
func (s *BaseBruenListener) EnterFunctionBlock(ctx *FunctionBlockContext) {}

// ExitFunctionBlock is called when production functionBlock is exited.
func (s *BaseBruenListener) ExitFunctionBlock(ctx *FunctionBlockContext) {}

// EnterReturnRule is called when production returnRule is entered.
func (s *BaseBruenListener) EnterReturnRule(ctx *ReturnRuleContext) {}

// ExitReturnRule is called when production returnRule is exited.
func (s *BaseBruenListener) ExitReturnRule(ctx *ReturnRuleContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseBruenListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseBruenListener) ExitBlock(ctx *BlockContext) {}

// EnterStatutesNoReturn is called when production statutesNoReturn is entered.
func (s *BaseBruenListener) EnterStatutesNoReturn(ctx *StatutesNoReturnContext) {}

// ExitStatutesNoReturn is called when production statutesNoReturn is exited.
func (s *BaseBruenListener) ExitStatutesNoReturn(ctx *StatutesNoReturnContext) {}

// EnterStatutes is called when production statutes is entered.
func (s *BaseBruenListener) EnterStatutes(ctx *StatutesContext) {}

// ExitStatutes is called when production statutes is exited.
func (s *BaseBruenListener) ExitStatutes(ctx *StatutesContext) {}

// EnterAssignation is called when production assignation is entered.
func (s *BaseBruenListener) EnterAssignation(ctx *AssignationContext) {}

// ExitAssignation is called when production assignation is exited.
func (s *BaseBruenListener) ExitAssignation(ctx *AssignationContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseBruenListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseBruenListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseBruenListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseBruenListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArguments2 is called when production arguments2 is entered.
func (s *BaseBruenListener) EnterArguments2(ctx *Arguments2Context) {}

// ExitArguments2 is called when production arguments2 is exited.
func (s *BaseBruenListener) ExitArguments2(ctx *Arguments2Context) {}

// EnterMethodCall is called when production methodCall is entered.
func (s *BaseBruenListener) EnterMethodCall(ctx *MethodCallContext) {}

// ExitMethodCall is called when production methodCall is exited.
func (s *BaseBruenListener) ExitMethodCall(ctx *MethodCallContext) {}

// EnterInitCall is called when production initCall is entered.
func (s *BaseBruenListener) EnterInitCall(ctx *InitCallContext) {}

// ExitInitCall is called when production initCall is exited.
func (s *BaseBruenListener) ExitInitCall(ctx *InitCallContext) {}

// EnterCall is called when production call is entered.
func (s *BaseBruenListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseBruenListener) ExitCall(ctx *CallContext) {}

// EnterRead is called when production read is entered.
func (s *BaseBruenListener) EnterRead(ctx *ReadContext) {}

// ExitRead is called when production read is exited.
func (s *BaseBruenListener) ExitRead(ctx *ReadContext) {}

// EnterRead2 is called when production read2 is entered.
func (s *BaseBruenListener) EnterRead2(ctx *Read2Context) {}

// ExitRead2 is called when production read2 is exited.
func (s *BaseBruenListener) ExitRead2(ctx *Read2Context) {}

// EnterRead3 is called when production read3 is entered.
func (s *BaseBruenListener) EnterRead3(ctx *Read3Context) {}

// ExitRead3 is called when production read3 is exited.
func (s *BaseBruenListener) ExitRead3(ctx *Read3Context) {}

// EnterWrite is called when production write is entered.
func (s *BaseBruenListener) EnterWrite(ctx *WriteContext) {}

// ExitWrite is called when production write is exited.
func (s *BaseBruenListener) ExitWrite(ctx *WriteContext) {}

// EnterW_arguments is called when production w_arguments is entered.
func (s *BaseBruenListener) EnterW_arguments(ctx *W_argumentsContext) {}

// ExitW_arguments is called when production w_arguments is exited.
func (s *BaseBruenListener) ExitW_arguments(ctx *W_argumentsContext) {}

// EnterW_arguments2 is called when production w_arguments2 is entered.
func (s *BaseBruenListener) EnterW_arguments2(ctx *W_arguments2Context) {}

// ExitW_arguments2 is called when production w_arguments2 is exited.
func (s *BaseBruenListener) ExitW_arguments2(ctx *W_arguments2Context) {}

// EnterConditional is called when production conditional is entered.
func (s *BaseBruenListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BaseBruenListener) ExitConditional(ctx *ConditionalContext) {}

// EnterConditional2 is called when production conditional2 is entered.
func (s *BaseBruenListener) EnterConditional2(ctx *Conditional2Context) {}

// ExitConditional2 is called when production conditional2 is exited.
func (s *BaseBruenListener) ExitConditional2(ctx *Conditional2Context) {}

// EnterConditional3 is called when production conditional3 is entered.
func (s *BaseBruenListener) EnterConditional3(ctx *Conditional3Context) {}

// ExitConditional3 is called when production conditional3 is exited.
func (s *BaseBruenListener) ExitConditional3(ctx *Conditional3Context) {}

// EnterConditional4 is called when production conditional4 is entered.
func (s *BaseBruenListener) EnterConditional4(ctx *Conditional4Context) {}

// ExitConditional4 is called when production conditional4 is exited.
func (s *BaseBruenListener) ExitConditional4(ctx *Conditional4Context) {}

// EnterForLoop is called when production forLoop is entered.
func (s *BaseBruenListener) EnterForLoop(ctx *ForLoopContext) {}

// ExitForLoop is called when production forLoop is exited.
func (s *BaseBruenListener) ExitForLoop(ctx *ForLoopContext) {}

// EnterForLoop2 is called when production forLoop2 is entered.
func (s *BaseBruenListener) EnterForLoop2(ctx *ForLoop2Context) {}

// ExitForLoop2 is called when production forLoop2 is exited.
func (s *BaseBruenListener) ExitForLoop2(ctx *ForLoop2Context) {}

// EnterForLoop3 is called when production forLoop3 is entered.
func (s *BaseBruenListener) EnterForLoop3(ctx *ForLoop3Context) {}

// ExitForLoop3 is called when production forLoop3 is exited.
func (s *BaseBruenListener) ExitForLoop3(ctx *ForLoop3Context) {}

// EnterWhileLoop is called when production whileLoop is entered.
func (s *BaseBruenListener) EnterWhileLoop(ctx *WhileLoopContext) {}

// ExitWhileLoop is called when production whileLoop is exited.
func (s *BaseBruenListener) ExitWhileLoop(ctx *WhileLoopContext) {}

// EnterWhileLoop2 is called when production whileLoop2 is entered.
func (s *BaseBruenListener) EnterWhileLoop2(ctx *WhileLoop2Context) {}

// ExitWhileLoop2 is called when production whileLoop2 is exited.
func (s *BaseBruenListener) ExitWhileLoop2(ctx *WhileLoop2Context) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBruenListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBruenListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseBruenListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseBruenListener) ExitExp(ctx *ExpContext) {}

// EnterExp2 is called when production exp2 is entered.
func (s *BaseBruenListener) EnterExp2(ctx *Exp2Context) {}

// ExitExp2 is called when production exp2 is exited.
func (s *BaseBruenListener) ExitExp2(ctx *Exp2Context) {}

// EnterT_exp is called when production t_exp is entered.
func (s *BaseBruenListener) EnterT_exp(ctx *T_expContext) {}

// ExitT_exp is called when production t_exp is exited.
func (s *BaseBruenListener) ExitT_exp(ctx *T_expContext) {}

// EnterT_exp2 is called when production t_exp2 is entered.
func (s *BaseBruenListener) EnterT_exp2(ctx *T_exp2Context) {}

// ExitT_exp2 is called when production t_exp2 is exited.
func (s *BaseBruenListener) ExitT_exp2(ctx *T_exp2Context) {}

// EnterG_exp is called when production g_exp is entered.
func (s *BaseBruenListener) EnterG_exp(ctx *G_expContext) {}

// ExitG_exp is called when production g_exp is exited.
func (s *BaseBruenListener) ExitG_exp(ctx *G_expContext) {}

// EnterG_exp2 is called when production g_exp2 is entered.
func (s *BaseBruenListener) EnterG_exp2(ctx *G_exp2Context) {}

// ExitG_exp2 is called when production g_exp2 is exited.
func (s *BaseBruenListener) ExitG_exp2(ctx *G_exp2Context) {}

// EnterM_exp is called when production m_exp is entered.
func (s *BaseBruenListener) EnterM_exp(ctx *M_expContext) {}

// ExitM_exp is called when production m_exp is exited.
func (s *BaseBruenListener) ExitM_exp(ctx *M_expContext) {}

// EnterM_exp2 is called when production m_exp2 is entered.
func (s *BaseBruenListener) EnterM_exp2(ctx *M_exp2Context) {}

// ExitM_exp2 is called when production m_exp2 is exited.
func (s *BaseBruenListener) ExitM_exp2(ctx *M_exp2Context) {}

// EnterTerm is called when production term is entered.
func (s *BaseBruenListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseBruenListener) ExitTerm(ctx *TermContext) {}

// EnterTerm2 is called when production term2 is entered.
func (s *BaseBruenListener) EnterTerm2(ctx *Term2Context) {}

// ExitTerm2 is called when production term2 is exited.
func (s *BaseBruenListener) ExitTerm2(ctx *Term2Context) {}

// EnterFactor is called when production factor is entered.
func (s *BaseBruenListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseBruenListener) ExitFactor(ctx *FactorContext) {}

// EnterFactor2 is called when production factor2 is entered.
func (s *BaseBruenListener) EnterFactor2(ctx *Factor2Context) {}

// ExitFactor2 is called when production factor2 is exited.
func (s *BaseBruenListener) ExitFactor2(ctx *Factor2Context) {}

// EnterVarCte is called when production varCte is entered.
func (s *BaseBruenListener) EnterVarCte(ctx *VarCteContext) {}

// ExitVarCte is called when production varCte is exited.
func (s *BaseBruenListener) ExitVarCte(ctx *VarCteContext) {}

// EnterCte_i is called when production cte_i is entered.
func (s *BaseBruenListener) EnterCte_i(ctx *Cte_iContext) {}

// ExitCte_i is called when production cte_i is exited.
func (s *BaseBruenListener) ExitCte_i(ctx *Cte_iContext) {}

// EnterCte_f is called when production cte_f is entered.
func (s *BaseBruenListener) EnterCte_f(ctx *Cte_fContext) {}

// ExitCte_f is called when production cte_f is exited.
func (s *BaseBruenListener) ExitCte_f(ctx *Cte_fContext) {}

// EnterCte_c is called when production cte_c is entered.
func (s *BaseBruenListener) EnterCte_c(ctx *Cte_cContext) {}

// ExitCte_c is called when production cte_c is exited.
func (s *BaseBruenListener) ExitCte_c(ctx *Cte_cContext) {}

// EnterCte_b is called when production cte_b is entered.
func (s *BaseBruenListener) EnterCte_b(ctx *Cte_bContext) {}

// ExitCte_b is called when production cte_b is exited.
func (s *BaseBruenListener) ExitCte_b(ctx *Cte_bContext) {}

// EnterCte_s is called when production cte_s is entered.
func (s *BaseBruenListener) EnterCte_s(ctx *Cte_sContext) {}

// ExitCte_s is called when production cte_s is exited.
func (s *BaseBruenListener) ExitCte_s(ctx *Cte_sContext) {}

// EnterMain is called when production main is entered.
func (s *BaseBruenListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BaseBruenListener) ExitMain(ctx *MainContext) {}

// EnterMainBlock is called when production mainBlock is entered.
func (s *BaseBruenListener) EnterMainBlock(ctx *MainBlockContext) {}

// ExitMainBlock is called when production mainBlock is exited.
func (s *BaseBruenListener) ExitMainBlock(ctx *MainBlockContext) {}

// EnterTypeRule is called when production typeRule is entered.
func (s *BaseBruenListener) EnterTypeRule(ctx *TypeRuleContext) {}

// ExitTypeRule is called when production typeRule is exited.
func (s *BaseBruenListener) ExitTypeRule(ctx *TypeRuleContext) {}

// EnterRelop is called when production relop is entered.
func (s *BaseBruenListener) EnterRelop(ctx *RelopContext) {}

// ExitRelop is called when production relop is exited.
func (s *BaseBruenListener) ExitRelop(ctx *RelopContext) {}
