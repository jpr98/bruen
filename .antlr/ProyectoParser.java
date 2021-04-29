// Generated from /Users/juanpabloramos/Documents/TEC/OctavoSemestre/Compiladores/Proyecto/Proyecto.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ProyectoParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, INT=11, FLOAT=12, CHAR=13, BOOL=14, WS=15, OR=16, AND=17, GT=18, 
		LT=19, EQ=20, NEQ=21, MUL=22, DIV=23, ADD=24, SUB=25, SEMI=26, LPAREN=27, 
		RPAREN=28, IF=29, ELSE=30, WHILE=31, FOR=32, IN=33, CLASS=34, ATTRIBUTES=35, 
		METHODS=36, WRITE=37, READ=38, FUNCTION=39, RETURN=40, MAIN=41, VAR=42, 
		INT_TYPE=43, FLOAT_TYPE=44, CHAR_TYPE=45, STRING_TYPE=46, BOOL_TYPE=47, 
		VOID=48, PROGRAM=49, ID=50;
	public static final int
		RULE_program = 0, RULE_classDef = 1, RULE_classBlock = 2, RULE_varsDec = 3, 
		RULE_varsTypeInit = 4, RULE_vars = 5, RULE_functions = 6, RULE_parameters = 7, 
		RULE_parameter = 8, RULE_functionBlock = 9, RULE_returnRule = 10, RULE_block = 11, 
		RULE_statutes = 12, RULE_assignation = 13, RULE_functionCall = 14, RULE_arguments = 15, 
		RULE_argument = 16, RULE_methodCall = 17, RULE_call = 18, RULE_read = 19, 
		RULE_write = 20, RULE_conditional = 21, RULE_forLoop = 22, RULE_whileLoop = 23, 
		RULE_exp = 24, RULE_exp2 = 25, RULE_t_exp = 26, RULE_t_exp2 = 27, RULE_g_exp = 28, 
		RULE_g_exp2 = 29, RULE_m_exp = 30, RULE_m_exp2 = 31, RULE_term = 32, RULE_term2 = 33, 
		RULE_factor = 34, RULE_varCte = 35, RULE_cte_i = 36, RULE_cte_f = 37, 
		RULE_cte_c = 38, RULE_cte_b = 39, RULE_cte_s = 40, RULE_main = 41, RULE_typeRule = 42, 
		RULE_relop = 43;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "classDef", "classBlock", "varsDec", "varsTypeInit", "vars", 
			"functions", "parameters", "parameter", "functionBlock", "returnRule", 
			"block", "statutes", "assignation", "functionCall", "arguments", "argument", 
			"methodCall", "call", "read", "write", "conditional", "forLoop", "whileLoop", 
			"exp", "exp2", "t_exp", "t_exp2", "g_exp", "g_exp2", "m_exp", "m_exp2", 
			"term", "term2", "factor", "varCte", "cte_i", "cte_f", "cte_c", "cte_b", 
			"cte_s", "main", "typeRule", "relop"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'{'", "'}'", "':'", "'['", "']'", "']['", "'='", "'.'", "','", 
			"'\"'", null, null, null, null, null, "'||'", "'&&'", "'>'", "'<'", "'=='", 
			"'!='", "'*'", "'/'", "'+'", "'-'", "';'", "'('", "')'", "'if'", "'else'", 
			"'while'", "'for'", "'in'", "'class'", "'attributes'", "'methods'", "'write'", 
			"'read'", "'function'", "'return'", "'main'", "'var'", "'int'", "'float'", 
			"'char'", "'string'", "'bool'", "'void'", "'program'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, "INT", 
			"FLOAT", "CHAR", "BOOL", "WS", "OR", "AND", "GT", "LT", "EQ", "NEQ", 
			"MUL", "DIV", "ADD", "SUB", "SEMI", "LPAREN", "RPAREN", "IF", "ELSE", 
			"WHILE", "FOR", "IN", "CLASS", "ATTRIBUTES", "METHODS", "WRITE", "READ", 
			"FUNCTION", "RETURN", "MAIN", "VAR", "INT_TYPE", "FLOAT_TYPE", "CHAR_TYPE", 
			"STRING_TYPE", "BOOL_TYPE", "VOID", "PROGRAM", "ID"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Proyecto.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public ProyectoParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgramContext extends ParserRuleContext {
		public TerminalNode PROGRAM() { return getToken(ProyectoParser.PROGRAM, 0); }
		public TerminalNode ID() { return getToken(ProyectoParser.ID, 0); }
		public TerminalNode SEMI() { return getToken(ProyectoParser.SEMI, 0); }
		public MainContext main() {
			return getRuleContext(MainContext.class,0);
		}
		public TerminalNode EOF() { return getToken(ProyectoParser.EOF, 0); }
		public List<ClassDefContext> classDef() {
			return getRuleContexts(ClassDefContext.class);
		}
		public ClassDefContext classDef(int i) {
			return getRuleContext(ClassDefContext.class,i);
		}
		public List<VarsDecContext> varsDec() {
			return getRuleContexts(VarsDecContext.class);
		}
		public VarsDecContext varsDec(int i) {
			return getRuleContext(VarsDecContext.class,i);
		}
		public List<FunctionsContext> functions() {
			return getRuleContexts(FunctionsContext.class);
		}
		public FunctionsContext functions(int i) {
			return getRuleContext(FunctionsContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(88);
			match(PROGRAM);
			setState(89);
			match(ID);
			setState(90);
			match(SEMI);
			setState(94);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CLASS) {
				{
				{
				setState(91);
				classDef();
				}
				}
				setState(96);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(100);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(97);
				varsDec();
				}
				}
				setState(102);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(106);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FUNCTION) {
				{
				{
				setState(103);
				functions();
				}
				}
				setState(108);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(109);
			main();
			setState(110);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ClassDefContext extends ParserRuleContext {
		public TerminalNode CLASS() { return getToken(ProyectoParser.CLASS, 0); }
		public List<TerminalNode> ID() { return getTokens(ProyectoParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(ProyectoParser.ID, i);
		}
		public ClassBlockContext classBlock() {
			return getRuleContext(ClassBlockContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(ProyectoParser.SEMI, 0); }
		public TerminalNode LT() { return getToken(ProyectoParser.LT, 0); }
		public TerminalNode GT() { return getToken(ProyectoParser.GT, 0); }
		public ClassDefContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classDef; }
	}

	public final ClassDefContext classDef() throws RecognitionException {
		ClassDefContext _localctx = new ClassDefContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_classDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(112);
			match(CLASS);
			setState(113);
			match(ID);
			setState(117);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(114);
				match(LT);
				setState(115);
				match(ID);
				setState(116);
				match(GT);
				}
			}

			setState(119);
			classBlock();
			setState(120);
			match(SEMI);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ClassBlockContext extends ParserRuleContext {
		public TerminalNode ATTRIBUTES() { return getToken(ProyectoParser.ATTRIBUTES, 0); }
		public TerminalNode METHODS() { return getToken(ProyectoParser.METHODS, 0); }
		public List<VarsDecContext> varsDec() {
			return getRuleContexts(VarsDecContext.class);
		}
		public VarsDecContext varsDec(int i) {
			return getRuleContext(VarsDecContext.class,i);
		}
		public List<FunctionsContext> functions() {
			return getRuleContexts(FunctionsContext.class);
		}
		public FunctionsContext functions(int i) {
			return getRuleContext(FunctionsContext.class,i);
		}
		public ClassBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classBlock; }
	}

	public final ClassBlockContext classBlock() throws RecognitionException {
		ClassBlockContext _localctx = new ClassBlockContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_classBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(122);
			match(T__0);
			setState(123);
			match(ATTRIBUTES);
			setState(127);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(124);
				varsDec();
				}
				}
				setState(129);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(130);
			match(METHODS);
			setState(134);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FUNCTION) {
				{
				{
				setState(131);
				functions();
				}
				}
				setState(136);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(137);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VarsDecContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(ProyectoParser.VAR, 0); }
		public TerminalNode ID() { return getToken(ProyectoParser.ID, 0); }
		public VarsTypeInitContext varsTypeInit() {
			return getRuleContext(VarsTypeInitContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(ProyectoParser.SEMI, 0); }
		public List<TerminalNode> INT() { return getTokens(ProyectoParser.INT); }
		public TerminalNode INT(int i) {
			return getToken(ProyectoParser.INT, i);
		}
		public VarsDecContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varsDec; }
	}

	public final VarsDecContext varsDec() throws RecognitionException {
		VarsDecContext _localctx = new VarsDecContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_varsDec);
		try {
			setState(165);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(139);
				match(VAR);
				setState(140);
				match(ID);
				setState(141);
				match(T__2);
				setState(142);
				varsTypeInit();
				setState(143);
				match(SEMI);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(145);
				match(VAR);
				setState(146);
				match(ID);
				setState(147);
				match(T__3);
				setState(148);
				match(INT);
				setState(149);
				match(T__4);
				setState(150);
				match(T__2);
				setState(151);
				varsTypeInit();
				setState(152);
				match(SEMI);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(154);
				match(VAR);
				setState(155);
				match(ID);
				setState(156);
				match(T__3);
				setState(157);
				match(INT);
				setState(158);
				match(T__5);
				setState(159);
				match(INT);
				setState(160);
				match(T__4);
				setState(161);
				match(T__2);
				setState(162);
				varsTypeInit();
				setState(163);
				match(SEMI);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VarsTypeInitContext extends ParserRuleContext {
		public TypeRuleContext typeRule() {
			return getRuleContext(TypeRuleContext.class,0);
		}
		public TerminalNode ID() { return getToken(ProyectoParser.ID, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public VarsTypeInitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varsTypeInit; }
	}

	public final VarsTypeInitContext varsTypeInit() throws RecognitionException {
		VarsTypeInitContext _localctx = new VarsTypeInitContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_varsTypeInit);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(169);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_TYPE:
			case FLOAT_TYPE:
			case CHAR_TYPE:
			case STRING_TYPE:
			case BOOL_TYPE:
				{
				setState(167);
				typeRule();
				}
				break;
			case ID:
				{
				setState(168);
				match(ID);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(173);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(171);
				match(T__6);
				setState(172);
				exp();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VarsContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(ProyectoParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(ProyectoParser.ID, i);
		}
		public List<ExpContext> exp() {
			return getRuleContexts(ExpContext.class);
		}
		public ExpContext exp(int i) {
			return getRuleContext(ExpContext.class,i);
		}
		public VarsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vars; }
	}

	public final VarsContext vars() throws RecognitionException {
		VarsContext _localctx = new VarsContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_vars);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(175);
			match(ID);
			setState(178);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__7) {
				{
				setState(176);
				match(T__7);
				setState(177);
				match(ID);
				}
			}

			setState(184);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				{
				setState(180);
				match(T__3);
				setState(181);
				exp();
				setState(182);
				match(T__4);
				}
				break;
			}
			setState(190);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(186);
				match(T__3);
				setState(187);
				exp();
				setState(188);
				match(T__4);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionsContext extends ParserRuleContext {
		public TerminalNode FUNCTION() { return getToken(ProyectoParser.FUNCTION, 0); }
		public TerminalNode ID() { return getToken(ProyectoParser.ID, 0); }
		public TerminalNode LPAREN() { return getToken(ProyectoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(ProyectoParser.RPAREN, 0); }
		public FunctionBlockContext functionBlock() {
			return getRuleContext(FunctionBlockContext.class,0);
		}
		public TypeRuleContext typeRule() {
			return getRuleContext(TypeRuleContext.class,0);
		}
		public TerminalNode VOID() { return getToken(ProyectoParser.VOID, 0); }
		public ParametersContext parameters() {
			return getRuleContext(ParametersContext.class,0);
		}
		public FunctionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functions; }
	}

	public final FunctionsContext functions() throws RecognitionException {
		FunctionsContext _localctx = new FunctionsContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_functions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(192);
			match(FUNCTION);
			setState(193);
			match(ID);
			setState(194);
			match(LPAREN);
			setState(196);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(195);
				parameters();
				}
			}

			setState(198);
			match(RPAREN);
			setState(201);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_TYPE:
			case FLOAT_TYPE:
			case CHAR_TYPE:
			case STRING_TYPE:
			case BOOL_TYPE:
				{
				setState(199);
				typeRule();
				}
				break;
			case VOID:
				{
				setState(200);
				match(VOID);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(203);
			functionBlock();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParametersContext extends ParserRuleContext {
		public List<ParameterContext> parameter() {
			return getRuleContexts(ParameterContext.class);
		}
		public ParameterContext parameter(int i) {
			return getRuleContext(ParameterContext.class,i);
		}
		public ParametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameters; }
	}

	public final ParametersContext parameters() throws RecognitionException {
		ParametersContext _localctx = new ParametersContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_parameters);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(205);
			parameter();
			setState(210);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__8) {
				{
				{
				setState(206);
				match(T__8);
				setState(207);
				parameter();
				}
				}
				setState(212);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParameterContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ProyectoParser.ID, 0); }
		public TypeRuleContext typeRule() {
			return getRuleContext(TypeRuleContext.class,0);
		}
		public ParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameter; }
	}

	public final ParameterContext parameter() throws RecognitionException {
		ParameterContext _localctx = new ParameterContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_parameter);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(213);
			match(ID);
			setState(214);
			match(T__2);
			setState(215);
			typeRule();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionBlockContext extends ParserRuleContext {
		public ReturnRuleContext returnRule() {
			return getRuleContext(ReturnRuleContext.class,0);
		}
		public List<VarsDecContext> varsDec() {
			return getRuleContexts(VarsDecContext.class);
		}
		public VarsDecContext varsDec(int i) {
			return getRuleContext(VarsDecContext.class,i);
		}
		public List<StatutesContext> statutes() {
			return getRuleContexts(StatutesContext.class);
		}
		public StatutesContext statutes(int i) {
			return getRuleContext(StatutesContext.class,i);
		}
		public FunctionBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionBlock; }
	}

	public final FunctionBlockContext functionBlock() throws RecognitionException {
		FunctionBlockContext _localctx = new FunctionBlockContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_functionBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(217);
			match(T__0);
			setState(221);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(218);
				varsDec();
				}
				}
				setState(223);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(227);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << ID))) != 0)) {
				{
				{
				setState(224);
				statutes();
				}
				}
				setState(229);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(230);
			returnRule();
			setState(231);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ReturnRuleContext extends ParserRuleContext {
		public TerminalNode RETURN() { return getToken(ProyectoParser.RETURN, 0); }
		public TerminalNode SEMI() { return getToken(ProyectoParser.SEMI, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public ReturnRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnRule; }
	}

	public final ReturnRuleContext returnRule() throws RecognitionException {
		ReturnRuleContext _localctx = new ReturnRuleContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_returnRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(233);
			match(RETURN);
			setState(235);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << ID))) != 0)) {
				{
				setState(234);
				exp();
				}
			}

			setState(237);
			match(SEMI);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BlockContext extends ParserRuleContext {
		public List<StatutesContext> statutes() {
			return getRuleContexts(StatutesContext.class);
		}
		public StatutesContext statutes(int i) {
			return getRuleContext(StatutesContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(239);
			match(T__0);
			setState(243);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << ID))) != 0)) {
				{
				{
				setState(240);
				statutes();
				}
				}
				setState(245);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(246);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatutesContext extends ParserRuleContext {
		public AssignationContext assignation() {
			return getRuleContext(AssignationContext.class,0);
		}
		public ReadContext read() {
			return getRuleContext(ReadContext.class,0);
		}
		public WriteContext write() {
			return getRuleContext(WriteContext.class,0);
		}
		public ConditionalContext conditional() {
			return getRuleContext(ConditionalContext.class,0);
		}
		public ForLoopContext forLoop() {
			return getRuleContext(ForLoopContext.class,0);
		}
		public WhileLoopContext whileLoop() {
			return getRuleContext(WhileLoopContext.class,0);
		}
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public StatutesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statutes; }
	}

	public final StatutesContext statutes() throws RecognitionException {
		StatutesContext _localctx = new StatutesContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_statutes);
		try {
			setState(255);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(248);
				assignation();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(249);
				read();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(250);
				write();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(251);
				conditional();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(252);
				forLoop();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(253);
				whileLoop();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(254);
				exp();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignationContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ProyectoParser.ID, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(ProyectoParser.SEMI, 0); }
		public AssignationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignation; }
	}

	public final AssignationContext assignation() throws RecognitionException {
		AssignationContext _localctx = new AssignationContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_assignation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(257);
			match(ID);
			setState(258);
			match(T__6);
			setState(259);
			exp();
			setState(260);
			match(SEMI);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionCallContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ProyectoParser.ID, 0); }
		public TerminalNode LPAREN() { return getToken(ProyectoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(ProyectoParser.RPAREN, 0); }
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public FunctionCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionCall; }
	}

	public final FunctionCallContext functionCall() throws RecognitionException {
		FunctionCallContext _localctx = new FunctionCallContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_functionCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(262);
			match(ID);
			setState(263);
			match(LPAREN);
			setState(265);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << ID))) != 0)) {
				{
				setState(264);
				arguments();
				}
			}

			setState(267);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArgumentsContext extends ParserRuleContext {
		public List<ArgumentContext> argument() {
			return getRuleContexts(ArgumentContext.class);
		}
		public ArgumentContext argument(int i) {
			return getRuleContext(ArgumentContext.class,i);
		}
		public ArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arguments; }
	}

	public final ArgumentsContext arguments() throws RecognitionException {
		ArgumentsContext _localctx = new ArgumentsContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_arguments);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(269);
			argument();
			setState(274);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__8) {
				{
				{
				setState(270);
				match(T__8);
				setState(271);
				argument();
				}
				}
				setState(276);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArgumentContext extends ParserRuleContext {
		public VarsContext vars() {
			return getRuleContext(VarsContext.class,0);
		}
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public ArgumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argument; }
	}

	public final ArgumentContext argument() throws RecognitionException {
		ArgumentContext _localctx = new ArgumentContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_argument);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(279);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				{
				setState(277);
				vars();
				}
				break;
			case 2:
				{
				setState(278);
				exp();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MethodCallContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(ProyectoParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(ProyectoParser.ID, i);
		}
		public TerminalNode LPAREN() { return getToken(ProyectoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(ProyectoParser.RPAREN, 0); }
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public MethodCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_methodCall; }
	}

	public final MethodCallContext methodCall() throws RecognitionException {
		MethodCallContext _localctx = new MethodCallContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_methodCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(281);
			match(ID);
			setState(282);
			match(T__7);
			setState(283);
			match(ID);
			setState(284);
			match(LPAREN);
			setState(286);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << ID))) != 0)) {
				{
				setState(285);
				arguments();
				}
			}

			setState(288);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CallContext extends ParserRuleContext {
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public MethodCallContext methodCall() {
			return getRuleContext(MethodCallContext.class,0);
		}
		public CallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_call; }
	}

	public final CallContext call() throws RecognitionException {
		CallContext _localctx = new CallContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_call);
		try {
			setState(292);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(290);
				functionCall();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(291);
				methodCall();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ReadContext extends ParserRuleContext {
		public TerminalNode READ() { return getToken(ProyectoParser.READ, 0); }
		public TerminalNode LPAREN() { return getToken(ProyectoParser.LPAREN, 0); }
		public List<VarsContext> vars() {
			return getRuleContexts(VarsContext.class);
		}
		public VarsContext vars(int i) {
			return getRuleContext(VarsContext.class,i);
		}
		public TerminalNode RPAREN() { return getToken(ProyectoParser.RPAREN, 0); }
		public TerminalNode SEMI() { return getToken(ProyectoParser.SEMI, 0); }
		public ReadContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_read; }
	}

	public final ReadContext read() throws RecognitionException {
		ReadContext _localctx = new ReadContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_read);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(294);
			match(READ);
			setState(295);
			match(LPAREN);
			setState(296);
			vars();
			setState(301);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__8) {
				{
				{
				setState(297);
				match(T__8);
				setState(298);
				vars();
				}
				}
				setState(303);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(304);
			match(RPAREN);
			setState(305);
			match(SEMI);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class WriteContext extends ParserRuleContext {
		public TerminalNode WRITE() { return getToken(ProyectoParser.WRITE, 0); }
		public TerminalNode LPAREN() { return getToken(ProyectoParser.LPAREN, 0); }
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(ProyectoParser.RPAREN, 0); }
		public TerminalNode SEMI() { return getToken(ProyectoParser.SEMI, 0); }
		public WriteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_write; }
	}

	public final WriteContext write() throws RecognitionException {
		WriteContext _localctx = new WriteContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_write);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(307);
			match(WRITE);
			setState(308);
			match(LPAREN);
			setState(309);
			arguments();
			setState(310);
			match(RPAREN);
			setState(311);
			match(SEMI);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConditionalContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(ProyectoParser.IF, 0); }
		public TerminalNode LPAREN() { return getToken(ProyectoParser.LPAREN, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(ProyectoParser.RPAREN, 0); }
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(ProyectoParser.ELSE, 0); }
		public ConditionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditional; }
	}

	public final ConditionalContext conditional() throws RecognitionException {
		ConditionalContext _localctx = new ConditionalContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_conditional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(313);
			match(IF);
			setState(314);
			match(LPAREN);
			setState(315);
			exp();
			setState(316);
			match(RPAREN);
			setState(317);
			block();
			setState(320);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(318);
				match(ELSE);
				setState(319);
				block();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ForLoopContext extends ParserRuleContext {
		public TerminalNode FOR() { return getToken(ProyectoParser.FOR, 0); }
		public TerminalNode ID() { return getToken(ProyectoParser.ID, 0); }
		public List<ExpContext> exp() {
			return getRuleContexts(ExpContext.class);
		}
		public ExpContext exp(int i) {
			return getRuleContext(ExpContext.class,i);
		}
		public TerminalNode IN() { return getToken(ProyectoParser.IN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ForLoopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forLoop; }
	}

	public final ForLoopContext forLoop() throws RecognitionException {
		ForLoopContext _localctx = new ForLoopContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_forLoop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(322);
			match(FOR);
			setState(323);
			match(ID);
			setState(324);
			match(T__6);
			setState(325);
			exp();
			setState(326);
			match(IN);
			setState(327);
			exp();
			setState(328);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class WhileLoopContext extends ParserRuleContext {
		public TerminalNode WHILE() { return getToken(ProyectoParser.WHILE, 0); }
		public TerminalNode LPAREN() { return getToken(ProyectoParser.LPAREN, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(ProyectoParser.RPAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public WhileLoopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whileLoop; }
	}

	public final WhileLoopContext whileLoop() throws RecognitionException {
		WhileLoopContext _localctx = new WhileLoopContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_whileLoop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(330);
			match(WHILE);
			setState(331);
			match(LPAREN);
			setState(332);
			exp();
			setState(333);
			match(RPAREN);
			setState(334);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpContext extends ParserRuleContext {
		public T_expContext t_exp() {
			return getRuleContext(T_expContext.class,0);
		}
		public List<Exp2Context> exp2() {
			return getRuleContexts(Exp2Context.class);
		}
		public Exp2Context exp2(int i) {
			return getRuleContext(Exp2Context.class,i);
		}
		public ExpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exp; }
	}

	public final ExpContext exp() throws RecognitionException {
		ExpContext _localctx = new ExpContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(336);
			t_exp();
			setState(340);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OR) {
				{
				{
				setState(337);
				exp2();
				}
				}
				setState(342);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Exp2Context extends ParserRuleContext {
		public TerminalNode OR() { return getToken(ProyectoParser.OR, 0); }
		public T_expContext t_exp() {
			return getRuleContext(T_expContext.class,0);
		}
		public Exp2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exp2; }
	}

	public final Exp2Context exp2() throws RecognitionException {
		Exp2Context _localctx = new Exp2Context(_ctx, getState());
		enterRule(_localctx, 50, RULE_exp2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(343);
			match(OR);
			setState(344);
			t_exp();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class T_expContext extends ParserRuleContext {
		public G_expContext g_exp() {
			return getRuleContext(G_expContext.class,0);
		}
		public List<T_exp2Context> t_exp2() {
			return getRuleContexts(T_exp2Context.class);
		}
		public T_exp2Context t_exp2(int i) {
			return getRuleContext(T_exp2Context.class,i);
		}
		public T_expContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_t_exp; }
	}

	public final T_expContext t_exp() throws RecognitionException {
		T_expContext _localctx = new T_expContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_t_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(346);
			g_exp();
			setState(350);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==AND) {
				{
				{
				setState(347);
				t_exp2();
				}
				}
				setState(352);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class T_exp2Context extends ParserRuleContext {
		public TerminalNode AND() { return getToken(ProyectoParser.AND, 0); }
		public G_expContext g_exp() {
			return getRuleContext(G_expContext.class,0);
		}
		public T_exp2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_t_exp2; }
	}

	public final T_exp2Context t_exp2() throws RecognitionException {
		T_exp2Context _localctx = new T_exp2Context(_ctx, getState());
		enterRule(_localctx, 54, RULE_t_exp2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(353);
			match(AND);
			setState(354);
			g_exp();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class G_expContext extends ParserRuleContext {
		public M_expContext m_exp() {
			return getRuleContext(M_expContext.class,0);
		}
		public G_exp2Context g_exp2() {
			return getRuleContext(G_exp2Context.class,0);
		}
		public G_expContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_g_exp; }
	}

	public final G_expContext g_exp() throws RecognitionException {
		G_expContext _localctx = new G_expContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_g_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(356);
			m_exp();
			setState(358);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << GT) | (1L << LT) | (1L << EQ) | (1L << NEQ))) != 0)) {
				{
				setState(357);
				g_exp2();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class G_exp2Context extends ParserRuleContext {
		public RelopContext relop() {
			return getRuleContext(RelopContext.class,0);
		}
		public M_expContext m_exp() {
			return getRuleContext(M_expContext.class,0);
		}
		public G_exp2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_g_exp2; }
	}

	public final G_exp2Context g_exp2() throws RecognitionException {
		G_exp2Context _localctx = new G_exp2Context(_ctx, getState());
		enterRule(_localctx, 58, RULE_g_exp2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(360);
			relop();
			setState(361);
			m_exp();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class M_expContext extends ParserRuleContext {
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public List<M_exp2Context> m_exp2() {
			return getRuleContexts(M_exp2Context.class);
		}
		public M_exp2Context m_exp2(int i) {
			return getRuleContext(M_exp2Context.class,i);
		}
		public M_expContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_m_exp; }
	}

	public final M_expContext m_exp() throws RecognitionException {
		M_expContext _localctx = new M_expContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_m_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			term();
			setState(367);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ADD || _la==SUB) {
				{
				{
				setState(364);
				m_exp2();
				}
				}
				setState(369);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class M_exp2Context extends ParserRuleContext {
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public TerminalNode ADD() { return getToken(ProyectoParser.ADD, 0); }
		public TerminalNode SUB() { return getToken(ProyectoParser.SUB, 0); }
		public M_exp2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_m_exp2; }
	}

	public final M_exp2Context m_exp2() throws RecognitionException {
		M_exp2Context _localctx = new M_exp2Context(_ctx, getState());
		enterRule(_localctx, 62, RULE_m_exp2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(370);
			_la = _input.LA(1);
			if ( !(_la==ADD || _la==SUB) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(371);
			term();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TermContext extends ParserRuleContext {
		public FactorContext factor() {
			return getRuleContext(FactorContext.class,0);
		}
		public List<Term2Context> term2() {
			return getRuleContexts(Term2Context.class);
		}
		public Term2Context term2(int i) {
			return getRuleContext(Term2Context.class,i);
		}
		public TermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_term; }
	}

	public final TermContext term() throws RecognitionException {
		TermContext _localctx = new TermContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_term);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(373);
			factor();
			setState(377);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==MUL || _la==DIV) {
				{
				{
				setState(374);
				term2();
				}
				}
				setState(379);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Term2Context extends ParserRuleContext {
		public FactorContext factor() {
			return getRuleContext(FactorContext.class,0);
		}
		public TerminalNode MUL() { return getToken(ProyectoParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(ProyectoParser.DIV, 0); }
		public Term2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_term2; }
	}

	public final Term2Context term2() throws RecognitionException {
		Term2Context _localctx = new Term2Context(_ctx, getState());
		enterRule(_localctx, 66, RULE_term2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(380);
			_la = _input.LA(1);
			if ( !(_la==MUL || _la==DIV) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(381);
			factor();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FactorContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(ProyectoParser.LPAREN, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(ProyectoParser.RPAREN, 0); }
		public VarCteContext varCte() {
			return getRuleContext(VarCteContext.class,0);
		}
		public VarsContext vars() {
			return getRuleContext(VarsContext.class,0);
		}
		public CallContext call() {
			return getRuleContext(CallContext.class,0);
		}
		public FactorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_factor; }
	}

	public final FactorContext factor() throws RecognitionException {
		FactorContext _localctx = new FactorContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_factor);
		try {
			setState(390);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(383);
				match(LPAREN);
				setState(384);
				exp();
				setState(385);
				match(RPAREN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(387);
				varCte();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(388);
				vars();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(389);
				call();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VarCteContext extends ParserRuleContext {
		public Cte_iContext cte_i() {
			return getRuleContext(Cte_iContext.class,0);
		}
		public Cte_fContext cte_f() {
			return getRuleContext(Cte_fContext.class,0);
		}
		public Cte_cContext cte_c() {
			return getRuleContext(Cte_cContext.class,0);
		}
		public Cte_sContext cte_s() {
			return getRuleContext(Cte_sContext.class,0);
		}
		public Cte_bContext cte_b() {
			return getRuleContext(Cte_bContext.class,0);
		}
		public VarCteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varCte; }
	}

	public final VarCteContext varCte() throws RecognitionException {
		VarCteContext _localctx = new VarCteContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_varCte);
		try {
			setState(397);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(392);
				cte_i();
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(393);
				cte_f();
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 3);
				{
				setState(394);
				cte_c();
				}
				break;
			case T__9:
				enterOuterAlt(_localctx, 4);
				{
				setState(395);
				cte_s();
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 5);
				{
				setState(396);
				cte_b();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Cte_iContext extends ParserRuleContext {
		public TerminalNode INT() { return getToken(ProyectoParser.INT, 0); }
		public Cte_iContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cte_i; }
	}

	public final Cte_iContext cte_i() throws RecognitionException {
		Cte_iContext _localctx = new Cte_iContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_cte_i);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(399);
			match(INT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Cte_fContext extends ParserRuleContext {
		public TerminalNode FLOAT() { return getToken(ProyectoParser.FLOAT, 0); }
		public Cte_fContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cte_f; }
	}

	public final Cte_fContext cte_f() throws RecognitionException {
		Cte_fContext _localctx = new Cte_fContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_cte_f);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(401);
			match(FLOAT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Cte_cContext extends ParserRuleContext {
		public TerminalNode CHAR() { return getToken(ProyectoParser.CHAR, 0); }
		public Cte_cContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cte_c; }
	}

	public final Cte_cContext cte_c() throws RecognitionException {
		Cte_cContext _localctx = new Cte_cContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_cte_c);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(403);
			match(CHAR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Cte_bContext extends ParserRuleContext {
		public TerminalNode BOOL() { return getToken(ProyectoParser.BOOL, 0); }
		public Cte_bContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cte_b; }
	}

	public final Cte_bContext cte_b() throws RecognitionException {
		Cte_bContext _localctx = new Cte_bContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_cte_b);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(405);
			match(BOOL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Cte_sContext extends ParserRuleContext {
		public Cte_sContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cte_s; }
	}

	public final Cte_sContext cte_s() throws RecognitionException {
		Cte_sContext _localctx = new Cte_sContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_cte_s);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(407);
			match(T__9);
			setState(411);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,34,_ctx);
			while ( _alt!=1 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1+1 ) {
					{
					{
					setState(408);
					matchWildcard();
					}
					} 
				}
				setState(413);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,34,_ctx);
			}
			setState(414);
			match(T__9);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MainContext extends ParserRuleContext {
		public TerminalNode MAIN() { return getToken(ProyectoParser.MAIN, 0); }
		public TerminalNode LPAREN() { return getToken(ProyectoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(ProyectoParser.RPAREN, 0); }
		public FunctionBlockContext functionBlock() {
			return getRuleContext(FunctionBlockContext.class,0);
		}
		public MainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_main; }
	}

	public final MainContext main() throws RecognitionException {
		MainContext _localctx = new MainContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_main);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(416);
			match(MAIN);
			setState(417);
			match(LPAREN);
			setState(418);
			match(RPAREN);
			setState(419);
			functionBlock();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeRuleContext extends ParserRuleContext {
		public TerminalNode INT_TYPE() { return getToken(ProyectoParser.INT_TYPE, 0); }
		public TerminalNode FLOAT_TYPE() { return getToken(ProyectoParser.FLOAT_TYPE, 0); }
		public TerminalNode CHAR_TYPE() { return getToken(ProyectoParser.CHAR_TYPE, 0); }
		public TerminalNode STRING_TYPE() { return getToken(ProyectoParser.STRING_TYPE, 0); }
		public TerminalNode BOOL_TYPE() { return getToken(ProyectoParser.BOOL_TYPE, 0); }
		public TypeRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeRule; }
	}

	public final TypeRuleContext typeRule() throws RecognitionException {
		TypeRuleContext _localctx = new TypeRuleContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_typeRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(421);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT_TYPE) | (1L << FLOAT_TYPE) | (1L << CHAR_TYPE) | (1L << STRING_TYPE) | (1L << BOOL_TYPE))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RelopContext extends ParserRuleContext {
		public TerminalNode GT() { return getToken(ProyectoParser.GT, 0); }
		public TerminalNode LT() { return getToken(ProyectoParser.LT, 0); }
		public TerminalNode EQ() { return getToken(ProyectoParser.EQ, 0); }
		public TerminalNode NEQ() { return getToken(ProyectoParser.NEQ, 0); }
		public RelopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relop; }
	}

	public final RelopContext relop() throws RecognitionException {
		RelopContext _localctx = new RelopContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_relop);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(423);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << GT) | (1L << LT) | (1L << EQ) | (1L << NEQ))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\64\u01ac\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\3\2\3\2\3\2\3\2\7\2_\n\2\f\2\16\2b\13\2\3\2\7\2e\n\2\f\2\16"+
		"\2h\13\2\3\2\7\2k\n\2\f\2\16\2n\13\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\5"+
		"\3x\n\3\3\3\3\3\3\3\3\4\3\4\3\4\7\4\u0080\n\4\f\4\16\4\u0083\13\4\3\4"+
		"\3\4\7\4\u0087\n\4\f\4\16\4\u008a\13\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\5\5\u00a8\n\5\3\6\3\6\5\6\u00ac\n\6\3\6\3\6\5\6\u00b0\n\6\3"+
		"\7\3\7\3\7\5\7\u00b5\n\7\3\7\3\7\3\7\3\7\5\7\u00bb\n\7\3\7\3\7\3\7\3\7"+
		"\5\7\u00c1\n\7\3\b\3\b\3\b\3\b\5\b\u00c7\n\b\3\b\3\b\3\b\5\b\u00cc\n\b"+
		"\3\b\3\b\3\t\3\t\3\t\7\t\u00d3\n\t\f\t\16\t\u00d6\13\t\3\n\3\n\3\n\3\n"+
		"\3\13\3\13\7\13\u00de\n\13\f\13\16\13\u00e1\13\13\3\13\7\13\u00e4\n\13"+
		"\f\13\16\13\u00e7\13\13\3\13\3\13\3\13\3\f\3\f\5\f\u00ee\n\f\3\f\3\f\3"+
		"\r\3\r\7\r\u00f4\n\r\f\r\16\r\u00f7\13\r\3\r\3\r\3\16\3\16\3\16\3\16\3"+
		"\16\3\16\3\16\5\16\u0102\n\16\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20"+
		"\5\20\u010c\n\20\3\20\3\20\3\21\3\21\3\21\7\21\u0113\n\21\f\21\16\21\u0116"+
		"\13\21\3\22\3\22\5\22\u011a\n\22\3\23\3\23\3\23\3\23\3\23\5\23\u0121\n"+
		"\23\3\23\3\23\3\24\3\24\5\24\u0127\n\24\3\25\3\25\3\25\3\25\3\25\7\25"+
		"\u012e\n\25\f\25\16\25\u0131\13\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26"+
		"\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\5\27\u0143\n\27\3\30\3\30"+
		"\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32"+
		"\7\32\u0155\n\32\f\32\16\32\u0158\13\32\3\33\3\33\3\33\3\34\3\34\7\34"+
		"\u015f\n\34\f\34\16\34\u0162\13\34\3\35\3\35\3\35\3\36\3\36\5\36\u0169"+
		"\n\36\3\37\3\37\3\37\3 \3 \7 \u0170\n \f \16 \u0173\13 \3!\3!\3!\3\"\3"+
		"\"\7\"\u017a\n\"\f\"\16\"\u017d\13\"\3#\3#\3#\3$\3$\3$\3$\3$\3$\3$\5$"+
		"\u0189\n$\3%\3%\3%\3%\3%\5%\u0190\n%\3&\3&\3\'\3\'\3(\3(\3)\3)\3*\3*\7"+
		"*\u019c\n*\f*\16*\u019f\13*\3*\3*\3+\3+\3+\3+\3+\3,\3,\3-\3-\3-\3\u019d"+
		"\2.\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@B"+
		"DFHJLNPRTVX\2\6\3\2\32\33\3\2\30\31\3\2-\61\3\2\24\27\2\u01ad\2Z\3\2\2"+
		"\2\4r\3\2\2\2\6|\3\2\2\2\b\u00a7\3\2\2\2\n\u00ab\3\2\2\2\f\u00b1\3\2\2"+
		"\2\16\u00c2\3\2\2\2\20\u00cf\3\2\2\2\22\u00d7\3\2\2\2\24\u00db\3\2\2\2"+
		"\26\u00eb\3\2\2\2\30\u00f1\3\2\2\2\32\u0101\3\2\2\2\34\u0103\3\2\2\2\36"+
		"\u0108\3\2\2\2 \u010f\3\2\2\2\"\u0119\3\2\2\2$\u011b\3\2\2\2&\u0126\3"+
		"\2\2\2(\u0128\3\2\2\2*\u0135\3\2\2\2,\u013b\3\2\2\2.\u0144\3\2\2\2\60"+
		"\u014c\3\2\2\2\62\u0152\3\2\2\2\64\u0159\3\2\2\2\66\u015c\3\2\2\28\u0163"+
		"\3\2\2\2:\u0166\3\2\2\2<\u016a\3\2\2\2>\u016d\3\2\2\2@\u0174\3\2\2\2B"+
		"\u0177\3\2\2\2D\u017e\3\2\2\2F\u0188\3\2\2\2H\u018f\3\2\2\2J\u0191\3\2"+
		"\2\2L\u0193\3\2\2\2N\u0195\3\2\2\2P\u0197\3\2\2\2R\u0199\3\2\2\2T\u01a2"+
		"\3\2\2\2V\u01a7\3\2\2\2X\u01a9\3\2\2\2Z[\7\63\2\2[\\\7\64\2\2\\`\7\34"+
		"\2\2]_\5\4\3\2^]\3\2\2\2_b\3\2\2\2`^\3\2\2\2`a\3\2\2\2af\3\2\2\2b`\3\2"+
		"\2\2ce\5\b\5\2dc\3\2\2\2eh\3\2\2\2fd\3\2\2\2fg\3\2\2\2gl\3\2\2\2hf\3\2"+
		"\2\2ik\5\16\b\2ji\3\2\2\2kn\3\2\2\2lj\3\2\2\2lm\3\2\2\2mo\3\2\2\2nl\3"+
		"\2\2\2op\5T+\2pq\7\2\2\3q\3\3\2\2\2rs\7$\2\2sw\7\64\2\2tu\7\25\2\2uv\7"+
		"\64\2\2vx\7\24\2\2wt\3\2\2\2wx\3\2\2\2xy\3\2\2\2yz\5\6\4\2z{\7\34\2\2"+
		"{\5\3\2\2\2|}\7\3\2\2}\u0081\7%\2\2~\u0080\5\b\5\2\177~\3\2\2\2\u0080"+
		"\u0083\3\2\2\2\u0081\177\3\2\2\2\u0081\u0082\3\2\2\2\u0082\u0084\3\2\2"+
		"\2\u0083\u0081\3\2\2\2\u0084\u0088\7&\2\2\u0085\u0087\5\16\b\2\u0086\u0085"+
		"\3\2\2\2\u0087\u008a\3\2\2\2\u0088\u0086\3\2\2\2\u0088\u0089\3\2\2\2\u0089"+
		"\u008b\3\2\2\2\u008a\u0088\3\2\2\2\u008b\u008c\7\4\2\2\u008c\7\3\2\2\2"+
		"\u008d\u008e\7,\2\2\u008e\u008f\7\64\2\2\u008f\u0090\7\5\2\2\u0090\u0091"+
		"\5\n\6\2\u0091\u0092\7\34\2\2\u0092\u00a8\3\2\2\2\u0093\u0094\7,\2\2\u0094"+
		"\u0095\7\64\2\2\u0095\u0096\7\6\2\2\u0096\u0097\7\r\2\2\u0097\u0098\7"+
		"\7\2\2\u0098\u0099\7\5\2\2\u0099\u009a\5\n\6\2\u009a\u009b\7\34\2\2\u009b"+
		"\u00a8\3\2\2\2\u009c\u009d\7,\2\2\u009d\u009e\7\64\2\2\u009e\u009f\7\6"+
		"\2\2\u009f\u00a0\7\r\2\2\u00a0\u00a1\7\b\2\2\u00a1\u00a2\7\r\2\2\u00a2"+
		"\u00a3\7\7\2\2\u00a3\u00a4\7\5\2\2\u00a4\u00a5\5\n\6\2\u00a5\u00a6\7\34"+
		"\2\2\u00a6\u00a8\3\2\2\2\u00a7\u008d\3\2\2\2\u00a7\u0093\3\2\2\2\u00a7"+
		"\u009c\3\2\2\2\u00a8\t\3\2\2\2\u00a9\u00ac\5V,\2\u00aa\u00ac\7\64\2\2"+
		"\u00ab\u00a9\3\2\2\2\u00ab\u00aa\3\2\2\2\u00ac\u00af\3\2\2\2\u00ad\u00ae"+
		"\7\t\2\2\u00ae\u00b0\5\62\32\2\u00af\u00ad\3\2\2\2\u00af\u00b0\3\2\2\2"+
		"\u00b0\13\3\2\2\2\u00b1\u00b4\7\64\2\2\u00b2\u00b3\7\n\2\2\u00b3\u00b5"+
		"\7\64\2\2\u00b4\u00b2\3\2\2\2\u00b4\u00b5\3\2\2\2\u00b5\u00ba\3\2\2\2"+
		"\u00b6\u00b7\7\6\2\2\u00b7\u00b8\5\62\32\2\u00b8\u00b9\7\7\2\2\u00b9\u00bb"+
		"\3\2\2\2\u00ba\u00b6\3\2\2\2\u00ba\u00bb\3\2\2\2\u00bb\u00c0\3\2\2\2\u00bc"+
		"\u00bd\7\6\2\2\u00bd\u00be\5\62\32\2\u00be\u00bf\7\7\2\2\u00bf\u00c1\3"+
		"\2\2\2\u00c0\u00bc\3\2\2\2\u00c0\u00c1\3\2\2\2\u00c1\r\3\2\2\2\u00c2\u00c3"+
		"\7)\2\2\u00c3\u00c4\7\64\2\2\u00c4\u00c6\7\35\2\2\u00c5\u00c7\5\20\t\2"+
		"\u00c6\u00c5\3\2\2\2\u00c6\u00c7\3\2\2\2\u00c7\u00c8\3\2\2\2\u00c8\u00cb"+
		"\7\36\2\2\u00c9\u00cc\5V,\2\u00ca\u00cc\7\62\2\2\u00cb\u00c9\3\2\2\2\u00cb"+
		"\u00ca\3\2\2\2\u00cc\u00cd\3\2\2\2\u00cd\u00ce\5\24\13\2\u00ce\17\3\2"+
		"\2\2\u00cf\u00d4\5\22\n\2\u00d0\u00d1\7\13\2\2\u00d1\u00d3\5\22\n\2\u00d2"+
		"\u00d0\3\2\2\2\u00d3\u00d6\3\2\2\2\u00d4\u00d2\3\2\2\2\u00d4\u00d5\3\2"+
		"\2\2\u00d5\21\3\2\2\2\u00d6\u00d4\3\2\2\2\u00d7\u00d8\7\64\2\2\u00d8\u00d9"+
		"\7\5\2\2\u00d9\u00da\5V,\2\u00da\23\3\2\2\2\u00db\u00df\7\3\2\2\u00dc"+
		"\u00de\5\b\5\2\u00dd\u00dc\3\2\2\2\u00de\u00e1\3\2\2\2\u00df\u00dd\3\2"+
		"\2\2\u00df\u00e0\3\2\2\2\u00e0\u00e5\3\2\2\2\u00e1\u00df\3\2\2\2\u00e2"+
		"\u00e4\5\32\16\2\u00e3\u00e2\3\2\2\2\u00e4\u00e7\3\2\2\2\u00e5\u00e3\3"+
		"\2\2\2\u00e5\u00e6\3\2\2\2\u00e6\u00e8\3\2\2\2\u00e7\u00e5\3\2\2\2\u00e8"+
		"\u00e9\5\26\f\2\u00e9\u00ea\7\4\2\2\u00ea\25\3\2\2\2\u00eb\u00ed\7*\2"+
		"\2\u00ec\u00ee\5\62\32\2\u00ed\u00ec\3\2\2\2\u00ed\u00ee\3\2\2\2\u00ee"+
		"\u00ef\3\2\2\2\u00ef\u00f0\7\34\2\2\u00f0\27\3\2\2\2\u00f1\u00f5\7\3\2"+
		"\2\u00f2\u00f4\5\32\16\2\u00f3\u00f2\3\2\2\2\u00f4\u00f7\3\2\2\2\u00f5"+
		"\u00f3\3\2\2\2\u00f5\u00f6\3\2\2\2\u00f6\u00f8\3\2\2\2\u00f7\u00f5\3\2"+
		"\2\2\u00f8\u00f9\7\4\2\2\u00f9\31\3\2\2\2\u00fa\u0102\5\34\17\2\u00fb"+
		"\u0102\5(\25\2\u00fc\u0102\5*\26\2\u00fd\u0102\5,\27\2\u00fe\u0102\5."+
		"\30\2\u00ff\u0102\5\60\31\2\u0100\u0102\5\62\32\2\u0101\u00fa\3\2\2\2"+
		"\u0101\u00fb\3\2\2\2\u0101\u00fc\3\2\2\2\u0101\u00fd\3\2\2\2\u0101\u00fe"+
		"\3\2\2\2\u0101\u00ff\3\2\2\2\u0101\u0100\3\2\2\2\u0102\33\3\2\2\2\u0103"+
		"\u0104\7\64\2\2\u0104\u0105\7\t\2\2\u0105\u0106\5\62\32\2\u0106\u0107"+
		"\7\34\2\2\u0107\35\3\2\2\2\u0108\u0109\7\64\2\2\u0109\u010b\7\35\2\2\u010a"+
		"\u010c\5 \21\2\u010b\u010a\3\2\2\2\u010b\u010c\3\2\2\2\u010c\u010d\3\2"+
		"\2\2\u010d\u010e\7\36\2\2\u010e\37\3\2\2\2\u010f\u0114\5\"\22\2\u0110"+
		"\u0111\7\13\2\2\u0111\u0113\5\"\22\2\u0112\u0110\3\2\2\2\u0113\u0116\3"+
		"\2\2\2\u0114\u0112\3\2\2\2\u0114\u0115\3\2\2\2\u0115!\3\2\2\2\u0116\u0114"+
		"\3\2\2\2\u0117\u011a\5\f\7\2\u0118\u011a\5\62\32\2\u0119\u0117\3\2\2\2"+
		"\u0119\u0118\3\2\2\2\u011a#\3\2\2\2\u011b\u011c\7\64\2\2\u011c\u011d\7"+
		"\n\2\2\u011d\u011e\7\64\2\2\u011e\u0120\7\35\2\2\u011f\u0121\5 \21\2\u0120"+
		"\u011f\3\2\2\2\u0120\u0121\3\2\2\2\u0121\u0122\3\2\2\2\u0122\u0123\7\36"+
		"\2\2\u0123%\3\2\2\2\u0124\u0127\5\36\20\2\u0125\u0127\5$\23\2\u0126\u0124"+
		"\3\2\2\2\u0126\u0125\3\2\2\2\u0127\'\3\2\2\2\u0128\u0129\7(\2\2\u0129"+
		"\u012a\7\35\2\2\u012a\u012f\5\f\7\2\u012b\u012c\7\13\2\2\u012c\u012e\5"+
		"\f\7\2\u012d\u012b\3\2\2\2\u012e\u0131\3\2\2\2\u012f\u012d\3\2\2\2\u012f"+
		"\u0130\3\2\2\2\u0130\u0132\3\2\2\2\u0131\u012f\3\2\2\2\u0132\u0133\7\36"+
		"\2\2\u0133\u0134\7\34\2\2\u0134)\3\2\2\2\u0135\u0136\7\'\2\2\u0136\u0137"+
		"\7\35\2\2\u0137\u0138\5 \21\2\u0138\u0139\7\36\2\2\u0139\u013a\7\34\2"+
		"\2\u013a+\3\2\2\2\u013b\u013c\7\37\2\2\u013c\u013d\7\35\2\2\u013d\u013e"+
		"\5\62\32\2\u013e\u013f\7\36\2\2\u013f\u0142\5\30\r\2\u0140\u0141\7 \2"+
		"\2\u0141\u0143\5\30\r\2\u0142\u0140\3\2\2\2\u0142\u0143\3\2\2\2\u0143"+
		"-\3\2\2\2\u0144\u0145\7\"\2\2\u0145\u0146\7\64\2\2\u0146\u0147\7\t\2\2"+
		"\u0147\u0148\5\62\32\2\u0148\u0149\7#\2\2\u0149\u014a\5\62\32\2\u014a"+
		"\u014b\5\30\r\2\u014b/\3\2\2\2\u014c\u014d\7!\2\2\u014d\u014e\7\35\2\2"+
		"\u014e\u014f\5\62\32\2\u014f\u0150\7\36\2\2\u0150\u0151\5\30\r\2\u0151"+
		"\61\3\2\2\2\u0152\u0156\5\66\34\2\u0153\u0155\5\64\33\2\u0154\u0153\3"+
		"\2\2\2\u0155\u0158\3\2\2\2\u0156\u0154\3\2\2\2\u0156\u0157\3\2\2\2\u0157"+
		"\63\3\2\2\2\u0158\u0156\3\2\2\2\u0159\u015a\7\22\2\2\u015a\u015b\5\66"+
		"\34\2\u015b\65\3\2\2\2\u015c\u0160\5:\36\2\u015d\u015f\58\35\2\u015e\u015d"+
		"\3\2\2\2\u015f\u0162\3\2\2\2\u0160\u015e\3\2\2\2\u0160\u0161\3\2\2\2\u0161"+
		"\67\3\2\2\2\u0162\u0160\3\2\2\2\u0163\u0164\7\23\2\2\u0164\u0165\5:\36"+
		"\2\u01659\3\2\2\2\u0166\u0168\5> \2\u0167\u0169\5<\37\2\u0168\u0167\3"+
		"\2\2\2\u0168\u0169\3\2\2\2\u0169;\3\2\2\2\u016a\u016b\5X-\2\u016b\u016c"+
		"\5> \2\u016c=\3\2\2\2\u016d\u0171\5B\"\2\u016e\u0170\5@!\2\u016f\u016e"+
		"\3\2\2\2\u0170\u0173\3\2\2\2\u0171\u016f\3\2\2\2\u0171\u0172\3\2\2\2\u0172"+
		"?\3\2\2\2\u0173\u0171\3\2\2\2\u0174\u0175\t\2\2\2\u0175\u0176\5B\"\2\u0176"+
		"A\3\2\2\2\u0177\u017b\5F$\2\u0178\u017a\5D#\2\u0179\u0178\3\2\2\2\u017a"+
		"\u017d\3\2\2\2\u017b\u0179\3\2\2\2\u017b\u017c\3\2\2\2\u017cC\3\2\2\2"+
		"\u017d\u017b\3\2\2\2\u017e\u017f\t\3\2\2\u017f\u0180\5F$\2\u0180E\3\2"+
		"\2\2\u0181\u0182\7\35\2\2\u0182\u0183\5\62\32\2\u0183\u0184\7\36\2\2\u0184"+
		"\u0189\3\2\2\2\u0185\u0189\5H%\2\u0186\u0189\5\f\7\2\u0187\u0189\5&\24"+
		"\2\u0188\u0181\3\2\2\2\u0188\u0185\3\2\2\2\u0188\u0186\3\2\2\2\u0188\u0187"+
		"\3\2\2\2\u0189G\3\2\2\2\u018a\u0190\5J&\2\u018b\u0190\5L\'\2\u018c\u0190"+
		"\5N(\2\u018d\u0190\5R*\2\u018e\u0190\5P)\2\u018f\u018a\3\2\2\2\u018f\u018b"+
		"\3\2\2\2\u018f\u018c\3\2\2\2\u018f\u018d\3\2\2\2\u018f\u018e\3\2\2\2\u0190"+
		"I\3\2\2\2\u0191\u0192\7\r\2\2\u0192K\3\2\2\2\u0193\u0194\7\16\2\2\u0194"+
		"M\3\2\2\2\u0195\u0196\7\17\2\2\u0196O\3\2\2\2\u0197\u0198\7\20\2\2\u0198"+
		"Q\3\2\2\2\u0199\u019d\7\f\2\2\u019a\u019c\13\2\2\2\u019b\u019a\3\2\2\2"+
		"\u019c\u019f\3\2\2\2\u019d\u019e\3\2\2\2\u019d\u019b\3\2\2\2\u019e\u01a0"+
		"\3\2\2\2\u019f\u019d\3\2\2\2\u01a0\u01a1\7\f\2\2\u01a1S\3\2\2\2\u01a2"+
		"\u01a3\7+\2\2\u01a3\u01a4\7\35\2\2\u01a4\u01a5\7\36\2\2\u01a5\u01a6\5"+
		"\24\13\2\u01a6U\3\2\2\2\u01a7\u01a8\t\4\2\2\u01a8W\3\2\2\2\u01a9\u01aa"+
		"\t\5\2\2\u01aaY\3\2\2\2%`flw\u0081\u0088\u00a7\u00ab\u00af\u00b4\u00ba"+
		"\u00c0\u00c6\u00cb\u00d4\u00df\u00e5\u00ed\u00f5\u0101\u010b\u0114\u0119"+
		"\u0120\u0126\u012f\u0142\u0156\u0160\u0168\u0171\u017b\u0188\u018f\u019d";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}