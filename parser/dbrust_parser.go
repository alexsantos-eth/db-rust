// Code generated from DBRust.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // DBRust

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import I "main/interfaces"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 66, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 21, 10, 3, 12, 3, 14, 3, 24, 11,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 45, 10, 6, 12, 6, 14,
	6, 48, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 58,
	10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 64, 10, 8, 3, 8, 2, 3, 10, 9, 2, 4,
	6, 8, 10, 12, 14, 2, 2, 2, 64, 2, 16, 3, 2, 2, 2, 4, 22, 3, 2, 2, 2, 6,
	27, 3, 2, 2, 2, 8, 30, 3, 2, 2, 2, 10, 35, 3, 2, 2, 2, 12, 57, 3, 2, 2,
	2, 14, 63, 3, 2, 2, 2, 16, 17, 5, 4, 3, 2, 17, 18, 8, 2, 1, 2, 18, 3, 3,
	2, 2, 2, 19, 21, 5, 6, 4, 2, 20, 19, 3, 2, 2, 2, 21, 24, 3, 2, 2, 2, 22,
	20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 25, 3, 2, 2, 2, 24, 22, 3, 2, 2,
	2, 25, 26, 8, 3, 1, 2, 26, 5, 3, 2, 2, 2, 27, 28, 5, 8, 5, 2, 28, 29, 8,
	4, 1, 2, 29, 7, 3, 2, 2, 2, 30, 31, 7, 5, 2, 2, 31, 32, 7, 6, 2, 2, 32,
	33, 5, 10, 6, 2, 33, 34, 8, 5, 1, 2, 34, 9, 3, 2, 2, 2, 35, 36, 8, 6, 1,
	2, 36, 37, 5, 14, 8, 2, 37, 38, 8, 6, 1, 2, 38, 46, 3, 2, 2, 2, 39, 40,
	12, 4, 2, 2, 40, 41, 5, 12, 7, 2, 41, 42, 5, 10, 6, 5, 42, 43, 8, 6, 1,
	2, 43, 45, 3, 2, 2, 2, 44, 39, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44,
	3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 11, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2,
	49, 50, 7, 7, 2, 2, 50, 58, 8, 7, 1, 2, 51, 52, 7, 8, 2, 2, 52, 58, 8,
	7, 1, 2, 53, 54, 7, 11, 2, 2, 54, 58, 8, 7, 1, 2, 55, 56, 7, 10, 2, 2,
	56, 58, 8, 7, 1, 2, 57, 49, 3, 2, 2, 2, 57, 51, 3, 2, 2, 2, 57, 53, 3,
	2, 2, 2, 57, 55, 3, 2, 2, 2, 58, 13, 3, 2, 2, 2, 59, 60, 7, 3, 2, 2, 60,
	64, 8, 8, 1, 2, 61, 62, 7, 4, 2, 2, 62, 64, 8, 8, 1, 2, 63, 59, 3, 2, 2,
	2, 63, 61, 3, 2, 2, 2, 64, 15, 3, 2, 2, 2, 6, 22, 46, 57, 63,
}
var literalNames = []string{
	"", "", "", "", "'='", "'*'", "'/'", "'%'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "NUMBER", "STRING", "ID", "EQUALS", "MUL", "DIV", "MOD", "ADD", "SUB",
	"WHITESPACE",
}

var ruleNames = []string{
	"start", "instructions", "instruction", "assignment", "expression", "expOp",
	"value",
}

type DBRustParser struct {
	*antlr.BaseParser
}

// NewDBRustParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *DBRustParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDBRustParser(input antlr.TokenStream) *DBRustParser {
	this := new(DBRustParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "DBRust.g4"

	return this
}

// DBRustParser tokens.
const (
	DBRustParserEOF        = antlr.TokenEOF
	DBRustParserNUMBER     = 1
	DBRustParserSTRING     = 2
	DBRustParserID         = 3
	DBRustParserEQUALS     = 4
	DBRustParserMUL        = 5
	DBRustParserDIV        = 6
	DBRustParserMOD        = 7
	DBRustParserADD        = 8
	DBRustParserSUB        = 9
	DBRustParserWHITESPACE = 10
)

