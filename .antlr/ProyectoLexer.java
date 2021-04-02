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
		T__9=10, T__10=11, T__11=12, INT=13, FLOAT=14, CHAR=15, WS=16, GT=17, 
		LT=18, EQ=19, NEQ=20, MUL=21, DIV=22, ADD=23, SUB=24, SEMI=25, IF=26, 
		ELSE=27, WHILE=28, FOR=29, IN=30, CLASS=31, ATTRIBUTES=32, METHODS=33, 
		WRITE=34, READ=35, FUNCTION=36, RETURN=37, MAIN=38, VAR=39, INT_TYPE=40, 
		FLOAT_TYPE=41, CHAR_TYPE=42, STRING_TYPE=43, VOID=44, PROGRAM=45, ID=46;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "T__11", "INT", "FLOAT", "CHAR", "WS", "GT", "LT", "EQ", 
			"NEQ", "MUL", "DIV", "ADD", "SUB", "SEMI", "IF", "ELSE", "WHILE", "FOR", 
			"IN", "CLASS", "ATTRIBUTES", "METHODS", "WRITE", "READ", "FUNCTION", 
			"RETURN", "MAIN", "VAR", "INT_TYPE", "FLOAT_TYPE", "CHAR_TYPE", "STRING_TYPE", 
			"VOID", "PROGRAM", "ID"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'{'", "'}'", "':'", "'['", "']'", "']['", "'='", "'('", "')'", 
			"','", "'.'", "'\"'", null, null, null, null, "'>'", "'<'", "'=='", "'!='", 
			"'*'", "'/'", "'+'", "'-'", "';'", "'if'", "'else'", "'while'", "'for'", 
			"'in'", "'class'", "'attributes'", "'methods'", "'write'", "'read'", 
			"'function'", "'return'", "'main'", "'var'", "'int'", "'float'", "'char'", 
			"'string'", "'void'", "'program'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, "INT", "FLOAT", "CHAR", "WS", "GT", "LT", "EQ", "NEQ", "MUL", "DIV", 
			"ADD", "SUB", "SEMI", "IF", "ELSE", "WHILE", "FOR", "IN", "CLASS", "ATTRIBUTES", 
			"METHODS", "WRITE", "READ", "FUNCTION", "RETURN", "MAIN", "VAR", "INT_TYPE", 
			"FLOAT_TYPE", "CHAR_TYPE", "STRING_TYPE", "VOID", "PROGRAM", "ID"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\60\u0124\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7"+
		"\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\6\16z"+
		"\n\16\r\16\16\16{\3\17\3\17\3\17\5\17\u0081\n\17\3\17\3\17\5\17\u0085"+
		"\n\17\3\17\5\17\u0088\n\17\3\20\3\20\3\20\3\20\3\21\6\21\u008f\n\21\r"+
		"\21\16\21\u0090\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3\24\3\25\3\25"+
		"\3\25\3\26\3\26\3\27\3\27\3\30\3\30\3\31\3\31\3\32\3\32\3\33\3\33\3\33"+
		"\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36"+
		"\3\36\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3"+
		"!\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3$\3$\3$\3$\3$\3%"+
		"\3%\3%\3%\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3(\3"+
		"(\3(\3(\3)\3)\3)\3)\3*\3*\3*\3*\3*\3*\3+\3+\3+\3+\3+\3,\3,\3,\3,\3,\3"+
		",\3,\3-\3-\3-\3-\3-\3.\3.\3.\3.\3.\3.\3.\3.\3/\3/\7/\u0120\n/\f/\16/\u0123"+
		"\13/\2\2\60\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33"+
		"\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67"+
		"\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60\3\2\b\3\2\62;\4\2--"+
		"//\4\2C\\c|\5\2\13\f\17\17\"\"\5\2C\\aac|\5\2\62;C\\c|\2\u0129\2\3\3\2"+
		"\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17"+
		"\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2"+
		"\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3"+
		"\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3"+
		"\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2"+
		"=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3"+
		"\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2"+
		"\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\3_\3\2\2\2\5a\3\2\2\2\7"+
		"c\3\2\2\2\te\3\2\2\2\13g\3\2\2\2\ri\3\2\2\2\17l\3\2\2\2\21n\3\2\2\2\23"+
		"p\3\2\2\2\25r\3\2\2\2\27t\3\2\2\2\31v\3\2\2\2\33y\3\2\2\2\35}\3\2\2\2"+
		"\37\u0089\3\2\2\2!\u008e\3\2\2\2#\u0094\3\2\2\2%\u0096\3\2\2\2\'\u0098"+
		"\3\2\2\2)\u009b\3\2\2\2+\u009e\3\2\2\2-\u00a0\3\2\2\2/\u00a2\3\2\2\2\61"+
		"\u00a4\3\2\2\2\63\u00a6\3\2\2\2\65\u00a8\3\2\2\2\67\u00ab\3\2\2\29\u00b0"+
		"\3\2\2\2;\u00b6\3\2\2\2=\u00ba\3\2\2\2?\u00bd\3\2\2\2A\u00c3\3\2\2\2C"+
		"\u00ce\3\2\2\2E\u00d6\3\2\2\2G\u00dc\3\2\2\2I\u00e1\3\2\2\2K\u00ea\3\2"+
		"\2\2M\u00f1\3\2\2\2O\u00f6\3\2\2\2Q\u00fa\3\2\2\2S\u00fe\3\2\2\2U\u0104"+
		"\3\2\2\2W\u0109\3\2\2\2Y\u0110\3\2\2\2[\u0115\3\2\2\2]\u011d\3\2\2\2_"+
		"`\7}\2\2`\4\3\2\2\2ab\7\177\2\2b\6\3\2\2\2cd\7<\2\2d\b\3\2\2\2ef\7]\2"+
		"\2f\n\3\2\2\2gh\7_\2\2h\f\3\2\2\2ij\7_\2\2jk\7]\2\2k\16\3\2\2\2lm\7?\2"+
		"\2m\20\3\2\2\2no\7*\2\2o\22\3\2\2\2pq\7+\2\2q\24\3\2\2\2rs\7.\2\2s\26"+
		"\3\2\2\2tu\7\60\2\2u\30\3\2\2\2vw\7$\2\2w\32\3\2\2\2xz\t\2\2\2yx\3\2\2"+
		"\2z{\3\2\2\2{y\3\2\2\2{|\3\2\2\2|\34\3\2\2\2}\u0080\5\33\16\2~\177\7\60"+
		"\2\2\177\u0081\5\33\16\2\u0080~\3\2\2\2\u0080\u0081\3\2\2\2\u0081\u0087"+
		"\3\2\2\2\u0082\u0084\7G\2\2\u0083\u0085\t\3\2\2\u0084\u0083\3\2\2\2\u0084"+
		"\u0085\3\2\2\2\u0085\u0086\3\2\2\2\u0086\u0088\5\33\16\2\u0087\u0082\3"+
		"\2\2\2\u0087\u0088\3\2\2\2\u0088\36\3\2\2\2\u0089\u008a\7$\2\2\u008a\u008b"+
		"\t\4\2\2\u008b\u008c\7$\2\2\u008c \3\2\2\2\u008d\u008f\t\5\2\2\u008e\u008d"+
		"\3\2\2\2\u008f\u0090\3\2\2\2\u0090\u008e\3\2\2\2\u0090\u0091\3\2\2\2\u0091"+
		"\u0092\3\2\2\2\u0092\u0093\b\21\2\2\u0093\"\3\2\2\2\u0094\u0095\7@\2\2"+
		"\u0095$\3\2\2\2\u0096\u0097\7>\2\2\u0097&\3\2\2\2\u0098\u0099\7?\2\2\u0099"+
		"\u009a\7?\2\2\u009a(\3\2\2\2\u009b\u009c\7#\2\2\u009c\u009d\7?\2\2\u009d"+
		"*\3\2\2\2\u009e\u009f\7,\2\2\u009f,\3\2\2\2\u00a0\u00a1\7\61\2\2\u00a1"+
		".\3\2\2\2\u00a2\u00a3\7-\2\2\u00a3\60\3\2\2\2\u00a4\u00a5\7/\2\2\u00a5"+
		"\62\3\2\2\2\u00a6\u00a7\7=\2\2\u00a7\64\3\2\2\2\u00a8\u00a9\7k\2\2\u00a9"+
		"\u00aa\7h\2\2\u00aa\66\3\2\2\2\u00ab\u00ac\7g\2\2\u00ac\u00ad\7n\2\2\u00ad"+
		"\u00ae\7u\2\2\u00ae\u00af\7g\2\2\u00af8\3\2\2\2\u00b0\u00b1\7y\2\2\u00b1"+
		"\u00b2\7j\2\2\u00b2\u00b3\7k\2\2\u00b3\u00b4\7n\2\2\u00b4\u00b5\7g\2\2"+
		"\u00b5:\3\2\2\2\u00b6\u00b7\7h\2\2\u00b7\u00b8\7q\2\2\u00b8\u00b9\7t\2"+
		"\2\u00b9<\3\2\2\2\u00ba\u00bb\7k\2\2\u00bb\u00bc\7p\2\2\u00bc>\3\2\2\2"+
		"\u00bd\u00be\7e\2\2\u00be\u00bf\7n\2\2\u00bf\u00c0\7c\2\2\u00c0\u00c1"+
		"\7u\2\2\u00c1\u00c2\7u\2\2\u00c2@\3\2\2\2\u00c3\u00c4\7c\2\2\u00c4\u00c5"+
		"\7v\2\2\u00c5\u00c6\7v\2\2\u00c6\u00c7\7t\2\2\u00c7\u00c8\7k\2\2\u00c8"+
		"\u00c9\7d\2\2\u00c9\u00ca\7w\2\2\u00ca\u00cb\7v\2\2\u00cb\u00cc\7g\2\2"+
		"\u00cc\u00cd\7u\2\2\u00cdB\3\2\2\2\u00ce\u00cf\7o\2\2\u00cf\u00d0\7g\2"+
		"\2\u00d0\u00d1\7v\2\2\u00d1\u00d2\7j\2\2\u00d2\u00d3\7q\2\2\u00d3\u00d4"+
		"\7f\2\2\u00d4\u00d5\7u\2\2\u00d5D\3\2\2\2\u00d6\u00d7\7y\2\2\u00d7\u00d8"+
		"\7t\2\2\u00d8\u00d9\7k\2\2\u00d9\u00da\7v\2\2\u00da\u00db\7g\2\2\u00db"+
		"F\3\2\2\2\u00dc\u00dd\7t\2\2\u00dd\u00de\7g\2\2\u00de\u00df\7c\2\2\u00df"+
		"\u00e0\7f\2\2\u00e0H\3\2\2\2\u00e1\u00e2\7h\2\2\u00e2\u00e3\7w\2\2\u00e3"+
		"\u00e4\7p\2\2\u00e4\u00e5\7e\2\2\u00e5\u00e6\7v\2\2\u00e6\u00e7\7k\2\2"+
		"\u00e7\u00e8\7q\2\2\u00e8\u00e9\7p\2\2\u00e9J\3\2\2\2\u00ea\u00eb\7t\2"+
		"\2\u00eb\u00ec\7g\2\2\u00ec\u00ed\7v\2\2\u00ed\u00ee\7w\2\2\u00ee\u00ef"+
		"\7t\2\2\u00ef\u00f0\7p\2\2\u00f0L\3\2\2\2\u00f1\u00f2\7o\2\2\u00f2\u00f3"+
		"\7c\2\2\u00f3\u00f4\7k\2\2\u00f4\u00f5\7p\2\2\u00f5N\3\2\2\2\u00f6\u00f7"+
		"\7x\2\2\u00f7\u00f8\7c\2\2\u00f8\u00f9\7t\2\2\u00f9P\3\2\2\2\u00fa\u00fb"+
		"\7k\2\2\u00fb\u00fc\7p\2\2\u00fc\u00fd\7v\2\2\u00fdR\3\2\2\2\u00fe\u00ff"+
		"\7h\2\2\u00ff\u0100\7n\2\2\u0100\u0101\7q\2\2\u0101\u0102\7c\2\2\u0102"+
		"\u0103\7v\2\2\u0103T\3\2\2\2\u0104\u0105\7e\2\2\u0105\u0106\7j\2\2\u0106"+
		"\u0107\7c\2\2\u0107\u0108\7t\2\2\u0108V\3\2\2\2\u0109\u010a\7u\2\2\u010a"+
		"\u010b\7v\2\2\u010b\u010c\7t\2\2\u010c\u010d\7k\2\2\u010d\u010e\7p\2\2"+
		"\u010e\u010f\7i\2\2\u010fX\3\2\2\2\u0110\u0111\7x\2\2\u0111\u0112\7q\2"+
		"\2\u0112\u0113\7k\2\2\u0113\u0114\7f\2\2\u0114Z\3\2\2\2\u0115\u0116\7"+
		"r\2\2\u0116\u0117\7t\2\2\u0117\u0118\7q\2\2\u0118\u0119\7i\2\2\u0119\u011a"+
		"\7t\2\2\u011a\u011b\7c\2\2\u011b\u011c\7o\2\2\u011c\\\3\2\2\2\u011d\u0121"+
		"\t\6\2\2\u011e\u0120\t\7\2\2\u011f\u011e\3\2\2\2\u0120\u0123\3\2\2\2\u0121"+
		"\u011f\3\2\2\2\u0121\u0122\3\2\2\2\u0122^\3\2\2\2\u0123\u0121\3\2\2\2"+
		"\t\2{\u0080\u0084\u0087\u0090\u0121\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}