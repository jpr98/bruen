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
		RULE_write = 36, RULE_w_arguments = 37, RULE_w_arguments2 = 38, RULE_conditional = 39, 
		RULE_conditional2 = 40, RULE_conditional3 = 41, RULE_conditional4 = 42, 
		RULE_forLoop = 43, RULE_forLoop2 = 44, RULE_forLoop3 = 45, RULE_whileLoop = 46, 
		RULE_whileLoop2 = 47, RULE_expression = 48, RULE_exp = 49, RULE_exp2 = 50, 
		RULE_t_exp = 51, RULE_t_exp2 = 52, RULE_g_exp = 53, RULE_g_exp2 = 54, 
		RULE_m_exp = 55, RULE_m_exp2 = 56, RULE_term = 57, RULE_term2 = 58, RULE_factor = 59, 
		RULE_factor2 = 60, RULE_varCte = 61, RULE_cte_i = 62, RULE_cte_f = 63, 
		RULE_cte_c = 64, RULE_cte_b = 65, RULE_cte_s = 66, RULE_main = 67, RULE_mainBlock = 68, 
		RULE_typeRule = 69, RULE_relop = 70;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "program2", "classDef", "classBlock", "classAttributes", "classMethod", 
			"classInit", "variableDeclaration", "varsDec", "varsDecArray", "varsDecMat", 
			"attributesDeclaration", "attributesDec", "varsTypeInit", "varsTypeInit2", 
			"vars", "varArray", "varMat", "functions", "parameters", "parameter", 
			"functionBlock", "returnRule", "block", "statutesNoReturn", "statutes", 
			"assignation", "functionCall", "arguments", "arguments2", "methodCall", 
			"initCall", "call", "read", "read2", "read3", "write", "w_arguments", 
			"w_arguments2", "conditional", "conditional2", "conditional3", "conditional4", 
			"forLoop", "forLoop2", "forLoop3", "whileLoop", "whileLoop2", "expression", 
			"exp", "exp2", "t_exp", "t_exp2", "g_exp", "g_exp2", "m_exp", "m_exp2", 
			"term", "term2", "factor", "factor2", "varCte", "cte_i", "cte_f", "cte_c", 
			"cte_b", "cte_s", "main", "mainBlock", "typeRule", "relop"
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
			setState(142);
			match(PROGRAM);
			setState(143);
			match(ID);
			setState(144);
			match(SEMI);
			setState(148);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(145);
				variableDeclaration();
				}
				}
				setState(150);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(151);
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
			setState(156);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CLASS) {
				{
				{
				setState(153);
				classDef();
				}
				}
				setState(158);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(162);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FUNCTION) {
				{
				{
				setState(159);
				functions();
				}
				}
				setState(164);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(165);
			main();
			setState(166);
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
			setState(168);
			match(CLASS);
			setState(169);
			match(ID);
			setState(170);
			classBlock();
			setState(171);
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
			setState(173);
			match(T__0);
			setState(174);
			classAttributes();
			setState(175);
			classInit();
			setState(176);
			match(METHODS);
			setState(180);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PRIVATE) | (1L << PUBLIC) | (1L << FUNCTION))) != 0)) {
				{
				{
				setState(177);
				classMethod();
				}
				}
				setState(182);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(183);
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
			setState(185);
			match(ATTRIBUTES);
			setState(189);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PRIVATE) | (1L << PUBLIC) | (1L << VAR))) != 0)) {
				{
				{
				setState(186);
				attributesDeclaration();
				}
				}
				setState(191);
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
			setState(196);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRIVATE:
				{
				setState(192);
				match(PRIVATE);
				}
				break;
			case PUBLIC:
			case FUNCTION:
				{
				setState(194);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUBLIC) {
					{
					setState(193);
					match(PUBLIC);
					}
				}

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(198);
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
			setState(200);
			match(INIT);
			setState(201);
			match(LPAREN);
			setState(202);
			match(RPAREN);
			setState(203);
			match(T__0);
			setState(207);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(204);
				variableDeclaration();
				}
				}
				setState(209);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(213);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << ID))) != 0)) {
				{
				{
				setState(210);
				statutesNoReturn();
				}
				}
				setState(215);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(216);
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
			setState(221);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(218);
				varsDec();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(219);
				varsDecArray();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(220);
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
			setState(223);
			match(VAR);
			setState(224);
			match(ID);
			setState(225);
			match(T__2);
			setState(226);
			varsTypeInit();
			setState(227);
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
			setState(229);
			match(VAR);
			setState(230);
			match(ID);
			setState(231);
			match(T__2);
			setState(232);
			match(T__3);
			setState(233);
			match(INT);
			setState(234);
			match(T__4);
			setState(235);
			typeRule();
			setState(236);
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
			setState(238);
			match(VAR);
			setState(239);
			match(ID);
			setState(240);
			match(T__2);
			setState(241);
			match(T__3);
			setState(242);
			match(INT);
			setState(243);
			match(T__5);
			setState(244);
			match(INT);
			setState(245);
			match(T__4);
			setState(246);
			typeRule();
			setState(247);
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
			setState(253);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRIVATE:
				{
				setState(249);
				match(PRIVATE);
				}
				break;
			case PUBLIC:
			case VAR:
				{
				setState(251);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUBLIC) {
					{
					setState(250);
					match(PUBLIC);
					}
				}

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(258);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				{
				setState(255);
				attributesDec();
				}
				break;
			case 2:
				{
				setState(256);
				varsDecArray();
				}
				break;
			case 3:
				{
				setState(257);
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
			setState(260);
			match(VAR);
			setState(261);
			match(ID);
			setState(262);
			match(T__2);
			setState(263);
			typeRule();
			setState(264);
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
			setState(268);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_TYPE:
			case FLOAT_TYPE:
			case CHAR_TYPE:
			case BOOL_TYPE:
				{
				setState(266);
				typeRule();
				}
				break;
			case ID:
				{
				setState(267);
				match(ID);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(271);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ASSIGN) {
				{
				setState(270);
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
			setState(273);
			match(ASSIGN);
			setState(274);
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
			setState(276);
			match(ID);
			setState(279);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(277);
				match(T__6);
				setState(278);
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
			setState(281);
			match(ID);
			setState(284);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(282);
				match(T__6);
				setState(283);
				match(ID);
				}
			}

			setState(286);
			match(T__3);
			setState(287);
			exp();
			setState(288);
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
			setState(290);
			match(ID);
			setState(293);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(291);
				match(T__6);
				setState(292);
				match(ID);
				}
			}

			setState(299);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				{
				setState(295);
				match(T__3);
				setState(296);
				exp();
				setState(297);
				match(T__4);
				}
				break;
			}
			setState(305);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(301);
				match(T__3);
				setState(302);
				exp();
				setState(303);
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
			setState(307);
			match(FUNCTION);
			setState(308);
			match(ID);
			setState(309);
			match(LPAREN);
			setState(311);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(310);
				parameters();
				}
			}

			setState(313);
			match(RPAREN);
			setState(316);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_TYPE:
			case FLOAT_TYPE:
			case CHAR_TYPE:
			case BOOL_TYPE:
				{
				setState(314);
				typeRule();
				}
				break;
			case VOID:
				{
				setState(315);
				match(VOID);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(318);
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
			setState(320);
			parameter();
			setState(325);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__7) {
				{
				{
				setState(321);
				match(T__7);
				setState(322);
				parameter();
				}
				}
				setState(327);
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
			setState(328);
			match(ID);
			setState(329);
			match(T__2);
			setState(330);
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
			setState(332);
			match(T__0);
			setState(336);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(333);
				variableDeclaration();
				}
				}
				setState(338);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(342);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << RETURN) | (1L << ID))) != 0)) {
				{
				{
				setState(339);
				statutes();
				}
				}
				setState(344);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(345);
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
			setState(347);
			match(RETURN);
			setState(349);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << ID))) != 0)) {
				{
				setState(348);
				exp();
				}
			}

			setState(351);
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
			setState(353);
			match(T__0);
			setState(357);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << RETURN) | (1L << ID))) != 0)) {
				{
				{
				setState(354);
				statutes();
				}
				}
				setState(359);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(360);
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
			setState(369);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(362);
				assignation();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(363);
				read();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(364);
				write();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(365);
				conditional();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(366);
				forLoop();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(367);
				whileLoop();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(368);
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
			setState(379);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(371);
				assignation();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(372);
				read();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(373);
				write();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(374);
				conditional();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(375);
				forLoop();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(376);
				whileLoop();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(377);
				expression();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(378);
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
			setState(384);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				{
				setState(381);
				vars();
				}
				break;
			case 2:
				{
				setState(382);
				varArray();
				}
				break;
			case 3:
				{
				setState(383);
				varMat();
				}
				break;
			}
			setState(386);
			match(ASSIGN);
			setState(387);
			exp();
			setState(388);
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
			setState(390);
			match(ID);
			setState(391);
			match(LPAREN);
			setState(393);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << ID))) != 0)) {
				{
				setState(392);
				arguments();
				}
			}

			setState(395);
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
			setState(397);
			exp();
			setState(398);
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
			setState(403);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__7:
				enterOuterAlt(_localctx, 1);
				{
				setState(400);
				match(T__7);
				setState(401);
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
			setState(405);
			match(ID);
			setState(406);
			match(T__6);
			setState(407);
			match(ID);
			setState(408);
			match(LPAREN);
			setState(410);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << ID))) != 0)) {
				{
				setState(409);
				arguments();
				}
			}

			setState(412);
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
			setState(414);
			match(T__8);
			setState(415);
			match(ID);
			setState(416);
			match(LPAREN);
			setState(417);
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
			setState(422);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(419);
				functionCall();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(420);
				methodCall();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(421);
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
			setState(424);
			match(READ);
			setState(425);
			match(LPAREN);
			setState(426);
			read2();
			setState(427);
			match(RPAREN);
			setState(428);
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
			setState(433);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				{
				setState(430);
				vars();
				}
				break;
			case 2:
				{
				setState(431);
				varArray();
				}
				break;
			case 3:
				{
				setState(432);
				varMat();
				}
				break;
			}
			setState(435);
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
			setState(440);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__7:
				enterOuterAlt(_localctx, 1);
				{
				setState(437);
				match(T__7);
				setState(438);
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
		public W_argumentsContext w_arguments() {
			return getRuleContext(W_argumentsContext.class,0);
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
			setState(442);
			match(WRITE);
			setState(443);
			match(LPAREN);
			setState(444);
			w_arguments();
			setState(445);
			match(RPAREN);
			setState(446);
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

	public static class W_argumentsContext extends ParserRuleContext {
		public ExpContext exp() {
			return getRuleContext(ExpContext.class,0);
		}
		public W_arguments2Context w_arguments2() {
			return getRuleContext(W_arguments2Context.class,0);
		}
		public W_argumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_w_arguments; }
	}

	public final W_argumentsContext w_arguments() throws RecognitionException {
		W_argumentsContext _localctx = new W_argumentsContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_w_arguments);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(448);
			exp();
			setState(449);
			w_arguments2();
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

	public static class W_arguments2Context extends ParserRuleContext {
		public W_argumentsContext w_arguments() {
			return getRuleContext(W_argumentsContext.class,0);
		}
		public W_arguments2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_w_arguments2; }
	}

	public final W_arguments2Context w_arguments2() throws RecognitionException {
		W_arguments2Context _localctx = new W_arguments2Context(_ctx, getState());
		enterRule(_localctx, 76, RULE_w_arguments2);
		try {
			setState(454);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__7:
				enterOuterAlt(_localctx, 1);
				{
				setState(451);
				match(T__7);
				setState(452);
				w_arguments();
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
		enterRule(_localctx, 78, RULE_conditional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(456);
			match(IF);
			setState(457);
			conditional2();
			setState(458);
			conditional3();
			setState(460);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(459);
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
		enterRule(_localctx, 80, RULE_conditional2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(462);
			match(LPAREN);
			setState(463);
			exp();
			setState(464);
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
		enterRule(_localctx, 82, RULE_conditional3);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(466);
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
		enterRule(_localctx, 84, RULE_conditional4);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(468);
			match(ELSE);
			setState(469);
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
		enterRule(_localctx, 86, RULE_forLoop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(471);
			match(FOR);
			setState(472);
			forLoop2();
			setState(473);
			match(IN);
			setState(474);
			forLoop3();
			setState(475);
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
		enterRule(_localctx, 88, RULE_forLoop2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(477);
			match(ID);
			setState(478);
			match(ASSIGN);
			setState(479);
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
		enterRule(_localctx, 90, RULE_forLoop3);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(481);
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
		enterRule(_localctx, 92, RULE_whileLoop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(483);
			match(WHILE);
			setState(484);
			whileLoop2();
			setState(485);
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
		enterRule(_localctx, 94, RULE_whileLoop2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(487);
			match(LPAREN);
			setState(488);
			exp();
			setState(489);
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
		enterRule(_localctx, 96, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(491);
			exp();
			setState(492);
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
		enterRule(_localctx, 98, RULE_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(494);
			t_exp();
			setState(498);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OR) {
				{
				{
				setState(495);
				exp2();
				}
				}
				setState(500);
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
		enterRule(_localctx, 100, RULE_exp2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(501);
			match(OR);
			setState(502);
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
		enterRule(_localctx, 102, RULE_t_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(504);
			g_exp();
			setState(508);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==AND) {
				{
				{
				setState(505);
				t_exp2();
				}
				}
				setState(510);
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
		enterRule(_localctx, 104, RULE_t_exp2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(511);
			match(AND);
			setState(512);
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
		enterRule(_localctx, 106, RULE_g_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(514);
			m_exp();
			setState(516);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << GT) | (1L << LT) | (1L << EQ) | (1L << NEQ))) != 0)) {
				{
				setState(515);
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
		enterRule(_localctx, 108, RULE_g_exp2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(518);
			relop();
			setState(519);
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
		enterRule(_localctx, 110, RULE_m_exp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(521);
			term();
			setState(525);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ADD || _la==SUB) {
				{
				{
				setState(522);
				m_exp2();
				}
				}
				setState(527);
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
		enterRule(_localctx, 112, RULE_m_exp2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(528);
			_la = _input.LA(1);
			if ( !(_la==ADD || _la==SUB) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(529);
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
		enterRule(_localctx, 114, RULE_term);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(531);
			factor();
			setState(535);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==MUL || _la==DIV) {
				{
				{
				setState(532);
				term2();
				}
				}
				setState(537);
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
		enterRule(_localctx, 116, RULE_term2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(538);
			_la = _input.LA(1);
			if ( !(_la==MUL || _la==DIV) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(539);
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
		enterRule(_localctx, 118, RULE_factor);
		try {
			setState(549);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,44,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(541);
				factor2();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(542);
				varCte();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(546);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,43,_ctx) ) {
				case 1:
					{
					setState(543);
					vars();
					}
					break;
				case 2:
					{
					setState(544);
					varArray();
					}
					break;
				case 3:
					{
					setState(545);
					varMat();
					}
					break;
				}
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(548);
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
		enterRule(_localctx, 120, RULE_factor2);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(551);
			match(LPAREN);
			setState(552);
			exp();
			setState(553);
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
		enterRule(_localctx, 122, RULE_varCte);
		try {
			setState(560);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(555);
				cte_i();
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(556);
				cte_f();
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 3);
				{
				setState(557);
				cte_c();
				}
				break;
			case T__9:
				enterOuterAlt(_localctx, 4);
				{
				setState(558);
				cte_s();
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 5);
				{
				setState(559);
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
		enterRule(_localctx, 124, RULE_cte_i);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(562);
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
		enterRule(_localctx, 126, RULE_cte_f);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(564);
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
		enterRule(_localctx, 128, RULE_cte_c);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(566);
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
		enterRule(_localctx, 130, RULE_cte_b);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(568);
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
		enterRule(_localctx, 132, RULE_cte_s);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(570);
			match(T__9);
			setState(574);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,46,_ctx);
			while ( _alt!=1 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1+1 ) {
					{
					{
					setState(571);
					matchWildcard();
					}
					} 
				}
				setState(576);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,46,_ctx);
			}
			setState(577);
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
		enterRule(_localctx, 134, RULE_main);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(579);
			match(MAIN);
			setState(580);
			match(LPAREN);
			setState(581);
			match(RPAREN);
			setState(582);
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
		enterRule(_localctx, 136, RULE_mainBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(584);
			match(T__0);
			setState(588);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(585);
				variableDeclaration();
				}
				}
				setState(590);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(594);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__9) | (1L << INT) | (1L << FLOAT) | (1L << CHAR) | (1L << BOOL) | (1L << LPAREN) | (1L << IF) | (1L << WHILE) | (1L << FOR) | (1L << WRITE) | (1L << READ) | (1L << RETURN) | (1L << ID))) != 0)) {
				{
				{
				setState(591);
				statutes();
				}
				}
				setState(596);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(597);
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
		enterRule(_localctx, 138, RULE_typeRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(599);
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
		enterRule(_localctx, 140, RULE_relop);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(601);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\67\u025e\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\3\2"+
		"\3\2\3\2\3\2\7\2\u0095\n\2\f\2\16\2\u0098\13\2\3\2\3\2\3\3\7\3\u009d\n"+
		"\3\f\3\16\3\u00a0\13\3\3\3\7\3\u00a3\n\3\f\3\16\3\u00a6\13\3\3\3\3\3\3"+
		"\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\7\5\u00b5\n\5\f\5\16\5\u00b8"+
		"\13\5\3\5\3\5\3\6\3\6\7\6\u00be\n\6\f\6\16\6\u00c1\13\6\3\7\3\7\5\7\u00c5"+
		"\n\7\5\7\u00c7\n\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\7\b\u00d0\n\b\f\b\16\b"+
		"\u00d3\13\b\3\b\7\b\u00d6\n\b\f\b\16\b\u00d9\13\b\3\b\3\b\3\t\3\t\3\t"+
		"\5\t\u00e0\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\5\r"+
		"\u00fe\n\r\5\r\u0100\n\r\3\r\3\r\3\r\5\r\u0105\n\r\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\17\3\17\5\17\u010f\n\17\3\17\5\17\u0112\n\17\3\20\3\20\3"+
		"\20\3\21\3\21\3\21\5\21\u011a\n\21\3\22\3\22\3\22\5\22\u011f\n\22\3\22"+
		"\3\22\3\22\3\22\3\23\3\23\3\23\5\23\u0128\n\23\3\23\3\23\3\23\3\23\5\23"+
		"\u012e\n\23\3\23\3\23\3\23\3\23\5\23\u0134\n\23\3\24\3\24\3\24\3\24\5"+
		"\24\u013a\n\24\3\24\3\24\3\24\5\24\u013f\n\24\3\24\3\24\3\25\3\25\3\25"+
		"\7\25\u0146\n\25\f\25\16\25\u0149\13\25\3\26\3\26\3\26\3\26\3\27\3\27"+
		"\7\27\u0151\n\27\f\27\16\27\u0154\13\27\3\27\7\27\u0157\n\27\f\27\16\27"+
		"\u015a\13\27\3\27\3\27\3\30\3\30\5\30\u0160\n\30\3\30\3\30\3\31\3\31\7"+
		"\31\u0166\n\31\f\31\16\31\u0169\13\31\3\31\3\31\3\32\3\32\3\32\3\32\3"+
		"\32\3\32\3\32\5\32\u0174\n\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\5\33\u017e\n\33\3\34\3\34\3\34\5\34\u0183\n\34\3\34\3\34\3\34\3\34\3"+
		"\35\3\35\3\35\5\35\u018c\n\35\3\35\3\35\3\36\3\36\3\36\3\37\3\37\3\37"+
		"\5\37\u0196\n\37\3 \3 \3 \3 \3 \5 \u019d\n \3 \3 \3!\3!\3!\3!\3!\3\"\3"+
		"\"\3\"\5\"\u01a9\n\"\3#\3#\3#\3#\3#\3#\3$\3$\3$\5$\u01b4\n$\3$\3$\3%\3"+
		"%\3%\5%\u01bb\n%\3&\3&\3&\3&\3&\3&\3\'\3\'\3\'\3(\3(\3(\5(\u01c9\n(\3"+
		")\3)\3)\3)\5)\u01cf\n)\3*\3*\3*\3*\3+\3+\3,\3,\3,\3-\3-\3-\3-\3-\3-\3"+
		".\3.\3.\3.\3/\3/\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\61\3\62\3\62\3\62"+
		"\3\63\3\63\7\63\u01f3\n\63\f\63\16\63\u01f6\13\63\3\64\3\64\3\64\3\65"+
		"\3\65\7\65\u01fd\n\65\f\65\16\65\u0200\13\65\3\66\3\66\3\66\3\67\3\67"+
		"\5\67\u0207\n\67\38\38\38\39\39\79\u020e\n9\f9\169\u0211\139\3:\3:\3:"+
		"\3;\3;\7;\u0218\n;\f;\16;\u021b\13;\3<\3<\3<\3=\3=\3=\3=\3=\5=\u0225\n"+
		"=\3=\5=\u0228\n=\3>\3>\3>\3>\3?\3?\3?\3?\3?\5?\u0233\n?\3@\3@\3A\3A\3"+
		"B\3B\3C\3C\3D\3D\7D\u023f\nD\fD\16D\u0242\13D\3D\3D\3E\3E\3E\3E\3E\3F"+
		"\3F\7F\u024d\nF\fF\16F\u0250\13F\3F\7F\u0253\nF\fF\16F\u0256\13F\3F\3"+
		"F\3G\3G\3H\3H\3H\3\u0240\2I\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \""+
		"$&(*,.\60\62\64\668:<>@BDFHJLNPRTVXZ\\^`bdfhjlnprtvxz|~\u0080\u0082\u0084"+
		"\u0086\u0088\u008a\u008c\u008e\2\6\3\2\32\33\3\2\30\31\3\2\61\64\3\2\24"+
		"\27\2\u025d\2\u0090\3\2\2\2\4\u009e\3\2\2\2\6\u00aa\3\2\2\2\b\u00af\3"+
		"\2\2\2\n\u00bb\3\2\2\2\f\u00c6\3\2\2\2\16\u00ca\3\2\2\2\20\u00df\3\2\2"+
		"\2\22\u00e1\3\2\2\2\24\u00e7\3\2\2\2\26\u00f0\3\2\2\2\30\u00ff\3\2\2\2"+
		"\32\u0106\3\2\2\2\34\u010e\3\2\2\2\36\u0113\3\2\2\2 \u0116\3\2\2\2\"\u011b"+
		"\3\2\2\2$\u0124\3\2\2\2&\u0135\3\2\2\2(\u0142\3\2\2\2*\u014a\3\2\2\2,"+
		"\u014e\3\2\2\2.\u015d\3\2\2\2\60\u0163\3\2\2\2\62\u0173\3\2\2\2\64\u017d"+
		"\3\2\2\2\66\u0182\3\2\2\28\u0188\3\2\2\2:\u018f\3\2\2\2<\u0195\3\2\2\2"+
		">\u0197\3\2\2\2@\u01a0\3\2\2\2B\u01a8\3\2\2\2D\u01aa\3\2\2\2F\u01b3\3"+
		"\2\2\2H\u01ba\3\2\2\2J\u01bc\3\2\2\2L\u01c2\3\2\2\2N\u01c8\3\2\2\2P\u01ca"+
		"\3\2\2\2R\u01d0\3\2\2\2T\u01d4\3\2\2\2V\u01d6\3\2\2\2X\u01d9\3\2\2\2Z"+
		"\u01df\3\2\2\2\\\u01e3\3\2\2\2^\u01e5\3\2\2\2`\u01e9\3\2\2\2b\u01ed\3"+
		"\2\2\2d\u01f0\3\2\2\2f\u01f7\3\2\2\2h\u01fa\3\2\2\2j\u0201\3\2\2\2l\u0204"+
		"\3\2\2\2n\u0208\3\2\2\2p\u020b\3\2\2\2r\u0212\3\2\2\2t\u0215\3\2\2\2v"+
		"\u021c\3\2\2\2x\u0227\3\2\2\2z\u0229\3\2\2\2|\u0232\3\2\2\2~\u0234\3\2"+
		"\2\2\u0080\u0236\3\2\2\2\u0082\u0238\3\2\2\2\u0084\u023a\3\2\2\2\u0086"+
		"\u023c\3\2\2\2\u0088\u0245\3\2\2\2\u008a\u024a\3\2\2\2\u008c\u0259\3\2"+
		"\2\2\u008e\u025b\3\2\2\2\u0090\u0091\7\66\2\2\u0091\u0092\7\67\2\2\u0092"+
		"\u0096\7\34\2\2\u0093\u0095\5\20\t\2\u0094\u0093\3\2\2\2\u0095\u0098\3"+
		"\2\2\2\u0096\u0094\3\2\2\2\u0096\u0097\3\2\2\2\u0097\u0099\3\2\2\2\u0098"+
		"\u0096\3\2\2\2\u0099\u009a\5\4\3\2\u009a\3\3\2\2\2\u009b\u009d\5\6\4\2"+
		"\u009c\u009b\3\2\2\2\u009d\u00a0\3\2\2\2\u009e\u009c\3\2\2\2\u009e\u009f"+
		"\3\2\2\2\u009f\u00a4\3\2\2\2\u00a0\u009e\3\2\2\2\u00a1\u00a3\5&\24\2\u00a2"+
		"\u00a1\3\2\2\2\u00a3\u00a6\3\2\2\2\u00a4\u00a2\3\2\2\2\u00a4\u00a5\3\2"+
		"\2\2\u00a5\u00a7\3\2\2\2\u00a6\u00a4\3\2\2\2\u00a7\u00a8\5\u0088E\2\u00a8"+
		"\u00a9\7\2\2\3\u00a9\5\3\2\2\2\u00aa\u00ab\7%\2\2\u00ab\u00ac\7\67\2\2"+
		"\u00ac\u00ad\5\b\5\2\u00ad\u00ae\7\34\2\2\u00ae\7\3\2\2\2\u00af\u00b0"+
		"\7\3\2\2\u00b0\u00b1\5\n\6\2\u00b1\u00b2\5\16\b\2\u00b2\u00b6\7\'\2\2"+
		"\u00b3\u00b5\5\f\7\2\u00b4\u00b3\3\2\2\2\u00b5\u00b8\3\2\2\2\u00b6\u00b4"+
		"\3\2\2\2\u00b6\u00b7\3\2\2\2\u00b7\u00b9\3\2\2\2\u00b8\u00b6\3\2\2\2\u00b9"+
		"\u00ba\7\4\2\2\u00ba\t\3\2\2\2\u00bb\u00bf\7&\2\2\u00bc\u00be\5\30\r\2"+
		"\u00bd\u00bc\3\2\2\2\u00be\u00c1\3\2\2\2\u00bf\u00bd\3\2\2\2\u00bf\u00c0"+
		"\3\2\2\2\u00c0\13\3\2\2\2\u00c1\u00bf\3\2\2\2\u00c2\u00c7\7)\2\2\u00c3"+
		"\u00c5\7*\2\2\u00c4\u00c3\3\2\2\2\u00c4\u00c5\3\2\2\2\u00c5\u00c7\3\2"+
		"\2\2\u00c6\u00c2\3\2\2\2\u00c6\u00c4\3\2\2\2\u00c7\u00c8\3\2\2\2\u00c8"+
		"\u00c9\5&\24\2\u00c9\r\3\2\2\2\u00ca\u00cb\7(\2\2\u00cb\u00cc\7\36\2\2"+
		"\u00cc\u00cd\7\37\2\2\u00cd\u00d1\7\3\2\2\u00ce\u00d0\5\20\t\2\u00cf\u00ce"+
		"\3\2\2\2\u00d0\u00d3\3\2\2\2\u00d1\u00cf\3\2\2\2\u00d1\u00d2\3\2\2\2\u00d2"+
		"\u00d7\3\2\2\2\u00d3\u00d1\3\2\2\2\u00d4\u00d6\5\62\32\2\u00d5\u00d4\3"+
		"\2\2\2\u00d6\u00d9\3\2\2\2\u00d7\u00d5\3\2\2\2\u00d7\u00d8\3\2\2\2\u00d8"+
		"\u00da\3\2\2\2\u00d9\u00d7\3\2\2\2\u00da\u00db\7\4\2\2\u00db\17\3\2\2"+
		"\2\u00dc\u00e0\5\22\n\2\u00dd\u00e0\5\24\13\2\u00de\u00e0\5\26\f\2\u00df"+
		"\u00dc\3\2\2\2\u00df\u00dd\3\2\2\2\u00df\u00de\3\2\2\2\u00e0\21\3\2\2"+
		"\2\u00e1\u00e2\7\60\2\2\u00e2\u00e3\7\67\2\2\u00e3\u00e4\7\5\2\2\u00e4"+
		"\u00e5\5\34\17\2\u00e5\u00e6\7\34\2\2\u00e6\23\3\2\2\2\u00e7\u00e8\7\60"+
		"\2\2\u00e8\u00e9\7\67\2\2\u00e9\u00ea\7\5\2\2\u00ea\u00eb\7\6\2\2\u00eb"+
		"\u00ec\7\r\2\2\u00ec\u00ed\7\7\2\2\u00ed\u00ee\5\u008cG\2\u00ee\u00ef"+
		"\7\34\2\2\u00ef\25\3\2\2\2\u00f0\u00f1\7\60\2\2\u00f1\u00f2\7\67\2\2\u00f2"+
		"\u00f3\7\5\2\2\u00f3\u00f4\7\6\2\2\u00f4\u00f5\7\r\2\2\u00f5\u00f6\7\b"+
		"\2\2\u00f6\u00f7\7\r\2\2\u00f7\u00f8\7\7\2\2\u00f8\u00f9\5\u008cG\2\u00f9"+
		"\u00fa\7\34\2\2\u00fa\27\3\2\2\2\u00fb\u0100\7)\2\2\u00fc\u00fe\7*\2\2"+
		"\u00fd\u00fc\3\2\2\2\u00fd\u00fe\3\2\2\2\u00fe\u0100\3\2\2\2\u00ff\u00fb"+
		"\3\2\2\2\u00ff\u00fd\3\2\2\2\u0100\u0104\3\2\2\2\u0101\u0105\5\32\16\2"+
		"\u0102\u0105\5\24\13\2\u0103\u0105\5\26\f\2\u0104\u0101\3\2\2\2\u0104"+
		"\u0102\3\2\2\2\u0104\u0103\3\2\2\2\u0105\31\3\2\2\2\u0106\u0107\7\60\2"+
		"\2\u0107\u0108\7\67\2\2\u0108\u0109\7\5\2\2\u0109\u010a\5\u008cG\2\u010a"+
		"\u010b\7\34\2\2\u010b\33\3\2\2\2\u010c\u010f\5\u008cG\2\u010d\u010f\7"+
		"\67\2\2\u010e\u010c\3\2\2\2\u010e\u010d\3\2\2\2\u010f\u0111\3\2\2\2\u0110"+
		"\u0112\5\36\20\2\u0111\u0110\3\2\2\2\u0111\u0112\3\2\2\2\u0112\35\3\2"+
		"\2\2\u0113\u0114\7\35\2\2\u0114\u0115\5d\63\2\u0115\37\3\2\2\2\u0116\u0119"+
		"\7\67\2\2\u0117\u0118\7\t\2\2\u0118\u011a\7\67\2\2\u0119\u0117\3\2\2\2"+
		"\u0119\u011a\3\2\2\2\u011a!\3\2\2\2\u011b\u011e\7\67\2\2\u011c\u011d\7"+
		"\t\2\2\u011d\u011f\7\67\2\2\u011e\u011c\3\2\2\2\u011e\u011f\3\2\2\2\u011f"+
		"\u0120\3\2\2\2\u0120\u0121\7\6\2\2\u0121\u0122\5d\63\2\u0122\u0123\7\7"+
		"\2\2\u0123#\3\2\2\2\u0124\u0127\7\67\2\2\u0125\u0126\7\t\2\2\u0126\u0128"+
		"\7\67\2\2\u0127\u0125\3\2\2\2\u0127\u0128\3\2\2\2\u0128\u012d\3\2\2\2"+
		"\u0129\u012a\7\6\2\2\u012a\u012b\5d\63\2\u012b\u012c\7\7\2\2\u012c\u012e"+
		"\3\2\2\2\u012d\u0129\3\2\2\2\u012d\u012e\3\2\2\2\u012e\u0133\3\2\2\2\u012f"+
		"\u0130\7\6\2\2\u0130\u0131\5d\63\2\u0131\u0132\7\7\2\2\u0132\u0134\3\2"+
		"\2\2\u0133\u012f\3\2\2\2\u0133\u0134\3\2\2\2\u0134%\3\2\2\2\u0135\u0136"+
		"\7-\2\2\u0136\u0137\7\67\2\2\u0137\u0139\7\36\2\2\u0138\u013a\5(\25\2"+
		"\u0139\u0138\3\2\2\2\u0139\u013a\3\2\2\2\u013a\u013b\3\2\2\2\u013b\u013e"+
		"\7\37\2\2\u013c\u013f\5\u008cG\2\u013d\u013f\7\65\2\2\u013e\u013c\3\2"+
		"\2\2\u013e\u013d\3\2\2\2\u013f\u0140\3\2\2\2\u0140\u0141\5,\27\2\u0141"+
		"\'\3\2\2\2\u0142\u0147\5*\26\2\u0143\u0144\7\n\2\2\u0144\u0146\5*\26\2"+
		"\u0145\u0143\3\2\2\2\u0146\u0149\3\2\2\2\u0147\u0145\3\2\2\2\u0147\u0148"+
		"\3\2\2\2\u0148)\3\2\2\2\u0149\u0147\3\2\2\2\u014a\u014b\7\67\2\2\u014b"+
		"\u014c\7\5\2\2\u014c\u014d\5\u008cG\2\u014d+\3\2\2\2\u014e\u0152\7\3\2"+
		"\2\u014f\u0151\5\20\t\2\u0150\u014f\3\2\2\2\u0151\u0154\3\2\2\2\u0152"+
		"\u0150\3\2\2\2\u0152\u0153\3\2\2\2\u0153\u0158\3\2\2\2\u0154\u0152\3\2"+
		"\2\2\u0155\u0157\5\64\33\2\u0156\u0155\3\2\2\2\u0157\u015a\3\2\2\2\u0158"+
		"\u0156\3\2\2\2\u0158\u0159\3\2\2\2\u0159\u015b\3\2\2\2\u015a\u0158\3\2"+
		"\2\2\u015b\u015c\7\4\2\2\u015c-\3\2\2\2\u015d\u015f\7.\2\2\u015e\u0160"+
		"\5d\63\2\u015f\u015e\3\2\2\2\u015f\u0160\3\2\2\2\u0160\u0161\3\2\2\2\u0161"+
		"\u0162\7\34\2\2\u0162/\3\2\2\2\u0163\u0167\7\3\2\2\u0164\u0166\5\64\33"+
		"\2\u0165\u0164\3\2\2\2\u0166\u0169\3\2\2\2\u0167\u0165\3\2\2\2\u0167\u0168"+
		"\3\2\2\2\u0168\u016a\3\2\2\2\u0169\u0167\3\2\2\2\u016a\u016b\7\4\2\2\u016b"+
		"\61\3\2\2\2\u016c\u0174\5\66\34\2\u016d\u0174\5D#\2\u016e\u0174\5J&\2"+
		"\u016f\u0174\5P)\2\u0170\u0174\5X-\2\u0171\u0174\5^\60\2\u0172\u0174\5"+
		"b\62\2\u0173\u016c\3\2\2\2\u0173\u016d\3\2\2\2\u0173\u016e\3\2\2\2\u0173"+
		"\u016f\3\2\2\2\u0173\u0170\3\2\2\2\u0173\u0171\3\2\2\2\u0173\u0172\3\2"+
		"\2\2\u0174\63\3\2\2\2\u0175\u017e\5\66\34\2\u0176\u017e\5D#\2\u0177\u017e"+
		"\5J&\2\u0178\u017e\5P)\2\u0179\u017e\5X-\2\u017a\u017e\5^\60\2\u017b\u017e"+
		"\5b\62\2\u017c\u017e\5.\30\2\u017d\u0175\3\2\2\2\u017d\u0176\3\2\2\2\u017d"+
		"\u0177\3\2\2\2\u017d\u0178\3\2\2\2\u017d\u0179\3\2\2\2\u017d\u017a\3\2"+
		"\2\2\u017d\u017b\3\2\2\2\u017d\u017c\3\2\2\2\u017e\65\3\2\2\2\u017f\u0183"+
		"\5 \21\2\u0180\u0183\5\"\22\2\u0181\u0183\5$\23\2\u0182\u017f\3\2\2\2"+
		"\u0182\u0180\3\2\2\2\u0182\u0181\3\2\2\2\u0183\u0184\3\2\2\2\u0184\u0185"+
		"\7\35\2\2\u0185\u0186\5d\63\2\u0186\u0187\7\34\2\2\u0187\67\3\2\2\2\u0188"+
		"\u0189\7\67\2\2\u0189\u018b\7\36\2\2\u018a\u018c\5:\36\2\u018b\u018a\3"+
		"\2\2\2\u018b\u018c\3\2\2\2\u018c\u018d\3\2\2\2\u018d\u018e\7\37\2\2\u018e"+
		"9\3\2\2\2\u018f\u0190\5d\63\2\u0190\u0191\5<\37\2\u0191;\3\2\2\2\u0192"+
		"\u0193\7\n\2\2\u0193\u0196\5:\36\2\u0194\u0196\3\2\2\2\u0195\u0192\3\2"+
		"\2\2\u0195\u0194\3\2\2\2\u0196=\3\2\2\2\u0197\u0198\7\67\2\2\u0198\u0199"+
		"\7\t\2\2\u0199\u019a\7\67\2\2\u019a\u019c\7\36\2\2\u019b\u019d\5:\36\2"+
		"\u019c\u019b\3\2\2\2\u019c\u019d\3\2\2\2\u019d\u019e\3\2\2\2\u019e\u019f"+
		"\7\37\2\2\u019f?\3\2\2\2\u01a0\u01a1\7\13\2\2\u01a1\u01a2\7\67\2\2\u01a2"+
		"\u01a3\7\36\2\2\u01a3\u01a4\7\37\2\2\u01a4A\3\2\2\2\u01a5\u01a9\58\35"+
		"\2\u01a6\u01a9\5> \2\u01a7\u01a9\5@!\2\u01a8\u01a5\3\2\2\2\u01a8\u01a6"+
		"\3\2\2\2\u01a8\u01a7\3\2\2\2\u01a9C\3\2\2\2\u01aa\u01ab\7,\2\2\u01ab\u01ac"+
		"\7\36\2\2\u01ac\u01ad\5F$\2\u01ad\u01ae\7\37\2\2\u01ae\u01af\7\34\2\2"+
		"\u01afE\3\2\2\2\u01b0\u01b4\5 \21\2\u01b1\u01b4\5\"\22\2\u01b2\u01b4\5"+
		"$\23\2\u01b3\u01b0\3\2\2\2\u01b3\u01b1\3\2\2\2\u01b3\u01b2\3\2\2\2\u01b4"+
		"\u01b5\3\2\2\2\u01b5\u01b6\5H%\2\u01b6G\3\2\2\2\u01b7\u01b8\7\n\2\2\u01b8"+
		"\u01bb\5F$\2\u01b9\u01bb\3\2\2\2\u01ba\u01b7\3\2\2\2\u01ba\u01b9\3\2\2"+
		"\2\u01bbI\3\2\2\2\u01bc\u01bd\7+\2\2\u01bd\u01be\7\36\2\2\u01be\u01bf"+
		"\5L\'\2\u01bf\u01c0\7\37\2\2\u01c0\u01c1\7\34\2\2\u01c1K\3\2\2\2\u01c2"+
		"\u01c3\5d\63\2\u01c3\u01c4\5N(\2\u01c4M\3\2\2\2\u01c5\u01c6\7\n\2\2\u01c6"+
		"\u01c9\5L\'\2\u01c7\u01c9\3\2\2\2\u01c8\u01c5\3\2\2\2\u01c8\u01c7\3\2"+
		"\2\2\u01c9O\3\2\2\2\u01ca\u01cb\7 \2\2\u01cb\u01cc\5R*\2\u01cc\u01ce\5"+
		"T+\2\u01cd\u01cf\5V,\2\u01ce\u01cd\3\2\2\2\u01ce\u01cf\3\2\2\2\u01cfQ"+
		"\3\2\2\2\u01d0\u01d1\7\36\2\2\u01d1\u01d2\5d\63\2\u01d2\u01d3\7\37\2\2"+
		"\u01d3S\3\2\2\2\u01d4\u01d5\5\60\31\2\u01d5U\3\2\2\2\u01d6\u01d7\7!\2"+
		"\2\u01d7\u01d8\5\60\31\2\u01d8W\3\2\2\2\u01d9\u01da\7#\2\2\u01da\u01db"+
		"\5Z.\2\u01db\u01dc\7$\2\2\u01dc\u01dd\5\\/\2\u01dd\u01de\5\60\31\2\u01de"+
		"Y\3\2\2\2\u01df\u01e0\7\67\2\2\u01e0\u01e1\7\35\2\2\u01e1\u01e2\5d\63"+
		"\2\u01e2[\3\2\2\2\u01e3\u01e4\5d\63\2\u01e4]\3\2\2\2\u01e5\u01e6\7\"\2"+
		"\2\u01e6\u01e7\5`\61\2\u01e7\u01e8\5\60\31\2\u01e8_\3\2\2\2\u01e9\u01ea"+
		"\7\36\2\2\u01ea\u01eb\5d\63\2\u01eb\u01ec\7\37\2\2\u01eca\3\2\2\2\u01ed"+
		"\u01ee\5d\63\2\u01ee\u01ef\7\34\2\2\u01efc\3\2\2\2\u01f0\u01f4\5h\65\2"+
		"\u01f1\u01f3\5f\64\2\u01f2\u01f1\3\2\2\2\u01f3\u01f6\3\2\2\2\u01f4\u01f2"+
		"\3\2\2\2\u01f4\u01f5\3\2\2\2\u01f5e\3\2\2\2\u01f6\u01f4\3\2\2\2\u01f7"+
		"\u01f8\7\22\2\2\u01f8\u01f9\5h\65\2\u01f9g\3\2\2\2\u01fa\u01fe\5l\67\2"+
		"\u01fb\u01fd\5j\66\2\u01fc\u01fb\3\2\2\2\u01fd\u0200\3\2\2\2\u01fe\u01fc"+
		"\3\2\2\2\u01fe\u01ff\3\2\2\2\u01ffi\3\2\2\2\u0200\u01fe\3\2\2\2\u0201"+
		"\u0202\7\23\2\2\u0202\u0203\5l\67\2\u0203k\3\2\2\2\u0204\u0206\5p9\2\u0205"+
		"\u0207\5n8\2\u0206\u0205\3\2\2\2\u0206\u0207\3\2\2\2\u0207m\3\2\2\2\u0208"+
		"\u0209\5\u008eH\2\u0209\u020a\5p9\2\u020ao\3\2\2\2\u020b\u020f\5t;\2\u020c"+
		"\u020e\5r:\2\u020d\u020c\3\2\2\2\u020e\u0211\3\2\2\2\u020f\u020d\3\2\2"+
		"\2\u020f\u0210\3\2\2\2\u0210q\3\2\2\2\u0211\u020f\3\2\2\2\u0212\u0213"+
		"\t\2\2\2\u0213\u0214\5t;\2\u0214s\3\2\2\2\u0215\u0219\5x=\2\u0216\u0218"+
		"\5v<\2\u0217\u0216\3\2\2\2\u0218\u021b\3\2\2\2\u0219\u0217\3\2\2\2\u0219"+
		"\u021a\3\2\2\2\u021au\3\2\2\2\u021b\u0219\3\2\2\2\u021c\u021d\t\3\2\2"+
		"\u021d\u021e\5x=\2\u021ew\3\2\2\2\u021f\u0228\5z>\2\u0220\u0228\5|?\2"+
		"\u0221\u0225\5 \21\2\u0222\u0225\5\"\22\2\u0223\u0225\5$\23\2\u0224\u0221"+
		"\3\2\2\2\u0224\u0222\3\2\2\2\u0224\u0223\3\2\2\2\u0225\u0228\3\2\2\2\u0226"+
		"\u0228\5B\"\2\u0227\u021f\3\2\2\2\u0227\u0220\3\2\2\2\u0227\u0224\3\2"+
		"\2\2\u0227\u0226\3\2\2\2\u0228y\3\2\2\2\u0229\u022a\7\36\2\2\u022a\u022b"+
		"\5d\63\2\u022b\u022c\7\37\2\2\u022c{\3\2\2\2\u022d\u0233\5~@\2\u022e\u0233"+
		"\5\u0080A\2\u022f\u0233\5\u0082B\2\u0230\u0233\5\u0086D\2\u0231\u0233"+
		"\5\u0084C\2\u0232\u022d\3\2\2\2\u0232\u022e\3\2\2\2\u0232\u022f\3\2\2"+
		"\2\u0232\u0230\3\2\2\2\u0232\u0231\3\2\2\2\u0233}\3\2\2\2\u0234\u0235"+
		"\7\r\2\2\u0235\177\3\2\2\2\u0236\u0237\7\16\2\2\u0237\u0081\3\2\2\2\u0238"+
		"\u0239\7\17\2\2\u0239\u0083\3\2\2\2\u023a\u023b\7\20\2\2\u023b\u0085\3"+
		"\2\2\2\u023c\u0240\7\f\2\2\u023d\u023f\13\2\2\2\u023e\u023d\3\2\2\2\u023f"+
		"\u0242\3\2\2\2\u0240\u0241\3\2\2\2\u0240\u023e\3\2\2\2\u0241\u0243\3\2"+
		"\2\2\u0242\u0240\3\2\2\2\u0243\u0244\7\f\2\2\u0244\u0087\3\2\2\2\u0245"+
		"\u0246\7/\2\2\u0246\u0247\7\36\2\2\u0247\u0248\7\37\2\2\u0248\u0249\5"+
		"\u008aF\2\u0249\u0089\3\2\2\2\u024a\u024e\7\3\2\2\u024b\u024d\5\20\t\2"+
		"\u024c\u024b\3\2\2\2\u024d\u0250\3\2\2\2\u024e\u024c\3\2\2\2\u024e\u024f"+
		"\3\2\2\2\u024f\u0254\3\2\2\2\u0250\u024e\3\2\2\2\u0251\u0253\5\64\33\2"+
		"\u0252\u0251\3\2\2\2\u0253\u0256\3\2\2\2\u0254\u0252\3\2\2\2\u0254\u0255"+
		"\3\2\2\2\u0255\u0257\3\2\2\2\u0256\u0254\3\2\2\2\u0257\u0258\7\4\2\2\u0258"+
		"\u008b\3\2\2\2\u0259\u025a\t\4\2\2\u025a\u008d\3\2\2\2\u025b\u025c\t\5"+
		"\2\2\u025c\u008f\3\2\2\2\63\u0096\u009e\u00a4\u00b6\u00bf\u00c4\u00c6"+
		"\u00d1\u00d7\u00df\u00fd\u00ff\u0104\u010e\u0111\u0119\u011e\u0127\u012d"+
		"\u0133\u0139\u013e\u0147\u0152\u0158\u015f\u0167\u0173\u017d\u0182\u018b"+
		"\u0195\u019c\u01a8\u01b3\u01ba\u01c8\u01ce\u01f4\u01fe\u0206\u020f\u0219"+
		"\u0224\u0227\u0232\u0240\u024e\u0254";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}