// DBRustParser rules.
const (
	DBRustParserRULE_start        = 0
	DBRustParserRULE_instructions = 1
	DBRustParserRULE_instruction  = 2
	DBRustParserRULE_assignment   = 3
	DBRustParserRULE_expression   = 4
	DBRustParserRULE_expOp        = 5
	DBRustParserRULE_value        = 6
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instructions returns the _instructions rule contexts.
	Get_instructions() IInstructionsContext

	// Set_instructions sets the _instructions rule contexts.
	Set_instructions(IInstructionsContext)

	// GetList returns the list attribute.
	GetList() *arrayList.List

	// SetList sets the list attribute.
	SetList(*arrayList.List)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	list          *arrayList.List
	_instructions IInstructionsContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_instructions() IInstructionsContext { return s._instructions }

func (s *StartContext) Set_instructions(v IInstructionsContext) { s._instructions = v }

func (s *StartContext) GetList() *arrayList.List { return s.list }

func (s *StartContext) SetList(v *arrayList.List) { s.list = v }

func (s *StartContext) Instructions() IInstructionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *DBRustParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DBRustParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)

		var _x = p.Instructions()

		localctx.(*StartContext)._instructions = _x
	}
	localctx.(*StartContext).list = localctx.(*StartContext).Get_instructions().GetL()

	return localctx
}

// IInstructionsContext is an interface to support dynamic dispatch.
type IInstructionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruction returns the _instruction rule contexts.
	Get_instruction() IInstructionContext

	// Set_instruction sets the _instruction rule contexts.
	Set_instruction(IInstructionContext)

	// GetE returns the e rule context list.
	GetE() []IInstructionContext

	// SetE sets the e rule context list.
	SetE([]IInstructionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstructionsContext differentiates from other interfaces.
	IsInstructionsContext()
}

type InstructionsContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            *arrayList.List
	_instruction IInstructionContext
	e            []IInstructionContext
}

func NewEmptyInstructionsContext() *InstructionsContext {
	var p = new(InstructionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_instructions
	return p
}

func (*InstructionsContext) IsInstructionsContext() {}

func NewInstructionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionsContext {
	var p = new(InstructionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_instructions

	return p
}

func (s *InstructionsContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionsContext) Get_instruction() IInstructionContext { return s._instruction }

func (s *InstructionsContext) Set_instruction(v IInstructionContext) { s._instruction = v }

func (s *InstructionsContext) GetE() []IInstructionContext { return s.e }

func (s *InstructionsContext) SetE(v []IInstructionContext) { s.e = v }

func (s *InstructionsContext) GetL() *arrayList.List { return s.l }

func (s *InstructionsContext) SetL(v *arrayList.List) { s.l = v }

func (s *InstructionsContext) AllInstruction() []IInstructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstructionContext)(nil)).Elem())
	var tst = make([]IInstructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstructionContext)
		}
	}

	return tst
}

func (s *InstructionsContext) Instruction(i int) IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *InstructionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterInstructions(s)
	}
}

func (s *InstructionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitInstructions(s)
	}
}

func (p *DBRustParser) Instructions() (localctx IInstructionsContext) {
	localctx = NewInstructionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DBRustParserRULE_instructions)

	localctx.(*InstructionsContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DBRustParserID {
		{
			p.SetState(17)

			var _x = p.Instruction()

			localctx.(*InstructionsContext)._instruction = _x
		}
		localctx.(*InstructionsContext).e = append(localctx.(*InstructionsContext).e, localctx.(*InstructionsContext)._instruction)

		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstructionsContext).GetE()
	for _, e := range listInt {
		localctx.(*InstructionsContext).l.Add(e.GetState())
	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAssign returns the assign rule contexts.
	GetAssign() IAssignmentContext

	// SetAssign sets the assign rule contexts.
	SetAssign(IAssignmentContext)

	// GetState returns the state attribute.
	GetState() I.IInstruction

	// SetState sets the state attribute.
	SetState(I.IInstruction)

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	state  I.IInstruction
	assign IAssignmentContext
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) GetAssign() IAssignmentContext { return s.assign }

func (s *InstructionContext) SetAssign(v IAssignmentContext) { s.assign = v }

func (s *InstructionContext) GetState() I.IInstruction { return s.state }

func (s *InstructionContext) SetState(v I.IInstruction) { s.state = v }

func (s *InstructionContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *DBRustParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DBRustParserRULE_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)

		var _x = p.Assignment()

		localctx.(*InstructionContext).assign = _x
	}

	localctx.(*InstructionContext).state = localctx.(*InstructionContext).GetAssign().GetState()

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdText returns the idText token.
	GetIdText() antlr.Token

	// SetIdText sets the idText token.
	SetIdText(antlr.Token)

	// GetExp returns the exp rule contexts.
	GetExp() IExpressionContext

	// SetExp sets the exp rule contexts.
	SetExp(IExpressionContext)

	// GetState returns the state attribute.
	GetState() I.Assignment

	// SetState sets the state attribute.
	SetState(I.Assignment)

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	state  I.Assignment
	idText antlr.Token
	exp    IExpressionContext
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) GetIdText() antlr.Token { return s.idText }

