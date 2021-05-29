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
		RULE_program = 0, RULE_program2 = 1, RULE_classDef = 2, RULE_classBlock = 3, 
		RULE_classAttributes = 4, RULE_variableDeclaration = 5, RULE_varsDec = 6, 
		RULE_varsDecArray = 7, RULE_varsDecMat = 8, RULE_varsTypeInit = 9, RULE_varsTypeInit2 = 10, 
		RULE_vars = 11, RULE_varArray = 12, RULE_varMat = 13, RULE_functions = 14, 
		RULE_parameters = 15, RULE_parameter = 16, RULE_functionBlock = 17, RULE_returnRule = 18, 
		RULE_block = 19, RULE_statutes = 20, RULE_assignation = 21, RULE_functionCall = 22, 
		RULE_arguments = 23, RULE_arguments2 = 24, RULE_methodCall = 25, RULE_call = 26, 
		RULE_read = 27, RULE_read2 = 28, RULE_read3 = 29, RULE_write = 30, RULE_conditional = 31, 
		RULE_conditional2 = 32, RULE_conditional3 = 33, RULE_conditional4 = 34, 
		RULE_forLoop = 35, RULE_forLoop2 = 36, RULE_forLoop3 = 37, RULE_whileLoop = 38, 
		RULE_whileLoop2 = 39, RULE_expression = 40, RULE_exp = 41, RULE_exp2 = 42, 
		RULE_t_exp = 43, RULE_t_exp2 = 44, RULE_g_exp = 45, RULE_g_exp2 = 46, 
		RULE_m_exp = 47, RULE_m_exp2 = 48, RULE_term = 49, RULE_term2 = 50, RULE_factor = 51, 
		RULE_factor2 = 52, RULE_varCte = 53, RULE_cte_i = 54, RULE_cte_f = 55, 
		RULE_cte_c = 56, RULE_cte_b = 57, RULE_cte_s = 58, RULE_main = 59, RULE_mainBlock = 60, 
		RULE_typeRule = 61, RULE_relop = 62;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "program2", "classDef", "classBlock", "classAttributes", "variableDeclaration", 
			"varsDec", "varsDecArray", "varsDecMat", "varsTypeInit", "varsTypeInit2", 
			"vars", "varArray", "varMat", "functions", "parameters", "parameter", 
			"functionBlock", "returnRule", "block", "statutes", "assignation", "functionCall", 
			"arguments", "arguments2", "methodCall", "call", "read", "read2", "read3", 
			"write", "conditional", "conditional2", "conditional3", "conditional4", 
			"forLoop", "forLoop2", "forLoop3", "whileLoop", "whileLoop2", "expression", 
			"exp", "exp2", "t_exp", "t_exp2", "g_exp", "g_exp2", "m_exp", "m_exp2", 
			"term", "term2", "factor", "factor2", "varCte", "cte_i", "cte_f", "cte_c", 
			"cte_b", "cte_s", "main", "mainBlock", "typeRule", "relop"
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
		public Program2Context program2() {
			return getRuleContext(Program2Context.class,0);
		}
		public List<ClassDefContext> classDef() {
			return getRuleContexts(ClassDefContext.class);
		}
		public ClassDefContext classDef(int i) {
			return getRuleContext(ClassDefContext.class,i);
		}
		public List<VariableDeclarationContext> variableDeclaration() {
			return getRuleContexts(VariableDeclarationContext.class);
		}
		public VariableDeclarationContext variableDeclaration(int i) {
			return getRuleContext(VariableDeclarationContext.class,i);
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
			setState(126);
			match(PROGRAM);
			setState(127);
			match(ID);
			setState(128);
			match(SEMI);
			setState(132);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CLASS) {
				{
				{
				setState(129);
				classDef();
				}
				}
				setState(134);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(138);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(135);
				variableDeclaration();
				}
				}
				setState(140);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(141);
			program2();
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

	public static class Program2Context extends ParserRuleContext {
		public MainContext main() {
			return getRuleContext(MainContext.class,0);
		}
		public TerminalNode EOF() { return getToken(ProyectoParser.EOF, 0); }
		public List<FunctionsContext> functions() {
			return getRuleContexts(FunctionsContext.class);
		}
		public FunctionsContext functions(int i) {
			return getRuleContext(FunctionsContext.class,i);
		}
		public Program2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program2; }
	}

	public final Program2Context program2() throws RecognitionException {
		Program2Context _localctx = new Program2Context(_ctx, getState());
		enterRule(_localctx, 2, RULE_program2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(146);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FUNCTION) {
				{
				{
				setState(143);
				functions();
				}
				}
				setState(148);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(149);
			main();
			setState(150);
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
		enterRule(_localctx, 4, RULE_classDef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			match(CLASS);
			setState(153);
			match(ID);
			setState(157);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LT) {
				{
				setState(154);
				match(LT);
				setState(155);
				match(ID);
				setState(156);
				match(GT);
				}
			}

			setState(159);
			classBlock();
			setState(160);
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
		public ClassAttributesContext classAttributes() {
			return getRuleContext(ClassAttributesContext.class,0);
		}
		public TerminalNode METHODS() { return getToken(ProyectoParser.METHODS, 0); }
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
		enterRule(_localctx, 6, RULE_classBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			match(T__0);
			setState(163);
			classAttributes();
			setState(164);
			match(METHODS);
			setState(168);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FUNCTION) {
				{
				{
				setState(165);
				functions();
				}
				}
				setState(170);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(171);
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

	public static class ClassAttributesContext extends ParserRuleContext {
		public TerminalNode ATTRIBUTES() { return getToken(ProyectoParser.ATTRIBUTES, 0); }
		public List<VariableDeclarationContext> variableDeclaration() {
			return getRuleContexts(VariableDeclarationContext.class);
		}
		public VariableDeclarationContext variableDeclaration(int i) {
			return getRuleContext(VariableDeclarationContext.class,i);
		}
		public ClassAttributesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classAttributes; }
	}

	public final ClassAttributesContext classAttributes() throws RecognitionException {
		ClassAttributesContext _localctx = new ClassAttributesContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_classAttributes);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(173);
			match(ATTRIBUTES);
			setState(177);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(174);
				variableDeclaration();
				}
				}
				setState(179);
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

	public static class VariableDeclarationContext extends ParserRuleContext {
		public VarsDecContext varsDec() {
			return getRuleContext(VarsDecContext.class,0);
		}
		public VarsDecArrayContext varsDecArray() {
			return getRuleContext(VarsDecArrayContext.class,0);
		}
		public VarsDecMatContext varsDecMat() {
			return getRuleContext(VarsDecMatContext.class,0);
		}
		public VariableDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variableDeclaration; }
	}

	public final VariableDeclarationContext variableDeclaration() throws RecognitionException {
		VariableDeclarationContext _localctx = new VariableDeclarationContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_variableDeclaration);
		try {
			setState(183);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(180);
				varsDec();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(181);
				varsDecArray();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(182);
				varsDecMat();
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

	public static class VarsDecContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(ProyectoParser.VAR, 0); }
		public TerminalNode ID() { return getToken(ProyectoParser.ID, 0); }
		public VarsTypeInitContext varsTypeInit() {
			return getRuleContext(VarsTypeInitContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(ProyectoParser.SEMI, 0); }
		public VarsDecContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varsDec; }
	}

	public final VarsDecContext varsDec() throws RecognitionException {
		VarsDecContext _localctx = new VarsDecContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_varsDec);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(185);
			match(VAR);
			setState(186);
			match(ID);
			setState(187);
			match(T__2);
			setState(188);
			varsTypeInit();
			setState(189);
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

	public static class VarsDecArrayContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(ProyectoParser.VAR, 0); }
		public TerminalNode ID() { return getToken(ProyectoParser.ID, 0); }
		public TerminalNode INT() { return getToken(ProyectoParser.INT, 0); }
		public TypeRuleContext typeRule() {
			return getRuleContext(TypeRuleContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(ProyectoParser.SEMI, 0); }
		public VarsDecArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varsDecArray; }
	}

	public final VarsDecArrayContext varsDecArray() throws RecognitionException {
		VarsDecArrayContext _localctx = new VarsDecArrayContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_varsDecArray);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(191);
			match(VAR);
			setState(192);
			match(ID);
			setState(193);
			match(T__2);
			setState(194);
			match(T__3);
			setState(195);
			match(INT);
			setState(196);
			match(T__4);
			setState(197);
			typeRule();
			setState(198);
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

	public static class VarsDecMatContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(ProyectoParser.VAR, 0); }
		public TerminalNode ID() { return getToken(ProyectoParser.ID, 0); }
		public List<TerminalNode> INT() { return getTokens(ProyectoParser.INT); }
		public TerminalNode INT(int i) {
			return getToken(ProyectoParser.INT, i);
		}
		public TypeRuleContext typeRule() {
			return getRuleContext(TypeRuleContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(ProyectoParser.SEMI, 0); }
		public VarsDecMatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varsDecMat; }
	}

	public final VarsDecMatContext varsDecMat() throws RecognitionException {
		VarsDecMatContext _localctx = new VarsDecMatContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_varsDecMat);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(200);
			match(VAR);
			setState(201);
			match(ID);
			setState(202);
			match(T__2);
			setState(203);
			match(T__3);
			setState(204);
			match(INT);
			setState(205);
			match(T__5);
			setState(206);
			match(INT);
			setState(207);
			match(T__4);
			setState(208);
			typeRule();
			setState(209);
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
		enterRule(_localctx, 18, RULE_varsTypeInit);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(213);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_TYPE:
			case FLOAT_TYPE:
			case CHAR_TYPE:
			case BOOL_TYPE:
				{
				setState(211);
				typeRule();
				}
				break;
			case ID:
				{
				setState(212);
				match(ID);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(216);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ASSIGN) {
				{
				setState(215);
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
		enterRule(_localctx, 20, RULE_varsTypeInit2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(218);
			match(ASSIGN);
			setState(219);
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
		public VarsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vars; }
	}

	public final VarsContext vars() throws RecognitionException {
		VarsContext _localctx = new VarsContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_vars);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(221);
			match(ID);
			setState(224);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(222);
				match(T__6);
				setState(223);
				match(ID);
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

	public static class VarArrayContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(ProyectoParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(ProyectoParser.ID, i);
		}
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public VarArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varArray; }
	}

	public final VarArrayContext varArray() throws RecognitionException {
		VarArrayContext _localctx = new VarArrayContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_varArray);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(226);
			match(ID);
			setState(229);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(227);
				match(T__6);
				setState(228);
				match(ID);
				}
			}

			setState(231);
			match(T__3);
			setState(232);
			exp();
			setState(233);
			match(T__4);
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

	public static class VarMatContext extends ParserRuleContext {
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
		public VarMatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varMat; }
	}

	public final VarMatContext varMat() throws RecognitionException {
		VarMatContext _localctx = new VarMatContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_varMat);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(235);
			match(ID);
			setState(238);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(236);
				match(T__6);
				setState(237);
				match(ID);
				}
			}

			setState(244);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				{
				setState(240);
				match(T__3);
				setState(241);
				exp();
				setState(242);
				match(T__4);
				}
				break;
			}
			setState(250);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(246);
				match(T__3);
				setState(247);
				exp();
				setState(248);
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
		enterRule(_localctx, 28, RULE_functions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(252);
			match(FUNCTION);
			setState(253);
			match(ID);
			setState(254);
			match(LPAREN);
			setState(256);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(255);
				parameters();
				}
			}

			setState(258);
			match(RPAREN);
			setState(261);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_TYPE:
			case FLOAT_TYPE:
			case CHAR_TYPE:
			case BOOL_TYPE:
				{
				setState(259);
				typeRule();
				}
				break;
			case VOID:
				{
				setState(260);
				match(VOID);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(263);
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
		enterRule(_localctx, 30, RULE_parameters);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(265);
			parameter();
			setState(270);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__7) {
				{
				{
				setState(266);
				match(T__7);
				setState(267);
				parameter();
				}
				}
				setState(272);
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
		enterRule(_localctx, 32, RULE_parameter);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(273);
			match(ID);
			setState(274);
			match(T__2);
			setState(275);
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
		public List<VariableDeclarationContext> variableDeclaration() {
			return getRuleContexts(VariableDeclarationContext.class);
		}
		public VariableDeclarationContext variableDeclaration(int i) {
			return getRuleContext(VariableDeclarationContext.class,i);
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
		enterRule(_localctx, 34, RULE_functionBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(277);
			match(T__0);
			setState(281);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(278);
				variableDeclaration();
				}
				}
				setState(283);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(287);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << RETURN) | (1L << ID))) != 0)) {
				{
				{
				setState(284);
				statutes();
				}
				}
				setState(289);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(290);
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
		enterRule(_localctx, 36, RULE_returnRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(292);
			match(RETURN);
			setState(294);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << ID))) != 0)) {
				{
				setState(293);
				exp();
				}
			}

			setState(296);
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
		enterRule(_localctx, 38, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(298);
			match(T__0);
			setState(302);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << RETURN) | (1L << ID))) != 0)) {
				{
				{
				setState(299);
				statutes();
				}
				}
				setState(304);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(305);
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
		public ReturnRuleContext returnRule() {
			return getRuleContext(ReturnRuleContext.class,0);
		}
		public StatutesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statutes; }
	}

	public final StatutesContext statutes() throws RecognitionException {
		StatutesContext _localctx = new StatutesContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_statutes);
		try {
			setState(315);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(307);
				assignation();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(308);
				read();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(309);
				write();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(310);
				conditional();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(311);
				forLoop();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(312);
				whileLoop();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(313);
				expression();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(314);
				returnRule();
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
		public TerminalNode ASSIGN() { return getToken(ProyectoParser.ASSIGN, 0); }
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(ProyectoParser.SEMI, 0); }
		public VarsContext vars() {
			return getRuleContext(VarsContext.class,0);
		}
		public VarArrayContext varArray() {
			return getRuleContext(VarArrayContext.class,0);
		}
		public VarMatContext varMat() {
			return getRuleContext(VarMatContext.class,0);
		}
		public AssignationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignation; }
	}

	public final AssignationContext assignation() throws RecognitionException {
		AssignationContext _localctx = new AssignationContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_assignation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(320);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				{
				setState(317);
				vars();
				}
				break;
			case 2:
				{
				setState(318);
				varArray();
				}
				break;
			case 3:
				{
				setState(319);
				varMat();
				}
				break;
			}
			setState(322);
			match(ASSIGN);
			setState(323);
			exp();
			setState(324);
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
		enterRule(_localctx, 44, RULE_functionCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(326);
			match(ID);
			setState(327);
			match(LPAREN);
			setState(329);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << ID))) != 0)) {
				{
				setState(328);
				arguments();
				}
			}

			setState(331);
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
		enterRule(_localctx, 46, RULE_arguments);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(333);
			exp();
			setState(334);
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
		enterRule(_localctx, 48, RULE_arguments2);
		try {
			setState(339);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__7:
				enterOuterAlt(_localctx, 1);
				{
				setState(336);
				match(T__7);
				setState(337);
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
		enterRule(_localctx, 50, RULE_methodCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(341);
			match(ID);
			setState(342);
			match(T__6);
			setState(343);
			match(ID);
			setState(344);
			match(LPAREN);
			setState(346);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << ID))) != 0)) {
				{
				setState(345);
				arguments();
				}
			}

			setState(348);
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
		enterRule(_localctx, 52, RULE_call);
		try {
			setState(352);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(350);
				functionCall();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(351);
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
		enterRule(_localctx, 54, RULE_read);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(354);
			match(READ);
			setState(355);
			match(LPAREN);
			setState(356);
			read2();
			setState(357);
			match(RPAREN);
			setState(358);
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
		public Read3Context read3() {
			return getRuleContext(Read3Context.class,0);
		}
		public VarsContext vars() {
			return getRuleContext(VarsContext.class,0);
		}
		public VarArrayContext varArray() {
			return getRuleContext(VarArrayContext.class,0);
		}
		public VarMatContext varMat() {
			return getRuleContext(VarMatContext.class,0);
		}
		public Read2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_read2; }
	}

	public final Read2Context read2() throws RecognitionException {
		Read2Context _localctx = new Read2Context(_ctx, getState());
		enterRule(_localctx, 56, RULE_read2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				{
				setState(360);
				vars();
				}
				break;
			case 2:
				{
				setState(361);
				varArray();
				}
				break;
			case 3:
				{
				setState(362);
				varMat();
				}
				break;
			}
			setState(365);
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
		enterRule(_localctx, 58, RULE_read3);
		try {
			setState(370);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__7:
				enterOuterAlt(_localctx, 1);
				{
				setState(367);
				match(T__7);
				setState(368);
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
		enterRule(_localctx, 60, RULE_write);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(372);
			match(WRITE);
			setState(373);
			match(LPAREN);
			setState(374);
			arguments();
			setState(375);
			match(RPAREN);
			setState(376);
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
		enterRule(_localctx, 62, RULE_conditional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(378);
			match(IF);
			setState(379);
			conditional2();
			setState(380);
			conditional3();
			setState(382);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(381);
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
		enterRule(_localctx, 64, RULE_conditional2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(384);
			match(LPAREN);
			setState(385);
			exp();
			setState(386);
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
		enterRule(_localctx, 66, RULE_conditional3);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(388);
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
		enterRule(_localctx, 68, RULE_conditional4);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(390);
			match(ELSE);
			setState(391);
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
		enterRule(_localctx, 70, RULE_forLoop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(393);
			match(FOR);
			setState(394);
			forLoop2();
			setState(395);
			match(IN);
			setState(396);
			forLoop3();
			setState(397);
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
		enterRule(_localctx, 72, RULE_forLoop2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(399);
			match(ID);
			setState(400);
			match(ASSIGN);
			setState(401);
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
		enterRule(_localctx, 74, RULE_forLoop3);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(403);
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
		enterRule(_localctx, 76, RULE_whileLoop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(405);
			match(WHILE);
			setState(406);
			whileLoop2();
			setState(407);
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
		enterRule(_localctx, 78, RULE_whileLoop2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(409);
			match(LPAREN);
			setState(410);
			exp();
			setState(411);
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
		enterRule(_localctx, 80, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(413);
			exp();
			setState(414);
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
		enterRule(_localctx, 82, RULE_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(416);
			t_exp();
			setState(420);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OR) {
				{
				{
				setState(417);
				exp2();
				}
				}
				setState(422);
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
		enterRule(_localctx, 84, RULE_exp2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(423);
			match(OR);
			setState(424);
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
		enterRule(_localctx, 86, RULE_t_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(426);
			g_exp();
			setState(430);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==AND) {
				{
				{
				setState(427);
				t_exp2();
				}
				}
				setState(432);
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
		enterRule(_localctx, 88, RULE_t_exp2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(433);
			match(AND);
			setState(434);
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
		enterRule(_localctx, 90, RULE_g_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(436);
			m_exp();
			setState(438);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << GT) | (1L << LT) | (1L << EQ) | (1L << NEQ))) != 0)) {
				{
				setState(437);
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
		enterRule(_localctx, 92, RULE_g_exp2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(440);
			relop();
			setState(441);
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
		enterRule(_localctx, 94, RULE_m_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(443);
			term();
			setState(447);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ADD || _la==SUB) {
				{
				{
				setState(444);
				m_exp2();
				}
				}
				setState(449);
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
		enterRule(_localctx, 96, RULE_m_exp2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(450);
			_la = _input.LA(1);
			if ( !(_la==ADD || _la==SUB) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(451);
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
		enterRule(_localctx, 98, RULE_term);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(453);
			factor();
			setState(457);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==MUL || _la==DIV) {
				{
				{
				setState(454);
				term2();
				}
				}
				setState(459);
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
		enterRule(_localctx, 100, RULE_term2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(460);
			_la = _input.LA(1);
			if ( !(_la==MUL || _la==DIV) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(461);
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
		public VarArrayContext varArray() {
			return getRuleContext(VarArrayContext.class,0);
		}
		public VarMatContext varMat() {
			return getRuleContext(VarMatContext.class,0);
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
		enterRule(_localctx, 102, RULE_factor);
		try {
			setState(471);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(463);
				factor2();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(464);
				varCte();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(468);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,35,_ctx) ) {
				case 1:
					{
					setState(465);
					vars();
					}
					break;
				case 2:
					{
					setState(466);
					varArray();
					}
					break;
				case 3:
					{
					setState(467);
					varMat();
					}
					break;
				}
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(470);
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
		enterRule(_localctx, 104, RULE_factor2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(473);
			match(LPAREN);
			setState(474);
			exp();
			setState(475);
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
		enterRule(_localctx, 106, RULE_varCte);
		try {
			setState(482);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(477);
				cte_i();
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(478);
				cte_f();
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 3);
				{
				setState(479);
				cte_c();
				}
				break;
			case T__8:
				enterOuterAlt(_localctx, 4);
				{
				setState(480);
				cte_s();
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 5);
				{
				setState(481);
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
		enterRule(_localctx, 108, RULE_cte_i);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(484);
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
		enterRule(_localctx, 110, RULE_cte_f);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(486);
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
		enterRule(_localctx, 112, RULE_cte_c);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(488);
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
		enterRule(_localctx, 114, RULE_cte_b);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(490);
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
		enterRule(_localctx, 116, RULE_cte_s);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(492);
			match(T__8);
			setState(496);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,38,_ctx);
			while ( _alt!=1 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1+1 ) {
					{
					{
					setState(493);
					matchWildcard();
					}
					} 
				}
				setState(498);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,38,_ctx);
			}
			setState(499);
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
		enterRule(_localctx, 118, RULE_main);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(501);
			match(MAIN);
			setState(502);
			match(LPAREN);
			setState(503);
			match(RPAREN);
			setState(504);
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
		public List<VariableDeclarationContext> variableDeclaration() {
			return getRuleContexts(VariableDeclarationContext.class);
		}
		public VariableDeclarationContext variableDeclaration(int i) {
			return getRuleContext(VariableDeclarationContext.class,i);
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
		enterRule(_localctx, 120, RULE_mainBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(506);
			match(T__0);
			setState(510);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(507);
				variableDeclaration();
				}
				}
				setState(512);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(516);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << RETURN) | (1L << ID))) != 0)) {
				{
				{
				setState(513);
				statutes();
				}
				}
				setState(518);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(519);
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
		enterRule(_localctx, 122, RULE_typeRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(521);
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
		enterRule(_localctx, 124, RULE_relop);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(523);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\63\u0210\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\3\2\3\2\3\2\3\2\7\2\u0085\n\2\f\2\16\2\u0088\13\2\3"+
		"\2\7\2\u008b\n\2\f\2\16\2\u008e\13\2\3\2\3\2\3\3\7\3\u0093\n\3\f\3\16"+
		"\3\u0096\13\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\5\4\u00a0\n\4\3\4\3\4\3"+
		"\4\3\5\3\5\3\5\3\5\7\5\u00a9\n\5\f\5\16\5\u00ac\13\5\3\5\3\5\3\6\3\6\7"+
		"\6\u00b2\n\6\f\6\16\6\u00b5\13\6\3\7\3\7\3\7\5\7\u00ba\n\7\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\5\13\u00d8\n\13\3\13\5\13\u00db\n\13"+
		"\3\f\3\f\3\f\3\r\3\r\3\r\5\r\u00e3\n\r\3\16\3\16\3\16\5\16\u00e8\n\16"+
		"\3\16\3\16\3\16\3\16\3\17\3\17\3\17\5\17\u00f1\n\17\3\17\3\17\3\17\3\17"+
		"\5\17\u00f7\n\17\3\17\3\17\3\17\3\17\5\17\u00fd\n\17\3\20\3\20\3\20\3"+
		"\20\5\20\u0103\n\20\3\20\3\20\3\20\5\20\u0108\n\20\3\20\3\20\3\21\3\21"+
		"\3\21\7\21\u010f\n\21\f\21\16\21\u0112\13\21\3\22\3\22\3\22\3\22\3\23"+
		"\3\23\7\23\u011a\n\23\f\23\16\23\u011d\13\23\3\23\7\23\u0120\n\23\f\23"+
		"\16\23\u0123\13\23\3\23\3\23\3\24\3\24\5\24\u0129\n\24\3\24\3\24\3\25"+
		"\3\25\7\25\u012f\n\25\f\25\16\25\u0132\13\25\3\25\3\25\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\5\26\u013e\n\26\3\27\3\27\3\27\5\27\u0143\n"+
		"\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\5\30\u014c\n\30\3\30\3\30\3\31"+
		"\3\31\3\31\3\32\3\32\3\32\5\32\u0156\n\32\3\33\3\33\3\33\3\33\3\33\5\33"+
		"\u015d\n\33\3\33\3\33\3\34\3\34\5\34\u0163\n\34\3\35\3\35\3\35\3\35\3"+
		"\35\3\35\3\36\3\36\3\36\5\36\u016e\n\36\3\36\3\36\3\37\3\37\3\37\5\37"+
		"\u0175\n\37\3 \3 \3 \3 \3 \3 \3!\3!\3!\3!\5!\u0181\n!\3\"\3\"\3\"\3\""+
		"\3#\3#\3$\3$\3$\3%\3%\3%\3%\3%\3%\3&\3&\3&\3&\3\'\3\'\3(\3(\3(\3(\3)\3"+
		")\3)\3)\3*\3*\3*\3+\3+\7+\u01a5\n+\f+\16+\u01a8\13+\3,\3,\3,\3-\3-\7-"+
		"\u01af\n-\f-\16-\u01b2\13-\3.\3.\3.\3/\3/\5/\u01b9\n/\3\60\3\60\3\60\3"+
		"\61\3\61\7\61\u01c0\n\61\f\61\16\61\u01c3\13\61\3\62\3\62\3\62\3\63\3"+
		"\63\7\63\u01ca\n\63\f\63\16\63\u01cd\13\63\3\64\3\64\3\64\3\65\3\65\3"+
		"\65\3\65\3\65\5\65\u01d7\n\65\3\65\5\65\u01da\n\65\3\66\3\66\3\66\3\66"+
		"\3\67\3\67\3\67\3\67\3\67\5\67\u01e5\n\67\38\38\39\39\3:\3:\3;\3;\3<\3"+
		"<\7<\u01f1\n<\f<\16<\u01f4\13<\3<\3<\3=\3=\3=\3=\3=\3>\3>\7>\u01ff\n>"+
		"\f>\16>\u0202\13>\3>\7>\u0205\n>\f>\16>\u0208\13>\3>\3>\3?\3?\3@\3@\3"+
		"@\3\u01f2\2A\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64"+
		"\668:<>@BDFHJLNPRTVXZ\\^`bdfhjlnprtvxz|~\2\6\3\2\31\32\3\2\27\30\3\2-"+
		"\60\3\2\23\26\2\u0208\2\u0080\3\2\2\2\4\u0094\3\2\2\2\6\u009a\3\2\2\2"+
		"\b\u00a4\3\2\2\2\n\u00af\3\2\2\2\f\u00b9\3\2\2\2\16\u00bb\3\2\2\2\20\u00c1"+
		"\3\2\2\2\22\u00ca\3\2\2\2\24\u00d7\3\2\2\2\26\u00dc\3\2\2\2\30\u00df\3"+
		"\2\2\2\32\u00e4\3\2\2\2\34\u00ed\3\2\2\2\36\u00fe\3\2\2\2 \u010b\3\2\2"+
		"\2\"\u0113\3\2\2\2$\u0117\3\2\2\2&\u0126\3\2\2\2(\u012c\3\2\2\2*\u013d"+
		"\3\2\2\2,\u0142\3\2\2\2.\u0148\3\2\2\2\60\u014f\3\2\2\2\62\u0155\3\2\2"+
		"\2\64\u0157\3\2\2\2\66\u0162\3\2\2\28\u0164\3\2\2\2:\u016d\3\2\2\2<\u0174"+
		"\3\2\2\2>\u0176\3\2\2\2@\u017c\3\2\2\2B\u0182\3\2\2\2D\u0186\3\2\2\2F"+
		"\u0188\3\2\2\2H\u018b\3\2\2\2J\u0191\3\2\2\2L\u0195\3\2\2\2N\u0197\3\2"+
		"\2\2P\u019b\3\2\2\2R\u019f\3\2\2\2T\u01a2\3\2\2\2V\u01a9\3\2\2\2X\u01ac"+
		"\3\2\2\2Z\u01b3\3\2\2\2\\\u01b6\3\2\2\2^\u01ba\3\2\2\2`\u01bd\3\2\2\2"+
		"b\u01c4\3\2\2\2d\u01c7\3\2\2\2f\u01ce\3\2\2\2h\u01d9\3\2\2\2j\u01db\3"+
		"\2\2\2l\u01e4\3\2\2\2n\u01e6\3\2\2\2p\u01e8\3\2\2\2r\u01ea\3\2\2\2t\u01ec"+
		"\3\2\2\2v\u01ee\3\2\2\2x\u01f7\3\2\2\2z\u01fc\3\2\2\2|\u020b\3\2\2\2~"+
		"\u020d\3\2\2\2\u0080\u0081\7\62\2\2\u0081\u0082\7\63\2\2\u0082\u0086\7"+
		"\33\2\2\u0083\u0085\5\6\4\2\u0084\u0083\3\2\2\2\u0085\u0088\3\2\2\2\u0086"+
		"\u0084\3\2\2\2\u0086\u0087\3\2\2\2\u0087\u008c\3\2\2\2\u0088\u0086\3\2"+
		"\2\2\u0089\u008b\5\f\7\2\u008a\u0089\3\2\2\2\u008b\u008e\3\2\2\2\u008c"+
		"\u008a\3\2\2\2\u008c\u008d\3\2\2\2\u008d\u008f\3\2\2\2\u008e\u008c\3\2"+
		"\2\2\u008f\u0090\5\4\3\2\u0090\3\3\2\2\2\u0091\u0093\5\36\20\2\u0092\u0091"+
		"\3\2\2\2\u0093\u0096\3\2\2\2\u0094\u0092\3\2\2\2\u0094\u0095\3\2\2\2\u0095"+
		"\u0097\3\2\2\2\u0096\u0094\3\2\2\2\u0097\u0098\5x=\2\u0098\u0099\7\2\2"+
		"\3\u0099\5\3\2\2\2\u009a\u009b\7$\2\2\u009b\u009f\7\63\2\2\u009c\u009d"+
		"\7\24\2\2\u009d\u009e\7\63\2\2\u009e\u00a0\7\23\2\2\u009f\u009c\3\2\2"+
		"\2\u009f\u00a0\3\2\2\2\u00a0\u00a1\3\2\2\2\u00a1\u00a2\5\b\5\2\u00a2\u00a3"+
		"\7\33\2\2\u00a3\7\3\2\2\2\u00a4\u00a5\7\3\2\2\u00a5\u00a6\5\n\6\2\u00a6"+
		"\u00aa\7&\2\2\u00a7\u00a9\5\36\20\2\u00a8\u00a7\3\2\2\2\u00a9\u00ac\3"+
		"\2\2\2\u00aa\u00a8\3\2\2\2\u00aa\u00ab\3\2\2\2\u00ab\u00ad\3\2\2\2\u00ac"+
		"\u00aa\3\2\2\2\u00ad\u00ae\7\4\2\2\u00ae\t\3\2\2\2\u00af\u00b3\7%\2\2"+
		"\u00b0\u00b2\5\f\7\2\u00b1\u00b0\3\2\2\2\u00b2\u00b5\3\2\2\2\u00b3\u00b1"+
		"\3\2\2\2\u00b3\u00b4\3\2\2\2\u00b4\13\3\2\2\2\u00b5\u00b3\3\2\2\2\u00b6"+
		"\u00ba\5\16\b\2\u00b7\u00ba\5\20\t\2\u00b8\u00ba\5\22\n\2\u00b9\u00b6"+
		"\3\2\2\2\u00b9\u00b7\3\2\2\2\u00b9\u00b8\3\2\2\2\u00ba\r\3\2\2\2\u00bb"+
		"\u00bc\7,\2\2\u00bc\u00bd\7\63\2\2\u00bd\u00be\7\5\2\2\u00be\u00bf\5\24"+
		"\13\2\u00bf\u00c0\7\33\2\2\u00c0\17\3\2\2\2\u00c1\u00c2\7,\2\2\u00c2\u00c3"+
		"\7\63\2\2\u00c3\u00c4\7\5\2\2\u00c4\u00c5\7\6\2\2\u00c5\u00c6\7\f\2\2"+
		"\u00c6\u00c7\7\7\2\2\u00c7\u00c8\5|?\2\u00c8\u00c9\7\33\2\2\u00c9\21\3"+
		"\2\2\2\u00ca\u00cb\7,\2\2\u00cb\u00cc\7\63\2\2\u00cc\u00cd\7\5\2\2\u00cd"+
		"\u00ce\7\6\2\2\u00ce\u00cf\7\f\2\2\u00cf\u00d0\7\b\2\2\u00d0\u00d1\7\f"+
		"\2\2\u00d1\u00d2\7\7\2\2\u00d2\u00d3\5|?\2\u00d3\u00d4\7\33\2\2\u00d4"+
		"\23\3\2\2\2\u00d5\u00d8\5|?\2\u00d6\u00d8\7\63\2\2\u00d7\u00d5\3\2\2\2"+
		"\u00d7\u00d6\3\2\2\2\u00d8\u00da\3\2\2\2\u00d9\u00db\5\26\f\2\u00da\u00d9"+
		"\3\2\2\2\u00da\u00db\3\2\2\2\u00db\25\3\2\2\2\u00dc\u00dd\7\34\2\2\u00dd"+
		"\u00de\5T+\2\u00de\27\3\2\2\2\u00df\u00e2\7\63\2\2\u00e0\u00e1\7\t\2\2"+
		"\u00e1\u00e3\7\63\2\2\u00e2\u00e0\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3\31"+
		"\3\2\2\2\u00e4\u00e7\7\63\2\2\u00e5\u00e6\7\t\2\2\u00e6\u00e8\7\63\2\2"+
		"\u00e7\u00e5\3\2\2\2\u00e7\u00e8\3\2\2\2\u00e8\u00e9\3\2\2\2\u00e9\u00ea"+
		"\7\6\2\2\u00ea\u00eb\5T+\2\u00eb\u00ec\7\7\2\2\u00ec\33\3\2\2\2\u00ed"+
		"\u00f0\7\63\2\2\u00ee\u00ef\7\t\2\2\u00ef\u00f1\7\63\2\2\u00f0\u00ee\3"+
		"\2\2\2\u00f0\u00f1\3\2\2\2\u00f1\u00f6\3\2\2\2\u00f2\u00f3\7\6\2\2\u00f3"+
		"\u00f4\5T+\2\u00f4\u00f5\7\7\2\2\u00f5\u00f7\3\2\2\2\u00f6\u00f2\3\2\2"+
		"\2\u00f6\u00f7\3\2\2\2\u00f7\u00fc\3\2\2\2\u00f8\u00f9\7\6\2\2\u00f9\u00fa"+
		"\5T+\2\u00fa\u00fb\7\7\2\2\u00fb\u00fd\3\2\2\2\u00fc\u00f8\3\2\2\2\u00fc"+
		"\u00fd\3\2\2\2\u00fd\35\3\2\2\2\u00fe\u00ff\7)\2\2\u00ff\u0100\7\63\2"+
		"\2\u0100\u0102\7\35\2\2\u0101\u0103\5 \21\2\u0102\u0101\3\2\2\2\u0102"+
		"\u0103\3\2\2\2\u0103\u0104\3\2\2\2\u0104\u0107\7\36\2\2\u0105\u0108\5"+
		"|?\2\u0106\u0108\7\61\2\2\u0107\u0105\3\2\2\2\u0107\u0106\3\2\2\2\u0108"+
		"\u0109\3\2\2\2\u0109\u010a\5$\23\2\u010a\37\3\2\2\2\u010b\u0110\5\"\22"+
		"\2\u010c\u010d\7\n\2\2\u010d\u010f\5\"\22\2\u010e\u010c\3\2\2\2\u010f"+
		"\u0112\3\2\2\2\u0110\u010e\3\2\2\2\u0110\u0111\3\2\2\2\u0111!\3\2\2\2"+
		"\u0112\u0110\3\2\2\2\u0113\u0114\7\63\2\2\u0114\u0115\7\5\2\2\u0115\u0116"+
		"\5|?\2\u0116#\3\2\2\2\u0117\u011b\7\3\2\2\u0118\u011a\5\f\7\2\u0119\u0118"+
		"\3\2\2\2\u011a\u011d\3\2\2\2\u011b\u0119\3\2\2\2\u011b\u011c\3\2\2\2\u011c"+
		"\u0121\3\2\2\2\u011d\u011b\3\2\2\2\u011e\u0120\5*\26\2\u011f\u011e\3\2"+
		"\2\2\u0120\u0123\3\2\2\2\u0121\u011f\3\2\2\2\u0121\u0122\3\2\2\2\u0122"+
		"\u0124\3\2\2\2\u0123\u0121\3\2\2\2\u0124\u0125\7\4\2\2\u0125%\3\2\2\2"+
		"\u0126\u0128\7*\2\2\u0127\u0129\5T+\2\u0128\u0127\3\2\2\2\u0128\u0129"+
		"\3\2\2\2\u0129\u012a\3\2\2\2\u012a\u012b\7\33\2\2\u012b\'\3\2\2\2\u012c"+
		"\u0130\7\3\2\2\u012d\u012f\5*\26\2\u012e\u012d\3\2\2\2\u012f\u0132\3\2"+
		"\2\2\u0130\u012e\3\2\2\2\u0130\u0131\3\2\2\2\u0131\u0133\3\2\2\2\u0132"+
		"\u0130\3\2\2\2\u0133\u0134\7\4\2\2\u0134)\3\2\2\2\u0135\u013e\5,\27\2"+
		"\u0136\u013e\58\35\2\u0137\u013e\5> \2\u0138\u013e\5@!\2\u0139\u013e\5"+
		"H%\2\u013a\u013e\5N(\2\u013b\u013e\5R*\2\u013c\u013e\5&\24\2\u013d\u0135"+
		"\3\2\2\2\u013d\u0136\3\2\2\2\u013d\u0137\3\2\2\2\u013d\u0138\3\2\2\2\u013d"+
		"\u0139\3\2\2\2\u013d\u013a\3\2\2\2\u013d\u013b\3\2\2\2\u013d\u013c\3\2"+
		"\2\2\u013e+\3\2\2\2\u013f\u0143\5\30\r\2\u0140\u0143\5\32\16\2\u0141\u0143"+
		"\5\34\17\2\u0142\u013f\3\2\2\2\u0142\u0140\3\2\2\2\u0142\u0141\3\2\2\2"+
		"\u0143\u0144\3\2\2\2\u0144\u0145\7\34\2\2\u0145\u0146\5T+\2\u0146\u0147"+
		"\7\33\2\2\u0147-\3\2\2\2\u0148\u0149\7\63\2\2\u0149\u014b\7\35\2\2\u014a"+
		"\u014c\5\60\31\2\u014b\u014a\3\2\2\2\u014b\u014c\3\2\2\2\u014c\u014d\3"+
		"\2\2\2\u014d\u014e\7\36\2\2\u014e/\3\2\2\2\u014f\u0150\5T+\2\u0150\u0151"+
		"\5\62\32\2\u0151\61\3\2\2\2\u0152\u0153\7\n\2\2\u0153\u0156\5\60\31\2"+
		"\u0154\u0156\3\2\2\2\u0155\u0152\3\2\2\2\u0155\u0154\3\2\2\2\u0156\63"+
		"\3\2\2\2\u0157\u0158\7\63\2\2\u0158\u0159\7\t\2\2\u0159\u015a\7\63\2\2"+
		"\u015a\u015c\7\35\2\2\u015b\u015d\5\60\31\2\u015c\u015b\3\2\2\2\u015c"+
		"\u015d\3\2\2\2\u015d\u015e\3\2\2\2\u015e\u015f\7\36\2\2\u015f\65\3\2\2"+
		"\2\u0160\u0163\5.\30\2\u0161\u0163\5\64\33\2\u0162\u0160\3\2\2\2\u0162"+
		"\u0161\3\2\2\2\u0163\67\3\2\2\2\u0164\u0165\7(\2\2\u0165\u0166\7\35\2"+
		"\2\u0166\u0167\5:\36\2\u0167\u0168\7\36\2\2\u0168\u0169\7\33\2\2\u0169"+
		"9\3\2\2\2\u016a\u016e\5\30\r\2\u016b\u016e\5\32\16\2\u016c\u016e\5\34"+
		"\17\2\u016d\u016a\3\2\2\2\u016d\u016b\3\2\2\2\u016d\u016c\3\2\2\2\u016e"+
		"\u016f\3\2\2\2\u016f\u0170\5<\37\2\u0170;\3\2\2\2\u0171\u0172\7\n\2\2"+
		"\u0172\u0175\5:\36\2\u0173\u0175\3\2\2\2\u0174\u0171\3\2\2\2\u0174\u0173"+
		"\3\2\2\2\u0175=\3\2\2\2\u0176\u0177\7\'\2\2\u0177\u0178\7\35\2\2\u0178"+
		"\u0179\5\60\31\2\u0179\u017a\7\36\2\2\u017a\u017b\7\33\2\2\u017b?\3\2"+
		"\2\2\u017c\u017d\7\37\2\2\u017d\u017e\5B\"\2\u017e\u0180\5D#\2\u017f\u0181"+
		"\5F$\2\u0180\u017f\3\2\2\2\u0180\u0181\3\2\2\2\u0181A\3\2\2\2\u0182\u0183"+
		"\7\35\2\2\u0183\u0184\5T+\2\u0184\u0185\7\36\2\2\u0185C\3\2\2\2\u0186"+
		"\u0187\5(\25\2\u0187E\3\2\2\2\u0188\u0189\7 \2\2\u0189\u018a\5(\25\2\u018a"+
		"G\3\2\2\2\u018b\u018c\7\"\2\2\u018c\u018d\5J&\2\u018d\u018e\7#\2\2\u018e"+
		"\u018f\5L\'\2\u018f\u0190\5(\25\2\u0190I\3\2\2\2\u0191\u0192\7\63\2\2"+
		"\u0192\u0193\7\34\2\2\u0193\u0194\5T+\2\u0194K\3\2\2\2\u0195\u0196\5T"+
		"+\2\u0196M\3\2\2\2\u0197\u0198\7!\2\2\u0198\u0199\5P)\2\u0199\u019a\5"+
		"(\25\2\u019aO\3\2\2\2\u019b\u019c\7\35\2\2\u019c\u019d\5T+\2\u019d\u019e"+
		"\7\36\2\2\u019eQ\3\2\2\2\u019f\u01a0\5T+\2\u01a0\u01a1\7\33\2\2\u01a1"+
		"S\3\2\2\2\u01a2\u01a6\5X-\2\u01a3\u01a5\5V,\2\u01a4\u01a3\3\2\2\2\u01a5"+
		"\u01a8\3\2\2\2\u01a6\u01a4\3\2\2\2\u01a6\u01a7\3\2\2\2\u01a7U\3\2\2\2"+
		"\u01a8\u01a6\3\2\2\2\u01a9\u01aa\7\21\2\2\u01aa\u01ab\5X-\2\u01abW\3\2"+
		"\2\2\u01ac\u01b0\5\\/\2\u01ad\u01af\5Z.\2\u01ae\u01ad\3\2\2\2\u01af\u01b2"+
		"\3\2\2\2\u01b0\u01ae\3\2\2\2\u01b0\u01b1\3\2\2\2\u01b1Y\3\2\2\2\u01b2"+
		"\u01b0\3\2\2\2\u01b3\u01b4\7\22\2\2\u01b4\u01b5\5\\/\2\u01b5[\3\2\2\2"+
		"\u01b6\u01b8\5`\61\2\u01b7\u01b9\5^\60\2\u01b8\u01b7\3\2\2\2\u01b8\u01b9"+
		"\3\2\2\2\u01b9]\3\2\2\2\u01ba\u01bb\5~@\2\u01bb\u01bc\5`\61\2\u01bc_\3"+
		"\2\2\2\u01bd\u01c1\5d\63\2\u01be\u01c0\5b\62\2\u01bf\u01be\3\2\2\2\u01c0"+
		"\u01c3\3\2\2\2\u01c1\u01bf\3\2\2\2\u01c1\u01c2\3\2\2\2\u01c2a\3\2\2\2"+
		"\u01c3\u01c1\3\2\2\2\u01c4\u01c5\t\2\2\2\u01c5\u01c6\5d\63\2\u01c6c\3"+
		"\2\2\2\u01c7\u01cb\5h\65\2\u01c8\u01ca\5f\64\2\u01c9\u01c8\3\2\2\2\u01ca"+
		"\u01cd\3\2\2\2\u01cb\u01c9\3\2\2\2\u01cb\u01cc\3\2\2\2\u01cce\3\2\2\2"+
		"\u01cd\u01cb\3\2\2\2\u01ce\u01cf\t\3\2\2\u01cf\u01d0\5h\65\2\u01d0g\3"+
		"\2\2\2\u01d1\u01da\5j\66\2\u01d2\u01da\5l\67\2\u01d3\u01d7\5\30\r\2\u01d4"+
		"\u01d7\5\32\16\2\u01d5\u01d7\5\34\17\2\u01d6\u01d3\3\2\2\2\u01d6\u01d4"+
		"\3\2\2\2\u01d6\u01d5\3\2\2\2\u01d7\u01da\3\2\2\2\u01d8\u01da\5\66\34\2"+
		"\u01d9\u01d1\3\2\2\2\u01d9\u01d2\3\2\2\2\u01d9\u01d6\3\2\2\2\u01d9\u01d8"+
		"\3\2\2\2\u01dai\3\2\2\2\u01db\u01dc\7\35\2\2\u01dc\u01dd\5T+\2\u01dd\u01de"+
		"\7\36\2\2\u01dek\3\2\2\2\u01df\u01e5\5n8\2\u01e0\u01e5\5p9\2\u01e1\u01e5"+
		"\5r:\2\u01e2\u01e5\5v<\2\u01e3\u01e5\5t;\2\u01e4\u01df\3\2\2\2\u01e4\u01e0"+
		"\3\2\2\2\u01e4\u01e1\3\2\2\2\u01e4\u01e2\3\2\2\2\u01e4\u01e3\3\2\2\2\u01e5"+
		"m\3\2\2\2\u01e6\u01e7\7\f\2\2\u01e7o\3\2\2\2\u01e8\u01e9\7\r\2\2\u01e9"+
		"q\3\2\2\2\u01ea\u01eb\7\16\2\2\u01ebs\3\2\2\2\u01ec\u01ed\7\17\2\2\u01ed"+
		"u\3\2\2\2\u01ee\u01f2\7\13\2\2\u01ef\u01f1\13\2\2\2\u01f0\u01ef\3\2\2"+
		"\2\u01f1\u01f4\3\2\2\2\u01f2\u01f3\3\2\2\2\u01f2\u01f0\3\2\2\2\u01f3\u01f5"+
		"\3\2\2\2\u01f4\u01f2\3\2\2\2\u01f5\u01f6\7\13\2\2\u01f6w\3\2\2\2\u01f7"+
		"\u01f8\7+\2\2\u01f8\u01f9\7\35\2\2\u01f9\u01fa\7\36\2\2\u01fa\u01fb\5"+
		"z>\2\u01fby\3\2\2\2\u01fc\u0200\7\3\2\2\u01fd\u01ff\5\f\7\2\u01fe\u01fd"+
		"\3\2\2\2\u01ff\u0202\3\2\2\2\u0200\u01fe\3\2\2\2\u0200\u0201\3\2\2\2\u0201"+
		"\u0206\3\2\2\2\u0202\u0200\3\2\2\2\u0203\u0205\5*\26\2\u0204\u0203\3\2"+
		"\2\2\u0205\u0208\3\2\2\2\u0206\u0204\3\2\2\2\u0206\u0207\3\2\2\2\u0207"+
		"\u0209\3\2\2\2\u0208\u0206\3\2\2\2\u0209\u020a\7\4\2\2\u020a{\3\2\2\2"+
		"\u020b\u020c\t\4\2\2\u020c}\3\2\2\2\u020d\u020e\t\5\2\2\u020e\177\3\2"+
		"\2\2+\u0086\u008c\u0094\u009f\u00aa\u00b3\u00b9\u00d7\u00da\u00e2\u00e7"+
		"\u00f0\u00f6\u00fc\u0102\u0107\u0110\u011b\u0121\u0128\u0130\u013d\u0142"+
		"\u014b\u0155\u015c\u0162\u016d\u0174\u0180\u01a6\u01b0\u01b8\u01c1\u01cb"+
		"\u01d6\u01d9\u01e4\u01f2\u0200\u0206";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}