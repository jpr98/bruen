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
		INT=10, FLOAT=11, CHAR=12, BOOL=13, WS=14, OR=15, AND=16, GT=17, LT=18, 
		EQ=19, NEQ=20, MUL=21, DIV=22, ADD=23, SUB=24, SEMI=25, ASSIGN=26, LPAREN=27, 
		RPAREN=28, IF=29, ELSE=30, WHILE=31, FOR=32, IN=33, CLASS=34, ATTRIBUTES=35, 
		METHODS=36, WRITE=37, READ=38, FUNCTION=39, RETURN=40, MAIN=41, VAR=42, 
		INT_TYPE=43, FLOAT_TYPE=44, CHAR_TYPE=45, BOOL_TYPE=46, VOID=47, PROGRAM=48, 
		ID=49;
	public static final int
		RULE_program = 0, RULE_classDef = 1, RULE_classBlock = 2, RULE_varsDec = 3, 
		RULE_varsTypeInit = 4, RULE_varsTypeInit2 = 5, RULE_vars = 6, RULE_functions = 7, 
		RULE_parameters = 8, RULE_parameter = 9, RULE_functionBlock = 10, RULE_returnRule = 11, 
		RULE_block = 12, RULE_statutes = 13, RULE_assignation = 14, RULE_functionCall = 15, 
		RULE_arguments = 16, RULE_arguments2 = 17, RULE_methodCall = 18, RULE_call = 19, 
		RULE_read = 20, RULE_read2 = 21, RULE_read3 = 22, RULE_write = 23, RULE_conditional = 24, 
		RULE_conditional2 = 25, RULE_conditional3 = 26, RULE_conditional4 = 27, 
		RULE_forLoop = 28, RULE_forLoop2 = 29, RULE_forLoop3 = 30, RULE_whileLoop = 31, 
		RULE_whileLoop2 = 32, RULE_expression = 33, RULE_exp = 34, RULE_exp2 = 35, 
		RULE_t_exp = 36, RULE_t_exp2 = 37, RULE_g_exp = 38, RULE_g_exp2 = 39, 
		RULE_m_exp = 40, RULE_m_exp2 = 41, RULE_term = 42, RULE_term2 = 43, RULE_factor = 44, 
		RULE_factor2 = 45, RULE_varCte = 46, RULE_cte_i = 47, RULE_cte_f = 48, 
		RULE_cte_c = 49, RULE_cte_b = 50, RULE_cte_s = 51, RULE_main = 52, RULE_mainBlock = 53, 
		RULE_typeRule = 54, RULE_relop = 55;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "classDef", "classBlock", "varsDec", "varsTypeInit", "varsTypeInit2", 
			"vars", "functions", "parameters", "parameter", "functionBlock", "returnRule", 
			"block", "statutes", "assignation", "functionCall", "arguments", "arguments2", 
			"methodCall", "call", "read", "read2", "read3", "write", "conditional", 
			"conditional2", "conditional3", "conditional4", "forLoop", "forLoop2", 
			"forLoop3", "whileLoop", "whileLoop2", "expression", "exp", "exp2", "t_exp", 
			"t_exp2", "g_exp", "g_exp2", "m_exp", "m_exp2", "term", "term2", "factor", 
			"factor2", "varCte", "cte_i", "cte_f", "cte_c", "cte_b", "cte_s", "main", 
			"mainBlock", "typeRule", "relop"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'{'", "'}'", "':'", "'['", "']'", "']['", "'.'", "','", "'\"'", 
			null, null, null, null, null, "'||'", "'&&'", "'>'", "'<'", "'=='", "'!='", 
			"'*'", "'/'", "'+'", "'-'", "';'", "'='", "'('", "')'", "'if'", "'else'", 
			"'while'", "'for'", "'in'", "'class'", "'attributes'", "'methods'", "'write'", 
			"'read'", "'function'", "'return'", "'main'", "'var'", "'int'", "'float'", 
			"'char'", "'bool'", "'void'", "'program'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, "INT", "FLOAT", 
			"CHAR", "BOOL", "WS", "OR", "AND", "GT", "LT", "EQ", "NEQ", "MUL", "DIV", 
			"ADD", "SUB", "SEMI", "ASSIGN", "LPAREN", "RPAREN", "IF", "ELSE", "WHILE", 
			"FOR", "IN", "CLASS", "ATTRIBUTES", "METHODS", "WRITE", "READ", "FUNCTION", 
			"RETURN", "MAIN", "VAR", "INT_TYPE", "FLOAT_TYPE", "CHAR_TYPE", "BOOL_TYPE", 
			"VOID", "PROGRAM", "ID"
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
			setState(112);
			match(PROGRAM);
			setState(113);
			match(ID);
			setState(114);
			match(SEMI);
			setState(118);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CLASS) {
				{
				{
				setState(115);
				classDef();
				}
				}
				setState(120);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(124);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(121);
				varsDec();
				}
				}
				setState(126);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(130);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FUNCTION) {
				{
				{
				setState(127);
				functions();
				}
				}
				setState(132);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(133);
			main();
			setState(134);
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
			setState(136);
			match(CLASS);
			setState(137);
			match(ID);
			setState(141);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(138);
				match(LT);
				setState(139);
				match(ID);
				setState(140);
				match(GT);
				}
			}

			setState(143);
			classBlock();
			setState(144);
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
			setState(146);
			match(T__0);
			setState(147);
			match(ATTRIBUTES);
			setState(151);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(148);
				varsDec();
				}
				}
				setState(153);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(154);
			match(METHODS);
			setState(158);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FUNCTION) {
				{
				{
				setState(155);
				functions();
				}
				}
				setState(160);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(161);
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
			setState(189);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(163);
				match(VAR);
				setState(164);
				match(ID);
				setState(165);
				match(T__2);
				setState(166);
				varsTypeInit();
				setState(167);
				match(SEMI);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(169);
				match(VAR);
				setState(170);
				match(ID);
				setState(171);
				match(T__3);
				setState(172);
				match(INT);
				setState(173);
				match(T__4);
				setState(174);
				match(T__2);
				setState(175);
				varsTypeInit();
				setState(176);
				match(SEMI);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(178);
				match(VAR);
				setState(179);
				match(ID);
				setState(180);
				match(T__3);
				setState(181);
				match(INT);
				setState(182);
				match(T__5);
				setState(183);
				match(INT);
				setState(184);
				match(T__4);
				setState(185);
				match(T__2);
				setState(186);
				varsTypeInit();
				setState(187);
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
		public VarsTypeInit2Context varsTypeInit2() {
			return getRuleContext(VarsTypeInit2Context.class,0);
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
			setState(193);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_TYPE:
			case FLOAT_TYPE:
			case CHAR_TYPE:
			case BOOL_TYPE:
				{
				setState(191);
				typeRule();
				}
				break;
			case ID:
				{
				setState(192);
				match(ID);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(196);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ASSIGN) {
				{
				setState(195);
				varsTypeInit2();
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

	public static class VarsTypeInit2Context extends ParserRuleContext {
		public TerminalNode ASSIGN() { return getToken(ProyectoParser.ASSIGN, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public VarsTypeInit2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varsTypeInit2; }
	}

	public final VarsTypeInit2Context varsTypeInit2() throws RecognitionException {
		VarsTypeInit2Context _localctx = new VarsTypeInit2Context(_ctx, getState());
		enterRule(_localctx, 10, RULE_varsTypeInit2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(198);
			match(ASSIGN);
			setState(199);
			exp();
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
		enterRule(_localctx, 12, RULE_vars);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(201);
			match(ID);
			setState(204);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(202);
				match(T__6);
				setState(203);
				match(ID);
				}
			}

			setState(210);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				{
				setState(206);
				match(T__3);
				setState(207);
				exp();
				setState(208);
				match(T__4);
				}
				break;
			}
			setState(216);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(212);
				match(T__3);
				setState(213);
				exp();
				setState(214);
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
		enterRule(_localctx, 14, RULE_functions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(218);
			match(FUNCTION);
			setState(219);
			match(ID);
			setState(220);
			match(LPAREN);
			setState(222);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(221);
				parameters();
				}
			}

			setState(224);
			match(RPAREN);
			setState(227);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_TYPE:
			case FLOAT_TYPE:
			case CHAR_TYPE:
			case BOOL_TYPE:
				{
				setState(225);
				typeRule();
				}
				break;
			case VOID:
				{
				setState(226);
				match(VOID);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(229);
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
		enterRule(_localctx, 16, RULE_parameters);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(231);
			parameter();
			setState(236);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__7) {
				{
				{
				setState(232);
				match(T__7);
				setState(233);
				parameter();
				}
				}
				setState(238);
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
		enterRule(_localctx, 18, RULE_parameter);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(239);
			match(ID);
			setState(240);
			match(T__2);
			setState(241);
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
		public ReturnRuleContext returnRule() {
			return getRuleContext(ReturnRuleContext.class,0);
		}
		public FunctionBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionBlock; }
	}

	public final FunctionBlockContext functionBlock() throws RecognitionException {
		FunctionBlockContext _localctx = new FunctionBlockContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_functionBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(243);
			match(T__0);
			setState(247);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(244);
				varsDec();
				}
				}
				setState(249);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(253);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << ID))) != 0)) {
				{
				{
				setState(250);
				statutes();
				}
				}
				setState(255);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(257);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==RETURN) {
				{
				setState(256);
				returnRule();
				}
			}

			setState(259);
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
		enterRule(_localctx, 22, RULE_returnRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(261);
			match(RETURN);
			setState(263);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << ID))) != 0)) {
				{
				setState(262);
				exp();
				}
			}

			setState(265);
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
		enterRule(_localctx, 24, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(267);
			match(T__0);
			setState(271);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << ID))) != 0)) {
				{
				{
				setState(268);
				statutes();
				}
				}
				setState(273);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(274);
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
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public StatutesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statutes; }
	}

	public final StatutesContext statutes() throws RecognitionException {
		StatutesContext _localctx = new StatutesContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_statutes);
		try {
			setState(283);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(276);
				assignation();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(277);
				read();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(278);
				write();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(279);
				conditional();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(280);
				forLoop();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(281);
				whileLoop();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(282);
				expression();
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
		public TerminalNode ASSIGN() { return getToken(ProyectoParser.ASSIGN, 0); }
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
		enterRule(_localctx, 28, RULE_assignation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(285);
			match(ID);
			setState(286);
			match(ASSIGN);
			setState(287);
			exp();
			setState(288);
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
		enterRule(_localctx, 30, RULE_functionCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(290);
			match(ID);
			setState(291);
			match(LPAREN);
			setState(293);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << ID))) != 0)) {
				{
				setState(292);
				arguments();
				}
			}

			setState(295);
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
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public Arguments2Context arguments2() {
			return getRuleContext(Arguments2Context.class,0);
		}
		public ArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arguments; }
	}

	public final ArgumentsContext arguments() throws RecognitionException {
		ArgumentsContext _localctx = new ArgumentsContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_arguments);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(297);
			exp();
			setState(298);
			arguments2();
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

	public static class Arguments2Context extends ParserRuleContext {
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public Arguments2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arguments2; }
	}

	public final Arguments2Context arguments2() throws RecognitionException {
		Arguments2Context _localctx = new Arguments2Context(_ctx, getState());
		enterRule(_localctx, 34, RULE_arguments2);
		try {
			setState(303);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__7:
				enterOuterAlt(_localctx, 1);
				{
				setState(300);
				match(T__7);
				setState(301);
				arguments();
				}
				break;
			case RPAREN:
				enterOuterAlt(_localctx, 2);
				{
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
		enterRule(_localctx, 36, RULE_methodCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(305);
			match(ID);
			setState(306);
			match(T__6);
			setState(307);
			match(ID);
			setState(308);
			match(LPAREN);
			setState(310);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << ID))) != 0)) {
				{
				setState(309);
				arguments();
				}
			}

			setState(312);
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
		enterRule(_localctx, 38, RULE_call);
		try {
			setState(316);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(314);
				functionCall();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(315);
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
		public Read2Context read2() {
			return getRuleContext(Read2Context.class,0);
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
		enterRule(_localctx, 40, RULE_read);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(318);
			match(READ);
			setState(319);
			match(LPAREN);
			setState(320);
			read2();
			setState(321);
			match(RPAREN);
			setState(322);
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

	public static class Read2Context extends ParserRuleContext {
		public VarsContext vars() {
			return getRuleContext(VarsContext.class,0);
		}
		public Read3Context read3() {
			return getRuleContext(Read3Context.class,0);
		}
		public Read2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_read2; }
	}

	public final Read2Context read2() throws RecognitionException {
		Read2Context _localctx = new Read2Context(_ctx, getState());
		enterRule(_localctx, 42, RULE_read2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(324);
			vars();
			setState(325);
			read3();
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

	public static class Read3Context extends ParserRuleContext {
		public Read2Context read2() {
			return getRuleContext(Read2Context.class,0);
		}
		public Read3Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_read3; }
	}

	public final Read3Context read3() throws RecognitionException {
		Read3Context _localctx = new Read3Context(_ctx, getState());
		enterRule(_localctx, 44, RULE_read3);
		try {
			setState(330);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__7:
				enterOuterAlt(_localctx, 1);
				{
				setState(327);
				match(T__7);
				setState(328);
				read2();
				}
				break;
			case RPAREN:
				enterOuterAlt(_localctx, 2);
				{
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
		enterRule(_localctx, 46, RULE_write);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(332);
			match(WRITE);
			setState(333);
			match(LPAREN);
			setState(334);
			arguments();
			setState(335);
			match(RPAREN);
			setState(336);
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
		public Conditional2Context conditional2() {
			return getRuleContext(Conditional2Context.class,0);
		}
		public Conditional3Context conditional3() {
			return getRuleContext(Conditional3Context.class,0);
		}
		public Conditional4Context conditional4() {
			return getRuleContext(Conditional4Context.class,0);
		}
		public ConditionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditional; }
	}

	public final ConditionalContext conditional() throws RecognitionException {
		ConditionalContext _localctx = new ConditionalContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_conditional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(338);
			match(IF);
			setState(339);
			conditional2();
			setState(340);
			conditional3();
			setState(342);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(341);
				conditional4();
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

	public static class Conditional2Context extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(ProyectoParser.LPAREN, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(ProyectoParser.RPAREN, 0); }
		public Conditional2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditional2; }
	}

	public final Conditional2Context conditional2() throws RecognitionException {
		Conditional2Context _localctx = new Conditional2Context(_ctx, getState());
		enterRule(_localctx, 50, RULE_conditional2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(344);
			match(LPAREN);
			setState(345);
			exp();
			setState(346);
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

	public static class Conditional3Context extends ParserRuleContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public Conditional3Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditional3; }
	}

	public final Conditional3Context conditional3() throws RecognitionException {
		Conditional3Context _localctx = new Conditional3Context(_ctx, getState());
		enterRule(_localctx, 52, RULE_conditional3);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(348);
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

	public static class Conditional4Context extends ParserRuleContext {
		public TerminalNode ELSE() { return getToken(ProyectoParser.ELSE, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public Conditional4Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditional4; }
	}

	public final Conditional4Context conditional4() throws RecognitionException {
		Conditional4Context _localctx = new Conditional4Context(_ctx, getState());
		enterRule(_localctx, 54, RULE_conditional4);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(350);
			match(ELSE);
			setState(351);
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

	public static class ForLoopContext extends ParserRuleContext {
		public TerminalNode FOR() { return getToken(ProyectoParser.FOR, 0); }
		public ForLoop2Context forLoop2() {
			return getRuleContext(ForLoop2Context.class,0);
		}
		public TerminalNode IN() { return getToken(ProyectoParser.IN, 0); }
		public ForLoop3Context forLoop3() {
			return getRuleContext(ForLoop3Context.class,0);
		}
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
		enterRule(_localctx, 56, RULE_forLoop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(353);
			match(FOR);
			setState(354);
			forLoop2();
			setState(355);
			match(IN);
			setState(356);
			forLoop3();
			setState(357);
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

	public static class ForLoop2Context extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ProyectoParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(ProyectoParser.ASSIGN, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public ForLoop2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forLoop2; }
	}

	public final ForLoop2Context forLoop2() throws RecognitionException {
		ForLoop2Context _localctx = new ForLoop2Context(_ctx, getState());
		enterRule(_localctx, 58, RULE_forLoop2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(359);
			match(ID);
			setState(360);
			match(ASSIGN);
			setState(361);
			exp();
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

	public static class ForLoop3Context extends ParserRuleContext {
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public ForLoop3Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forLoop3; }
	}

	public final ForLoop3Context forLoop3() throws RecognitionException {
		ForLoop3Context _localctx = new ForLoop3Context(_ctx, getState());
		enterRule(_localctx, 60, RULE_forLoop3);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			exp();
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
		public WhileLoop2Context whileLoop2() {
			return getRuleContext(WhileLoop2Context.class,0);
		}
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
		enterRule(_localctx, 62, RULE_whileLoop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(365);
			match(WHILE);
			setState(366);
			whileLoop2();
			setState(367);
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

	public static class WhileLoop2Context extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(ProyectoParser.LPAREN, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(ProyectoParser.RPAREN, 0); }
		public WhileLoop2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whileLoop2; }
	}

	public final WhileLoop2Context whileLoop2() throws RecognitionException {
		WhileLoop2Context _localctx = new WhileLoop2Context(_ctx, getState());
		enterRule(_localctx, 64, RULE_whileLoop2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(369);
			match(LPAREN);
			setState(370);
			exp();
			setState(371);
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

	public static class ExpressionContext extends ParserRuleContext {
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(ProyectoParser.SEMI, 0); }
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(373);
			exp();
			setState(374);
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
		enterRule(_localctx, 68, RULE_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(376);
			t_exp();
			setState(380);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OR) {
				{
				{
				setState(377);
				exp2();
				}
				}
				setState(382);
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
		enterRule(_localctx, 70, RULE_exp2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(383);
			match(OR);
			setState(384);
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
		enterRule(_localctx, 72, RULE_t_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(386);
			g_exp();
			setState(390);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==AND) {
				{
				{
				setState(387);
				t_exp2();
				}
				}
				setState(392);
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
		enterRule(_localctx, 74, RULE_t_exp2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(393);
			match(AND);
			setState(394);
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
		enterRule(_localctx, 76, RULE_g_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(396);
			m_exp();
			setState(398);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << GT) | (1L << LT) | (1L << EQ) | (1L << NEQ))) != 0)) {
				{
				setState(397);
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
		enterRule(_localctx, 78, RULE_g_exp2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(400);
			relop();
			setState(401);
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
		enterRule(_localctx, 80, RULE_m_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(403);
			term();
			setState(407);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ADD || _la==SUB) {
				{
				{
				setState(404);
				m_exp2();
				}
				}
				setState(409);
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
		enterRule(_localctx, 82, RULE_m_exp2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(410);
			_la = _input.LA(1);
			if ( !(_la==ADD || _la==SUB) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(411);
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
		enterRule(_localctx, 84, RULE_term);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(413);
			factor();
			setState(417);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==MUL || _la==DIV) {
				{
				{
				setState(414);
				term2();
				}
				}
				setState(419);
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
		enterRule(_localctx, 86, RULE_term2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(420);
			_la = _input.LA(1);
			if ( !(_la==MUL || _la==DIV) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(421);
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
		public Factor2Context factor2() {
			return getRuleContext(Factor2Context.class,0);
		}
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
		enterRule(_localctx, 88, RULE_factor);
		try {
			setState(427);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(423);
				factor2();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(424);
				varCte();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(425);
				vars();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(426);
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

	public static class Factor2Context extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(ProyectoParser.LPAREN, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(ProyectoParser.RPAREN, 0); }
		public Factor2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_factor2; }
	}

	public final Factor2Context factor2() throws RecognitionException {
		Factor2Context _localctx = new Factor2Context(_ctx, getState());
		enterRule(_localctx, 90, RULE_factor2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(429);
			match(LPAREN);
			setState(430);
			exp();
			setState(431);
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
		enterRule(_localctx, 92, RULE_varCte);
		try {
			setState(438);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(433);
				cte_i();
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(434);
				cte_f();
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 3);
				{
				setState(435);
				cte_c();
				}
				break;
			case T__8:
				enterOuterAlt(_localctx, 4);
				{
				setState(436);
				cte_s();
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 5);
				{
				setState(437);
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
		enterRule(_localctx, 94, RULE_cte_i);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(440);
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
		enterRule(_localctx, 96, RULE_cte_f);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(442);
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
		enterRule(_localctx, 98, RULE_cte_c);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(444);
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
		enterRule(_localctx, 100, RULE_cte_b);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(446);
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
		enterRule(_localctx, 102, RULE_cte_s);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(448);
			match(T__8);
			setState(452);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,34,_ctx);
			while ( _alt!=1 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1+1 ) {
					{
					{
					setState(449);
					matchWildcard();
					}
					} 
				}
				setState(454);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,34,_ctx);
			}
			setState(455);
			match(T__8);
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
		public MainBlockContext mainBlock() {
			return getRuleContext(MainBlockContext.class,0);
		}
		public MainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_main; }
	}

	public final MainContext main() throws RecognitionException {
		MainContext _localctx = new MainContext(_ctx, getState());
		enterRule(_localctx, 104, RULE_main);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(457);
			match(MAIN);
			setState(458);
			match(LPAREN);
			setState(459);
			match(RPAREN);
			setState(460);
			mainBlock();
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

	public static class MainBlockContext extends ParserRuleContext {
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
		public MainBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mainBlock; }
	}

	public final MainBlockContext mainBlock() throws RecognitionException {
		MainBlockContext _localctx = new MainBlockContext(_ctx, getState());
		enterRule(_localctx, 106, RULE_mainBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(462);
			match(T__0);
			setState(466);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(463);
				varsDec();
				}
				}
				setState(468);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(472);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << ID))) != 0)) {
				{
				{
				setState(469);
				statutes();
				}
				}
				setState(474);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(475);
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

	public static class TypeRuleContext extends ParserRuleContext {
		public TerminalNode INT_TYPE() { return getToken(ProyectoParser.INT_TYPE, 0); }
		public TerminalNode FLOAT_TYPE() { return getToken(ProyectoParser.FLOAT_TYPE, 0); }
		public TerminalNode CHAR_TYPE() { return getToken(ProyectoParser.CHAR_TYPE, 0); }
		public TerminalNode BOOL_TYPE() { return getToken(ProyectoParser.BOOL_TYPE, 0); }
		public TypeRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeRule; }
	}

	public final TypeRuleContext typeRule() throws RecognitionException {
		TypeRuleContext _localctx = new TypeRuleContext(_ctx, getState());
		enterRule(_localctx, 108, RULE_typeRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(477);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT_TYPE) | (1L << FLOAT_TYPE) | (1L << CHAR_TYPE) | (1L << BOOL_TYPE))) != 0)) ) {
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
		enterRule(_localctx, 110, RULE_relop);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(479);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\63\u01e4\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\3\2\3\2\3\2\3\2\7\2w\n\2"+
		"\f\2\16\2z\13\2\3\2\7\2}\n\2\f\2\16\2\u0080\13\2\3\2\7\2\u0083\n\2\f\2"+
		"\16\2\u0086\13\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\5\3\u0090\n\3\3\3\3\3"+
		"\3\3\3\4\3\4\3\4\7\4\u0098\n\4\f\4\16\4\u009b\13\4\3\4\3\4\7\4\u009f\n"+
		"\4\f\4\16\4\u00a2\13\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5\u00c0"+
		"\n\5\3\6\3\6\5\6\u00c4\n\6\3\6\5\6\u00c7\n\6\3\7\3\7\3\7\3\b\3\b\3\b\5"+
		"\b\u00cf\n\b\3\b\3\b\3\b\3\b\5\b\u00d5\n\b\3\b\3\b\3\b\3\b\5\b\u00db\n"+
		"\b\3\t\3\t\3\t\3\t\5\t\u00e1\n\t\3\t\3\t\3\t\5\t\u00e6\n\t\3\t\3\t\3\n"+
		"\3\n\3\n\7\n\u00ed\n\n\f\n\16\n\u00f0\13\n\3\13\3\13\3\13\3\13\3\f\3\f"+
		"\7\f\u00f8\n\f\f\f\16\f\u00fb\13\f\3\f\7\f\u00fe\n\f\f\f\16\f\u0101\13"+
		"\f\3\f\5\f\u0104\n\f\3\f\3\f\3\r\3\r\5\r\u010a\n\r\3\r\3\r\3\16\3\16\7"+
		"\16\u0110\n\16\f\16\16\16\u0113\13\16\3\16\3\16\3\17\3\17\3\17\3\17\3"+
		"\17\3\17\3\17\5\17\u011e\n\17\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21"+
		"\5\21\u0128\n\21\3\21\3\21\3\22\3\22\3\22\3\23\3\23\3\23\5\23\u0132\n"+
		"\23\3\24\3\24\3\24\3\24\3\24\5\24\u0139\n\24\3\24\3\24\3\25\3\25\5\25"+
		"\u013f\n\25\3\26\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\30\3\30\3\30"+
		"\5\30\u014d\n\30\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\5\32"+
		"\u0159\n\32\3\33\3\33\3\33\3\33\3\34\3\34\3\35\3\35\3\35\3\36\3\36\3\36"+
		"\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3 \3 \3!\3!\3!\3!\3\"\3\"\3\"\3\""+
		"\3#\3#\3#\3$\3$\7$\u017d\n$\f$\16$\u0180\13$\3%\3%\3%\3&\3&\7&\u0187\n"+
		"&\f&\16&\u018a\13&\3\'\3\'\3\'\3(\3(\5(\u0191\n(\3)\3)\3)\3*\3*\7*\u0198"+
		"\n*\f*\16*\u019b\13*\3+\3+\3+\3,\3,\7,\u01a2\n,\f,\16,\u01a5\13,\3-\3"+
		"-\3-\3.\3.\3.\3.\5.\u01ae\n.\3/\3/\3/\3/\3\60\3\60\3\60\3\60\3\60\5\60"+
		"\u01b9\n\60\3\61\3\61\3\62\3\62\3\63\3\63\3\64\3\64\3\65\3\65\7\65\u01c5"+
		"\n\65\f\65\16\65\u01c8\13\65\3\65\3\65\3\66\3\66\3\66\3\66\3\66\3\67\3"+
		"\67\7\67\u01d3\n\67\f\67\16\67\u01d6\13\67\3\67\7\67\u01d9\n\67\f\67\16"+
		"\67\u01dc\13\67\3\67\3\67\38\38\39\39\39\3\u01c6\2:\2\4\6\b\n\f\16\20"+
		"\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDFHJLNPRTVXZ\\^`bdfhj"+
		"lnp\2\6\3\2\31\32\3\2\27\30\3\2-\60\3\2\23\26\2\u01db\2r\3\2\2\2\4\u008a"+
		"\3\2\2\2\6\u0094\3\2\2\2\b\u00bf\3\2\2\2\n\u00c3\3\2\2\2\f\u00c8\3\2\2"+
		"\2\16\u00cb\3\2\2\2\20\u00dc\3\2\2\2\22\u00e9\3\2\2\2\24\u00f1\3\2\2\2"+
		"\26\u00f5\3\2\2\2\30\u0107\3\2\2\2\32\u010d\3\2\2\2\34\u011d\3\2\2\2\36"+
		"\u011f\3\2\2\2 \u0124\3\2\2\2\"\u012b\3\2\2\2$\u0131\3\2\2\2&\u0133\3"+
		"\2\2\2(\u013e\3\2\2\2*\u0140\3\2\2\2,\u0146\3\2\2\2.\u014c\3\2\2\2\60"+
		"\u014e\3\2\2\2\62\u0154\3\2\2\2\64\u015a\3\2\2\2\66\u015e\3\2\2\28\u0160"+
		"\3\2\2\2:\u0163\3\2\2\2<\u0169\3\2\2\2>\u016d\3\2\2\2@\u016f\3\2\2\2B"+
		"\u0173\3\2\2\2D\u0177\3\2\2\2F\u017a\3\2\2\2H\u0181\3\2\2\2J\u0184\3\2"+
		"\2\2L\u018b\3\2\2\2N\u018e\3\2\2\2P\u0192\3\2\2\2R\u0195\3\2\2\2T\u019c"+
		"\3\2\2\2V\u019f\3\2\2\2X\u01a6\3\2\2\2Z\u01ad\3\2\2\2\\\u01af\3\2\2\2"+
		"^\u01b8\3\2\2\2`\u01ba\3\2\2\2b\u01bc\3\2\2\2d\u01be\3\2\2\2f\u01c0\3"+
		"\2\2\2h\u01c2\3\2\2\2j\u01cb\3\2\2\2l\u01d0\3\2\2\2n\u01df\3\2\2\2p\u01e1"+
		"\3\2\2\2rs\7\62\2\2st\7\63\2\2tx\7\33\2\2uw\5\4\3\2vu\3\2\2\2wz\3\2\2"+
		"\2xv\3\2\2\2xy\3\2\2\2y~\3\2\2\2zx\3\2\2\2{}\5\b\5\2|{\3\2\2\2}\u0080"+
		"\3\2\2\2~|\3\2\2\2~\177\3\2\2\2\177\u0084\3\2\2\2\u0080~\3\2\2\2\u0081"+
		"\u0083\5\20\t\2\u0082\u0081\3\2\2\2\u0083\u0086\3\2\2\2\u0084\u0082\3"+
		"\2\2\2\u0084\u0085\3\2\2\2\u0085\u0087\3\2\2\2\u0086\u0084\3\2\2\2\u0087"+
		"\u0088\5j\66\2\u0088\u0089\7\2\2\3\u0089\3\3\2\2\2\u008a\u008b\7$\2\2"+
		"\u008b\u008f\7\63\2\2\u008c\u008d\7\24\2\2\u008d\u008e\7\63\2\2\u008e"+
		"\u0090\7\23\2\2\u008f\u008c\3\2\2\2\u008f\u0090\3\2\2\2\u0090\u0091\3"+
		"\2\2\2\u0091\u0092\5\6\4\2\u0092\u0093\7\33\2\2\u0093\5\3\2\2\2\u0094"+
		"\u0095\7\3\2\2\u0095\u0099\7%\2\2\u0096\u0098\5\b\5\2\u0097\u0096\3\2"+
		"\2\2\u0098\u009b\3\2\2\2\u0099\u0097\3\2\2\2\u0099\u009a\3\2\2\2\u009a"+
		"\u009c\3\2\2\2\u009b\u0099\3\2\2\2\u009c\u00a0\7&\2\2\u009d\u009f\5\20"+
		"\t\2\u009e\u009d\3\2\2\2\u009f\u00a2\3\2\2\2\u00a0\u009e\3\2\2\2\u00a0"+
		"\u00a1\3\2\2\2\u00a1\u00a3\3\2\2\2\u00a2\u00a0\3\2\2\2\u00a3\u00a4\7\4"+
		"\2\2\u00a4\7\3\2\2\2\u00a5\u00a6\7,\2\2\u00a6\u00a7\7\63\2\2\u00a7\u00a8"+
		"\7\5\2\2\u00a8\u00a9\5\n\6\2\u00a9\u00aa\7\33\2\2\u00aa\u00c0\3\2\2\2"+
		"\u00ab\u00ac\7,\2\2\u00ac\u00ad\7\63\2\2\u00ad\u00ae\7\6\2\2\u00ae\u00af"+
		"\7\f\2\2\u00af\u00b0\7\7\2\2\u00b0\u00b1\7\5\2\2\u00b1\u00b2\5\n\6\2\u00b2"+
		"\u00b3\7\33\2\2\u00b3\u00c0\3\2\2\2\u00b4\u00b5\7,\2\2\u00b5\u00b6\7\63"+
		"\2\2\u00b6\u00b7\7\6\2\2\u00b7\u00b8\7\f\2\2\u00b8\u00b9\7\b\2\2\u00b9"+
		"\u00ba\7\f\2\2\u00ba\u00bb\7\7\2\2\u00bb\u00bc\7\5\2\2\u00bc\u00bd\5\n"+
		"\6\2\u00bd\u00be\7\33\2\2\u00be\u00c0\3\2\2\2\u00bf\u00a5\3\2\2\2\u00bf"+
		"\u00ab\3\2\2\2\u00bf\u00b4\3\2\2\2\u00c0\t\3\2\2\2\u00c1\u00c4\5n8\2\u00c2"+
		"\u00c4\7\63\2\2\u00c3\u00c1\3\2\2\2\u00c3\u00c2\3\2\2\2\u00c4\u00c6\3"+
		"\2\2\2\u00c5\u00c7\5\f\7\2\u00c6\u00c5\3\2\2\2\u00c6\u00c7\3\2\2\2\u00c7"+
		"\13\3\2\2\2\u00c8\u00c9\7\34\2\2\u00c9\u00ca\5F$\2\u00ca\r\3\2\2\2\u00cb"+
		"\u00ce\7\63\2\2\u00cc\u00cd\7\t\2\2\u00cd\u00cf\7\63\2\2\u00ce\u00cc\3"+
		"\2\2\2\u00ce\u00cf\3\2\2\2\u00cf\u00d4\3\2\2\2\u00d0\u00d1\7\6\2\2\u00d1"+
		"\u00d2\5F$\2\u00d2\u00d3\7\7\2\2\u00d3\u00d5\3\2\2\2\u00d4\u00d0\3\2\2"+
		"\2\u00d4\u00d5\3\2\2\2\u00d5\u00da\3\2\2\2\u00d6\u00d7\7\6\2\2\u00d7\u00d8"+
		"\5F$\2\u00d8\u00d9\7\7\2\2\u00d9\u00db\3\2\2\2\u00da\u00d6\3\2\2\2\u00da"+
		"\u00db\3\2\2\2\u00db\17\3\2\2\2\u00dc\u00dd\7)\2\2\u00dd\u00de\7\63\2"+
		"\2\u00de\u00e0\7\35\2\2\u00df\u00e1\5\22\n\2\u00e0\u00df\3\2\2\2\u00e0"+
		"\u00e1\3\2\2\2\u00e1\u00e2\3\2\2\2\u00e2\u00e5\7\36\2\2\u00e3\u00e6\5"+
		"n8\2\u00e4\u00e6\7\61\2\2\u00e5\u00e3\3\2\2\2\u00e5\u00e4\3\2\2\2\u00e6"+
		"\u00e7\3\2\2\2\u00e7\u00e8\5\26\f\2\u00e8\21\3\2\2\2\u00e9\u00ee\5\24"+
		"\13\2\u00ea\u00eb\7\n\2\2\u00eb\u00ed\5\24\13\2\u00ec\u00ea\3\2\2\2\u00ed"+
		"\u00f0\3\2\2\2\u00ee\u00ec\3\2\2\2\u00ee\u00ef\3\2\2\2\u00ef\23\3\2\2"+
		"\2\u00f0\u00ee\3\2\2\2\u00f1\u00f2\7\63\2\2\u00f2\u00f3\7\5\2\2\u00f3"+
		"\u00f4\5n8\2\u00f4\25\3\2\2\2\u00f5\u00f9\7\3\2\2\u00f6\u00f8\5\b\5\2"+
		"\u00f7\u00f6\3\2\2\2\u00f8\u00fb\3\2\2\2\u00f9\u00f7\3\2\2\2\u00f9\u00fa"+
		"\3\2\2\2\u00fa\u00ff\3\2\2\2\u00fb\u00f9\3\2\2\2\u00fc\u00fe\5\34\17\2"+
		"\u00fd\u00fc\3\2\2\2\u00fe\u0101\3\2\2\2\u00ff\u00fd\3\2\2\2\u00ff\u0100"+
		"\3\2\2\2\u0100\u0103\3\2\2\2\u0101\u00ff\3\2\2\2\u0102\u0104\5\30\r\2"+
		"\u0103\u0102\3\2\2\2\u0103\u0104\3\2\2\2\u0104\u0105\3\2\2\2\u0105\u0106"+
		"\7\4\2\2\u0106\27\3\2\2\2\u0107\u0109\7*\2\2\u0108\u010a\5F$\2\u0109\u0108"+
		"\3\2\2\2\u0109\u010a\3\2\2\2\u010a\u010b\3\2\2\2\u010b\u010c\7\33\2\2"+
		"\u010c\31\3\2\2\2\u010d\u0111\7\3\2\2\u010e\u0110\5\34\17\2\u010f\u010e"+
		"\3\2\2\2\u0110\u0113\3\2\2\2\u0111\u010f\3\2\2\2\u0111\u0112\3\2\2\2\u0112"+
		"\u0114\3\2\2\2\u0113\u0111\3\2\2\2\u0114\u0115\7\4\2\2\u0115\33\3\2\2"+
		"\2\u0116\u011e\5\36\20\2\u0117\u011e\5*\26\2\u0118\u011e\5\60\31\2\u0119"+
		"\u011e\5\62\32\2\u011a\u011e\5:\36\2\u011b\u011e\5@!\2\u011c\u011e\5D"+
		"#\2\u011d\u0116\3\2\2\2\u011d\u0117\3\2\2\2\u011d\u0118\3\2\2\2\u011d"+
		"\u0119\3\2\2\2\u011d\u011a\3\2\2\2\u011d\u011b\3\2\2\2\u011d\u011c\3\2"+
		"\2\2\u011e\35\3\2\2\2\u011f\u0120\7\63\2\2\u0120\u0121\7\34\2\2\u0121"+
		"\u0122\5F$\2\u0122\u0123\7\33\2\2\u0123\37\3\2\2\2\u0124\u0125\7\63\2"+
		"\2\u0125\u0127\7\35\2\2\u0126\u0128\5\"\22\2\u0127\u0126\3\2\2\2\u0127"+
		"\u0128\3\2\2\2\u0128\u0129\3\2\2\2\u0129\u012a\7\36\2\2\u012a!\3\2\2\2"+
		"\u012b\u012c\5F$\2\u012c\u012d\5$\23\2\u012d#\3\2\2\2\u012e\u012f\7\n"+
		"\2\2\u012f\u0132\5\"\22\2\u0130\u0132\3\2\2\2\u0131\u012e\3\2\2\2\u0131"+
		"\u0130\3\2\2\2\u0132%\3\2\2\2\u0133\u0134\7\63\2\2\u0134\u0135\7\t\2\2"+
		"\u0135\u0136\7\63\2\2\u0136\u0138\7\35\2\2\u0137\u0139\5\"\22\2\u0138"+
		"\u0137\3\2\2\2\u0138\u0139\3\2\2\2\u0139\u013a\3\2\2\2\u013a\u013b\7\36"+
		"\2\2\u013b\'\3\2\2\2\u013c\u013f\5 \21\2\u013d\u013f\5&\24\2\u013e\u013c"+
		"\3\2\2\2\u013e\u013d\3\2\2\2\u013f)\3\2\2\2\u0140\u0141\7(\2\2\u0141\u0142"+
		"\7\35\2\2\u0142\u0143\5,\27\2\u0143\u0144\7\36\2\2\u0144\u0145\7\33\2"+
		"\2\u0145+\3\2\2\2\u0146\u0147\5\16\b\2\u0147\u0148\5.\30\2\u0148-\3\2"+
		"\2\2\u0149\u014a\7\n\2\2\u014a\u014d\5,\27\2\u014b\u014d\3\2\2\2\u014c"+
		"\u0149\3\2\2\2\u014c\u014b\3\2\2\2\u014d/\3\2\2\2\u014e\u014f\7\'\2\2"+
		"\u014f\u0150\7\35\2\2\u0150\u0151\5\"\22\2\u0151\u0152\7\36\2\2\u0152"+
		"\u0153\7\33\2\2\u0153\61\3\2\2\2\u0154\u0155\7\37\2\2\u0155\u0156\5\64"+
		"\33\2\u0156\u0158\5\66\34\2\u0157\u0159\58\35\2\u0158\u0157\3\2\2\2\u0158"+
		"\u0159\3\2\2\2\u0159\63\3\2\2\2\u015a\u015b\7\35\2\2\u015b\u015c\5F$\2"+
		"\u015c\u015d\7\36\2\2\u015d\65\3\2\2\2\u015e\u015f\5\32\16\2\u015f\67"+
		"\3\2\2\2\u0160\u0161\7 \2\2\u0161\u0162\5\32\16\2\u01629\3\2\2\2\u0163"+
		"\u0164\7\"\2\2\u0164\u0165\5<\37\2\u0165\u0166\7#\2\2\u0166\u0167\5> "+
		"\2\u0167\u0168\5\32\16\2\u0168;\3\2\2\2\u0169\u016a\7\63\2\2\u016a\u016b"+
		"\7\34\2\2\u016b\u016c\5F$\2\u016c=\3\2\2\2\u016d\u016e\5F$\2\u016e?\3"+
		"\2\2\2\u016f\u0170\7!\2\2\u0170\u0171\5B\"\2\u0171\u0172\5\32\16\2\u0172"+
		"A\3\2\2\2\u0173\u0174\7\35\2\2\u0174\u0175\5F$\2\u0175\u0176\7\36\2\2"+
		"\u0176C\3\2\2\2\u0177\u0178\5F$\2\u0178\u0179\7\33\2\2\u0179E\3\2\2\2"+
		"\u017a\u017e\5J&\2\u017b\u017d\5H%\2\u017c\u017b\3\2\2\2\u017d\u0180\3"+
		"\2\2\2\u017e\u017c\3\2\2\2\u017e\u017f\3\2\2\2\u017fG\3\2\2\2\u0180\u017e"+
		"\3\2\2\2\u0181\u0182\7\21\2\2\u0182\u0183\5J&\2\u0183I\3\2\2\2\u0184\u0188"+
		"\5N(\2\u0185\u0187\5L\'\2\u0186\u0185\3\2\2\2\u0187\u018a\3\2\2\2\u0188"+
		"\u0186\3\2\2\2\u0188\u0189\3\2\2\2\u0189K\3\2\2\2\u018a\u0188\3\2\2\2"+
		"\u018b\u018c\7\22\2\2\u018c\u018d\5N(\2\u018dM\3\2\2\2\u018e\u0190\5R"+
		"*\2\u018f\u0191\5P)\2\u0190\u018f\3\2\2\2\u0190\u0191\3\2\2\2\u0191O\3"+
		"\2\2\2\u0192\u0193\5p9\2\u0193\u0194\5R*\2\u0194Q\3\2\2\2\u0195\u0199"+
		"\5V,\2\u0196\u0198\5T+\2\u0197\u0196\3\2\2\2\u0198\u019b\3\2\2\2\u0199"+
		"\u0197\3\2\2\2\u0199\u019a\3\2\2\2\u019aS\3\2\2\2\u019b\u0199\3\2\2\2"+
		"\u019c\u019d\t\2\2\2\u019d\u019e\5V,\2\u019eU\3\2\2\2\u019f\u01a3\5Z."+
		"\2\u01a0\u01a2\5X-\2\u01a1\u01a0\3\2\2\2\u01a2\u01a5\3\2\2\2\u01a3\u01a1"+
		"\3\2\2\2\u01a3\u01a4\3\2\2\2\u01a4W\3\2\2\2\u01a5\u01a3\3\2\2\2\u01a6"+
		"\u01a7\t\3\2\2\u01a7\u01a8\5Z.\2\u01a8Y\3\2\2\2\u01a9\u01ae\5\\/\2\u01aa"+
		"\u01ae\5^\60\2\u01ab\u01ae\5\16\b\2\u01ac\u01ae\5(\25\2\u01ad\u01a9\3"+
		"\2\2\2\u01ad\u01aa\3\2\2\2\u01ad\u01ab\3\2\2\2\u01ad\u01ac\3\2\2\2\u01ae"+
		"[\3\2\2\2\u01af\u01b0\7\35\2\2\u01b0\u01b1\5F$\2\u01b1\u01b2\7\36\2\2"+
		"\u01b2]\3\2\2\2\u01b3\u01b9\5`\61\2\u01b4\u01b9\5b\62\2\u01b5\u01b9\5"+
		"d\63\2\u01b6\u01b9\5h\65\2\u01b7\u01b9\5f\64\2\u01b8\u01b3\3\2\2\2\u01b8"+
		"\u01b4\3\2\2\2\u01b8\u01b5\3\2\2\2\u01b8\u01b6\3\2\2\2\u01b8\u01b7\3\2"+
		"\2\2\u01b9_\3\2\2\2\u01ba\u01bb\7\f\2\2\u01bba\3\2\2\2\u01bc\u01bd\7\r"+
		"\2\2\u01bdc\3\2\2\2\u01be\u01bf\7\16\2\2\u01bfe\3\2\2\2\u01c0\u01c1\7"+
		"\17\2\2\u01c1g\3\2\2\2\u01c2\u01c6\7\13\2\2\u01c3\u01c5\13\2\2\2\u01c4"+
		"\u01c3\3\2\2\2\u01c5\u01c8\3\2\2\2\u01c6\u01c7\3\2\2\2\u01c6\u01c4\3\2"+
		"\2\2\u01c7\u01c9\3\2\2\2\u01c8\u01c6\3\2\2\2\u01c9\u01ca\7\13\2\2\u01ca"+
		"i\3\2\2\2\u01cb\u01cc\7+\2\2\u01cc\u01cd\7\35\2\2\u01cd\u01ce\7\36\2\2"+
		"\u01ce\u01cf\5l\67\2\u01cfk\3\2\2\2\u01d0\u01d4\7\3\2\2\u01d1\u01d3\5"+
		"\b\5\2\u01d2\u01d1\3\2\2\2\u01d3\u01d6\3\2\2\2\u01d4\u01d2\3\2\2\2\u01d4"+
		"\u01d5\3\2\2\2\u01d5\u01da\3\2\2\2\u01d6\u01d4\3\2\2\2\u01d7\u01d9\5\34"+
		"\17\2\u01d8\u01d7\3\2\2\2\u01d9\u01dc\3\2\2\2\u01da\u01d8\3\2\2\2\u01da"+
		"\u01db\3\2\2\2\u01db\u01dd\3\2\2\2\u01dc\u01da\3\2\2\2\u01dd\u01de\7\4"+
		"\2\2\u01dem\3\2\2\2\u01df\u01e0\t\4\2\2\u01e0o\3\2\2\2\u01e1\u01e2\t\5"+
		"\2\2\u01e2q\3\2\2\2\'x~\u0084\u008f\u0099\u00a0\u00bf\u00c3\u00c6\u00ce"+
		"\u00d4\u00da\u00e0\u00e5\u00ee\u00f9\u00ff\u0103\u0109\u0111\u011d\u0127"+
		"\u0131\u0138\u013e\u014c\u0158\u017e\u0188\u0190\u0199\u01a3\u01ad\u01b8"+
		"\u01c6\u01d4\u01da";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}