func (s *AssignmentContext) SetIdText(v antlr.Token) { s.idText = v }

func (s *AssignmentContext) GetExp() IExpressionContext { return s.exp }

func (s *AssignmentContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *AssignmentContext) GetState() I.Assignment { return s.state }

func (s *AssignmentContext) SetState(v I.Assignment) { s.state = v }

func (s *AssignmentContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(DBRustParserEQUALS, 0)
}

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(DBRustParserID, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *DBRustParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DBRustParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)

		var _m = p.Match(DBRustParserID)

		localctx.(*AssignmentContext).idText = _m
	}
	{
		p.SetState(29)
		p.Match(DBRustParserEQUALS)
	}
	{
		p.SetState(30)

		var _x = p.expression(0)

		localctx.(*AssignmentContext).exp = _x
	}

	expPoint := localctx.(*AssignmentContext).GetExp().GetState()
	localctx.(*AssignmentContext).state = I.Assignment{(func() string {
		if localctx.(*AssignmentContext).GetIdText() == nil {
			return ""
		} else {
			return localctx.(*AssignmentContext).GetIdText().GetText()
		}
	}()), &expPoint}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeftExp returns the leftExp rule contexts.
	GetLeftExp() IExpressionContext

	// Get_value returns the _value rule contexts.
	Get_value() IValueContext

	// Get_expOp returns the _expOp rule contexts.
	Get_expOp() IExpOpContext

	// GetRightExp returns the rightExp rule contexts.
	GetRightExp() IExpressionContext

	// SetLeftExp sets the leftExp rule contexts.
	SetLeftExp(IExpressionContext)

	// Set_value sets the _value rule contexts.
	Set_value(IValueContext)

	// Set_expOp sets the _expOp rule contexts.
	Set_expOp(IExpOpContext)

	// SetRightExp sets the rightExp rule contexts.
	SetRightExp(IExpressionContext)

	// GetState returns the state attribute.
	GetState() I.Expression

	// SetState sets the state attribute.
	SetState(I.Expression)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	state    I.Expression
	leftExp  IExpressionContext
	_value   IValueContext
	_expOp   IExpOpContext
	rightExp IExpressionContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetLeftExp() IExpressionContext { return s.leftExp }

func (s *ExpressionContext) Get_value() IValueContext { return s._value }

func (s *ExpressionContext) Get_expOp() IExpOpContext { return s._expOp }

func (s *ExpressionContext) GetRightExp() IExpressionContext { return s.rightExp }

func (s *ExpressionContext) SetLeftExp(v IExpressionContext) { s.leftExp = v }

func (s *ExpressionContext) Set_value(v IValueContext) { s._value = v }

func (s *ExpressionContext) Set_expOp(v IExpOpContext) { s._expOp = v }

func (s *ExpressionContext) SetRightExp(v IExpressionContext) { s.rightExp = v }

func (s *ExpressionContext) GetState() I.Expression { return s.state }

func (s *ExpressionContext) SetState(v I.Expression) { s.state = v }

