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
		LT=19, EQ=20, NEQ=21, MUL=22, DIV=23, ADD=24, SUB=25, SEMI=26, ASSIGN=27, 
		LPAREN=28, RPAREN=29, IF=30, ELSE=31, WHILE=32, FOR=33, IN=34, CLASS=35, 
		ATTRIBUTES=36, METHODS=37, INIT=38, PRIVATE=39, PUBLIC=40, WRITE=41, READ=42, 
		FUNCTION=43, RETURN=44, MAIN=45, VAR=46, INT_TYPE=47, FLOAT_TYPE=48, CHAR_TYPE=49, 
		BOOL_TYPE=50, VOID=51, PROGRAM=52, ID=53;
	public static final int
		RULE_program = 0, RULE_program2 = 1, RULE_classDef = 2, RULE_classBlock = 3, 
		RULE_classAttributes = 4, RULE_classMethod = 5, RULE_classInit = 6, RULE_variableDeclaration = 7, 
		RULE_varsDec = 8, RULE_varsDecArray = 9, RULE_varsDecMat = 10, RULE_attributesDeclaration = 11, 
		RULE_attributesDec = 12, RULE_varsTypeInit = 13, RULE_varsTypeInit2 = 14, 
		RULE_vars = 15, RULE_varArray = 16, RULE_varMat = 17, RULE_functions = 18, 
		RULE_parameters = 19, RULE_parameter = 20, RULE_functionBlock = 21, RULE_returnRule = 22, 
		RULE_block = 23, RULE_statutesNoReturn = 24, RULE_statutes = 25, RULE_assignation = 26, 
		RULE_functionCall = 27, RULE_arguments = 28, RULE_arguments2 = 29, RULE_methodCall = 30, 
		RULE_initCall = 31, RULE_call = 32, RULE_read = 33, RULE_read2 = 34, RULE_read3 = 35, 
		RULE_write = 36, RULE_conditional = 37, RULE_conditional2 = 38, RULE_conditional3 = 39, 
		RULE_conditional4 = 40, RULE_forLoop = 41, RULE_forLoop2 = 42, RULE_forLoop3 = 43, 
		RULE_whileLoop = 44, RULE_whileLoop2 = 45, RULE_expression = 46, RULE_exp = 47, 
		RULE_exp2 = 48, RULE_t_exp = 49, RULE_t_exp2 = 50, RULE_g_exp = 51, RULE_g_exp2 = 52, 
		RULE_m_exp = 53, RULE_m_exp2 = 54, RULE_term = 55, RULE_term2 = 56, RULE_factor = 57, 
		RULE_factor2 = 58, RULE_varCte = 59, RULE_cte_i = 60, RULE_cte_f = 61, 
		RULE_cte_c = 62, RULE_cte_b = 63, RULE_cte_s = 64, RULE_main = 65, RULE_mainBlock = 66, 
		RULE_typeRule = 67, RULE_relop = 68;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "program2", "classDef", "classBlock", "classAttributes", "classMethod", 
			"classInit", "variableDeclaration", "varsDec", "varsDecArray", "varsDecMat", 
			"attributesDeclaration", "attributesDec", "varsTypeInit", "varsTypeInit2", 
			"vars", "varArray", "varMat", "functions", "parameters", "parameter", 
			"functionBlock", "returnRule", "block", "statutesNoReturn", "statutes", 
			"assignation", "functionCall", "arguments", "arguments2", "methodCall", 
			"initCall", "call", "read", "read2", "read3", "write", "conditional", 
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
			null, "'{'", "'}'", "':'", "'['", "']'", "']['", "'.'", "','", "'new'", 
			"'\"'", null, null, null, null, null, "'||'", "'&&'", "'>'", "'<'", "'=='", 
			"'!='", "'*'", "'/'", "'+'", "'-'", "';'", "'='", "'('", "')'", "'if'", 
			"'else'", "'while'", "'for'", "'in'", "'class'", "'attributes'", "'methods'", 
			"'init'", "'private'", "'public'", "'write'", "'read'", "'function'", 
			"'return'", "'main'", "'var'", "'int'", "'float'", "'char'", "'bool'", 
			"'void'", "'program'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, "INT", 
			"FLOAT", "CHAR", "BOOL", "WS", "OR", "AND", "GT", "LT", "EQ", "NEQ", 
			"MUL", "DIV", "ADD", "SUB", "SEMI", "ASSIGN", "LPAREN", "RPAREN", "IF", 
			"ELSE", "WHILE", "FOR", "IN", "CLASS", "ATTRIBUTES", "METHODS", "INIT", 
			"PRIVATE", "PUBLIC", "WRITE", "READ", "FUNCTION", "RETURN", "MAIN", "VAR", 
			"INT_TYPE", "FLOAT_TYPE", "CHAR_TYPE", "BOOL_TYPE", "VOID", "PROGRAM", 
			"ID"
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
			setState(138);
			match(PROGRAM);
			setState(139);
			match(ID);
			setState(140);
			match(SEMI);
			setState(144);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(141);
				variableDeclaration();
				}
				}
				setState(146);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(147);
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
		public List<ClassDefContext> classDef() {
			return getRuleContexts(ClassDefContext.class);
		}
		public ClassDefContext classDef(int i) {
			return getRuleContext(ClassDefContext.class,i);
		}
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
			setState(152);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CLASS) {
				{
				{
				setState(149);
				classDef();
				}
				}
				setState(154);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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
			main();
			setState(162);
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
		public TerminalNode ID() { return getToken(ProyectoParser.ID, 0); }
		public ClassBlockContext classBlock() {
			return getRuleContext(ClassBlockContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(ProyectoParser.SEMI, 0); }
		public ClassDefContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classDef; }
	}

	public final ClassDefContext classDef() throws RecognitionException {
		ClassDefContext _localctx = new ClassDefContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_classDef);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(164);
			match(CLASS);
			setState(165);
			match(ID);
			setState(166);
			classBlock();
			setState(167);
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
		public ClassInitContext classInit() {
			return getRuleContext(ClassInitContext.class,0);
		}
		public TerminalNode METHODS() { return getToken(ProyectoParser.METHODS, 0); }
		public List<ClassMethodContext> classMethod() {
			return getRuleContexts(ClassMethodContext.class);
		}
		public ClassMethodContext classMethod(int i) {
			return getRuleContext(ClassMethodContext.class,i);
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
			setState(169);
			match(T__0);
			setState(170);
			classAttributes();
			setState(171);
			classInit();
			setState(172);
			match(METHODS);
			setState(176);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PRIVATE) | (1L << PUBLIC) | (1L << FUNCTION))) != 0)) {
				{
				{
				setState(173);
				classMethod();
				}
				}
				setState(178);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(179);
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
		public List<AttributesDeclarationContext> attributesDeclaration() {
			return getRuleContexts(AttributesDeclarationContext.class);
		}
		public AttributesDeclarationContext attributesDeclaration(int i) {
			return getRuleContext(AttributesDeclarationContext.class,i);
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
			setState(181);
			match(ATTRIBUTES);
			setState(185);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PRIVATE) | (1L << PUBLIC) | (1L << VAR))) != 0)) {
				{
				{
				setState(182);
				attributesDeclaration();
				}
				}
				setState(187);
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

	public static class ClassMethodContext extends ParserRuleContext {
		public FunctionsContext functions() {
			return getRuleContext(FunctionsContext.class,0);
		}
		public TerminalNode PRIVATE() { return getToken(ProyectoParser.PRIVATE, 0); }
		public TerminalNode PUBLIC() { return getToken(ProyectoParser.PUBLIC, 0); }
		public ClassMethodContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classMethod; }
	}

	public final ClassMethodContext classMethod() throws RecognitionException {
		ClassMethodContext _localctx = new ClassMethodContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_classMethod);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(192);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRIVATE:
				{
				setState(188);
				match(PRIVATE);
				}
				break;
			case PUBLIC:
			case FUNCTION:
				{
				setState(190);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUBLIC) {
					{
					setState(189);
					match(PUBLIC);
					}
				}

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(194);
			functions();
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

	public static class ClassInitContext extends ParserRuleContext {
		public TerminalNode INIT() { return getToken(ProyectoParser.INIT, 0); }
		public TerminalNode LPAREN() { return getToken(ProyectoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(ProyectoParser.RPAREN, 0); }
		public List<VariableDeclarationContext> variableDeclaration() {
			return getRuleContexts(VariableDeclarationContext.class);
		}
		public VariableDeclarationContext variableDeclaration(int i) {
			return getRuleContext(VariableDeclarationContext.class,i);
		}
		public List<StatutesNoReturnContext> statutesNoReturn() {
			return getRuleContexts(StatutesNoReturnContext.class);
		}
		public StatutesNoReturnContext statutesNoReturn(int i) {
			return getRuleContext(StatutesNoReturnContext.class,i);
		}
		public ClassInitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classInit; }
	}

	public final ClassInitContext classInit() throws RecognitionException {
		ClassInitContext _localctx = new ClassInitContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_classInit);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(196);
			match(INIT);
			setState(197);
			match(LPAREN);
			setState(198);
			match(RPAREN);
			setState(199);
			match(T__0);
			setState(203);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(200);
				variableDeclaration();
				}
				}
				setState(205);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(209);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << ID))) != 0)) {
				{
				{
				setState(206);
				statutesNoReturn();
				}
				}
				setState(211);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(212);
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
		enterRule(_localctx, 14, RULE_variableDeclaration);
		try {
			setState(217);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(214);
				varsDec();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(215);
				varsDecArray();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(216);
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
		enterRule(_localctx, 16, RULE_varsDec);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(219);
			match(VAR);
			setState(220);
			match(ID);
			setState(221);
			match(T__2);
			setState(222);
			varsTypeInit();
			setState(223);
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
		enterRule(_localctx, 18, RULE_varsDecArray);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(225);
			match(VAR);
			setState(226);
			match(ID);
			setState(227);
			match(T__2);
			setState(228);
			match(T__3);
			setState(229);
			match(INT);
			setState(230);
			match(T__4);
			setState(231);
			typeRule();
			setState(232);
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
		enterRule(_localctx, 20, RULE_varsDecMat);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(234);
			match(VAR);
			setState(235);
			match(ID);
			setState(236);
			match(T__2);
			setState(237);
			match(T__3);
			setState(238);
			match(INT);
			setState(239);
			match(T__5);
			setState(240);
			match(INT);
			setState(241);
			match(T__4);
			setState(242);
			typeRule();
			setState(243);
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

	public static class AttributesDeclarationContext extends ParserRuleContext {
		public TerminalNode PRIVATE() { return getToken(ProyectoParser.PRIVATE, 0); }
		public AttributesDecContext attributesDec() {
			return getRuleContext(AttributesDecContext.class,0);
		}
		public VarsDecArrayContext varsDecArray() {
			return getRuleContext(VarsDecArrayContext.class,0);
		}
		public VarsDecMatContext varsDecMat() {
			return getRuleContext(VarsDecMatContext.class,0);
		}
		public TerminalNode PUBLIC() { return getToken(ProyectoParser.PUBLIC, 0); }
		public AttributesDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributesDeclaration; }
	}

	public final AttributesDeclarationContext attributesDeclaration() throws RecognitionException {
		AttributesDeclarationContext _localctx = new AttributesDeclarationContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_attributesDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(249);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRIVATE:
				{
				setState(245);
				match(PRIVATE);
				}
				break;
			case PUBLIC:
			case VAR:
				{
				setState(247);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUBLIC) {
					{
					setState(246);
					match(PUBLIC);
					}
				}

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(254);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				{
				setState(251);
				attributesDec();
				}
				break;
			case 2:
				{
				setState(252);
				varsDecArray();
				}
				break;
			case 3:
				{
				setState(253);
				varsDecMat();
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

	public static class AttributesDecContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(ProyectoParser.VAR, 0); }
		public TerminalNode ID() { return getToken(ProyectoParser.ID, 0); }
		public TypeRuleContext typeRule() {
			return getRuleContext(TypeRuleContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(ProyectoParser.SEMI, 0); }
		public AttributesDecContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributesDec; }
	}

	public final AttributesDecContext attributesDec() throws RecognitionException {
		AttributesDecContext _localctx = new AttributesDecContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_attributesDec);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(256);
			match(VAR);
			setState(257);
			match(ID);
			setState(258);
			match(T__2);
			setState(259);
			typeRule();
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
		enterRule(_localctx, 26, RULE_varsTypeInit);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(264);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_TYPE:
			case FLOAT_TYPE:
			case CHAR_TYPE:
			case BOOL_TYPE:
				{
				setState(262);
				typeRule();
				}
				break;
			case ID:
				{
				setState(263);
				match(ID);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(267);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ASSIGN) {
				{
				setState(266);
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
		enterRule(_localctx, 28, RULE_varsTypeInit2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(269);
			match(ASSIGN);
			setState(270);
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
		enterRule(_localctx, 30, RULE_vars);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(272);
			match(ID);
			setState(275);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(273);
				match(T__6);
				setState(274);
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
		enterRule(_localctx, 32, RULE_varArray);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(277);
			match(ID);
			setState(280);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(278);
				match(T__6);
				setState(279);
				match(ID);
				}
			}

			setState(282);
			match(T__3);
			setState(283);
			exp();
			setState(284);
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
		enterRule(_localctx, 34, RULE_varMat);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(286);
			match(ID);
			setState(289);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(287);
				match(T__6);
				setState(288);
				match(ID);
				}
			}

			setState(295);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				{
				setState(291);
				match(T__3);
				setState(292);
				exp();
				setState(293);
				match(T__4);
				}
				break;
			}
			setState(301);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(297);
				match(T__3);
				setState(298);
				exp();
				setState(299);
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
		enterRule(_localctx, 36, RULE_functions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(303);
			match(FUNCTION);
			setState(304);
			match(ID);
			setState(305);
			match(LPAREN);
			setState(307);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(306);
				parameters();
				}
			}

			setState(309);
			match(RPAREN);
			setState(312);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_TYPE:
			case FLOAT_TYPE:
			case CHAR_TYPE:
			case BOOL_TYPE:
				{
				setState(310);
				typeRule();
				}
				break;
			case VOID:
				{
				setState(311);
				match(VOID);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(314);
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
		enterRule(_localctx, 38, RULE_parameters);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(316);
			parameter();
			setState(321);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__7) {
				{
				{
				setState(317);
				match(T__7);
				setState(318);
				parameter();
				}
				}
				setState(323);
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
		enterRule(_localctx, 40, RULE_parameter);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(324);
			match(ID);
			setState(325);
			match(T__2);
			setState(326);
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
		enterRule(_localctx, 42, RULE_functionBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(328);
			match(T__0);
			setState(332);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(329);
				variableDeclaration();
				}
				}
				setState(334);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(338);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << RETURN) | (1L << ID))) != 0)) {
				{
				{
				setState(335);
				statutes();
				}
				}
				setState(340);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(341);
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
		enterRule(_localctx, 44, RULE_returnRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(343);
			match(RETURN);
			setState(345);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << ID))) != 0)) {
				{
				setState(344);
				exp();
				}
			}

			setState(347);
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
		enterRule(_localctx, 46, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(349);
			match(T__0);
			setState(353);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << RETURN) | (1L << ID))) != 0)) {
				{
				{
				setState(350);
				statutes();
				}
				}
				setState(355);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(356);
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

	public static class StatutesNoReturnContext extends ParserRuleContext {
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
		public StatutesNoReturnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statutesNoReturn; }
	}

	public final StatutesNoReturnContext statutesNoReturn() throws RecognitionException {
		StatutesNoReturnContext _localctx = new StatutesNoReturnContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_statutesNoReturn);
		try {
			setState(365);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(358);
				assignation();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(359);
				read();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(360);
				write();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(361);
				conditional();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(362);
				forLoop();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(363);
				whileLoop();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(364);
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
		enterRule(_localctx, 50, RULE_statutes);
		try {
			setState(375);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(367);
				assignation();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(368);
				read();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(369);
				write();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(370);
				conditional();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(371);
				forLoop();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(372);
				whileLoop();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(373);
				expression();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(374);
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
		enterRule(_localctx, 52, RULE_assignation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(380);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				{
				setState(377);
				vars();
				}
				break;
			case 2:
				{
				setState(378);
				varArray();
				}
				break;
			case 3:
				{
				setState(379);
				varMat();
				}
				break;
			}
			setState(382);
			match(ASSIGN);
			setState(383);
			exp();
			setState(384);
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
		enterRule(_localctx, 54, RULE_functionCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(386);
			match(ID);
			setState(387);
			match(LPAREN);
			setState(389);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << ID))) != 0)) {
				{
				setState(388);
				arguments();
				}
			}

			setState(391);
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
		enterRule(_localctx, 56, RULE_arguments);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(393);
			exp();
			setState(394);
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
		enterRule(_localctx, 58, RULE_arguments2);
		try {
			setState(399);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__7:
				enterOuterAlt(_localctx, 1);
				{
				setState(396);
				match(T__7);
				setState(397);
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
		enterRule(_localctx, 60, RULE_methodCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(401);
			match(ID);
			setState(402);
			match(T__6);
			setState(403);
			match(ID);
			setState(404);
			match(LPAREN);
			setState(406);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << ID))) != 0)) {
				{
				setState(405);
				arguments();
				}
			}

			setState(408);
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

	public static class InitCallContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ProyectoParser.ID, 0); }
		public TerminalNode LPAREN() { return getToken(ProyectoParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(ProyectoParser.RPAREN, 0); }
		public InitCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_initCall; }
	}

	public final InitCallContext initCall() throws RecognitionException {
		InitCallContext _localctx = new InitCallContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_initCall);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(410);
			match(T__8);
			setState(411);
			match(ID);
			setState(412);
			match(LPAREN);
			setState(413);
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
		public InitCallContext initCall() {
			return getRuleContext(InitCallContext.class,0);
		}
		public CallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_call; }
	}

	public final CallContext call() throws RecognitionException {
		CallContext _localctx = new CallContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_call);
		try {
			setState(418);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(415);
				functionCall();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(416);
				methodCall();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(417);
				initCall();
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
		enterRule(_localctx, 66, RULE_read);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(420);
			match(READ);
			setState(421);
			match(LPAREN);
			setState(422);
			read2();
			setState(423);
			match(RPAREN);
			setState(424);
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
		enterRule(_localctx, 68, RULE_read2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(429);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				{
				setState(426);
				vars();
				}
				break;
			case 2:
				{
				setState(427);
				varArray();
				}
				break;
			case 3:
				{
				setState(428);
				varMat();
				}
				break;
			}
			setState(431);
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
		enterRule(_localctx, 70, RULE_read3);
		try {
			setState(436);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__7:
				enterOuterAlt(_localctx, 1);
				{
				setState(433);
				match(T__7);
				setState(434);
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
		enterRule(_localctx, 72, RULE_write);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(438);
			match(WRITE);
			setState(439);
			match(LPAREN);
			setState(440);
			arguments();
			setState(441);
			match(RPAREN);
			setState(442);
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
		enterRule(_localctx, 74, RULE_conditional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(444);
			match(IF);
			setState(445);
			conditional2();
			setState(446);
			conditional3();
			setState(448);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(447);
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
		enterRule(_localctx, 76, RULE_conditional2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(450);
			match(LPAREN);
			setState(451);
			exp();
			setState(452);
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
		enterRule(_localctx, 78, RULE_conditional3);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(454);
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
		enterRule(_localctx, 80, RULE_conditional4);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(456);
			match(ELSE);
			setState(457);
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
		enterRule(_localctx, 82, RULE_forLoop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(459);
			match(FOR);
			setState(460);
			forLoop2();
			setState(461);
			match(IN);
			setState(462);
			forLoop3();
			setState(463);
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
		enterRule(_localctx, 84, RULE_forLoop2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(465);
			match(ID);
			setState(466);
			match(ASSIGN);
			setState(467);
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
		enterRule(_localctx, 86, RULE_forLoop3);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(469);
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
		enterRule(_localctx, 88, RULE_whileLoop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(471);
			match(WHILE);
			setState(472);
			whileLoop2();
			setState(473);
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
		enterRule(_localctx, 90, RULE_whileLoop2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(475);
			match(LPAREN);
			setState(476);
			exp();
			setState(477);
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
		enterRule(_localctx, 92, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(479);
			exp();
			setState(480);
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
		enterRule(_localctx, 94, RULE_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(482);
			t_exp();
			setState(486);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OR) {
				{
				{
				setState(483);
				exp2();
				}
				}
				setState(488);
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
		enterRule(_localctx, 96, RULE_exp2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(489);
			match(OR);
			setState(490);
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
		enterRule(_localctx, 98, RULE_t_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(492);
			g_exp();
			setState(496);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==AND) {
				{
				{
				setState(493);
				t_exp2();
				}
				}
				setState(498);
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
		enterRule(_localctx, 100, RULE_t_exp2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(499);
			match(AND);
			setState(500);
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
		enterRule(_localctx, 102, RULE_g_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(502);
			m_exp();
			setState(504);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << GT) | (1L << LT) | (1L << EQ) | (1L << NEQ))) != 0)) {
				{
				setState(503);
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
		enterRule(_localctx, 104, RULE_g_exp2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(506);
			relop();
			setState(507);
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
		enterRule(_localctx, 106, RULE_m_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(509);
			term();
			setState(513);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ADD || _la==SUB) {
				{
				{
				setState(510);
				m_exp2();
				}
				}
				setState(515);
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
		enterRule(_localctx, 108, RULE_m_exp2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(516);
			_la = _input.LA(1);
			if ( !(_la==ADD || _la==SUB) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(517);
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
		enterRule(_localctx, 110, RULE_term);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(519);
			factor();
			setState(523);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==MUL || _la==DIV) {
				{
				{
				setState(520);
				term2();
				}
				}
				setState(525);
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
		enterRule(_localctx, 112, RULE_term2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(526);
			_la = _input.LA(1);
			if ( !(_la==MUL || _la==DIV) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(527);
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
		enterRule(_localctx, 114, RULE_factor);
		try {
			setState(537);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,43,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(529);
				factor2();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(530);
				varCte();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(534);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,42,_ctx) ) {
				case 1:
					{
					setState(531);
					vars();
					}
					break;
				case 2:
					{
					setState(532);
					varArray();
					}
					break;
				case 3:
					{
					setState(533);
					varMat();
					}
					break;
				}
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(536);
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
		enterRule(_localctx, 116, RULE_factor2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(539);
			match(LPAREN);
			setState(540);
			exp();
			setState(541);
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
		enterRule(_localctx, 118, RULE_varCte);
		try {
			setState(548);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(543);
				cte_i();
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(544);
				cte_f();
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 3);
				{
				setState(545);
				cte_c();
				}
				break;
			case T__9:
				enterOuterAlt(_localctx, 4);
				{
				setState(546);
				cte_s();
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 5);
				{
				setState(547);
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
		enterRule(_localctx, 120, RULE_cte_i);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(550);
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
		enterRule(_localctx, 122, RULE_cte_f);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(552);
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
		enterRule(_localctx, 124, RULE_cte_c);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(554);
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
		enterRule(_localctx, 126, RULE_cte_b);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(556);
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
		enterRule(_localctx, 128, RULE_cte_s);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(558);
			match(T__9);
			setState(562);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,45,_ctx);
			while ( _alt!=1 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1+1 ) {
					{
					{
					setState(559);
					matchWildcard();
					}
					} 
				}
				setState(564);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,45,_ctx);
			}
			setState(565);
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
		enterRule(_localctx, 130, RULE_main);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(567);
			match(MAIN);
			setState(568);
			match(LPAREN);
			setState(569);
			match(RPAREN);
			setState(570);
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
		enterRule(_localctx, 132, RULE_mainBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(572);
			match(T__0);
			setState(576);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(573);
				variableDeclaration();
				}
				}
				setState(578);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(582);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << RETURN) | (1L << ID))) != 0)) {
				{
				{
				setState(579);
				statutes();
				}
				}
				setState(584);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(585);
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
		enterRule(_localctx, 134, RULE_typeRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(587);
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
		enterRule(_localctx, 136, RULE_relop);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(589);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\67\u0252\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\3\2\3\2\3\2\3\2"+
		"\7\2\u0091\n\2\f\2\16\2\u0094\13\2\3\2\3\2\3\3\7\3\u0099\n\3\f\3\16\3"+
		"\u009c\13\3\3\3\7\3\u009f\n\3\f\3\16\3\u00a2\13\3\3\3\3\3\3\3\3\4\3\4"+
		"\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\7\5\u00b1\n\5\f\5\16\5\u00b4\13\5\3\5"+
		"\3\5\3\6\3\6\7\6\u00ba\n\6\f\6\16\6\u00bd\13\6\3\7\3\7\5\7\u00c1\n\7\5"+
		"\7\u00c3\n\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\7\b\u00cc\n\b\f\b\16\b\u00cf"+
		"\13\b\3\b\7\b\u00d2\n\b\f\b\16\b\u00d5\13\b\3\b\3\b\3\t\3\t\3\t\5\t\u00dc"+
		"\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\5\r\u00fa\n\r"+
		"\5\r\u00fc\n\r\3\r\3\r\3\r\5\r\u0101\n\r\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\17\3\17\5\17\u010b\n\17\3\17\5\17\u010e\n\17\3\20\3\20\3\20\3\21\3"+
		"\21\3\21\5\21\u0116\n\21\3\22\3\22\3\22\5\22\u011b\n\22\3\22\3\22\3\22"+
		"\3\22\3\23\3\23\3\23\5\23\u0124\n\23\3\23\3\23\3\23\3\23\5\23\u012a\n"+
		"\23\3\23\3\23\3\23\3\23\5\23\u0130\n\23\3\24\3\24\3\24\3\24\5\24\u0136"+
		"\n\24\3\24\3\24\3\24\5\24\u013b\n\24\3\24\3\24\3\25\3\25\3\25\7\25\u0142"+
		"\n\25\f\25\16\25\u0145\13\25\3\26\3\26\3\26\3\26\3\27\3\27\7\27\u014d"+
		"\n\27\f\27\16\27\u0150\13\27\3\27\7\27\u0153\n\27\f\27\16\27\u0156\13"+
		"\27\3\27\3\27\3\30\3\30\5\30\u015c\n\30\3\30\3\30\3\31\3\31\7\31\u0162"+
		"\n\31\f\31\16\31\u0165\13\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3"+
		"\32\5\32\u0170\n\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\5\33\u017a"+
		"\n\33\3\34\3\34\3\34\5\34\u017f\n\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35"+
		"\5\35\u0188\n\35\3\35\3\35\3\36\3\36\3\36\3\37\3\37\3\37\5\37\u0192\n"+
		"\37\3 \3 \3 \3 \3 \5 \u0199\n \3 \3 \3!\3!\3!\3!\3!\3\"\3\"\3\"\5\"\u01a5"+
		"\n\"\3#\3#\3#\3#\3#\3#\3$\3$\3$\5$\u01b0\n$\3$\3$\3%\3%\3%\5%\u01b7\n"+
		"%\3&\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\5\'\u01c3\n\'\3(\3(\3(\3(\3)\3)\3"+
		"*\3*\3*\3+\3+\3+\3+\3+\3+\3,\3,\3,\3,\3-\3-\3.\3.\3.\3.\3/\3/\3/\3/\3"+
		"\60\3\60\3\60\3\61\3\61\7\61\u01e7\n\61\f\61\16\61\u01ea\13\61\3\62\3"+
		"\62\3\62\3\63\3\63\7\63\u01f1\n\63\f\63\16\63\u01f4\13\63\3\64\3\64\3"+
		"\64\3\65\3\65\5\65\u01fb\n\65\3\66\3\66\3\66\3\67\3\67\7\67\u0202\n\67"+
		"\f\67\16\67\u0205\13\67\38\38\38\39\39\79\u020c\n9\f9\169\u020f\139\3"+
		":\3:\3:\3;\3;\3;\3;\3;\5;\u0219\n;\3;\5;\u021c\n;\3<\3<\3<\3<\3=\3=\3"+
		"=\3=\3=\5=\u0227\n=\3>\3>\3?\3?\3@\3@\3A\3A\3B\3B\7B\u0233\nB\fB\16B\u0236"+
		"\13B\3B\3B\3C\3C\3C\3C\3C\3D\3D\7D\u0241\nD\fD\16D\u0244\13D\3D\7D\u0247"+
		"\nD\fD\16D\u024a\13D\3D\3D\3E\3E\3F\3F\3F\3\u0234\2G\2\4\6\b\n\f\16\20"+
		"\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDFHJLNPRTVXZ\\^`bdfhj"+
		"lnprtvxz|~\u0080\u0082\u0084\u0086\u0088\u008a\2\6\3\2\32\33\3\2\30\31"+
		"\3\2\61\64\3\2\24\27\2\u0252\2\u008c\3\2\2\2\4\u009a\3\2\2\2\6\u00a6\3"+
		"\2\2\2\b\u00ab\3\2\2\2\n\u00b7\3\2\2\2\f\u00c2\3\2\2\2\16\u00c6\3\2\2"+
		"\2\20\u00db\3\2\2\2\22\u00dd\3\2\2\2\24\u00e3\3\2\2\2\26\u00ec\3\2\2\2"+
		"\30\u00fb\3\2\2\2\32\u0102\3\2\2\2\34\u010a\3\2\2\2\36\u010f\3\2\2\2 "+
		"\u0112\3\2\2\2\"\u0117\3\2\2\2$\u0120\3\2\2\2&\u0131\3\2\2\2(\u013e\3"+
		"\2\2\2*\u0146\3\2\2\2,\u014a\3\2\2\2.\u0159\3\2\2\2\60\u015f\3\2\2\2\62"+
		"\u016f\3\2\2\2\64\u0179\3\2\2\2\66\u017e\3\2\2\28\u0184\3\2\2\2:\u018b"+
		"\3\2\2\2<\u0191\3\2\2\2>\u0193\3\2\2\2@\u019c\3\2\2\2B\u01a4\3\2\2\2D"+
		"\u01a6\3\2\2\2F\u01af\3\2\2\2H\u01b6\3\2\2\2J\u01b8\3\2\2\2L\u01be\3\2"+
		"\2\2N\u01c4\3\2\2\2P\u01c8\3\2\2\2R\u01ca\3\2\2\2T\u01cd\3\2\2\2V\u01d3"+
		"\3\2\2\2X\u01d7\3\2\2\2Z\u01d9\3\2\2\2\\\u01dd\3\2\2\2^\u01e1\3\2\2\2"+
		"`\u01e4\3\2\2\2b\u01eb\3\2\2\2d\u01ee\3\2\2\2f\u01f5\3\2\2\2h\u01f8\3"+
		"\2\2\2j\u01fc\3\2\2\2l\u01ff\3\2\2\2n\u0206\3\2\2\2p\u0209\3\2\2\2r\u0210"+
		"\3\2\2\2t\u021b\3\2\2\2v\u021d\3\2\2\2x\u0226\3\2\2\2z\u0228\3\2\2\2|"+
		"\u022a\3\2\2\2~\u022c\3\2\2\2\u0080\u022e\3\2\2\2\u0082\u0230\3\2\2\2"+
		"\u0084\u0239\3\2\2\2\u0086\u023e\3\2\2\2\u0088\u024d\3\2\2\2\u008a\u024f"+
		"\3\2\2\2\u008c\u008d\7\66\2\2\u008d\u008e\7\67\2\2\u008e\u0092\7\34\2"+
		"\2\u008f\u0091\5\20\t\2\u0090\u008f\3\2\2\2\u0091\u0094\3\2\2\2\u0092"+
		"\u0090\3\2\2\2\u0092\u0093\3\2\2\2\u0093\u0095\3\2\2\2\u0094\u0092\3\2"+
		"\2\2\u0095\u0096\5\4\3\2\u0096\3\3\2\2\2\u0097\u0099\5\6\4\2\u0098\u0097"+
		"\3\2\2\2\u0099\u009c\3\2\2\2\u009a\u0098\3\2\2\2\u009a\u009b\3\2\2\2\u009b"+
		"\u00a0\3\2\2\2\u009c\u009a\3\2\2\2\u009d\u009f\5&\24\2\u009e\u009d\3\2"+
		"\2\2\u009f\u00a2\3\2\2\2\u00a0\u009e\3\2\2\2\u00a0\u00a1\3\2\2\2\u00a1"+
		"\u00a3\3\2\2\2\u00a2\u00a0\3\2\2\2\u00a3\u00a4\5\u0084C\2\u00a4\u00a5"+
		"\7\2\2\3\u00a5\5\3\2\2\2\u00a6\u00a7\7%\2\2\u00a7\u00a8\7\67\2\2\u00a8"+
		"\u00a9\5\b\5\2\u00a9\u00aa\7\34\2\2\u00aa\7\3\2\2\2\u00ab\u00ac\7\3\2"+
		"\2\u00ac\u00ad\5\n\6\2\u00ad\u00ae\5\16\b\2\u00ae\u00b2\7\'\2\2\u00af"+
		"\u00b1\5\f\7\2\u00b0\u00af\3\2\2\2\u00b1\u00b4\3\2\2\2\u00b2\u00b0\3\2"+
		"\2\2\u00b2\u00b3\3\2\2\2\u00b3\u00b5\3\2\2\2\u00b4\u00b2\3\2\2\2\u00b5"+
		"\u00b6\7\4\2\2\u00b6\t\3\2\2\2\u00b7\u00bb\7&\2\2\u00b8\u00ba\5\30\r\2"+
		"\u00b9\u00b8\3\2\2\2\u00ba\u00bd\3\2\2\2\u00bb\u00b9\3\2\2\2\u00bb\u00bc"+
		"\3\2\2\2\u00bc\13\3\2\2\2\u00bd\u00bb\3\2\2\2\u00be\u00c3\7)\2\2\u00bf"+
		"\u00c1\7*\2\2\u00c0\u00bf\3\2\2\2\u00c0\u00c1\3\2\2\2\u00c1\u00c3\3\2"+
		"\2\2\u00c2\u00be\3\2\2\2\u00c2\u00c0\3\2\2\2\u00c3\u00c4\3\2\2\2\u00c4"+
		"\u00c5\5&\24\2\u00c5\r\3\2\2\2\u00c6\u00c7\7(\2\2\u00c7\u00c8\7\36\2\2"+
		"\u00c8\u00c9\7\37\2\2\u00c9\u00cd\7\3\2\2\u00ca\u00cc\5\20\t\2\u00cb\u00ca"+
		"\3\2\2\2\u00cc\u00cf\3\2\2\2\u00cd\u00cb\3\2\2\2\u00cd\u00ce\3\2\2\2\u00ce"+
		"\u00d3\3\2\2\2\u00cf\u00cd\3\2\2\2\u00d0\u00d2\5\62\32\2\u00d1\u00d0\3"+
		"\2\2\2\u00d2\u00d5\3\2\2\2\u00d3\u00d1\3\2\2\2\u00d3\u00d4\3\2\2\2\u00d4"+
		"\u00d6\3\2\2\2\u00d5\u00d3\3\2\2\2\u00d6\u00d7\7\4\2\2\u00d7\17\3\2\2"+
		"\2\u00d8\u00dc\5\22\n\2\u00d9\u00dc\5\24\13\2\u00da\u00dc\5\26\f\2\u00db"+
		"\u00d8\3\2\2\2\u00db\u00d9\3\2\2\2\u00db\u00da\3\2\2\2\u00dc\21\3\2\2"+
		"\2\u00dd\u00de\7\60\2\2\u00de\u00df\7\67\2\2\u00df\u00e0\7\5\2\2\u00e0"+
		"\u00e1\5\34\17\2\u00e1\u00e2\7\34\2\2\u00e2\23\3\2\2\2\u00e3\u00e4\7\60"+
		"\2\2\u00e4\u00e5\7\67\2\2\u00e5\u00e6\7\5\2\2\u00e6\u00e7\7\6\2\2\u00e7"+
		"\u00e8\7\r\2\2\u00e8\u00e9\7\7\2\2\u00e9\u00ea\5\u0088E\2\u00ea\u00eb"+
		"\7\34\2\2\u00eb\25\3\2\2\2\u00ec\u00ed\7\60\2\2\u00ed\u00ee\7\67\2\2\u00ee"+
		"\u00ef\7\5\2\2\u00ef\u00f0\7\6\2\2\u00f0\u00f1\7\r\2\2\u00f1\u00f2\7\b"+
		"\2\2\u00f2\u00f3\7\r\2\2\u00f3\u00f4\7\7\2\2\u00f4\u00f5\5\u0088E\2\u00f5"+
		"\u00f6\7\34\2\2\u00f6\27\3\2\2\2\u00f7\u00fc\7)\2\2\u00f8\u00fa\7*\2\2"+
		"\u00f9\u00f8\3\2\2\2\u00f9\u00fa\3\2\2\2\u00fa\u00fc\3\2\2\2\u00fb\u00f7"+
		"\3\2\2\2\u00fb\u00f9\3\2\2\2\u00fc\u0100\3\2\2\2\u00fd\u0101\5\32\16\2"+
		"\u00fe\u0101\5\24\13\2\u00ff\u0101\5\26\f\2\u0100\u00fd\3\2\2\2\u0100"+
		"\u00fe\3\2\2\2\u0100\u00ff\3\2\2\2\u0101\31\3\2\2\2\u0102\u0103\7\60\2"+
		"\2\u0103\u0104\7\67\2\2\u0104\u0105\7\5\2\2\u0105\u0106\5\u0088E\2\u0106"+
		"\u0107\7\34\2\2\u0107\33\3\2\2\2\u0108\u010b\5\u0088E\2\u0109\u010b\7"+
		"\67\2\2\u010a\u0108\3\2\2\2\u010a\u0109\3\2\2\2\u010b\u010d\3\2\2\2\u010c"+
		"\u010e\5\36\20\2\u010d\u010c\3\2\2\2\u010d\u010e\3\2\2\2\u010e\35\3\2"+
		"\2\2\u010f\u0110\7\35\2\2\u0110\u0111\5`\61\2\u0111\37\3\2\2\2\u0112\u0115"+
		"\7\67\2\2\u0113\u0114\7\t\2\2\u0114\u0116\7\67\2\2\u0115\u0113\3\2\2\2"+
		"\u0115\u0116\3\2\2\2\u0116!\3\2\2\2\u0117\u011a\7\67\2\2\u0118\u0119\7"+
		"\t\2\2\u0119\u011b\7\67\2\2\u011a\u0118\3\2\2\2\u011a\u011b\3\2\2\2\u011b"+
		"\u011c\3\2\2\2\u011c\u011d\7\6\2\2\u011d\u011e\5`\61\2\u011e\u011f\7\7"+
		"\2\2\u011f#\3\2\2\2\u0120\u0123\7\67\2\2\u0121\u0122\7\t\2\2\u0122\u0124"+
		"\7\67\2\2\u0123\u0121\3\2\2\2\u0123\u0124\3\2\2\2\u0124\u0129\3\2\2\2"+
		"\u0125\u0126\7\6\2\2\u0126\u0127\5`\61\2\u0127\u0128\7\7\2\2\u0128\u012a"+
		"\3\2\2\2\u0129\u0125\3\2\2\2\u0129\u012a\3\2\2\2\u012a\u012f\3\2\2\2\u012b"+
		"\u012c\7\6\2\2\u012c\u012d\5`\61\2\u012d\u012e\7\7\2\2\u012e\u0130\3\2"+
		"\2\2\u012f\u012b\3\2\2\2\u012f\u0130\3\2\2\2\u0130%\3\2\2\2\u0131\u0132"+
		"\7-\2\2\u0132\u0133\7\67\2\2\u0133\u0135\7\36\2\2\u0134\u0136\5(\25\2"+
		"\u0135\u0134\3\2\2\2\u0135\u0136\3\2\2\2\u0136\u0137\3\2\2\2\u0137\u013a"+
		"\7\37\2\2\u0138\u013b\5\u0088E\2\u0139\u013b\7\65\2\2\u013a\u0138\3\2"+
		"\2\2\u013a\u0139\3\2\2\2\u013b\u013c\3\2\2\2\u013c\u013d\5,\27\2\u013d"+
		"\'\3\2\2\2\u013e\u0143\5*\26\2\u013f\u0140\7\n\2\2\u0140\u0142\5*\26\2"+
		"\u0141\u013f\3\2\2\2\u0142\u0145\3\2\2\2\u0143\u0141\3\2\2\2\u0143\u0144"+
		"\3\2\2\2\u0144)\3\2\2\2\u0145\u0143\3\2\2\2\u0146\u0147\7\67\2\2\u0147"+
		"\u0148\7\5\2\2\u0148\u0149\5\u0088E\2\u0149+\3\2\2\2\u014a\u014e\7\3\2"+
		"\2\u014b\u014d\5\20\t\2\u014c\u014b\3\2\2\2\u014d\u0150\3\2\2\2\u014e"+
		"\u014c\3\2\2\2\u014e\u014f\3\2\2\2\u014f\u0154\3\2\2\2\u0150\u014e\3\2"+
		"\2\2\u0151\u0153\5\64\33\2\u0152\u0151\3\2\2\2\u0153\u0156\3\2\2\2\u0154"+
		"\u0152\3\2\2\2\u0154\u0155\3\2\2\2\u0155\u0157\3\2\2\2\u0156\u0154\3\2"+
		"\2\2\u0157\u0158\7\4\2\2\u0158-\3\2\2\2\u0159\u015b\7.\2\2\u015a\u015c"+
		"\5`\61\2\u015b\u015a\3\2\2\2\u015b\u015c\3\2\2\2\u015c\u015d\3\2\2\2\u015d"+
		"\u015e\7\34\2\2\u015e/\3\2\2\2\u015f\u0163\7\3\2\2\u0160\u0162\5\64\33"+
		"\2\u0161\u0160\3\2\2\2\u0162\u0165\3\2\2\2\u0163\u0161\3\2\2\2\u0163\u0164"+
		"\3\2\2\2\u0164\u0166\3\2\2\2\u0165\u0163\3\2\2\2\u0166\u0167\7\4\2\2\u0167"+
		"\61\3\2\2\2\u0168\u0170\5\66\34\2\u0169\u0170\5D#\2\u016a\u0170\5J&\2"+
		"\u016b\u0170\5L\'\2\u016c\u0170\5T+\2\u016d\u0170\5Z.\2\u016e\u0170\5"+
		"^\60\2\u016f\u0168\3\2\2\2\u016f\u0169\3\2\2\2\u016f\u016a\3\2\2\2\u016f"+
		"\u016b\3\2\2\2\u016f\u016c\3\2\2\2\u016f\u016d\3\2\2\2\u016f\u016e\3\2"+
		"\2\2\u0170\63\3\2\2\2\u0171\u017a\5\66\34\2\u0172\u017a\5D#\2\u0173\u017a"+
		"\5J&\2\u0174\u017a\5L\'\2\u0175\u017a\5T+\2\u0176\u017a\5Z.\2\u0177\u017a"+
		"\5^\60\2\u0178\u017a\5.\30\2\u0179\u0171\3\2\2\2\u0179\u0172\3\2\2\2\u0179"+
		"\u0173\3\2\2\2\u0179\u0174\3\2\2\2\u0179\u0175\3\2\2\2\u0179\u0176\3\2"+
		"\2\2\u0179\u0177\3\2\2\2\u0179\u0178\3\2\2\2\u017a\65\3\2\2\2\u017b\u017f"+
		"\5 \21\2\u017c\u017f\5\"\22\2\u017d\u017f\5$\23\2\u017e\u017b\3\2\2\2"+
		"\u017e\u017c\3\2\2\2\u017e\u017d\3\2\2\2\u017f\u0180\3\2\2\2\u0180\u0181"+
		"\7\35\2\2\u0181\u0182\5`\61\2\u0182\u0183\7\34\2\2\u0183\67\3\2\2\2\u0184"+
		"\u0185\7\67\2\2\u0185\u0187\7\36\2\2\u0186\u0188\5:\36\2\u0187\u0186\3"+
		"\2\2\2\u0187\u0188\3\2\2\2\u0188\u0189\3\2\2\2\u0189\u018a\7\37\2\2\u018a"+
		"9\3\2\2\2\u018b\u018c\5`\61\2\u018c\u018d\5<\37\2\u018d;\3\2\2\2\u018e"+
		"\u018f\7\n\2\2\u018f\u0192\5:\36\2\u0190\u0192\3\2\2\2\u0191\u018e\3\2"+
		"\2\2\u0191\u0190\3\2\2\2\u0192=\3\2\2\2\u0193\u0194\7\67\2\2\u0194\u0195"+
		"\7\t\2\2\u0195\u0196\7\67\2\2\u0196\u0198\7\36\2\2\u0197\u0199\5:\36\2"+
		"\u0198\u0197\3\2\2\2\u0198\u0199\3\2\2\2\u0199\u019a\3\2\2\2\u019a\u019b"+
		"\7\37\2\2\u019b?\3\2\2\2\u019c\u019d\7\13\2\2\u019d\u019e\7\67\2\2\u019e"+
		"\u019f\7\36\2\2\u019f\u01a0\7\37\2\2\u01a0A\3\2\2\2\u01a1\u01a5\58\35"+
		"\2\u01a2\u01a5\5> \2\u01a3\u01a5\5@!\2\u01a4\u01a1\3\2\2\2\u01a4\u01a2"+
		"\3\2\2\2\u01a4\u01a3\3\2\2\2\u01a5C\3\2\2\2\u01a6\u01a7\7,\2\2\u01a7\u01a8"+
		"\7\36\2\2\u01a8\u01a9\5F$\2\u01a9\u01aa\7\37\2\2\u01aa\u01ab\7\34\2\2"+
		"\u01abE\3\2\2\2\u01ac\u01b0\5 \21\2\u01ad\u01b0\5\"\22\2\u01ae\u01b0\5"+
		"$\23\2\u01af\u01ac\3\2\2\2\u01af\u01ad\3\2\2\2\u01af\u01ae\3\2\2\2\u01b0"+
		"\u01b1\3\2\2\2\u01b1\u01b2\5H%\2\u01b2G\3\2\2\2\u01b3\u01b4\7\n\2\2\u01b4"+
		"\u01b7\5F$\2\u01b5\u01b7\3\2\2\2\u01b6\u01b3\3\2\2\2\u01b6\u01b5\3\2\2"+
		"\2\u01b7I\3\2\2\2\u01b8\u01b9\7+\2\2\u01b9\u01ba\7\36\2\2\u01ba\u01bb"+
		"\5:\36\2\u01bb\u01bc\7\37\2\2\u01bc\u01bd\7\34\2\2\u01bdK\3\2\2\2\u01be"+
		"\u01bf\7 \2\2\u01bf\u01c0\5N(\2\u01c0\u01c2\5P)\2\u01c1\u01c3\5R*\2\u01c2"+
		"\u01c1\3\2\2\2\u01c2\u01c3\3\2\2\2\u01c3M\3\2\2\2\u01c4\u01c5\7\36\2\2"+
		"\u01c5\u01c6\5`\61\2\u01c6\u01c7\7\37\2\2\u01c7O\3\2\2\2\u01c8\u01c9\5"+
		"\60\31\2\u01c9Q\3\2\2\2\u01ca\u01cb\7!\2\2\u01cb\u01cc\5\60\31\2\u01cc"+
		"S\3\2\2\2\u01cd\u01ce\7#\2\2\u01ce\u01cf\5V,\2\u01cf\u01d0\7$\2\2\u01d0"+
		"\u01d1\5X-\2\u01d1\u01d2\5\60\31\2\u01d2U\3\2\2\2\u01d3\u01d4\7\67\2\2"+
		"\u01d4\u01d5\7\35\2\2\u01d5\u01d6\5`\61\2\u01d6W\3\2\2\2\u01d7\u01d8\5"+
		"`\61\2\u01d8Y\3\2\2\2\u01d9\u01da\7\"\2\2\u01da\u01db\5\\/\2\u01db\u01dc"+
		"\5\60\31\2\u01dc[\3\2\2\2\u01dd\u01de\7\36\2\2\u01de\u01df\5`\61\2\u01df"+
		"\u01e0\7\37\2\2\u01e0]\3\2\2\2\u01e1\u01e2\5`\61\2\u01e2\u01e3\7\34\2"+
		"\2\u01e3_\3\2\2\2\u01e4\u01e8\5d\63\2\u01e5\u01e7\5b\62\2\u01e6\u01e5"+
		"\3\2\2\2\u01e7\u01ea\3\2\2\2\u01e8\u01e6\3\2\2\2\u01e8\u01e9\3\2\2\2\u01e9"+
		"a\3\2\2\2\u01ea\u01e8\3\2\2\2\u01eb\u01ec\7\22\2\2\u01ec\u01ed\5d\63\2"+
		"\u01edc\3\2\2\2\u01ee\u01f2\5h\65\2\u01ef\u01f1\5f\64\2\u01f0\u01ef\3"+
		"\2\2\2\u01f1\u01f4\3\2\2\2\u01f2\u01f0\3\2\2\2\u01f2\u01f3\3\2\2\2\u01f3"+
		"e\3\2\2\2\u01f4\u01f2\3\2\2\2\u01f5\u01f6\7\23\2\2\u01f6\u01f7\5h\65\2"+
		"\u01f7g\3\2\2\2\u01f8\u01fa\5l\67\2\u01f9\u01fb\5j\66\2\u01fa\u01f9\3"+
		"\2\2\2\u01fa\u01fb\3\2\2\2\u01fbi\3\2\2\2\u01fc\u01fd\5\u008aF\2\u01fd"+
		"\u01fe\5l\67\2\u01fek\3\2\2\2\u01ff\u0203\5p9\2\u0200\u0202\5n8\2\u0201"+
		"\u0200\3\2\2\2\u0202\u0205\3\2\2\2\u0203\u0201\3\2\2\2\u0203\u0204\3\2"+
		"\2\2\u0204m\3\2\2\2\u0205\u0203\3\2\2\2\u0206\u0207\t\2\2\2\u0207\u0208"+
		"\5p9\2\u0208o\3\2\2\2\u0209\u020d\5t;\2\u020a\u020c\5r:\2\u020b\u020a"+
		"\3\2\2\2\u020c\u020f\3\2\2\2\u020d\u020b\3\2\2\2\u020d\u020e\3\2\2\2\u020e"+
		"q\3\2\2\2\u020f\u020d\3\2\2\2\u0210\u0211\t\3\2\2\u0211\u0212\5t;\2\u0212"+
		"s\3\2\2\2\u0213\u021c\5v<\2\u0214\u021c\5x=\2\u0215\u0219\5 \21\2\u0216"+
		"\u0219\5\"\22\2\u0217\u0219\5$\23\2\u0218\u0215\3\2\2\2\u0218\u0216\3"+
		"\2\2\2\u0218\u0217\3\2\2\2\u0219\u021c\3\2\2\2\u021a\u021c\5B\"\2\u021b"+
		"\u0213\3\2\2\2\u021b\u0214\3\2\2\2\u021b\u0218\3\2\2\2\u021b\u021a\3\2"+
		"\2\2\u021cu\3\2\2\2\u021d\u021e\7\36\2\2\u021e\u021f\5`\61\2\u021f\u0220"+
		"\7\37\2\2\u0220w\3\2\2\2\u0221\u0227\5z>\2\u0222\u0227\5|?\2\u0223\u0227"+
		"\5~@\2\u0224\u0227\5\u0082B\2\u0225\u0227\5\u0080A\2\u0226\u0221\3\2\2"+
		"\2\u0226\u0222\3\2\2\2\u0226\u0223\3\2\2\2\u0226\u0224\3\2\2\2\u0226\u0225"+
		"\3\2\2\2\u0227y\3\2\2\2\u0228\u0229\7\r\2\2\u0229{\3\2\2\2\u022a\u022b"+
		"\7\16\2\2\u022b}\3\2\2\2\u022c\u022d\7\17\2\2\u022d\177\3\2\2\2\u022e"+
		"\u022f\7\20\2\2\u022f\u0081\3\2\2\2\u0230\u0234\7\f\2\2\u0231\u0233\13"+
		"\2\2\2\u0232\u0231\3\2\2\2\u0233\u0236\3\2\2\2\u0234\u0235\3\2\2\2\u0234"+
		"\u0232\3\2\2\2\u0235\u0237\3\2\2\2\u0236\u0234\3\2\2\2\u0237\u0238\7\f"+
		"\2\2\u0238\u0083\3\2\2\2\u0239\u023a\7/\2\2\u023a\u023b\7\36\2\2\u023b"+
		"\u023c\7\37\2\2\u023c\u023d\5\u0086D\2\u023d\u0085\3\2\2\2\u023e\u0242"+
		"\7\3\2\2\u023f\u0241\5\20\t\2\u0240\u023f\3\2\2\2\u0241\u0244\3\2\2\2"+
		"\u0242\u0240\3\2\2\2\u0242\u0243\3\2\2\2\u0243\u0248\3\2\2\2\u0244\u0242"+
		"\3\2\2\2\u0245\u0247\5\64\33\2\u0246\u0245\3\2\2\2\u0247\u024a\3\2\2\2"+
		"\u0248\u0246\3\2\2\2\u0248\u0249\3\2\2\2\u0249\u024b\3\2\2\2\u024a\u0248"+
		"\3\2\2\2\u024b\u024c\7\4\2\2\u024c\u0087\3\2\2\2\u024d\u024e\t\4\2\2\u024e"+
		"\u0089\3\2\2\2\u024f\u0250\t\5\2\2\u0250\u008b\3\2\2\2\62\u0092\u009a"+
		"\u00a0\u00b2\u00bb\u00c0\u00c2\u00cd\u00d3\u00db\u00f9\u00fb\u0100\u010a"+
		"\u010d\u0115\u011a\u0123\u0129\u012f\u0135\u013a\u0143\u014e\u0154\u015b"+
		"\u0163\u016f\u0179\u017e\u0187\u0191\u0198\u01a4\u01af\u01b6\u01c2\u01e8"+
		"\u01f2\u01fa\u0203\u020d\u0218\u021b\u0226\u0234\u0242\u0248";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}