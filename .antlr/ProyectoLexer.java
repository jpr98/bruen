// Generated from /Users/juanpabloramos/Documents/TEC/OctavoSemestre/Compiladores/Proyecto/Proyecto.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ProyectoLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "INT", "FLOAT", "CHAR", "BOOL", "WS", "OR", "AND", "GT", "LT", 
			"EQ", "NEQ", "MUL", "DIV", "ADD", "SUB", "SEMI", "LPAREN", "RPAREN", 
			"IF", "ELSE", "WHILE", "FOR", "IN", "CLASS", "ATTRIBUTES", "METHODS", 
			"WRITE", "READ", "FUNCTION", "RETURN", "MAIN", "VAR", "INT_TYPE", "FLOAT_TYPE", 
			"CHAR_TYPE", "STRING_TYPE", "BOOL_TYPE", "VOID", "PROGRAM", "ID"
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


	public ProyectoLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Proyecto.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\64\u0142\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\3\2"+
		"\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3"+
		"\n\3\13\3\13\3\f\6\f~\n\f\r\f\16\f\177\3\r\3\r\3\r\5\r\u0085\n\r\3\r\3"+
		"\r\5\r\u0089\n\r\3\r\5\r\u008c\n\r\3\16\3\16\3\16\3\16\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\5\17\u009b\n\17\3\20\6\20\u009e\n\20\r"+
		"\20\16\20\u009f\3\20\3\20\3\21\3\21\3\21\3\22\3\22\3\22\3\23\3\23\3\24"+
		"\3\24\3\25\3\25\3\25\3\26\3\26\3\26\3\27\3\27\3\30\3\30\3\31\3\31\3\32"+
		"\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\36\3\36\3\36\3\37\3\37\3\37\3\37"+
		"\3\37\3 \3 \3 \3 \3 \3 \3!\3!\3!\3!\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3$\3"+
		"$\3$\3$\3$\3$\3$\3$\3$\3$\3$\3%\3%\3%\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3"+
		"&\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3(\3(\3(\3(\3(\3)\3)\3)\3)\3)\3)\3)"+
		"\3*\3*\3*\3*\3*\3+\3+\3+\3+\3,\3,\3,\3,\3-\3-\3-\3-\3-\3-\3.\3.\3.\3."+
		"\3.\3/\3/\3/\3/\3/\3/\3/\3\60\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\61"+
		"\3\61\3\62\3\62\3\62\3\62\3\62\3\62\3\62\3\62\3\63\3\63\7\63\u013e\n\63"+
		"\f\63\16\63\u0141\13\63\2\2\64\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13"+
		"\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61"+
		"\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61"+
		"a\62c\63e\64\3\2\b\3\2\62;\4\2--//\4\2C\\c|\5\2\13\f\17\17\"\"\5\2C\\"+
		"aac|\5\2\62;C\\c|\2\u0148\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2"+
		"\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2"+
		"\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3"+
		"\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2"+
		"\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67"+
		"\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2"+
		"\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2"+
		"\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]"+
		"\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\3g\3\2\2\2\5i\3\2"+
		"\2\2\7k\3\2\2\2\tm\3\2\2\2\13o\3\2\2\2\rq\3\2\2\2\17t\3\2\2\2\21v\3\2"+
		"\2\2\23x\3\2\2\2\25z\3\2\2\2\27}\3\2\2\2\31\u0081\3\2\2\2\33\u008d\3\2"+
		"\2\2\35\u009a\3\2\2\2\37\u009d\3\2\2\2!\u00a3\3\2\2\2#\u00a6\3\2\2\2%"+
		"\u00a9\3\2\2\2\'\u00ab\3\2\2\2)\u00ad\3\2\2\2+\u00b0\3\2\2\2-\u00b3\3"+
		"\2\2\2/\u00b5\3\2\2\2\61\u00b7\3\2\2\2\63\u00b9\3\2\2\2\65\u00bb\3\2\2"+
		"\2\67\u00bd\3\2\2\29\u00bf\3\2\2\2;\u00c1\3\2\2\2=\u00c4\3\2\2\2?\u00c9"+
		"\3\2\2\2A\u00cf\3\2\2\2C\u00d3\3\2\2\2E\u00d6\3\2\2\2G\u00dc\3\2\2\2I"+
		"\u00e7\3\2\2\2K\u00ef\3\2\2\2M\u00f5\3\2\2\2O\u00fa\3\2\2\2Q\u0103\3\2"+
		"\2\2S\u010a\3\2\2\2U\u010f\3\2\2\2W\u0113\3\2\2\2Y\u0117\3\2\2\2[\u011d"+
		"\3\2\2\2]\u0122\3\2\2\2_\u0129\3\2\2\2a\u012e\3\2\2\2c\u0133\3\2\2\2e"+
		"\u013b\3\2\2\2gh\7}\2\2h\4\3\2\2\2ij\7\177\2\2j\6\3\2\2\2kl\7<\2\2l\b"+
		"\3\2\2\2mn\7]\2\2n\n\3\2\2\2op\7_\2\2p\f\3\2\2\2qr\7_\2\2rs\7]\2\2s\16"+
		"\3\2\2\2tu\7?\2\2u\20\3\2\2\2vw\7\60\2\2w\22\3\2\2\2xy\7.\2\2y\24\3\2"+
		"\2\2z{\7$\2\2{\26\3\2\2\2|~\t\2\2\2}|\3\2\2\2~\177\3\2\2\2\177}\3\2\2"+
		"\2\177\u0080\3\2\2\2\u0080\30\3\2\2\2\u0081\u0084\5\27\f\2\u0082\u0083"+
		"\7\60\2\2\u0083\u0085\5\27\f\2\u0084\u0082\3\2\2\2\u0084\u0085\3\2\2\2"+
		"\u0085\u008b\3\2\2\2\u0086\u0088\7G\2\2\u0087\u0089\t\3\2\2\u0088\u0087"+
		"\3\2\2\2\u0088\u0089\3\2\2\2\u0089\u008a\3\2\2\2\u008a\u008c\5\27\f\2"+
		"\u008b\u0086\3\2\2\2\u008b\u008c\3\2\2\2\u008c\32\3\2\2\2\u008d\u008e"+
		"\7$\2\2\u008e\u008f\t\4\2\2\u008f\u0090\7$\2\2\u0090\34\3\2\2\2\u0091"+
		"\u0092\7v\2\2\u0092\u0093\7t\2\2\u0093\u0094\7w\2\2\u0094\u009b\7g\2\2"+
		"\u0095\u0096\7h\2\2\u0096\u0097\7c\2\2\u0097\u0098\7n\2\2\u0098\u0099"+
		"\7u\2\2\u0099\u009b\7g\2\2\u009a\u0091\3\2\2\2\u009a\u0095\3\2\2\2\u009b"+
		"\36\3\2\2\2\u009c\u009e\t\5\2\2\u009d\u009c\3\2\2\2\u009e\u009f\3\2\2"+
		"\2\u009f\u009d\3\2\2\2\u009f\u00a0\3\2\2\2\u00a0\u00a1\3\2\2\2\u00a1\u00a2"+
		"\b\20\2\2\u00a2 \3\2\2\2\u00a3\u00a4\7~\2\2\u00a4\u00a5\7~\2\2\u00a5\""+
		"\3\2\2\2\u00a6\u00a7\7(\2\2\u00a7\u00a8\7(\2\2\u00a8$\3\2\2\2\u00a9\u00aa"+
		"\7@\2\2\u00aa&\3\2\2\2\u00ab\u00ac\7>\2\2\u00ac(\3\2\2\2\u00ad\u00ae\7"+
		"?\2\2\u00ae\u00af\7?\2\2\u00af*\3\2\2\2\u00b0\u00b1\7#\2\2\u00b1\u00b2"+
		"\7?\2\2\u00b2,\3\2\2\2\u00b3\u00b4\7,\2\2\u00b4.\3\2\2\2\u00b5\u00b6\7"+
		"\61\2\2\u00b6\60\3\2\2\2\u00b7\u00b8\7-\2\2\u00b8\62\3\2\2\2\u00b9\u00ba"+
		"\7/\2\2\u00ba\64\3\2\2\2\u00bb\u00bc\7=\2\2\u00bc\66\3\2\2\2\u00bd\u00be"+
		"\7*\2\2\u00be8\3\2\2\2\u00bf\u00c0\7+\2\2\u00c0:\3\2\2\2\u00c1\u00c2\7"+
		"k\2\2\u00c2\u00c3\7h\2\2\u00c3<\3\2\2\2\u00c4\u00c5\7g\2\2\u00c5\u00c6"+
		"\7n\2\2\u00c6\u00c7\7u\2\2\u00c7\u00c8\7g\2\2\u00c8>\3\2\2\2\u00c9\u00ca"+
		"\7y\2\2\u00ca\u00cb\7j\2\2\u00cb\u00cc\7k\2\2\u00cc\u00cd\7n\2\2\u00cd"+
		"\u00ce\7g\2\2\u00ce@\3\2\2\2\u00cf\u00d0\7h\2\2\u00d0\u00d1\7q\2\2\u00d1"+
		"\u00d2\7t\2\2\u00d2B\3\2\2\2\u00d3\u00d4\7k\2\2\u00d4\u00d5\7p\2\2\u00d5"+
		"D\3\2\2\2\u00d6\u00d7\7e\2\2\u00d7\u00d8\7n\2\2\u00d8\u00d9\7c\2\2\u00d9"+
		"\u00da\7u\2\2\u00da\u00db\7u\2\2\u00dbF\3\2\2\2\u00dc\u00dd\7c\2\2\u00dd"+
		"\u00de\7v\2\2\u00de\u00df\7v\2\2\u00df\u00e0\7t\2\2\u00e0\u00e1\7k\2\2"+
		"\u00e1\u00e2\7d\2\2\u00e2\u00e3\7w\2\2\u00e3\u00e4\7v\2\2\u00e4\u00e5"+
		"\7g\2\2\u00e5\u00e6\7u\2\2\u00e6H\3\2\2\2\u00e7\u00e8\7o\2\2\u00e8\u00e9"+
		"\7g\2\2\u00e9\u00ea\7v\2\2\u00ea\u00eb\7j\2\2\u00eb\u00ec\7q\2\2\u00ec"+
		"\u00ed\7f\2\2\u00ed\u00ee\7u\2\2\u00eeJ\3\2\2\2\u00ef\u00f0\7y\2\2\u00f0"+
		"\u00f1\7t\2\2\u00f1\u00f2\7k\2\2\u00f2\u00f3\7v\2\2\u00f3\u00f4\7g\2\2"+
		"\u00f4L\3\2\2\2\u00f5\u00f6\7t\2\2\u00f6\u00f7\7g\2\2\u00f7\u00f8\7c\2"+
		"\2\u00f8\u00f9\7f\2\2\u00f9N\3\2\2\2\u00fa\u00fb\7h\2\2\u00fb\u00fc\7"+
		"w\2\2\u00fc\u00fd\7p\2\2\u00fd\u00fe\7e\2\2\u00fe\u00ff\7v\2\2\u00ff\u0100"+
		"\7k\2\2\u0100\u0101\7q\2\2\u0101\u0102\7p\2\2\u0102P\3\2\2\2\u0103\u0104"+
		"\7t\2\2\u0104\u0105\7g\2\2\u0105\u0106\7v\2\2\u0106\u0107\7w\2\2\u0107"+
		"\u0108\7t\2\2\u0108\u0109\7p\2\2\u0109R\3\2\2\2\u010a\u010b\7o\2\2\u010b"+
		"\u010c\7c\2\2\u010c\u010d\7k\2\2\u010d\u010e\7p\2\2\u010eT\3\2\2\2\u010f"+
		"\u0110\7x\2\2\u0110\u0111\7c\2\2\u0111\u0112\7t\2\2\u0112V\3\2\2\2\u0113"+
		"\u0114\7k\2\2\u0114\u0115\7p\2\2\u0115\u0116\7v\2\2\u0116X\3\2\2\2\u0117"+
		"\u0118\7h\2\2\u0118\u0119\7n\2\2\u0119\u011a\7q\2\2\u011a\u011b\7c\2\2"+
		"\u011b\u011c\7v\2\2\u011cZ\3\2\2\2\u011d\u011e\7e\2\2\u011e\u011f\7j\2"+
		"\2\u011f\u0120\7c\2\2\u0120\u0121\7t\2\2\u0121\\\3\2\2\2\u0122\u0123\7"+
		"u\2\2\u0123\u0124\7v\2\2\u0124\u0125\7t\2\2\u0125\u0126\7k\2\2\u0126\u0127"+
		"\7p\2\2\u0127\u0128\7i\2\2\u0128^\3\2\2\2\u0129\u012a\7d\2\2\u012a\u012b"+
		"\7q\2\2\u012b\u012c\7q\2\2\u012c\u012d\7n\2\2\u012d`\3\2\2\2\u012e\u012f"+
		"\7x\2\2\u012f\u0130\7q\2\2\u0130\u0131\7k\2\2\u0131\u0132\7f\2\2\u0132"+
		"b\3\2\2\2\u0133\u0134\7r\2\2\u0134\u0135\7t\2\2\u0135\u0136\7q\2\2\u0136"+
		"\u0137\7i\2\2\u0137\u0138\7t\2\2\u0138\u0139\7c\2\2\u0139\u013a\7o\2\2"+
		"\u013ad\3\2\2\2\u013b\u013f\t\6\2\2\u013c\u013e\t\7\2\2\u013d\u013c\3"+
		"\2\2\2\u013e\u0141\3\2\2\2\u013f\u013d\3\2\2\2\u013f\u0140\3\2\2\2\u0140"+
		"f\3\2\2\2\u0141\u013f\3\2\2\2\n\2\177\u0084\u0088\u008b\u009a\u009f\u013f"+
		"\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}