func (s *ExpressionContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ExpressionContext) ExpOp() IExpOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpOpContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DBRustParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *DBRustParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, DBRustParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)

		var _x = p.Value()

		localctx.(*ExpressionContext)._value = _x
	}

	sym := localctx.(*ExpressionContext).Get_value().GetState()
	localctx.(*ExpressionContext).state = I.Expression{
		Value: &sym, Left: nil, Right: nil, Operation: I.NOOP}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			localctx.(*ExpressionContext).leftExp = _prevctx
			p.PushNewRecursionContext(localctx, _startState, DBRustParserRULE_expression)
			p.SetState(37)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(38)

				var _x = p.ExpOp()

				localctx.(*ExpressionContext)._expOp = _x
			}
			{
				p.SetState(39)

				var _x = p.expression(3)

				localctx.(*ExpressionContext).rightExp = _x
			}

			left, right := localctx.(*ExpressionContext).GetLeftExp().GetState(), localctx.(*ExpressionContext).GetRightExp().GetState()
			localctx.(*ExpressionContext).state = I.Expression{
				Value: nil, Left: &left, Right: &right, Operation: localctx.(*ExpressionContext).Get_expOp().GetState()}

		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IExpOpContext is an interface to support dynamic dispatch.
type IExpOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetState returns the state attribute.
	GetState() I.Operation

	// SetState sets the state attribute.
	SetState(I.Operation)

	// IsExpOpContext differentiates from other interfaces.
	IsExpOpContext()
}

type ExpOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	state  I.Operation
}

func NewEmptyExpOpContext() *ExpOpContext {
	var p = new(ExpOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_expOp
	return p
}

func (*ExpOpContext) IsExpOpContext() {}

func NewExpOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpOpContext {
	var p = new(ExpOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_expOp

	return p
}

func (s *ExpOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpOpContext) GetState() I.Operation { return s.state }

func (s *ExpOpContext) SetState(v I.Operation) { s.state = v }

func (s *ExpOpContext) MUL() antlr.TerminalNode {
	return s.GetToken(DBRustParserMUL, 0)
}

func (s *ExpOpContext) DIV() antlr.TerminalNode {
	return s.GetToken(DBRustParserDIV, 0)
}

func (s *ExpOpContext) SUB() antlr.TerminalNode {
	return s.GetToken(DBRustParserSUB, 0)
}

func (s *ExpOpContext) ADD() antlr.TerminalNode {
	return s.GetToken(DBRustParserADD, 0)
}

func (s *ExpOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterExpOp(s)
	}
}

func (s *ExpOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitExpOp(s)
	}
}

func (p *DBRustParser) ExpOp() (localctx IExpOpContext) {
	localctx = NewExpOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DBRustParserRULE_expOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DBRustParserMUL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.Match(DBRustParserMUL)
		}

		localctx.(*ExpOpContext).state = I.MUL

	case DBRustParserDIV:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Match(DBRustParserDIV)
		}

		localctx.(*ExpOpContext).state = I.DIV

	case DBRustParserSUB:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(51)
			p.Match(DBRustParserSUB)
		}

		localctx.(*ExpOpContext).state = I.SUB

	case DBRustParserADD:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(53)
			p.Match(DBRustParserADD)
		}

		localctx.(*ExpOpContext).state = I.ADD

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// GetState returns the state attribute.
	GetState() I.Value

	// SetState sets the state attribute.
	SetState(I.Value)

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	state   I.Value
	_NUMBER antlr.Token
	_STRING antlr.Token
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *ValueContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ValueContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *ValueContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ValueContext) GetState() I.Value { return s.state }

func (s *ValueContext) SetState(v I.Value) { s.state = v }

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(DBRustParserNUMBER, 0)
}

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(DBRustParserSTRING, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *DBRustParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DBRustParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(61)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DBRustParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)

			var _m = p.Match(DBRustParserNUMBER)

			localctx.(*ValueContext)._NUMBER = _m
		}

		localctx.(*ValueContext).state = I.Value{I.INTEGER, (func() string {
			if localctx.(*ValueContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).Get_NUMBER().GetText()
			}
		}())}

	case DBRustParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)

			var _m = p.Match(DBRustParserSTRING)

			localctx.(*ValueContext)._STRING = _m
		}

		localctx.(*ValueContext).state = I.Value{I.INTEGER, (func() string {
			if localctx.(*ValueContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).Get_STRING().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ValueContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).Get_STRING().GetText()
			}
		}()))-1]}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *DBRustParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *DBRustParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}