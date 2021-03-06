// Code generated from DBRust.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // DBRust

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDBRustListener is a complete listener for a parse tree produced by DBRustParser.
type BaseDBRustListener struct{}

var _ DBRustListener = &BaseDBRustListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDBRustListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDBRustListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDBRustListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDBRustListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseDBRustListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseDBRustListener) ExitStart(ctx *StartContext) {}

// EnterInstructionsBlock is called when production instructionsBlock is entered.
func (s *BaseDBRustListener) EnterInstructionsBlock(ctx *InstructionsBlockContext) {}

// ExitInstructionsBlock is called when production instructionsBlock is exited.
func (s *BaseDBRustListener) ExitInstructionsBlock(ctx *InstructionsBlockContext) {}

// EnterInstructions is called when production instructions is entered.
func (s *BaseDBRustListener) EnterInstructions(ctx *InstructionsContext) {}

// ExitInstructions is called when production instructions is exited.
func (s *BaseDBRustListener) ExitInstructions(ctx *InstructionsContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseDBRustListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseDBRustListener) ExitInstruction(ctx *InstructionContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseDBRustListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseDBRustListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseDBRustListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseDBRustListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpList is called when production expList is entered.
func (s *BaseDBRustListener) EnterExpList(ctx *ExpListContext) {}

// ExitExpList is called when production expList is exited.
func (s *BaseDBRustListener) ExitExpList(ctx *ExpListContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDBRustListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDBRustListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpOpAlgb1 is called when production expOpAlgb1 is entered.
func (s *BaseDBRustListener) EnterExpOpAlgb1(ctx *ExpOpAlgb1Context) {}

// ExitExpOpAlgb1 is called when production expOpAlgb1 is exited.
func (s *BaseDBRustListener) ExitExpOpAlgb1(ctx *ExpOpAlgb1Context) {}

// EnterExpOpAlgb2 is called when production expOpAlgb2 is entered.
func (s *BaseDBRustListener) EnterExpOpAlgb2(ctx *ExpOpAlgb2Context) {}

// ExitExpOpAlgb2 is called when production expOpAlgb2 is exited.
func (s *BaseDBRustListener) ExitExpOpAlgb2(ctx *ExpOpAlgb2Context) {}

// EnterExpOpRel1 is called when production expOpRel1 is entered.
func (s *BaseDBRustListener) EnterExpOpRel1(ctx *ExpOpRel1Context) {}

// ExitExpOpRel1 is called when production expOpRel1 is exited.
func (s *BaseDBRustListener) ExitExpOpRel1(ctx *ExpOpRel1Context) {}

// EnterValueType is called when production valueType is entered.
func (s *BaseDBRustListener) EnterValueType(ctx *ValueTypeContext) {}

// ExitValueType is called when production valueType is exited.
func (s *BaseDBRustListener) ExitValueType(ctx *ValueTypeContext) {}

// EnterValue is called when production value is entered.
func (s *BaseDBRustListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseDBRustListener) ExitValue(ctx *ValueContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseDBRustListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseDBRustListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterMethods is called when production methods is entered.
func (s *BaseDBRustListener) EnterMethods(ctx *MethodsContext) {}

// ExitMethods is called when production methods is exited.
func (s *BaseDBRustListener) ExitMethods(ctx *MethodsContext) {}

// EnterPrintlnCall is called when production printlnCall is entered.
func (s *BaseDBRustListener) EnterPrintlnCall(ctx *PrintlnCallContext) {}

// ExitPrintlnCall is called when production printlnCall is exited.
func (s *BaseDBRustListener) ExitPrintlnCall(ctx *PrintlnCallContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseDBRustListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseDBRustListener) ExitParamList(ctx *ParamListContext) {}

// EnterParam is called when production param is entered.
func (s *BaseDBRustListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseDBRustListener) ExitParam(ctx *ParamContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseDBRustListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseDBRustListener) ExitFunction(ctx *FunctionContext) {}

// EnterReturnValue is called when production returnValue is entered.
func (s *BaseDBRustListener) EnterReturnValue(ctx *ReturnValueContext) {}

// ExitReturnValue is called when production returnValue is exited.
func (s *BaseDBRustListener) ExitReturnValue(ctx *ReturnValueContext) {}

// EnterConditions is called when production conditions is entered.
func (s *BaseDBRustListener) EnterConditions(ctx *ConditionsContext) {}

// ExitConditions is called when production conditions is exited.
func (s *BaseDBRustListener) ExitConditions(ctx *ConditionsContext) {}

// EnterTernaryConditions is called when production ternaryConditions is entered.
func (s *BaseDBRustListener) EnterTernaryConditions(ctx *TernaryConditionsContext) {}

// ExitTernaryConditions is called when production ternaryConditions is exited.
func (s *BaseDBRustListener) ExitTernaryConditions(ctx *TernaryConditionsContext) {}

// EnterConditionList is called when production conditionList is entered.
func (s *BaseDBRustListener) EnterConditionList(ctx *ConditionListContext) {}

// ExitConditionList is called when production conditionList is exited.
func (s *BaseDBRustListener) ExitConditionList(ctx *ConditionListContext) {}

// EnterTernConditionList is called when production ternConditionList is entered.
func (s *BaseDBRustListener) EnterTernConditionList(ctx *TernConditionListContext) {}

// ExitTernConditionList is called when production ternConditionList is exited.
func (s *BaseDBRustListener) ExitTernConditionList(ctx *TernConditionListContext) {}

// EnterElseIf is called when production elseIf is entered.
func (s *BaseDBRustListener) EnterElseIf(ctx *ElseIfContext) {}

// ExitElseIf is called when production elseIf is exited.
func (s *BaseDBRustListener) ExitElseIf(ctx *ElseIfContext) {}

// EnterTernElseIf is called when production ternElseIf is entered.
func (s *BaseDBRustListener) EnterTernElseIf(ctx *TernElseIfContext) {}

// ExitTernElseIf is called when production ternElseIf is exited.
func (s *BaseDBRustListener) ExitTernElseIf(ctx *TernElseIfContext) {}

// EnterMatchExp is called when production matchExp is entered.
func (s *BaseDBRustListener) EnterMatchExp(ctx *MatchExpContext) {}

// ExitMatchExp is called when production matchExp is exited.
func (s *BaseDBRustListener) ExitMatchExp(ctx *MatchExpContext) {}

// EnterMatchCaseList is called when production matchCaseList is entered.
func (s *BaseDBRustListener) EnterMatchCaseList(ctx *MatchCaseListContext) {}

// ExitMatchCaseList is called when production matchCaseList is exited.
func (s *BaseDBRustListener) ExitMatchCaseList(ctx *MatchCaseListContext) {}

// EnterMatchCase is called when production matchCase is entered.
func (s *BaseDBRustListener) EnterMatchCase(ctx *MatchCaseContext) {}

// ExitMatchCase is called when production matchCase is exited.
func (s *BaseDBRustListener) ExitMatchCase(ctx *MatchCaseContext) {}
