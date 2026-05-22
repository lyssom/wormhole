// Generated from d:\wormhole\server\query\SqlParser.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SqlParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		WS=1, CREATE=2, DATABASE=3, DEFAULT=4, EXISTS=5, FROM=6, IF=7, INF=8, 
		INSERT=9, INTO=10, NAN_SQL=11, NOT=12, NULL_SQL=13, SELECT=14, SET=15, 
		TABLE=16, USE=17, VALUE=18, VALUES=19, SECOND=20, MINUTE=21, HOUR=22, 
		DAY=23, WEEK=24, MONTH=25, QUARTER=26, YEAR=27, IDENTIFIER=28, FLOATING_LITERAL=29, 
		OCTAL_LITERAL=30, DECIMAL_LITERAL=31, HEXADECIMAL_LITERAL=32, EXPONENT_NUM_PART=33, 
		STRING_LITERAL=34, ARROW=35, ASTERISK=36, BACKQUOTE=37, BACKSLASH=38, 
		COLON=39, COMMA=40, CONCAT=41, DASH=42, DOT=43, EQ_DOUBLE=44, EQ_SINGLE=45, 
		GE=46, GT=47, LBRACE=48, LBRACKET=49, LE=50, LPAREN=51, LT=52, NOT_EQ=53, 
		PERCENT=54, PLUS=55, QUERY=56, QUOTE_DOUBLE=57, QUOTE_SINGLE=58, RBRACE=59, 
		RBRACKET=60, RPAREN=61, SEMICOLON=62, SLASH=63, UNDERSCORE=64, MULTI_LINE_COMMENT=65, 
		SINGLE_LINE_COMMENT=66, WHITESPACE=67;
	public static final int
		RULE_statement = 0, RULE_dmlStatement = 1, RULE_ddlStatement = 2, RULE_systemStatement = 3, 
		RULE_createDatabase = 4, RULE_createTable = 5, RULE_tableSchemaClause = 6, 
		RULE_tableElementExpr = 7, RULE_tableColumnDfnt = 8, RULE_insertStatement = 9, 
		RULE_columnsClause = 10, RULE_insertValuesSpec = 11, RULE_insertMultiValue = 12, 
		RULE_dataValue = 13, RULE_selectStatement = 14, RULE_fromClause = 15, 
		RULE_joinExpr = 16, RULE_useStatement = 17, RULE_columnTypeExpr = 18, 
		RULE_columnExprList = 19, RULE_columnsExpr = 20, RULE_columnExpr = 21, 
		RULE_columnIdentifier = 22, RULE_nestedIdentifier = 23, RULE_constant = 24, 
		RULE_realLiteral = 25, RULE_tableExpr = 26, RULE_tableIdentifier = 27, 
		RULE_databaseIdentifier = 28, RULE_floatingLiteral = 29, RULE_numberLiteral = 30, 
		RULE_literal = 31, RULE_interval = 32, RULE_keyword = 33, RULE_identifier = 34, 
		RULE_enumValue = 35, RULE_ifNotExists = 36;
	private static String[] makeRuleNames() {
		return new String[] {
			"statement", "dmlStatement", "ddlStatement", "systemStatement", "createDatabase", 
			"createTable", "tableSchemaClause", "tableElementExpr", "tableColumnDfnt", 
			"insertStatement", "columnsClause", "insertValuesSpec", "insertMultiValue", 
			"dataValue", "selectStatement", "fromClause", "joinExpr", "useStatement", 
			"columnTypeExpr", "columnExprList", "columnsExpr", "columnExpr", "columnIdentifier", 
			"nestedIdentifier", "constant", "realLiteral", "tableExpr", "tableIdentifier", 
			"databaseIdentifier", "floatingLiteral", "numberLiteral", "literal", 
			"interval", "keyword", "identifier", "enumValue", "ifNotExists"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, "'->'", 
			"'*'", "'`'", "'\\'", "':'", "','", "'||'", "'-'", "'.'", "'=='", "'='", 
			"'>='", "'>'", "'{'", "'['", "'<='", "'('", "'<'", null, "'%'", "'+'", 
			"'?'", "'\"'", "'''", "'}'", "']'", "')'", "';'", "'/'", "'_'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "WS", "CREATE", "DATABASE", "DEFAULT", "EXISTS", "FROM", "IF", 
			"INF", "INSERT", "INTO", "NAN_SQL", "NOT", "NULL_SQL", "SELECT", "SET", 
			"TABLE", "USE", "VALUE", "VALUES", "SECOND", "MINUTE", "HOUR", "DAY", 
			"WEEK", "MONTH", "QUARTER", "YEAR", "IDENTIFIER", "FLOATING_LITERAL", 
			"OCTAL_LITERAL", "DECIMAL_LITERAL", "HEXADECIMAL_LITERAL", "EXPONENT_NUM_PART", 
			"STRING_LITERAL", "ARROW", "ASTERISK", "BACKQUOTE", "BACKSLASH", "COLON", 
			"COMMA", "CONCAT", "DASH", "DOT", "EQ_DOUBLE", "EQ_SINGLE", "GE", "GT", 
			"LBRACE", "LBRACKET", "LE", "LPAREN", "LT", "NOT_EQ", "PERCENT", "PLUS", 
			"QUERY", "QUOTE_DOUBLE", "QUOTE_SINGLE", "RBRACE", "RBRACKET", "RPAREN", 
			"SEMICOLON", "SLASH", "UNDERSCORE", "MULTI_LINE_COMMENT", "SINGLE_LINE_COMMENT", 
			"WHITESPACE"
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
	public String getGrammarFileName() { return "SqlParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SqlParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StatementContext extends ParserRuleContext {
		public DdlStatementContext ddlStatement() {
			return getRuleContext(DdlStatementContext.class,0);
		}
		public DmlStatementContext dmlStatement() {
			return getRuleContext(DmlStatementContext.class,0);
		}
		public SystemStatementContext systemStatement() {
			return getRuleContext(SystemStatementContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_statement);
		try {
			setState(77);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CREATE:
				enterOuterAlt(_localctx, 1);
				{
				setState(74);
				ddlStatement();
				}
				break;
			case INSERT:
			case SELECT:
				enterOuterAlt(_localctx, 2);
				{
				setState(75);
				dmlStatement();
				}
				break;
			case USE:
				enterOuterAlt(_localctx, 3);
				{
				setState(76);
				systemStatement();
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

	public static class DmlStatementContext extends ParserRuleContext {
		public SelectStatementContext selectStatement() {
			return getRuleContext(SelectStatementContext.class,0);
		}
		public InsertStatementContext insertStatement() {
			return getRuleContext(InsertStatementContext.class,0);
		}
		public DmlStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dmlStatement; }
	}

	public final DmlStatementContext dmlStatement() throws RecognitionException {
		DmlStatementContext _localctx = new DmlStatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_dmlStatement);
		try {
			setState(81);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SELECT:
				enterOuterAlt(_localctx, 1);
				{
				setState(79);
				selectStatement();
				}
				break;
			case INSERT:
				enterOuterAlt(_localctx, 2);
				{
				setState(80);
				insertStatement();
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

	public static class DdlStatementContext extends ParserRuleContext {
		public CreateDatabaseContext createDatabase() {
			return getRuleContext(CreateDatabaseContext.class,0);
		}
		public CreateTableContext createTable() {
			return getRuleContext(CreateTableContext.class,0);
		}
		public DdlStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ddlStatement; }
	}

	public final DdlStatementContext ddlStatement() throws RecognitionException {
		DdlStatementContext _localctx = new DdlStatementContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_ddlStatement);
		try {
			setState(85);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(83);
				createDatabase();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(84);
				createTable();
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

	public static class SystemStatementContext extends ParserRuleContext {
		public UseStatementContext useStatement() {
			return getRuleContext(UseStatementContext.class,0);
		}
		public SystemStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_systemStatement; }
	}

	public final SystemStatementContext systemStatement() throws RecognitionException {
		SystemStatementContext _localctx = new SystemStatementContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_systemStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(87);
			useStatement();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CreateDatabaseContext extends ParserRuleContext {
		public TerminalNode CREATE() { return getToken(SqlParser.CREATE, 0); }
		public TerminalNode DATABASE() { return getToken(SqlParser.DATABASE, 0); }
		public DatabaseIdentifierContext databaseIdentifier() {
			return getRuleContext(DatabaseIdentifierContext.class,0);
		}
		public IfNotExistsContext ifNotExists() {
			return getRuleContext(IfNotExistsContext.class,0);
		}
		public CreateDatabaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_createDatabase; }
	}

	public final CreateDatabaseContext createDatabase() throws RecognitionException {
		CreateDatabaseContext _localctx = new CreateDatabaseContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_createDatabase);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			match(CREATE);
			setState(90);
			match(DATABASE);
			setState(92);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				{
				setState(91);
				ifNotExists();
				}
				break;
			}
			setState(94);
			databaseIdentifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CreateTableContext extends ParserRuleContext {
		public TerminalNode CREATE() { return getToken(SqlParser.CREATE, 0); }
		public TerminalNode TABLE() { return getToken(SqlParser.TABLE, 0); }
		public TableIdentifierContext tableIdentifier() {
			return getRuleContext(TableIdentifierContext.class,0);
		}
		public TableSchemaClauseContext tableSchemaClause() {
			return getRuleContext(TableSchemaClauseContext.class,0);
		}
		public IfNotExistsContext ifNotExists() {
			return getRuleContext(IfNotExistsContext.class,0);
		}
		public CreateTableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_createTable; }
	}

	public final CreateTableContext createTable() throws RecognitionException {
		CreateTableContext _localctx = new CreateTableContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_createTable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(96);
			match(CREATE);
			setState(97);
			match(TABLE);
			setState(99);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				{
				setState(98);
				ifNotExists();
				}
				break;
			}
			setState(101);
			tableIdentifier();
			setState(102);
			tableSchemaClause();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TableSchemaClauseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(SqlParser.LPAREN, 0); }
		public List<TableElementExprContext> tableElementExpr() {
			return getRuleContexts(TableElementExprContext.class);
		}
		public TableElementExprContext tableElementExpr(int i) {
			return getRuleContext(TableElementExprContext.class,i);
		}
		public TerminalNode RPAREN() { return getToken(SqlParser.RPAREN, 0); }
		public List<TerminalNode> COMMA() { return getTokens(SqlParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(SqlParser.COMMA, i);
		}
		public TableSchemaClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tableSchemaClause; }
	}

	public final TableSchemaClauseContext tableSchemaClause() throws RecognitionException {
		TableSchemaClauseContext _localctx = new TableSchemaClauseContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_tableSchemaClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(104);
			match(LPAREN);
			setState(105);
			tableElementExpr();
			setState(110);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(106);
				match(COMMA);
				setState(107);
				tableElementExpr();
				}
				}
				setState(112);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(113);
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

	public static class TableElementExprContext extends ParserRuleContext {
		public TableColumnDfntContext tableColumnDfnt() {
			return getRuleContext(TableColumnDfntContext.class,0);
		}
		public TableElementExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tableElementExpr; }
	}

	public final TableElementExprContext tableElementExpr() throws RecognitionException {
		TableElementExprContext _localctx = new TableElementExprContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_tableElementExpr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(115);
			tableColumnDfnt();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TableColumnDfntContext extends ParserRuleContext {
		public NestedIdentifierContext nestedIdentifier() {
			return getRuleContext(NestedIdentifierContext.class,0);
		}
		public ColumnTypeExprContext columnTypeExpr() {
			return getRuleContext(ColumnTypeExprContext.class,0);
		}
		public TableColumnDfntContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tableColumnDfnt; }
	}

	public final TableColumnDfntContext tableColumnDfnt() throws RecognitionException {
		TableColumnDfntContext _localctx = new TableColumnDfntContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_tableColumnDfnt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			nestedIdentifier();
			setState(118);
			columnTypeExpr();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InsertStatementContext extends ParserRuleContext {
		public TerminalNode INSERT() { return getToken(SqlParser.INSERT, 0); }
		public TerminalNode INTO() { return getToken(SqlParser.INTO, 0); }
		public TerminalNode VALUES() { return getToken(SqlParser.VALUES, 0); }
		public InsertValuesSpecContext insertValuesSpec() {
			return getRuleContext(InsertValuesSpecContext.class,0);
		}
		public TableIdentifierContext tableIdentifier() {
			return getRuleContext(TableIdentifierContext.class,0);
		}
		public TerminalNode TABLE() { return getToken(SqlParser.TABLE, 0); }
		public ColumnsClauseContext columnsClause() {
			return getRuleContext(ColumnsClauseContext.class,0);
		}
		public InsertStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_insertStatement; }
	}

	public final InsertStatementContext insertStatement() throws RecognitionException {
		InsertStatementContext _localctx = new InsertStatementContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_insertStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(120);
			match(INSERT);
			setState(121);
			match(INTO);
			setState(123);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				{
				setState(122);
				match(TABLE);
				}
				break;
			}
			{
			setState(125);
			tableIdentifier();
			}
			setState(127);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LPAREN) {
				{
				setState(126);
				columnsClause();
				}
			}

			setState(129);
			match(VALUES);
			setState(130);
			insertValuesSpec();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ColumnsClauseContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(SqlParser.LPAREN, 0); }
		public List<NestedIdentifierContext> nestedIdentifier() {
			return getRuleContexts(NestedIdentifierContext.class);
		}
		public NestedIdentifierContext nestedIdentifier(int i) {
			return getRuleContext(NestedIdentifierContext.class,i);
		}
		public TerminalNode RPAREN() { return getToken(SqlParser.RPAREN, 0); }
		public List<TerminalNode> COMMA() { return getTokens(SqlParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(SqlParser.COMMA, i);
		}
		public ColumnsClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnsClause; }
	}

	public final ColumnsClauseContext columnsClause() throws RecognitionException {
		ColumnsClauseContext _localctx = new ColumnsClauseContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_columnsClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(132);
			match(LPAREN);
			setState(133);
			nestedIdentifier();
			setState(138);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(134);
				match(COMMA);
				setState(135);
				nestedIdentifier();
				}
				}
				setState(140);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(141);
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

	public static class InsertValuesSpecContext extends ParserRuleContext {
		public List<InsertMultiValueContext> insertMultiValue() {
			return getRuleContexts(InsertMultiValueContext.class);
		}
		public InsertMultiValueContext insertMultiValue(int i) {
			return getRuleContext(InsertMultiValueContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(SqlParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(SqlParser.COMMA, i);
		}
		public InsertValuesSpecContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_insertValuesSpec; }
	}

	public final InsertValuesSpecContext insertValuesSpec() throws RecognitionException {
		InsertValuesSpecContext _localctx = new InsertValuesSpecContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_insertValuesSpec);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(149);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA || _la==LPAREN) {
				{
				{
				setState(144);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COMMA) {
					{
					setState(143);
					match(COMMA);
					}
				}

				setState(146);
				insertMultiValue();
				}
				}
				setState(151);
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

	public static class InsertMultiValueContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(SqlParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(SqlParser.RPAREN, 0); }
		public List<DataValueContext> dataValue() {
			return getRuleContexts(DataValueContext.class);
		}
		public DataValueContext dataValue(int i) {
			return getRuleContext(DataValueContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(SqlParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(SqlParser.COMMA, i);
		}
		public InsertMultiValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_insertMultiValue; }
	}

	public final InsertMultiValueContext insertMultiValue() throws RecognitionException {
		InsertMultiValueContext _localctx = new InsertMultiValueContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_insertMultiValue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			match(LPAREN);
			setState(157); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(153);
				dataValue();
				setState(155);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COMMA) {
					{
					setState(154);
					match(COMMA);
					}
				}

				}
				}
				setState(159); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << DECIMAL_LITERAL) | (1L << EXPONENT_NUM_PART) | (1L << STRING_LITERAL) | (1L << DASH) | (1L << DOT) | (1L << LPAREN) | (1L << PLUS))) != 0) );
			setState(161);
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

	public static class DataValueContext extends ParserRuleContext {
		public List<ConstantContext> constant() {
			return getRuleContexts(ConstantContext.class);
		}
		public ConstantContext constant(int i) {
			return getRuleContext(ConstantContext.class,i);
		}
		public TerminalNode LPAREN() { return getToken(SqlParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(SqlParser.RPAREN, 0); }
		public List<TerminalNode> COMMA() { return getTokens(SqlParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(SqlParser.COMMA, i);
		}
		public DataValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dataValue; }
	}

	public final DataValueContext dataValue() throws RecognitionException {
		DataValueContext _localctx = new DataValueContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_dataValue);
		int _la;
		try {
			setState(174);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DECIMAL_LITERAL:
			case EXPONENT_NUM_PART:
			case STRING_LITERAL:
			case DASH:
			case DOT:
			case PLUS:
				enterOuterAlt(_localctx, 1);
				{
				setState(163);
				constant();
				}
				break;
			case LPAREN:
				enterOuterAlt(_localctx, 2);
				{
				setState(164);
				match(LPAREN);
				setState(165);
				constant();
				setState(168); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(166);
					match(COMMA);
					setState(167);
					constant();
					}
					}
					setState(170); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==COMMA );
				setState(172);
				match(RPAREN);
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

	public static class SelectStatementContext extends ParserRuleContext {
		public TerminalNode SELECT() { return getToken(SqlParser.SELECT, 0); }
		public ColumnExprListContext columnExprList() {
			return getRuleContext(ColumnExprListContext.class,0);
		}
		public FromClauseContext fromClause() {
			return getRuleContext(FromClauseContext.class,0);
		}
		public SelectStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selectStatement; }
	}

	public final SelectStatementContext selectStatement() throws RecognitionException {
		SelectStatementContext _localctx = new SelectStatementContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_selectStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(176);
			match(SELECT);
			setState(177);
			columnExprList();
			setState(179);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FROM) {
				{
				setState(178);
				fromClause();
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

	public static class FromClauseContext extends ParserRuleContext {
		public TerminalNode FROM() { return getToken(SqlParser.FROM, 0); }
		public JoinExprContext joinExpr() {
			return getRuleContext(JoinExprContext.class,0);
		}
		public FromClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fromClause; }
	}

	public final FromClauseContext fromClause() throws RecognitionException {
		FromClauseContext _localctx = new FromClauseContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_fromClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(181);
			match(FROM);
			setState(182);
			joinExpr();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class JoinExprContext extends ParserRuleContext {
		public TableExprContext tableExpr() {
			return getRuleContext(TableExprContext.class,0);
		}
		public JoinExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_joinExpr; }
	}

	public final JoinExprContext joinExpr() throws RecognitionException {
		JoinExprContext _localctx = new JoinExprContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_joinExpr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(184);
			tableExpr();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class UseStatementContext extends ParserRuleContext {
		public TerminalNode USE() { return getToken(SqlParser.USE, 0); }
		public DatabaseIdentifierContext databaseIdentifier() {
			return getRuleContext(DatabaseIdentifierContext.class,0);
		}
		public UseStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_useStatement; }
	}

	public final UseStatementContext useStatement() throws RecognitionException {
		UseStatementContext _localctx = new UseStatementContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_useStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(186);
			match(USE);
			setState(187);
			databaseIdentifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ColumnTypeExprContext extends ParserRuleContext {
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public TerminalNode LPAREN() { return getToken(SqlParser.LPAREN, 0); }
		public List<ColumnTypeExprContext> columnTypeExpr() {
			return getRuleContexts(ColumnTypeExprContext.class);
		}
		public ColumnTypeExprContext columnTypeExpr(int i) {
			return getRuleContext(ColumnTypeExprContext.class,i);
		}
		public TerminalNode RPAREN() { return getToken(SqlParser.RPAREN, 0); }
		public List<TerminalNode> COMMA() { return getTokens(SqlParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(SqlParser.COMMA, i);
		}
		public List<EnumValueContext> enumValue() {
			return getRuleContexts(EnumValueContext.class);
		}
		public EnumValueContext enumValue(int i) {
			return getRuleContext(EnumValueContext.class,i);
		}
		public ColumnExprListContext columnExprList() {
			return getRuleContext(ColumnExprListContext.class,0);
		}
		public ColumnTypeExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnTypeExpr; }
	}

	public final ColumnTypeExprContext columnTypeExpr() throws RecognitionException {
		ColumnTypeExprContext _localctx = new ColumnTypeExprContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_columnTypeExpr);
		int _la;
		try {
			setState(236);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(189);
				identifier();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(190);
				identifier();
				setState(191);
				match(LPAREN);
				setState(192);
				identifier();
				setState(193);
				columnTypeExpr();
				setState(200);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(194);
					match(COMMA);
					setState(195);
					identifier();
					setState(196);
					columnTypeExpr();
					}
					}
					setState(202);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(203);
				match(RPAREN);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(205);
				identifier();
				setState(206);
				match(LPAREN);
				setState(207);
				enumValue();
				setState(212);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(208);
					match(COMMA);
					setState(209);
					enumValue();
					}
					}
					setState(214);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(215);
				match(RPAREN);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(217);
				identifier();
				setState(218);
				match(LPAREN);
				setState(219);
				columnTypeExpr();
				setState(224);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(220);
					match(COMMA);
					setState(221);
					columnTypeExpr();
					}
					}
					setState(226);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(227);
				match(RPAREN);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(229);
				identifier();
				setState(230);
				match(LPAREN);
				setState(232);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << CREATE) | (1L << DATABASE) | (1L << DEFAULT) | (1L << EXISTS) | (1L << FROM) | (1L << IF) | (1L << INF) | (1L << INSERT) | (1L << INTO) | (1L << NAN_SQL) | (1L << NOT) | (1L << NULL_SQL) | (1L << SELECT) | (1L << SET) | (1L << TABLE) | (1L << VALUE) | (1L << VALUES) | (1L << SECOND) | (1L << MINUTE) | (1L << HOUR) | (1L << DAY) | (1L << WEEK) | (1L << MONTH) | (1L << QUARTER) | (1L << YEAR) | (1L << IDENTIFIER) | (1L << FLOATING_LITERAL) | (1L << OCTAL_LITERAL) | (1L << DECIMAL_LITERAL) | (1L << HEXADECIMAL_LITERAL) | (1L << STRING_LITERAL) | (1L << DASH) | (1L << DOT) | (1L << PLUS))) != 0)) {
					{
					setState(231);
					columnExprList();
					}
				}

				setState(234);
				match(RPAREN);
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

	public static class ColumnExprListContext extends ParserRuleContext {
		public List<ColumnsExprContext> columnsExpr() {
			return getRuleContexts(ColumnsExprContext.class);
		}
		public ColumnsExprContext columnsExpr(int i) {
			return getRuleContext(ColumnsExprContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(SqlParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(SqlParser.COMMA, i);
		}
		public ColumnExprListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnExprList; }
	}

	public final ColumnExprListContext columnExprList() throws RecognitionException {
		ColumnExprListContext _localctx = new ColumnExprListContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_columnExprList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(238);
			columnsExpr();
			setState(243);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(239);
				match(COMMA);
				setState(240);
				columnsExpr();
				}
				}
				setState(245);
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

	public static class ColumnsExprContext extends ParserRuleContext {
		public ColumnExprContext columnExpr() {
			return getRuleContext(ColumnExprContext.class,0);
		}
		public ColumnsExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnsExpr; }
	}

	public final ColumnsExprContext columnsExpr() throws RecognitionException {
		ColumnsExprContext _localctx = new ColumnsExprContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_columnsExpr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(246);
			columnExpr();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ColumnExprContext extends ParserRuleContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public ColumnIdentifierContext columnIdentifier() {
			return getRuleContext(ColumnIdentifierContext.class,0);
		}
		public ColumnExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnExpr; }
	}

	public final ColumnExprContext columnExpr() throws RecognitionException {
		ColumnExprContext _localctx = new ColumnExprContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_columnExpr);
		try {
			setState(250);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(248);
				literal();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(249);
				columnIdentifier();
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

	public static class ColumnIdentifierContext extends ParserRuleContext {
		public NestedIdentifierContext nestedIdentifier() {
			return getRuleContext(NestedIdentifierContext.class,0);
		}
		public TableIdentifierContext tableIdentifier() {
			return getRuleContext(TableIdentifierContext.class,0);
		}
		public TerminalNode DOT() { return getToken(SqlParser.DOT, 0); }
		public ColumnIdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnIdentifier; }
	}

	public final ColumnIdentifierContext columnIdentifier() throws RecognitionException {
		ColumnIdentifierContext _localctx = new ColumnIdentifierContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_columnIdentifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(255);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				{
				setState(252);
				tableIdentifier();
				setState(253);
				match(DOT);
				}
				break;
			}
			setState(257);
			nestedIdentifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NestedIdentifierContext extends ParserRuleContext {
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public TerminalNode DOT() { return getToken(SqlParser.DOT, 0); }
		public NestedIdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nestedIdentifier; }
	}

	public final NestedIdentifierContext nestedIdentifier() throws RecognitionException {
		NestedIdentifierContext _localctx = new NestedIdentifierContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_nestedIdentifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(259);
			identifier();
			setState(262);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOT) {
				{
				setState(260);
				match(DOT);
				setState(261);
				identifier();
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

	public static class ConstantContext extends ParserRuleContext {
		public RealLiteralContext realLiteral() {
			return getRuleContext(RealLiteralContext.class,0);
		}
		public TerminalNode DASH() { return getToken(SqlParser.DASH, 0); }
		public TerminalNode PLUS() { return getToken(SqlParser.PLUS, 0); }
		public TerminalNode DECIMAL_LITERAL() { return getToken(SqlParser.DECIMAL_LITERAL, 0); }
		public TerminalNode STRING_LITERAL() { return getToken(SqlParser.STRING_LITERAL, 0); }
		public ConstantContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constant; }
	}

	public final ConstantContext constant() throws RecognitionException {
		ConstantContext _localctx = new ConstantContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_constant);
		int _la;
		try {
			setState(273);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(265);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DASH || _la==PLUS) {
					{
					setState(264);
					_la = _input.LA(1);
					if ( !(_la==DASH || _la==PLUS) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
				}

				setState(267);
				realLiteral();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(269);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DASH || _la==PLUS) {
					{
					setState(268);
					_la = _input.LA(1);
					if ( !(_la==DASH || _la==PLUS) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
				}

				setState(271);
				match(DECIMAL_LITERAL);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(272);
				match(STRING_LITERAL);
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

	public static class RealLiteralContext extends ParserRuleContext {
		public List<TerminalNode> DECIMAL_LITERAL() { return getTokens(SqlParser.DECIMAL_LITERAL); }
		public TerminalNode DECIMAL_LITERAL(int i) {
			return getToken(SqlParser.DECIMAL_LITERAL, i);
		}
		public TerminalNode DOT() { return getToken(SqlParser.DOT, 0); }
		public TerminalNode EXPONENT_NUM_PART() { return getToken(SqlParser.EXPONENT_NUM_PART, 0); }
		public RealLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_realLiteral; }
	}

	public final RealLiteralContext realLiteral() throws RecognitionException {
		RealLiteralContext _localctx = new RealLiteralContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_realLiteral);
		int _la;
		try {
			setState(283);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DECIMAL_LITERAL:
				enterOuterAlt(_localctx, 1);
				{
				setState(275);
				match(DECIMAL_LITERAL);
				setState(276);
				match(DOT);
				setState(278);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
				case 1:
					{
					setState(277);
					_la = _input.LA(1);
					if ( !(_la==DECIMAL_LITERAL || _la==EXPONENT_NUM_PART) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					break;
				}
				}
				break;
			case DOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(280);
				match(DOT);
				setState(281);
				_la = _input.LA(1);
				if ( !(_la==DECIMAL_LITERAL || _la==EXPONENT_NUM_PART) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case EXPONENT_NUM_PART:
				enterOuterAlt(_localctx, 3);
				{
				setState(282);
				match(EXPONENT_NUM_PART);
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

	public static class TableExprContext extends ParserRuleContext {
		public TableIdentifierContext tableIdentifier() {
			return getRuleContext(TableIdentifierContext.class,0);
		}
		public TableExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tableExpr; }
	}

	public final TableExprContext tableExpr() throws RecognitionException {
		TableExprContext _localctx = new TableExprContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_tableExpr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(285);
			tableIdentifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TableIdentifierContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public DatabaseIdentifierContext databaseIdentifier() {
			return getRuleContext(DatabaseIdentifierContext.class,0);
		}
		public TerminalNode DOT() { return getToken(SqlParser.DOT, 0); }
		public TableIdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tableIdentifier; }
	}

	public final TableIdentifierContext tableIdentifier() throws RecognitionException {
		TableIdentifierContext _localctx = new TableIdentifierContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_tableIdentifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(290);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				{
				setState(287);
				databaseIdentifier();
				setState(288);
				match(DOT);
				}
				break;
			}
			setState(292);
			identifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DatabaseIdentifierContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public DatabaseIdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_databaseIdentifier; }
	}

	public final DatabaseIdentifierContext databaseIdentifier() throws RecognitionException {
		DatabaseIdentifierContext _localctx = new DatabaseIdentifierContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_databaseIdentifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(294);
			identifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FloatingLiteralContext extends ParserRuleContext {
		public TerminalNode FLOATING_LITERAL() { return getToken(SqlParser.FLOATING_LITERAL, 0); }
		public TerminalNode DOT() { return getToken(SqlParser.DOT, 0); }
		public List<TerminalNode> DECIMAL_LITERAL() { return getTokens(SqlParser.DECIMAL_LITERAL); }
		public TerminalNode DECIMAL_LITERAL(int i) {
			return getToken(SqlParser.DECIMAL_LITERAL, i);
		}
		public TerminalNode OCTAL_LITERAL() { return getToken(SqlParser.OCTAL_LITERAL, 0); }
		public FloatingLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_floatingLiteral; }
	}

	public final FloatingLiteralContext floatingLiteral() throws RecognitionException {
		FloatingLiteralContext _localctx = new FloatingLiteralContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_floatingLiteral);
		int _la;
		try {
			setState(304);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FLOATING_LITERAL:
				enterOuterAlt(_localctx, 1);
				{
				setState(296);
				match(FLOATING_LITERAL);
				}
				break;
			case DOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(297);
				match(DOT);
				setState(298);
				_la = _input.LA(1);
				if ( !(_la==OCTAL_LITERAL || _la==DECIMAL_LITERAL) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case DECIMAL_LITERAL:
				enterOuterAlt(_localctx, 3);
				{
				setState(299);
				match(DECIMAL_LITERAL);
				setState(300);
				match(DOT);
				setState(302);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==OCTAL_LITERAL || _la==DECIMAL_LITERAL) {
					{
					setState(301);
					_la = _input.LA(1);
					if ( !(_la==OCTAL_LITERAL || _la==DECIMAL_LITERAL) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
				}

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

	public static class NumberLiteralContext extends ParserRuleContext {
		public FloatingLiteralContext floatingLiteral() {
			return getRuleContext(FloatingLiteralContext.class,0);
		}
		public TerminalNode OCTAL_LITERAL() { return getToken(SqlParser.OCTAL_LITERAL, 0); }
		public TerminalNode DECIMAL_LITERAL() { return getToken(SqlParser.DECIMAL_LITERAL, 0); }
		public TerminalNode HEXADECIMAL_LITERAL() { return getToken(SqlParser.HEXADECIMAL_LITERAL, 0); }
		public TerminalNode INF() { return getToken(SqlParser.INF, 0); }
		public TerminalNode NAN_SQL() { return getToken(SqlParser.NAN_SQL, 0); }
		public TerminalNode PLUS() { return getToken(SqlParser.PLUS, 0); }
		public TerminalNode DASH() { return getToken(SqlParser.DASH, 0); }
		public NumberLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_numberLiteral; }
	}

	public final NumberLiteralContext numberLiteral() throws RecognitionException {
		NumberLiteralContext _localctx = new NumberLiteralContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_numberLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(307);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DASH || _la==PLUS) {
				{
				setState(306);
				_la = _input.LA(1);
				if ( !(_la==DASH || _la==PLUS) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			setState(315);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				{
				setState(309);
				floatingLiteral();
				}
				break;
			case 2:
				{
				setState(310);
				match(OCTAL_LITERAL);
				}
				break;
			case 3:
				{
				setState(311);
				match(DECIMAL_LITERAL);
				}
				break;
			case 4:
				{
				setState(312);
				match(HEXADECIMAL_LITERAL);
				}
				break;
			case 5:
				{
				setState(313);
				match(INF);
				}
				break;
			case 6:
				{
				setState(314);
				match(NAN_SQL);
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

	public static class LiteralContext extends ParserRuleContext {
		public NumberLiteralContext numberLiteral() {
			return getRuleContext(NumberLiteralContext.class,0);
		}
		public TerminalNode STRING_LITERAL() { return getToken(SqlParser.STRING_LITERAL, 0); }
		public TerminalNode NULL_SQL() { return getToken(SqlParser.NULL_SQL, 0); }
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_literal);
		try {
			setState(320);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INF:
			case NAN_SQL:
			case FLOATING_LITERAL:
			case OCTAL_LITERAL:
			case DECIMAL_LITERAL:
			case HEXADECIMAL_LITERAL:
			case DASH:
			case DOT:
			case PLUS:
				enterOuterAlt(_localctx, 1);
				{
				setState(317);
				numberLiteral();
				}
				break;
			case STRING_LITERAL:
				enterOuterAlt(_localctx, 2);
				{
				setState(318);
				match(STRING_LITERAL);
				}
				break;
			case NULL_SQL:
				enterOuterAlt(_localctx, 3);
				{
				setState(319);
				match(NULL_SQL);
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

	public static class IntervalContext extends ParserRuleContext {
		public TerminalNode SECOND() { return getToken(SqlParser.SECOND, 0); }
		public TerminalNode MINUTE() { return getToken(SqlParser.MINUTE, 0); }
		public TerminalNode HOUR() { return getToken(SqlParser.HOUR, 0); }
		public TerminalNode DAY() { return getToken(SqlParser.DAY, 0); }
		public TerminalNode WEEK() { return getToken(SqlParser.WEEK, 0); }
		public TerminalNode MONTH() { return getToken(SqlParser.MONTH, 0); }
		public TerminalNode QUARTER() { return getToken(SqlParser.QUARTER, 0); }
		public TerminalNode YEAR() { return getToken(SqlParser.YEAR, 0); }
		public IntervalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interval; }
	}

	public final IntervalContext interval() throws RecognitionException {
		IntervalContext _localctx = new IntervalContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_interval);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(322);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << SECOND) | (1L << MINUTE) | (1L << HOUR) | (1L << DAY) | (1L << WEEK) | (1L << MONTH) | (1L << QUARTER) | (1L << YEAR))) != 0)) ) {
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

	public static class KeywordContext extends ParserRuleContext {
		public TerminalNode CREATE() { return getToken(SqlParser.CREATE, 0); }
		public TerminalNode DATABASE() { return getToken(SqlParser.DATABASE, 0); }
		public TerminalNode DEFAULT() { return getToken(SqlParser.DEFAULT, 0); }
		public TerminalNode EXISTS() { return getToken(SqlParser.EXISTS, 0); }
		public TerminalNode FROM() { return getToken(SqlParser.FROM, 0); }
		public TerminalNode IF() { return getToken(SqlParser.IF, 0); }
		public TerminalNode INSERT() { return getToken(SqlParser.INSERT, 0); }
		public TerminalNode INTO() { return getToken(SqlParser.INTO, 0); }
		public TerminalNode NOT() { return getToken(SqlParser.NOT, 0); }
		public TerminalNode NULL_SQL() { return getToken(SqlParser.NULL_SQL, 0); }
		public TerminalNode SELECT() { return getToken(SqlParser.SELECT, 0); }
		public TerminalNode SET() { return getToken(SqlParser.SET, 0); }
		public TerminalNode TABLE() { return getToken(SqlParser.TABLE, 0); }
		public TerminalNode VALUE() { return getToken(SqlParser.VALUE, 0); }
		public TerminalNode VALUES() { return getToken(SqlParser.VALUES, 0); }
		public KeywordContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyword; }
	}

	public final KeywordContext keyword() throws RecognitionException {
		KeywordContext _localctx = new KeywordContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_keyword);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(324);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << CREATE) | (1L << DATABASE) | (1L << DEFAULT) | (1L << EXISTS) | (1L << FROM) | (1L << IF) | (1L << INSERT) | (1L << INTO) | (1L << NOT) | (1L << NULL_SQL) | (1L << SELECT) | (1L << SET) | (1L << TABLE) | (1L << VALUE) | (1L << VALUES))) != 0)) ) {
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

	public static class IdentifierContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(SqlParser.IDENTIFIER, 0); }
		public IntervalContext interval() {
			return getRuleContext(IntervalContext.class,0);
		}
		public KeywordContext keyword() {
			return getRuleContext(KeywordContext.class,0);
		}
		public IdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifier; }
	}

	public final IdentifierContext identifier() throws RecognitionException {
		IdentifierContext _localctx = new IdentifierContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_identifier);
		try {
			setState(329);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(326);
				match(IDENTIFIER);
				}
				break;
			case SECOND:
			case MINUTE:
			case HOUR:
			case DAY:
			case WEEK:
			case MONTH:
			case QUARTER:
			case YEAR:
				enterOuterAlt(_localctx, 2);
				{
				setState(327);
				interval();
				}
				break;
			case CREATE:
			case DATABASE:
			case DEFAULT:
			case EXISTS:
			case FROM:
			case IF:
			case INSERT:
			case INTO:
			case NOT:
			case NULL_SQL:
			case SELECT:
			case SET:
			case TABLE:
			case VALUE:
			case VALUES:
				enterOuterAlt(_localctx, 3);
				{
				setState(328);
				keyword();
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

	public static class EnumValueContext extends ParserRuleContext {
		public TerminalNode STRING_LITERAL() { return getToken(SqlParser.STRING_LITERAL, 0); }
		public TerminalNode EQ_SINGLE() { return getToken(SqlParser.EQ_SINGLE, 0); }
		public NumberLiteralContext numberLiteral() {
			return getRuleContext(NumberLiteralContext.class,0);
		}
		public EnumValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumValue; }
	}

	public final EnumValueContext enumValue() throws RecognitionException {
		EnumValueContext _localctx = new EnumValueContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_enumValue);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(331);
			match(STRING_LITERAL);
			setState(332);
			match(EQ_SINGLE);
			setState(333);
			numberLiteral();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IfNotExistsContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(SqlParser.IF, 0); }
		public TerminalNode NOT() { return getToken(SqlParser.NOT, 0); }
		public TerminalNode EXISTS() { return getToken(SqlParser.EXISTS, 0); }
		public IfNotExistsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifNotExists; }
	}

	public final IfNotExistsContext ifNotExists() throws RecognitionException {
		IfNotExistsContext _localctx = new IfNotExistsContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_ifNotExists);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(335);
			match(IF);
			setState(336);
			match(NOT);
			setState(337);
			match(EXISTS);
			}
		}
		catch (RecognitionException re) {
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3E\u0156\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\3\2\3\2\3\2\5\2P\n\2\3\3\3\3\5\3T"+
		"\n\3\3\4\3\4\5\4X\n\4\3\5\3\5\3\6\3\6\3\6\5\6_\n\6\3\6\3\6\3\7\3\7\3\7"+
		"\5\7f\n\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\7\bo\n\b\f\b\16\br\13\b\3\b\3\b"+
		"\3\t\3\t\3\n\3\n\3\n\3\13\3\13\3\13\5\13~\n\13\3\13\3\13\5\13\u0082\n"+
		"\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\7\f\u008b\n\f\f\f\16\f\u008e\13\f\3"+
		"\f\3\f\3\r\5\r\u0093\n\r\3\r\7\r\u0096\n\r\f\r\16\r\u0099\13\r\3\16\3"+
		"\16\3\16\5\16\u009e\n\16\6\16\u00a0\n\16\r\16\16\16\u00a1\3\16\3\16\3"+
		"\17\3\17\3\17\3\17\3\17\6\17\u00ab\n\17\r\17\16\17\u00ac\3\17\3\17\5\17"+
		"\u00b1\n\17\3\20\3\20\3\20\5\20\u00b6\n\20\3\21\3\21\3\21\3\22\3\22\3"+
		"\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\7\24\u00c9"+
		"\n\24\f\24\16\24\u00cc\13\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\7\24\u00d5"+
		"\n\24\f\24\16\24\u00d8\13\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\7\24\u00e1"+
		"\n\24\f\24\16\24\u00e4\13\24\3\24\3\24\3\24\3\24\3\24\5\24\u00eb\n\24"+
		"\3\24\3\24\5\24\u00ef\n\24\3\25\3\25\3\25\7\25\u00f4\n\25\f\25\16\25\u00f7"+
		"\13\25\3\26\3\26\3\27\3\27\5\27\u00fd\n\27\3\30\3\30\3\30\5\30\u0102\n"+
		"\30\3\30\3\30\3\31\3\31\3\31\5\31\u0109\n\31\3\32\5\32\u010c\n\32\3\32"+
		"\3\32\5\32\u0110\n\32\3\32\3\32\5\32\u0114\n\32\3\33\3\33\3\33\5\33\u0119"+
		"\n\33\3\33\3\33\3\33\5\33\u011e\n\33\3\34\3\34\3\35\3\35\3\35\5\35\u0125"+
		"\n\35\3\35\3\35\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3\37\5\37\u0131\n\37"+
		"\5\37\u0133\n\37\3 \5 \u0136\n \3 \3 \3 \3 \3 \3 \5 \u013e\n \3!\3!\3"+
		"!\5!\u0143\n!\3\"\3\"\3#\3#\3$\3$\3$\5$\u014c\n$\3%\3%\3%\3%\3&\3&\3&"+
		"\3&\3&\2\2\'\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64"+
		"\668:<>@BDFHJ\2\7\4\2,,99\4\2!!##\3\2 !\3\2\26\35\6\2\4\t\13\f\16\22\24"+
		"\25\2\u0162\2O\3\2\2\2\4S\3\2\2\2\6W\3\2\2\2\bY\3\2\2\2\n[\3\2\2\2\fb"+
		"\3\2\2\2\16j\3\2\2\2\20u\3\2\2\2\22w\3\2\2\2\24z\3\2\2\2\26\u0086\3\2"+
		"\2\2\30\u0097\3\2\2\2\32\u009a\3\2\2\2\34\u00b0\3\2\2\2\36\u00b2\3\2\2"+
		"\2 \u00b7\3\2\2\2\"\u00ba\3\2\2\2$\u00bc\3\2\2\2&\u00ee\3\2\2\2(\u00f0"+
		"\3\2\2\2*\u00f8\3\2\2\2,\u00fc\3\2\2\2.\u0101\3\2\2\2\60\u0105\3\2\2\2"+
		"\62\u0113\3\2\2\2\64\u011d\3\2\2\2\66\u011f\3\2\2\28\u0124\3\2\2\2:\u0128"+
		"\3\2\2\2<\u0132\3\2\2\2>\u0135\3\2\2\2@\u0142\3\2\2\2B\u0144\3\2\2\2D"+
		"\u0146\3\2\2\2F\u014b\3\2\2\2H\u014d\3\2\2\2J\u0151\3\2\2\2LP\5\6\4\2"+
		"MP\5\4\3\2NP\5\b\5\2OL\3\2\2\2OM\3\2\2\2ON\3\2\2\2P\3\3\2\2\2QT\5\36\20"+
		"\2RT\5\24\13\2SQ\3\2\2\2SR\3\2\2\2T\5\3\2\2\2UX\5\n\6\2VX\5\f\7\2WU\3"+
		"\2\2\2WV\3\2\2\2X\7\3\2\2\2YZ\5$\23\2Z\t\3\2\2\2[\\\7\4\2\2\\^\7\5\2\2"+
		"]_\5J&\2^]\3\2\2\2^_\3\2\2\2_`\3\2\2\2`a\5:\36\2a\13\3\2\2\2bc\7\4\2\2"+
		"ce\7\22\2\2df\5J&\2ed\3\2\2\2ef\3\2\2\2fg\3\2\2\2gh\58\35\2hi\5\16\b\2"+
		"i\r\3\2\2\2jk\7\65\2\2kp\5\20\t\2lm\7*\2\2mo\5\20\t\2nl\3\2\2\2or\3\2"+
		"\2\2pn\3\2\2\2pq\3\2\2\2qs\3\2\2\2rp\3\2\2\2st\7?\2\2t\17\3\2\2\2uv\5"+
		"\22\n\2v\21\3\2\2\2wx\5\60\31\2xy\5&\24\2y\23\3\2\2\2z{\7\13\2\2{}\7\f"+
		"\2\2|~\7\22\2\2}|\3\2\2\2}~\3\2\2\2~\177\3\2\2\2\177\u0081\58\35\2\u0080"+
		"\u0082\5\26\f\2\u0081\u0080\3\2\2\2\u0081\u0082\3\2\2\2\u0082\u0083\3"+
		"\2\2\2\u0083\u0084\7\25\2\2\u0084\u0085\5\30\r\2\u0085\25\3\2\2\2\u0086"+
		"\u0087\7\65\2\2\u0087\u008c\5\60\31\2\u0088\u0089\7*\2\2\u0089\u008b\5"+
		"\60\31\2\u008a\u0088\3\2\2\2\u008b\u008e\3\2\2\2\u008c\u008a\3\2\2\2\u008c"+
		"\u008d\3\2\2\2\u008d\u008f\3\2\2\2\u008e\u008c\3\2\2\2\u008f\u0090\7?"+
		"\2\2\u0090\27\3\2\2\2\u0091\u0093\7*\2\2\u0092\u0091\3\2\2\2\u0092\u0093"+
		"\3\2\2\2\u0093\u0094\3\2\2\2\u0094\u0096\5\32\16\2\u0095\u0092\3\2\2\2"+
		"\u0096\u0099\3\2\2\2\u0097\u0095\3\2\2\2\u0097\u0098\3\2\2\2\u0098\31"+
		"\3\2\2\2\u0099\u0097\3\2\2\2\u009a\u009f\7\65\2\2\u009b\u009d\5\34\17"+
		"\2\u009c\u009e\7*\2\2\u009d\u009c\3\2\2\2\u009d\u009e\3\2\2\2\u009e\u00a0"+
		"\3\2\2\2\u009f\u009b\3\2\2\2\u00a0\u00a1\3\2\2\2\u00a1\u009f\3\2\2\2\u00a1"+
		"\u00a2\3\2\2\2\u00a2\u00a3\3\2\2\2\u00a3\u00a4\7?\2\2\u00a4\33\3\2\2\2"+
		"\u00a5\u00b1\5\62\32\2\u00a6\u00a7\7\65\2\2\u00a7\u00aa\5\62\32\2\u00a8"+
		"\u00a9\7*\2\2\u00a9\u00ab\5\62\32\2\u00aa\u00a8\3\2\2\2\u00ab\u00ac\3"+
		"\2\2\2\u00ac\u00aa\3\2\2\2\u00ac\u00ad\3\2\2\2\u00ad\u00ae\3\2\2\2\u00ae"+
		"\u00af\7?\2\2\u00af\u00b1\3\2\2\2\u00b0\u00a5\3\2\2\2\u00b0\u00a6\3\2"+
		"\2\2\u00b1\35\3\2\2\2\u00b2\u00b3\7\20\2\2\u00b3\u00b5\5(\25\2\u00b4\u00b6"+
		"\5 \21\2\u00b5\u00b4\3\2\2\2\u00b5\u00b6\3\2\2\2\u00b6\37\3\2\2\2\u00b7"+
		"\u00b8\7\b\2\2\u00b8\u00b9\5\"\22\2\u00b9!\3\2\2\2\u00ba\u00bb\5\66\34"+
		"\2\u00bb#\3\2\2\2\u00bc\u00bd\7\23\2\2\u00bd\u00be\5:\36\2\u00be%\3\2"+
		"\2\2\u00bf\u00ef\5F$\2\u00c0\u00c1\5F$\2\u00c1\u00c2\7\65\2\2\u00c2\u00c3"+
		"\5F$\2\u00c3\u00ca\5&\24\2\u00c4\u00c5\7*\2\2\u00c5\u00c6\5F$\2\u00c6"+
		"\u00c7\5&\24\2\u00c7\u00c9\3\2\2\2\u00c8\u00c4\3\2\2\2\u00c9\u00cc\3\2"+
		"\2\2\u00ca\u00c8\3\2\2\2\u00ca\u00cb\3\2\2\2\u00cb\u00cd\3\2\2\2\u00cc"+
		"\u00ca\3\2\2\2\u00cd\u00ce\7?\2\2\u00ce\u00ef\3\2\2\2\u00cf\u00d0\5F$"+
		"\2\u00d0\u00d1\7\65\2\2\u00d1\u00d6\5H%\2\u00d2\u00d3\7*\2\2\u00d3\u00d5"+
		"\5H%\2\u00d4\u00d2\3\2\2\2\u00d5\u00d8\3\2\2\2\u00d6\u00d4\3\2\2\2\u00d6"+
		"\u00d7\3\2\2\2\u00d7\u00d9\3\2\2\2\u00d8\u00d6\3\2\2\2\u00d9\u00da\7?"+
		"\2\2\u00da\u00ef\3\2\2\2\u00db\u00dc\5F$\2\u00dc\u00dd\7\65\2\2\u00dd"+
		"\u00e2\5&\24\2\u00de\u00df\7*\2\2\u00df\u00e1\5&\24\2\u00e0\u00de\3\2"+
		"\2\2\u00e1\u00e4\3\2\2\2\u00e2\u00e0\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3"+
		"\u00e5\3\2\2\2\u00e4\u00e2\3\2\2\2\u00e5\u00e6\7?\2\2\u00e6\u00ef\3\2"+
		"\2\2\u00e7\u00e8\5F$\2\u00e8\u00ea\7\65\2\2\u00e9\u00eb\5(\25\2\u00ea"+
		"\u00e9\3\2\2\2\u00ea\u00eb\3\2\2\2\u00eb\u00ec\3\2\2\2\u00ec\u00ed\7?"+
		"\2\2\u00ed\u00ef\3\2\2\2\u00ee\u00bf\3\2\2\2\u00ee\u00c0\3\2\2\2\u00ee"+
		"\u00cf\3\2\2\2\u00ee\u00db\3\2\2\2\u00ee\u00e7\3\2\2\2\u00ef\'\3\2\2\2"+
		"\u00f0\u00f5\5*\26\2\u00f1\u00f2\7*\2\2\u00f2\u00f4\5*\26\2\u00f3\u00f1"+
		"\3\2\2\2\u00f4\u00f7\3\2\2\2\u00f5\u00f3\3\2\2\2\u00f5\u00f6\3\2\2\2\u00f6"+
		")\3\2\2\2\u00f7\u00f5\3\2\2\2\u00f8\u00f9\5,\27\2\u00f9+\3\2\2\2\u00fa"+
		"\u00fd\5@!\2\u00fb\u00fd\5.\30\2\u00fc\u00fa\3\2\2\2\u00fc\u00fb\3\2\2"+
		"\2\u00fd-\3\2\2\2\u00fe\u00ff\58\35\2\u00ff\u0100\7-\2\2\u0100\u0102\3"+
		"\2\2\2\u0101\u00fe\3\2\2\2\u0101\u0102\3\2\2\2\u0102\u0103\3\2\2\2\u0103"+
		"\u0104\5\60\31\2\u0104/\3\2\2\2\u0105\u0108\5F$\2\u0106\u0107\7-\2\2\u0107"+
		"\u0109\5F$\2\u0108\u0106\3\2\2\2\u0108\u0109\3\2\2\2\u0109\61\3\2\2\2"+
		"\u010a\u010c\t\2\2\2\u010b\u010a\3\2\2\2\u010b\u010c\3\2\2\2\u010c\u010d"+
		"\3\2\2\2\u010d\u0114\5\64\33\2\u010e\u0110\t\2\2\2\u010f\u010e\3\2\2\2"+
		"\u010f\u0110\3\2\2\2\u0110\u0111\3\2\2\2\u0111\u0114\7!\2\2\u0112\u0114"+
		"\7$\2\2\u0113\u010b\3\2\2\2\u0113\u010f\3\2\2\2\u0113\u0112\3\2\2\2\u0114"+
		"\63\3\2\2\2\u0115\u0116\7!\2\2\u0116\u0118\7-\2\2\u0117\u0119\t\3\2\2"+
		"\u0118\u0117\3\2\2\2\u0118\u0119\3\2\2\2\u0119\u011e\3\2\2\2\u011a\u011b"+
		"\7-\2\2\u011b\u011e\t\3\2\2\u011c\u011e\7#\2\2\u011d\u0115\3\2\2\2\u011d"+
		"\u011a\3\2\2\2\u011d\u011c\3\2\2\2\u011e\65\3\2\2\2\u011f\u0120\58\35"+
		"\2\u0120\67\3\2\2\2\u0121\u0122\5:\36\2\u0122\u0123\7-\2\2\u0123\u0125"+
		"\3\2\2\2\u0124\u0121\3\2\2\2\u0124\u0125\3\2\2\2\u0125\u0126\3\2\2\2\u0126"+
		"\u0127\5F$\2\u01279\3\2\2\2\u0128\u0129\5F$\2\u0129;\3\2\2\2\u012a\u0133"+
		"\7\37\2\2\u012b\u012c\7-\2\2\u012c\u0133\t\4\2\2\u012d\u012e\7!\2\2\u012e"+
		"\u0130\7-\2\2\u012f\u0131\t\4\2\2\u0130\u012f\3\2\2\2\u0130\u0131\3\2"+
		"\2\2\u0131\u0133\3\2\2\2\u0132\u012a\3\2\2\2\u0132\u012b\3\2\2\2\u0132"+
		"\u012d\3\2\2\2\u0133=\3\2\2\2\u0134\u0136\t\2\2\2\u0135\u0134\3\2\2\2"+
		"\u0135\u0136\3\2\2\2\u0136\u013d\3\2\2\2\u0137\u013e\5<\37\2\u0138\u013e"+
		"\7 \2\2\u0139\u013e\7!\2\2\u013a\u013e\7\"\2\2\u013b\u013e\7\n\2\2\u013c"+
		"\u013e\7\r\2\2\u013d\u0137\3\2\2\2\u013d\u0138\3\2\2\2\u013d\u0139\3\2"+
		"\2\2\u013d\u013a\3\2\2\2\u013d\u013b\3\2\2\2\u013d\u013c\3\2\2\2\u013e"+
		"?\3\2\2\2\u013f\u0143\5> \2\u0140\u0143\7$\2\2\u0141\u0143\7\17\2\2\u0142"+
		"\u013f\3\2\2\2\u0142\u0140\3\2\2\2\u0142\u0141\3\2\2\2\u0143A\3\2\2\2"+
		"\u0144\u0145\t\5\2\2\u0145C\3\2\2\2\u0146\u0147\t\6\2\2\u0147E\3\2\2\2"+
		"\u0148\u014c\7\36\2\2\u0149\u014c\5B\"\2\u014a\u014c\5D#\2\u014b\u0148"+
		"\3\2\2\2\u014b\u0149\3\2\2\2\u014b\u014a\3\2\2\2\u014cG\3\2\2\2\u014d"+
		"\u014e\7$\2\2\u014e\u014f\7/\2\2\u014f\u0150\5> \2\u0150I\3\2\2\2\u0151"+
		"\u0152\7\t\2\2\u0152\u0153\7\16\2\2\u0153\u0154\7\7\2\2\u0154K\3\2\2\2"+
		"\'OSW^ep}\u0081\u008c\u0092\u0097\u009d\u00a1\u00ac\u00b0\u00b5\u00ca"+
		"\u00d6\u00e2\u00ea\u00ee\u00f5\u00fc\u0101\u0108\u010b\u010f\u0113\u0118"+
		"\u011d\u0124\u0130\u0132\u0135\u013d\u0142\u014b";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}