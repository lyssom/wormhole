// Generated from SqlParser.g4 by ANTLR 4.7.2
package parser;
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
	static { RuntimeMetaData.checkVersion("4.7.2", RuntimeMetaData.VERSION); }

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
		SINGLE_LINE_COMMENT=66, WHITESPACE=67, WHERE=68, AND=69, OR=70;
	public static final int
		RULE_statement = 0, RULE_dmlStatement = 1, RULE_ddlStatement = 2, RULE_systemStatement = 3, 
		RULE_createDatabase = 4, RULE_createTable = 5, RULE_tableSchemaClause = 6, 
		RULE_tableElementExpr = 7, RULE_tableColumnDfnt = 8, RULE_insertStatement = 9, 
		RULE_columnsClause = 10, RULE_insertValuesSpec = 11, RULE_insertMultiValue = 12, 
		RULE_dataValue = 13, RULE_selectStatement = 14, RULE_fromClause = 15, 
		RULE_whereClause = 16, RULE_searchCondition = 17, RULE_predicate = 18, 
		RULE_expression = 19, RULE_joinExpr = 20, RULE_useStatement = 21, RULE_columnTypeExpr = 22, 
		RULE_columnExprList = 23, RULE_columnsExpr = 24, RULE_columnExpr = 25, 
		RULE_columnIdentifier = 26, RULE_nestedIdentifier = 27, RULE_constant = 28, 
		RULE_realLiteral = 29, RULE_tableExpr = 30, RULE_tableIdentifier = 31, 
		RULE_databaseIdentifier = 32, RULE_floatingLiteral = 33, RULE_numberLiteral = 34, 
		RULE_literal = 35, RULE_interval = 36, RULE_keyword = 37, RULE_identifier = 38, 
		RULE_enumValue = 39, RULE_ifNotExists = 40;
	private static String[] makeRuleNames() {
		return new String[] {
			"statement", "dmlStatement", "ddlStatement", "systemStatement", "createDatabase", 
			"createTable", "tableSchemaClause", "tableElementExpr", "tableColumnDfnt", 
			"insertStatement", "columnsClause", "insertValuesSpec", "insertMultiValue", 
			"dataValue", "selectStatement", "fromClause", "whereClause", "searchCondition", 
			"predicate", "expression", "joinExpr", "useStatement", "columnTypeExpr", 
			"columnExprList", "columnsExpr", "columnExpr", "columnIdentifier", "nestedIdentifier", 
			"constant", "realLiteral", "tableExpr", "tableIdentifier", "databaseIdentifier", 
			"floatingLiteral", "numberLiteral", "literal", "interval", "keyword", 
			"identifier", "enumValue", "ifNotExists"
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
			"WHITESPACE", "WHERE", "AND", "OR"
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitStatement(this);
		}
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_statement);
		try {
			setState(85);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CREATE:
				enterOuterAlt(_localctx, 1);
				{
				setState(82);
				ddlStatement();
				}
				break;
			case INSERT:
			case SELECT:
				enterOuterAlt(_localctx, 2);
				{
				setState(83);
				dmlStatement();
				}
				break;
			case USE:
				enterOuterAlt(_localctx, 3);
				{
				setState(84);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterDmlStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitDmlStatement(this);
		}
	}

	public final DmlStatementContext dmlStatement() throws RecognitionException {
		DmlStatementContext _localctx = new DmlStatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_dmlStatement);
		try {
			setState(89);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SELECT:
				enterOuterAlt(_localctx, 1);
				{
				setState(87);
				selectStatement();
				}
				break;
			case INSERT:
				enterOuterAlt(_localctx, 2);
				{
				setState(88);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterDdlStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitDdlStatement(this);
		}
	}

	public final DdlStatementContext ddlStatement() throws RecognitionException {
		DdlStatementContext _localctx = new DdlStatementContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_ddlStatement);
		try {
			setState(93);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(91);
				createDatabase();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(92);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterSystemStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitSystemStatement(this);
		}
	}

	public final SystemStatementContext systemStatement() throws RecognitionException {
		SystemStatementContext _localctx = new SystemStatementContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_systemStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(95);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterCreateDatabase(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitCreateDatabase(this);
		}
	}

	public final CreateDatabaseContext createDatabase() throws RecognitionException {
		CreateDatabaseContext _localctx = new CreateDatabaseContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_createDatabase);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(97);
			match(CREATE);
			setState(98);
			match(DATABASE);
			setState(100);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				{
				setState(99);
				ifNotExists();
				}
				break;
			}
			setState(102);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterCreateTable(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitCreateTable(this);
		}
	}

	public final CreateTableContext createTable() throws RecognitionException {
		CreateTableContext _localctx = new CreateTableContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_createTable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(104);
			match(CREATE);
			setState(105);
			match(TABLE);
			setState(107);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				{
				setState(106);
				ifNotExists();
				}
				break;
			}
			setState(109);
			tableIdentifier();
			setState(110);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterTableSchemaClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitTableSchemaClause(this);
		}
	}

	public final TableSchemaClauseContext tableSchemaClause() throws RecognitionException {
		TableSchemaClauseContext _localctx = new TableSchemaClauseContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_tableSchemaClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(112);
			match(LPAREN);
			setState(113);
			tableElementExpr();
			setState(118);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(114);
				match(COMMA);
				setState(115);
				tableElementExpr();
				}
				}
				setState(120);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(121);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterTableElementExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitTableElementExpr(this);
		}
	}

	public final TableElementExprContext tableElementExpr() throws RecognitionException {
		TableElementExprContext _localctx = new TableElementExprContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_tableElementExpr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(123);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterTableColumnDfnt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitTableColumnDfnt(this);
		}
	}

	public final TableColumnDfntContext tableColumnDfnt() throws RecognitionException {
		TableColumnDfntContext _localctx = new TableColumnDfntContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_tableColumnDfnt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(125);
			nestedIdentifier();
			setState(126);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterInsertStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitInsertStatement(this);
		}
	}

	public final InsertStatementContext insertStatement() throws RecognitionException {
		InsertStatementContext _localctx = new InsertStatementContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_insertStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			match(INSERT);
			setState(129);
			match(INTO);
			setState(131);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				{
				setState(130);
				match(TABLE);
				}
				break;
			}
			{
			setState(133);
			tableIdentifier();
			}
			setState(135);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LPAREN) {
				{
				setState(134);
				columnsClause();
				}
			}

			setState(137);
			match(VALUES);
			setState(138);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterColumnsClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitColumnsClause(this);
		}
	}

	public final ColumnsClauseContext columnsClause() throws RecognitionException {
		ColumnsClauseContext _localctx = new ColumnsClauseContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_columnsClause);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(140);
			match(LPAREN);
			setState(141);
			nestedIdentifier();
			setState(146);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(142);
				match(COMMA);
				setState(143);
				nestedIdentifier();
				}
				}
				setState(148);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(149);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterInsertValuesSpec(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitInsertValuesSpec(this);
		}
	}

	public final InsertValuesSpecContext insertValuesSpec() throws RecognitionException {
		InsertValuesSpecContext _localctx = new InsertValuesSpecContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_insertValuesSpec);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(157);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA || _la==LPAREN) {
				{
				{
				setState(152);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COMMA) {
					{
					setState(151);
					match(COMMA);
					}
				}

				setState(154);
				insertMultiValue();
				}
				}
				setState(159);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterInsertMultiValue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitInsertMultiValue(this);
		}
	}

	public final InsertMultiValueContext insertMultiValue() throws RecognitionException {
		InsertMultiValueContext _localctx = new InsertMultiValueContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_insertMultiValue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(160);
			match(LPAREN);
			setState(165); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(161);
				dataValue();
				setState(163);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COMMA) {
					{
					setState(162);
					match(COMMA);
					}
				}

				}
				}
				setState(167); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << DECIMAL_LITERAL) | (1L << EXPONENT_NUM_PART) | (1L << STRING_LITERAL) | (1L << DASH) | (1L << DOT) | (1L << LPAREN) | (1L << PLUS))) != 0) );
			setState(169);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterDataValue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitDataValue(this);
		}
	}

	public final DataValueContext dataValue() throws RecognitionException {
		DataValueContext _localctx = new DataValueContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_dataValue);
		int _la;
		try {
			setState(182);
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
				setState(171);
				constant();
				}
				break;
			case LPAREN:
				enterOuterAlt(_localctx, 2);
				{
				setState(172);
				match(LPAREN);
				setState(173);
				constant();
				setState(176); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(174);
					match(COMMA);
					setState(175);
					constant();
					}
					}
					setState(178); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==COMMA );
				setState(180);
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
		public WhereClauseContext whereClause() {
			return getRuleContext(WhereClauseContext.class,0);
		}
		public SelectStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selectStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterSelectStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitSelectStatement(this);
		}
	}

	public final SelectStatementContext selectStatement() throws RecognitionException {
		SelectStatementContext _localctx = new SelectStatementContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_selectStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(184);
			match(SELECT);
			setState(185);
			columnExprList();
			setState(187);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==FROM) {
				{
				setState(186);
				fromClause();
				}
			}

			setState(190);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHERE) {
				{
				setState(189);
				whereClause();
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterFromClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitFromClause(this);
		}
	}

	public final FromClauseContext fromClause() throws RecognitionException {
		FromClauseContext _localctx = new FromClauseContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_fromClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(192);
			match(FROM);
			setState(193);
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

	public static class WhereClauseContext extends ParserRuleContext {
		public TerminalNode WHERE() { return getToken(SqlParser.WHERE, 0); }
		public SearchConditionContext searchCondition() {
			return getRuleContext(SearchConditionContext.class,0);
		}
		public WhereClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whereClause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterWhereClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitWhereClause(this);
		}
	}

	public final WhereClauseContext whereClause() throws RecognitionException {
		WhereClauseContext _localctx = new WhereClauseContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_whereClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(195);
			match(WHERE);
			setState(196);
			searchCondition(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SearchConditionContext extends ParserRuleContext {
		public TerminalNode NOT() { return getToken(SqlParser.NOT, 0); }
		public List<SearchConditionContext> searchCondition() {
			return getRuleContexts(SearchConditionContext.class);
		}
		public SearchConditionContext searchCondition(int i) {
			return getRuleContext(SearchConditionContext.class,i);
		}
		public PredicateContext predicate() {
			return getRuleContext(PredicateContext.class,0);
		}
		public TerminalNode AND() { return getToken(SqlParser.AND, 0); }
		public TerminalNode OR() { return getToken(SqlParser.OR, 0); }
		public SearchConditionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_searchCondition; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterSearchCondition(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitSearchCondition(this);
		}
	}

	public final SearchConditionContext searchCondition() throws RecognitionException {
		return searchCondition(0);
	}

	private SearchConditionContext searchCondition(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		SearchConditionContext _localctx = new SearchConditionContext(_ctx, _parentState);
		SearchConditionContext _prevctx = _localctx;
		int _startState = 34;
		enterRecursionRule(_localctx, 34, RULE_searchCondition, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(202);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				{
				setState(199);
				match(NOT);
				setState(200);
				searchCondition(2);
				}
				break;
			case 2:
				{
				setState(201);
				predicate();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(212);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,19,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(210);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
					case 1:
						{
						_localctx = new SearchConditionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_searchCondition);
						setState(204);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(205);
						match(AND);
						setState(206);
						searchCondition(5);
						}
						break;
					case 2:
						{
						_localctx = new SearchConditionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_searchCondition);
						setState(207);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(208);
						match(OR);
						setState(209);
						searchCondition(4);
						}
						break;
					}
					} 
				}
				setState(214);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,19,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class PredicateContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode EQ_SINGLE() { return getToken(SqlParser.EQ_SINGLE, 0); }
		public TerminalNode NOT_EQ() { return getToken(SqlParser.NOT_EQ, 0); }
		public TerminalNode LT() { return getToken(SqlParser.LT, 0); }
		public TerminalNode LE() { return getToken(SqlParser.LE, 0); }
		public TerminalNode GT() { return getToken(SqlParser.GT, 0); }
		public TerminalNode GE() { return getToken(SqlParser.GE, 0); }
		public PredicateContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_predicate; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterPredicate(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitPredicate(this);
		}
	}

	public final PredicateContext predicate() throws RecognitionException {
		PredicateContext _localctx = new PredicateContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_predicate);
		try {
			setState(239);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(215);
				expression(0);
				setState(216);
				match(EQ_SINGLE);
				setState(217);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(219);
				expression(0);
				setState(220);
				match(NOT_EQ);
				setState(221);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(223);
				expression(0);
				setState(224);
				match(LT);
				setState(225);
				expression(0);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(227);
				expression(0);
				setState(228);
				match(LE);
				setState(229);
				expression(0);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(231);
				expression(0);
				setState(232);
				match(GT);
				setState(233);
				expression(0);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(235);
				expression(0);
				setState(236);
				match(GE);
				setState(237);
				expression(0);
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

	public static class ExpressionContext extends ParserRuleContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public ColumnIdentifierContext columnIdentifier() {
			return getRuleContext(ColumnIdentifierContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode ASTERISK() { return getToken(SqlParser.ASTERISK, 0); }
		public TerminalNode SLASH() { return getToken(SqlParser.SLASH, 0); }
		public TerminalNode PERCENT() { return getToken(SqlParser.PERCENT, 0); }
		public TerminalNode PLUS() { return getToken(SqlParser.PLUS, 0); }
		public TerminalNode DASH() { return getToken(SqlParser.DASH, 0); }
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitExpression(this);
		}
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 38;
		enterRecursionRule(_localctx, 38, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(244);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				{
				setState(242);
				literal();
				}
				break;
			case 2:
				{
				setState(243);
				columnIdentifier();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(263);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(261);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(246);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(247);
						match(ASTERISK);
						setState(248);
						expression(6);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(249);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(250);
						match(SLASH);
						setState(251);
						expression(5);
						}
						break;
					case 3:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(252);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(253);
						match(PERCENT);
						setState(254);
						expression(4);
						}
						break;
					case 4:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(255);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(256);
						match(PLUS);
						setState(257);
						expression(3);
						}
						break;
					case 5:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(258);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(259);
						match(DASH);
						setState(260);
						expression(2);
						}
						break;
					}
					} 
				}
				setState(265);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterJoinExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitJoinExpr(this);
		}
	}

	public final JoinExprContext joinExpr() throws RecognitionException {
		JoinExprContext _localctx = new JoinExprContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_joinExpr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(266);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterUseStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitUseStatement(this);
		}
	}

	public final UseStatementContext useStatement() throws RecognitionException {
		UseStatementContext _localctx = new UseStatementContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_useStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(268);
			match(USE);
			setState(269);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterColumnTypeExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitColumnTypeExpr(this);
		}
	}

	public final ColumnTypeExprContext columnTypeExpr() throws RecognitionException {
		ColumnTypeExprContext _localctx = new ColumnTypeExprContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_columnTypeExpr);
		int _la;
		try {
			setState(318);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(271);
				identifier();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(272);
				identifier();
				setState(273);
				match(LPAREN);
				setState(274);
				identifier();
				setState(275);
				columnTypeExpr();
				setState(282);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(276);
					match(COMMA);
					setState(277);
					identifier();
					setState(278);
					columnTypeExpr();
					}
					}
					setState(284);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(285);
				match(RPAREN);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(287);
				identifier();
				setState(288);
				match(LPAREN);
				setState(289);
				enumValue();
				setState(294);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(290);
					match(COMMA);
					setState(291);
					enumValue();
					}
					}
					setState(296);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(297);
				match(RPAREN);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(299);
				identifier();
				setState(300);
				match(LPAREN);
				setState(301);
				columnTypeExpr();
				setState(306);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(302);
					match(COMMA);
					setState(303);
					columnTypeExpr();
					}
					}
					setState(308);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(309);
				match(RPAREN);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(311);
				identifier();
				setState(312);
				match(LPAREN);
				setState(314);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << CREATE) | (1L << DATABASE) | (1L << DEFAULT) | (1L << EXISTS) | (1L << FROM) | (1L << IF) | (1L << INF) | (1L << INSERT) | (1L << INTO) | (1L << NAN_SQL) | (1L << NOT) | (1L << NULL_SQL) | (1L << SELECT) | (1L << SET) | (1L << TABLE) | (1L << VALUE) | (1L << VALUES) | (1L << SECOND) | (1L << MINUTE) | (1L << HOUR) | (1L << DAY) | (1L << WEEK) | (1L << MONTH) | (1L << QUARTER) | (1L << YEAR) | (1L << IDENTIFIER) | (1L << FLOATING_LITERAL) | (1L << OCTAL_LITERAL) | (1L << DECIMAL_LITERAL) | (1L << HEXADECIMAL_LITERAL) | (1L << STRING_LITERAL) | (1L << DASH) | (1L << DOT) | (1L << PLUS))) != 0)) {
					{
					setState(313);
					columnExprList();
					}
				}

				setState(316);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterColumnExprList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitColumnExprList(this);
		}
	}

	public final ColumnExprListContext columnExprList() throws RecognitionException {
		ColumnExprListContext _localctx = new ColumnExprListContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_columnExprList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(320);
			columnsExpr();
			setState(325);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(321);
				match(COMMA);
				setState(322);
				columnsExpr();
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

	public static class ColumnsExprContext extends ParserRuleContext {
		public ColumnExprContext columnExpr() {
			return getRuleContext(ColumnExprContext.class,0);
		}
		public ColumnsExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnsExpr; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterColumnsExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitColumnsExpr(this);
		}
	}

	public final ColumnsExprContext columnsExpr() throws RecognitionException {
		ColumnsExprContext _localctx = new ColumnsExprContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_columnsExpr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(328);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterColumnExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitColumnExpr(this);
		}
	}

	public final ColumnExprContext columnExpr() throws RecognitionException {
		ColumnExprContext _localctx = new ColumnExprContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_columnExpr);
		try {
			setState(332);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(330);
				literal();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(331);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterColumnIdentifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitColumnIdentifier(this);
		}
	}

	public final ColumnIdentifierContext columnIdentifier() throws RecognitionException {
		ColumnIdentifierContext _localctx = new ColumnIdentifierContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_columnIdentifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(337);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				{
				setState(334);
				tableIdentifier();
				setState(335);
				match(DOT);
				}
				break;
			}
			setState(339);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterNestedIdentifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitNestedIdentifier(this);
		}
	}

	public final NestedIdentifierContext nestedIdentifier() throws RecognitionException {
		NestedIdentifierContext _localctx = new NestedIdentifierContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_nestedIdentifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(341);
			identifier();
			setState(344);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				{
				setState(342);
				match(DOT);
				setState(343);
				identifier();
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterConstant(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitConstant(this);
		}
	}

	public final ConstantContext constant() throws RecognitionException {
		ConstantContext _localctx = new ConstantContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_constant);
		int _la;
		try {
			setState(355);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,35,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(347);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DASH || _la==PLUS) {
					{
					setState(346);
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

				setState(349);
				realLiteral();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(351);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DASH || _la==PLUS) {
					{
					setState(350);
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

				setState(353);
				match(DECIMAL_LITERAL);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(354);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterRealLiteral(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitRealLiteral(this);
		}
	}

	public final RealLiteralContext realLiteral() throws RecognitionException {
		RealLiteralContext _localctx = new RealLiteralContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_realLiteral);
		int _la;
		try {
			setState(365);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DECIMAL_LITERAL:
				enterOuterAlt(_localctx, 1);
				{
				setState(357);
				match(DECIMAL_LITERAL);
				setState(358);
				match(DOT);
				setState(360);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
				case 1:
					{
					setState(359);
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
				setState(362);
				match(DOT);
				setState(363);
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
				setState(364);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterTableExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitTableExpr(this);
		}
	}

	public final TableExprContext tableExpr() throws RecognitionException {
		TableExprContext _localctx = new TableExprContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_tableExpr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(367);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterTableIdentifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitTableIdentifier(this);
		}
	}

	public final TableIdentifierContext tableIdentifier() throws RecognitionException {
		TableIdentifierContext _localctx = new TableIdentifierContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_tableIdentifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(372);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,38,_ctx) ) {
			case 1:
				{
				setState(369);
				databaseIdentifier();
				setState(370);
				match(DOT);
				}
				break;
			}
			setState(374);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterDatabaseIdentifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitDatabaseIdentifier(this);
		}
	}

	public final DatabaseIdentifierContext databaseIdentifier() throws RecognitionException {
		DatabaseIdentifierContext _localctx = new DatabaseIdentifierContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_databaseIdentifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(376);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterFloatingLiteral(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitFloatingLiteral(this);
		}
	}

	public final FloatingLiteralContext floatingLiteral() throws RecognitionException {
		FloatingLiteralContext _localctx = new FloatingLiteralContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_floatingLiteral);
		int _la;
		try {
			setState(386);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FLOATING_LITERAL:
				enterOuterAlt(_localctx, 1);
				{
				setState(378);
				match(FLOATING_LITERAL);
				}
				break;
			case DOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(379);
				match(DOT);
				setState(380);
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
				setState(381);
				match(DECIMAL_LITERAL);
				setState(382);
				match(DOT);
				setState(384);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
				case 1:
					{
					setState(383);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterNumberLiteral(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitNumberLiteral(this);
		}
	}

	public final NumberLiteralContext numberLiteral() throws RecognitionException {
		NumberLiteralContext _localctx = new NumberLiteralContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_numberLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(389);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DASH || _la==PLUS) {
				{
				setState(388);
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

			setState(397);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,42,_ctx) ) {
			case 1:
				{
				setState(391);
				floatingLiteral();
				}
				break;
			case 2:
				{
				setState(392);
				match(OCTAL_LITERAL);
				}
				break;
			case 3:
				{
				setState(393);
				match(DECIMAL_LITERAL);
				}
				break;
			case 4:
				{
				setState(394);
				match(HEXADECIMAL_LITERAL);
				}
				break;
			case 5:
				{
				setState(395);
				match(INF);
				}
				break;
			case 6:
				{
				setState(396);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterLiteral(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitLiteral(this);
		}
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_literal);
		try {
			setState(402);
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
				setState(399);
				numberLiteral();
				}
				break;
			case STRING_LITERAL:
				enterOuterAlt(_localctx, 2);
				{
				setState(400);
				match(STRING_LITERAL);
				}
				break;
			case NULL_SQL:
				enterOuterAlt(_localctx, 3);
				{
				setState(401);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterInterval(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitInterval(this);
		}
	}

	public final IntervalContext interval() throws RecognitionException {
		IntervalContext _localctx = new IntervalContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_interval);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(404);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterKeyword(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitKeyword(this);
		}
	}

	public final KeywordContext keyword() throws RecognitionException {
		KeywordContext _localctx = new KeywordContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_keyword);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(406);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterIdentifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitIdentifier(this);
		}
	}

	public final IdentifierContext identifier() throws RecognitionException {
		IdentifierContext _localctx = new IdentifierContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_identifier);
		try {
			setState(411);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(408);
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
				setState(409);
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
				setState(410);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterEnumValue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitEnumValue(this);
		}
	}

	public final EnumValueContext enumValue() throws RecognitionException {
		EnumValueContext _localctx = new EnumValueContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_enumValue);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(413);
			match(STRING_LITERAL);
			setState(414);
			match(EQ_SINGLE);
			setState(415);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).enterIfNotExists(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SqlParserListener ) ((SqlParserListener)listener).exitIfNotExists(this);
		}
	}

	public final IfNotExistsContext ifNotExists() throws RecognitionException {
		IfNotExistsContext _localctx = new IfNotExistsContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_ifNotExists);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(417);
			match(IF);
			setState(418);
			match(NOT);
			setState(419);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 17:
			return searchCondition_sempred((SearchConditionContext)_localctx, predIndex);
		case 19:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean searchCondition_sempred(SearchConditionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 4);
		case 1:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 5);
		case 3:
			return precpred(_ctx, 4);
		case 4:
			return precpred(_ctx, 3);
		case 5:
			return precpred(_ctx, 2);
		case 6:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3H\u01a8\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\3\2\3\2"+
		"\3\2\5\2X\n\2\3\3\3\3\5\3\\\n\3\3\4\3\4\5\4`\n\4\3\5\3\5\3\6\3\6\3\6\5"+
		"\6g\n\6\3\6\3\6\3\7\3\7\3\7\5\7n\n\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\7\bw"+
		"\n\b\f\b\16\bz\13\b\3\b\3\b\3\t\3\t\3\n\3\n\3\n\3\13\3\13\3\13\5\13\u0086"+
		"\n\13\3\13\3\13\5\13\u008a\n\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\7\f\u0093"+
		"\n\f\f\f\16\f\u0096\13\f\3\f\3\f\3\r\5\r\u009b\n\r\3\r\7\r\u009e\n\r\f"+
		"\r\16\r\u00a1\13\r\3\16\3\16\3\16\5\16\u00a6\n\16\6\16\u00a8\n\16\r\16"+
		"\16\16\u00a9\3\16\3\16\3\17\3\17\3\17\3\17\3\17\6\17\u00b3\n\17\r\17\16"+
		"\17\u00b4\3\17\3\17\5\17\u00b9\n\17\3\20\3\20\3\20\5\20\u00be\n\20\3\20"+
		"\5\20\u00c1\n\20\3\21\3\21\3\21\3\22\3\22\3\22\3\23\3\23\3\23\3\23\5\23"+
		"\u00cd\n\23\3\23\3\23\3\23\3\23\3\23\3\23\7\23\u00d5\n\23\f\23\16\23\u00d8"+
		"\13\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u00f2\n\24"+
		"\3\25\3\25\3\25\5\25\u00f7\n\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\3\25\3\25\3\25\3\25\3\25\7\25\u0108\n\25\f\25\16\25\u010b\13"+
		"\25\3\26\3\26\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3"+
		"\30\7\30\u011b\n\30\f\30\16\30\u011e\13\30\3\30\3\30\3\30\3\30\3\30\3"+
		"\30\3\30\7\30\u0127\n\30\f\30\16\30\u012a\13\30\3\30\3\30\3\30\3\30\3"+
		"\30\3\30\3\30\7\30\u0133\n\30\f\30\16\30\u0136\13\30\3\30\3\30\3\30\3"+
		"\30\3\30\5\30\u013d\n\30\3\30\3\30\5\30\u0141\n\30\3\31\3\31\3\31\7\31"+
		"\u0146\n\31\f\31\16\31\u0149\13\31\3\32\3\32\3\33\3\33\5\33\u014f\n\33"+
		"\3\34\3\34\3\34\5\34\u0154\n\34\3\34\3\34\3\35\3\35\3\35\5\35\u015b\n"+
		"\35\3\36\5\36\u015e\n\36\3\36\3\36\5\36\u0162\n\36\3\36\3\36\5\36\u0166"+
		"\n\36\3\37\3\37\3\37\5\37\u016b\n\37\3\37\3\37\3\37\5\37\u0170\n\37\3"+
		" \3 \3!\3!\3!\5!\u0177\n!\3!\3!\3\"\3\"\3#\3#\3#\3#\3#\3#\5#\u0183\n#"+
		"\5#\u0185\n#\3$\5$\u0188\n$\3$\3$\3$\3$\3$\3$\5$\u0190\n$\3%\3%\3%\5%"+
		"\u0195\n%\3&\3&\3\'\3\'\3(\3(\3(\5(\u019e\n(\3)\3)\3)\3)\3*\3*\3*\3*\3"+
		"*\2\4$(+\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668"+
		":<>@BDFHJLNPR\2\7\4\2,,99\4\2!!##\3\2 !\3\2\26\35\6\2\4\t\13\f\16\22\24"+
		"\25\2\u01bf\2W\3\2\2\2\4[\3\2\2\2\6_\3\2\2\2\ba\3\2\2\2\nc\3\2\2\2\fj"+
		"\3\2\2\2\16r\3\2\2\2\20}\3\2\2\2\22\177\3\2\2\2\24\u0082\3\2\2\2\26\u008e"+
		"\3\2\2\2\30\u009f\3\2\2\2\32\u00a2\3\2\2\2\34\u00b8\3\2\2\2\36\u00ba\3"+
		"\2\2\2 \u00c2\3\2\2\2\"\u00c5\3\2\2\2$\u00cc\3\2\2\2&\u00f1\3\2\2\2(\u00f6"+
		"\3\2\2\2*\u010c\3\2\2\2,\u010e\3\2\2\2.\u0140\3\2\2\2\60\u0142\3\2\2\2"+
		"\62\u014a\3\2\2\2\64\u014e\3\2\2\2\66\u0153\3\2\2\28\u0157\3\2\2\2:\u0165"+
		"\3\2\2\2<\u016f\3\2\2\2>\u0171\3\2\2\2@\u0176\3\2\2\2B\u017a\3\2\2\2D"+
		"\u0184\3\2\2\2F\u0187\3\2\2\2H\u0194\3\2\2\2J\u0196\3\2\2\2L\u0198\3\2"+
		"\2\2N\u019d\3\2\2\2P\u019f\3\2\2\2R\u01a3\3\2\2\2TX\5\6\4\2UX\5\4\3\2"+
		"VX\5\b\5\2WT\3\2\2\2WU\3\2\2\2WV\3\2\2\2X\3\3\2\2\2Y\\\5\36\20\2Z\\\5"+
		"\24\13\2[Y\3\2\2\2[Z\3\2\2\2\\\5\3\2\2\2]`\5\n\6\2^`\5\f\7\2_]\3\2\2\2"+
		"_^\3\2\2\2`\7\3\2\2\2ab\5,\27\2b\t\3\2\2\2cd\7\4\2\2df\7\5\2\2eg\5R*\2"+
		"fe\3\2\2\2fg\3\2\2\2gh\3\2\2\2hi\5B\"\2i\13\3\2\2\2jk\7\4\2\2km\7\22\2"+
		"\2ln\5R*\2ml\3\2\2\2mn\3\2\2\2no\3\2\2\2op\5@!\2pq\5\16\b\2q\r\3\2\2\2"+
		"rs\7\65\2\2sx\5\20\t\2tu\7*\2\2uw\5\20\t\2vt\3\2\2\2wz\3\2\2\2xv\3\2\2"+
		"\2xy\3\2\2\2y{\3\2\2\2zx\3\2\2\2{|\7?\2\2|\17\3\2\2\2}~\5\22\n\2~\21\3"+
		"\2\2\2\177\u0080\58\35\2\u0080\u0081\5.\30\2\u0081\23\3\2\2\2\u0082\u0083"+
		"\7\13\2\2\u0083\u0085\7\f\2\2\u0084\u0086\7\22\2\2\u0085\u0084\3\2\2\2"+
		"\u0085\u0086\3\2\2\2\u0086\u0087\3\2\2\2\u0087\u0089\5@!\2\u0088\u008a"+
		"\5\26\f\2\u0089\u0088\3\2\2\2\u0089\u008a\3\2\2\2\u008a\u008b\3\2\2\2"+
		"\u008b\u008c\7\25\2\2\u008c\u008d\5\30\r\2\u008d\25\3\2\2\2\u008e\u008f"+
		"\7\65\2\2\u008f\u0094\58\35\2\u0090\u0091\7*\2\2\u0091\u0093\58\35\2\u0092"+
		"\u0090\3\2\2\2\u0093\u0096\3\2\2\2\u0094\u0092\3\2\2\2\u0094\u0095\3\2"+
		"\2\2\u0095\u0097\3\2\2\2\u0096\u0094\3\2\2\2\u0097\u0098\7?\2\2\u0098"+
		"\27\3\2\2\2\u0099\u009b\7*\2\2\u009a\u0099\3\2\2\2\u009a\u009b\3\2\2\2"+
		"\u009b\u009c\3\2\2\2\u009c\u009e\5\32\16\2\u009d\u009a\3\2\2\2\u009e\u00a1"+
		"\3\2\2\2\u009f\u009d\3\2\2\2\u009f\u00a0\3\2\2\2\u00a0\31\3\2\2\2\u00a1"+
		"\u009f\3\2\2\2\u00a2\u00a7\7\65\2\2\u00a3\u00a5\5\34\17\2\u00a4\u00a6"+
		"\7*\2\2\u00a5\u00a4\3\2\2\2\u00a5\u00a6\3\2\2\2\u00a6\u00a8\3\2\2\2\u00a7"+
		"\u00a3\3\2\2\2\u00a8\u00a9\3\2\2\2\u00a9\u00a7\3\2\2\2\u00a9\u00aa\3\2"+
		"\2\2\u00aa\u00ab\3\2\2\2\u00ab\u00ac\7?\2\2\u00ac\33\3\2\2\2\u00ad\u00b9"+
		"\5:\36\2\u00ae\u00af\7\65\2\2\u00af\u00b2\5:\36\2\u00b0\u00b1\7*\2\2\u00b1"+
		"\u00b3\5:\36\2\u00b2\u00b0\3\2\2\2\u00b3\u00b4\3\2\2\2\u00b4\u00b2\3\2"+
		"\2\2\u00b4\u00b5\3\2\2\2\u00b5\u00b6\3\2\2\2\u00b6\u00b7\7?\2\2\u00b7"+
		"\u00b9\3\2\2\2\u00b8\u00ad\3\2\2\2\u00b8\u00ae\3\2\2\2\u00b9\35\3\2\2"+
		"\2\u00ba\u00bb\7\20\2\2\u00bb\u00bd\5\60\31\2\u00bc\u00be\5 \21\2\u00bd"+
		"\u00bc\3\2\2\2\u00bd\u00be\3\2\2\2\u00be\u00c0\3\2\2\2\u00bf\u00c1\5\""+
		"\22\2\u00c0\u00bf\3\2\2\2\u00c0\u00c1\3\2\2\2\u00c1\37\3\2\2\2\u00c2\u00c3"+
		"\7\b\2\2\u00c3\u00c4\5*\26\2\u00c4!\3\2\2\2\u00c5\u00c6\7F\2\2\u00c6\u00c7"+
		"\5$\23\2\u00c7#\3\2\2\2\u00c8\u00c9\b\23\1\2\u00c9\u00ca\7\16\2\2\u00ca"+
		"\u00cd\5$\23\4\u00cb\u00cd\5&\24\2\u00cc\u00c8\3\2\2\2\u00cc\u00cb\3\2"+
		"\2\2\u00cd\u00d6\3\2\2\2\u00ce\u00cf\f\6\2\2\u00cf\u00d0\7G\2\2\u00d0"+
		"\u00d5\5$\23\7\u00d1\u00d2\f\5\2\2\u00d2\u00d3\7H\2\2\u00d3\u00d5\5$\23"+
		"\6\u00d4\u00ce\3\2\2\2\u00d4\u00d1\3\2\2\2\u00d5\u00d8\3\2\2\2\u00d6\u00d4"+
		"\3\2\2\2\u00d6\u00d7\3\2\2\2\u00d7%\3\2\2\2\u00d8\u00d6\3\2\2\2\u00d9"+
		"\u00da\5(\25\2\u00da\u00db\7/\2\2\u00db\u00dc\5(\25\2\u00dc\u00f2\3\2"+
		"\2\2\u00dd\u00de\5(\25\2\u00de\u00df\7\67\2\2\u00df\u00e0\5(\25\2\u00e0"+
		"\u00f2\3\2\2\2\u00e1\u00e2\5(\25\2\u00e2\u00e3\7\66\2\2\u00e3\u00e4\5"+
		"(\25\2\u00e4\u00f2\3\2\2\2\u00e5\u00e6\5(\25\2\u00e6\u00e7\7\64\2\2\u00e7"+
		"\u00e8\5(\25\2\u00e8\u00f2\3\2\2\2\u00e9\u00ea\5(\25\2\u00ea\u00eb\7\61"+
		"\2\2\u00eb\u00ec\5(\25\2\u00ec\u00f2\3\2\2\2\u00ed\u00ee\5(\25\2\u00ee"+
		"\u00ef\7\60\2\2\u00ef\u00f0\5(\25\2\u00f0\u00f2\3\2\2\2\u00f1\u00d9\3"+
		"\2\2\2\u00f1\u00dd\3\2\2\2\u00f1\u00e1\3\2\2\2\u00f1\u00e5\3\2\2\2\u00f1"+
		"\u00e9\3\2\2\2\u00f1\u00ed\3\2\2\2\u00f2\'\3\2\2\2\u00f3\u00f4\b\25\1"+
		"\2\u00f4\u00f7\5H%\2\u00f5\u00f7\5\66\34\2\u00f6\u00f3\3\2\2\2\u00f6\u00f5"+
		"\3\2\2\2\u00f7\u0109\3\2\2\2\u00f8\u00f9\f\7\2\2\u00f9\u00fa\7&\2\2\u00fa"+
		"\u0108\5(\25\b\u00fb\u00fc\f\6\2\2\u00fc\u00fd\7A\2\2\u00fd\u0108\5(\25"+
		"\7\u00fe\u00ff\f\5\2\2\u00ff\u0100\78\2\2\u0100\u0108\5(\25\6\u0101\u0102"+
		"\f\4\2\2\u0102\u0103\79\2\2\u0103\u0108\5(\25\5\u0104\u0105\f\3\2\2\u0105"+
		"\u0106\7,\2\2\u0106\u0108\5(\25\4\u0107\u00f8\3\2\2\2\u0107\u00fb\3\2"+
		"\2\2\u0107\u00fe\3\2\2\2\u0107\u0101\3\2\2\2\u0107\u0104\3\2\2\2\u0108"+
		"\u010b\3\2\2\2\u0109\u0107\3\2\2\2\u0109\u010a\3\2\2\2\u010a)\3\2\2\2"+
		"\u010b\u0109\3\2\2\2\u010c\u010d\5> \2\u010d+\3\2\2\2\u010e\u010f\7\23"+
		"\2\2\u010f\u0110\5B\"\2\u0110-\3\2\2\2\u0111\u0141\5N(\2\u0112\u0113\5"+
		"N(\2\u0113\u0114\7\65\2\2\u0114\u0115\5N(\2\u0115\u011c\5.\30\2\u0116"+
		"\u0117\7*\2\2\u0117\u0118\5N(\2\u0118\u0119\5.\30\2\u0119\u011b\3\2\2"+
		"\2\u011a\u0116\3\2\2\2\u011b\u011e\3\2\2\2\u011c\u011a\3\2\2\2\u011c\u011d"+
		"\3\2\2\2\u011d\u011f\3\2\2\2\u011e\u011c\3\2\2\2\u011f\u0120\7?\2\2\u0120"+
		"\u0141\3\2\2\2\u0121\u0122\5N(\2\u0122\u0123\7\65\2\2\u0123\u0128\5P)"+
		"\2\u0124\u0125\7*\2\2\u0125\u0127\5P)\2\u0126\u0124\3\2\2\2\u0127\u012a"+
		"\3\2\2\2\u0128\u0126\3\2\2\2\u0128\u0129\3\2\2\2\u0129\u012b\3\2\2\2\u012a"+
		"\u0128\3\2\2\2\u012b\u012c\7?\2\2\u012c\u0141\3\2\2\2\u012d\u012e\5N("+
		"\2\u012e\u012f\7\65\2\2\u012f\u0134\5.\30\2\u0130\u0131\7*\2\2\u0131\u0133"+
		"\5.\30\2\u0132\u0130\3\2\2\2\u0133\u0136\3\2\2\2\u0134\u0132\3\2\2\2\u0134"+
		"\u0135\3\2\2\2\u0135\u0137\3\2\2\2\u0136\u0134\3\2\2\2\u0137\u0138\7?"+
		"\2\2\u0138\u0141\3\2\2\2\u0139\u013a\5N(\2\u013a\u013c\7\65\2\2\u013b"+
		"\u013d\5\60\31\2\u013c\u013b\3\2\2\2\u013c\u013d\3\2\2\2\u013d\u013e\3"+
		"\2\2\2\u013e\u013f\7?\2\2\u013f\u0141\3\2\2\2\u0140\u0111\3\2\2\2\u0140"+
		"\u0112\3\2\2\2\u0140\u0121\3\2\2\2\u0140\u012d\3\2\2\2\u0140\u0139\3\2"+
		"\2\2\u0141/\3\2\2\2\u0142\u0147\5\62\32\2\u0143\u0144\7*\2\2\u0144\u0146"+
		"\5\62\32\2\u0145\u0143\3\2\2\2\u0146\u0149\3\2\2\2\u0147\u0145\3\2\2\2"+
		"\u0147\u0148\3\2\2\2\u0148\61\3\2\2\2\u0149\u0147\3\2\2\2\u014a\u014b"+
		"\5\64\33\2\u014b\63\3\2\2\2\u014c\u014f\5H%\2\u014d\u014f\5\66\34\2\u014e"+
		"\u014c\3\2\2\2\u014e\u014d\3\2\2\2\u014f\65\3\2\2\2\u0150\u0151\5@!\2"+
		"\u0151\u0152\7-\2\2\u0152\u0154\3\2\2\2\u0153\u0150\3\2\2\2\u0153\u0154"+
		"\3\2\2\2\u0154\u0155\3\2\2\2\u0155\u0156\58\35\2\u0156\67\3\2\2\2\u0157"+
		"\u015a\5N(\2\u0158\u0159\7-\2\2\u0159\u015b\5N(\2\u015a\u0158\3\2\2\2"+
		"\u015a\u015b\3\2\2\2\u015b9\3\2\2\2\u015c\u015e\t\2\2\2\u015d\u015c\3"+
		"\2\2\2\u015d\u015e\3\2\2\2\u015e\u015f\3\2\2\2\u015f\u0166\5<\37\2\u0160"+
		"\u0162\t\2\2\2\u0161\u0160\3\2\2\2\u0161\u0162\3\2\2\2\u0162\u0163\3\2"+
		"\2\2\u0163\u0166\7!\2\2\u0164\u0166\7$\2\2\u0165\u015d\3\2\2\2\u0165\u0161"+
		"\3\2\2\2\u0165\u0164\3\2\2\2\u0166;\3\2\2\2\u0167\u0168\7!\2\2\u0168\u016a"+
		"\7-\2\2\u0169\u016b\t\3\2\2\u016a\u0169\3\2\2\2\u016a\u016b\3\2\2\2\u016b"+
		"\u0170\3\2\2\2\u016c\u016d\7-\2\2\u016d\u0170\t\3\2\2\u016e\u0170\7#\2"+
		"\2\u016f\u0167\3\2\2\2\u016f\u016c\3\2\2\2\u016f\u016e\3\2\2\2\u0170="+
		"\3\2\2\2\u0171\u0172\5@!\2\u0172?\3\2\2\2\u0173\u0174\5B\"\2\u0174\u0175"+
		"\7-\2\2\u0175\u0177\3\2\2\2\u0176\u0173\3\2\2\2\u0176\u0177\3\2\2\2\u0177"+
		"\u0178\3\2\2\2\u0178\u0179\5N(\2\u0179A\3\2\2\2\u017a\u017b\5N(\2\u017b"+
		"C\3\2\2\2\u017c\u0185\7\37\2\2\u017d\u017e\7-\2\2\u017e\u0185\t\4\2\2"+
		"\u017f\u0180\7!\2\2\u0180\u0182\7-\2\2\u0181\u0183\t\4\2\2\u0182\u0181"+
		"\3\2\2\2\u0182\u0183\3\2\2\2\u0183\u0185\3\2\2\2\u0184\u017c\3\2\2\2\u0184"+
		"\u017d\3\2\2\2\u0184\u017f\3\2\2\2\u0185E\3\2\2\2\u0186\u0188\t\2\2\2"+
		"\u0187\u0186\3\2\2\2\u0187\u0188\3\2\2\2\u0188\u018f\3\2\2\2\u0189\u0190"+
		"\5D#\2\u018a\u0190\7 \2\2\u018b\u0190\7!\2\2\u018c\u0190\7\"\2\2\u018d"+
		"\u0190\7\n\2\2\u018e\u0190\7\r\2\2\u018f\u0189\3\2\2\2\u018f\u018a\3\2"+
		"\2\2\u018f\u018b\3\2\2\2\u018f\u018c\3\2\2\2\u018f\u018d\3\2\2\2\u018f"+
		"\u018e\3\2\2\2\u0190G\3\2\2\2\u0191\u0195\5F$\2\u0192\u0195\7$\2\2\u0193"+
		"\u0195\7\17\2\2\u0194\u0191\3\2\2\2\u0194\u0192\3\2\2\2\u0194\u0193\3"+
		"\2\2\2\u0195I\3\2\2\2\u0196\u0197\t\5\2\2\u0197K\3\2\2\2\u0198\u0199\t"+
		"\6\2\2\u0199M\3\2\2\2\u019a\u019e\7\36\2\2\u019b\u019e\5J&\2\u019c\u019e"+
		"\5L\'\2\u019d\u019a\3\2\2\2\u019d\u019b\3\2\2\2\u019d\u019c\3\2\2\2\u019e"+
		"O\3\2\2\2\u019f\u01a0\7$\2\2\u01a0\u01a1\7/\2\2\u01a1\u01a2\5F$\2\u01a2"+
		"Q\3\2\2\2\u01a3\u01a4\7\t\2\2\u01a4\u01a5\7\16\2\2\u01a5\u01a6\7\7\2\2"+
		"\u01a6S\3\2\2\2/W[_fmx\u0085\u0089\u0094\u009a\u009f\u00a5\u00a9\u00b4"+
		"\u00b8\u00bd\u00c0\u00cc\u00d4\u00d6\u00f1\u00f6\u0107\u0109\u011c\u0128"+
		"\u0134\u013c\u0140\u0147\u014e\u0153\u015a\u015d\u0161\u0165\u016a\u016f"+
		"\u0176\u0182\u0184\u0187\u018f\u0194\u019d";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}