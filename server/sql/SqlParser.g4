parser grammar SqlParser;

options { tokenVocab=SqlLexer; }


/**
 * 1. Top Level Description
 */
statement
    : ddlStatement
    | dmlStatement
    | systemStatement
    ;

dmlStatement
    : selectStatement
    | insertStatement
    | updateStatement
    | deleteStatement
    ;

ddlStatement
    : createDatabase 
    | createTable
    ;

systemStatement
    : useStatement
    ;
/**
 * 2. Data Definition Language (DDL)
 */

createDatabase
    : CREATE DATABASE ifNotExists? databaseIdentifier
    ;

createTable
    : CREATE TABLE ifNotExists?
       tableIdentifier tableSchemaClause
    ;

tableSchemaClause
    : LPAREN tableElementExpr (COMMA tableElementExpr)* RPAREN
    ;

tableElementExpr
    : tableColumnDfnt 
    ;

tableColumnDfnt
    : nestedIdentifier columnTypeExpr 
    ;

/**
 * 3. Data Manipulation Language (DML)
 */

insertStatement: INSERT INTO TABLE? (tableIdentifier) columnsClause? VALUES insertValuesSpec;

updateStatement: UPDATE tableIdentifier SET setClause (COMMA setClause)* whereClause?;

deleteStatement: DELETE FROM tableIdentifier whereClause?;

setClause: columnIdentifier EQ_SINGLE expression;

columnsClause: LPAREN nestedIdentifier (COMMA nestedIdentifier)* RPAREN;

insertValuesSpec
    : (COMMA? insertMultiValue)*
    ;

insertMultiValue
    : LPAREN (dataValue COMMA?)+ RPAREN
    ;

dataValue
    : constant
    | LPAREN constant (COMMA constant)+ RPAREN
    ;


// Select Statement
selectStatement
    : SELECT columnExprList
      fromClause?
      whereClause?
      groupByClause?
      havingClause?
      orderByClause?
      limitClause?
      offsetClause?
    ;

fromClause: FROM joinExpr;

whereClause: WHERE searchCondition;

groupByClause: GROUP BY columnIdentifier (COMMA columnIdentifier)*;

havingClause: HAVING searchCondition;

orderByClause: ORDER BY orderByExpr (COMMA orderByExpr)*;

orderByExpr: columnIdentifier (ASC | DESC)?;

limitClause: LIMIT DECIMAL_LITERAL;

offsetClause: OFFSET DECIMAL_LITERAL;

searchCondition
    : searchCondition AND searchCondition
    | searchCondition OR searchCondition
    | NOT searchCondition
    | predicate
    ;

predicate
    : expression EQ_SINGLE expression
    | expression NOT_EQ expression
    | expression LT expression
    | expression LE expression
    | expression GT expression
    | expression GE expression
    ;

expression
    : literal
    | columnIdentifier
    | expression ASTERISK expression
    | expression SLASH expression
    | expression PERCENT expression
    | expression PLUS expression
    | expression DASH expression
    ;


joinExpr
    : tableExpr (joinOp tableExpr ON searchCondition)*
    ;

joinOp
    : INNER? JOIN
    | LEFT JOIN
    ;




/**
 * 4. system
 */
useStatement: USE databaseIdentifier;


// Columns
columnTypeExpr
    : identifier 
    | identifier LPAREN identifier columnTypeExpr (COMMA identifier columnTypeExpr)* RPAREN
    | identifier LPAREN enumValue (COMMA enumValue)* RPAREN
    | identifier LPAREN columnTypeExpr (COMMA columnTypeExpr)* RPAREN
    | identifier LPAREN columnExprList? RPAREN
    ;

columnExprList: columnsExpr (COMMA columnsExpr)*;

columnsExpr
    : columnExpr
    ;


columnExpr
    : literal
    | columnIdentifier
    | aggregateFunction
    ;

aggregateFunction
    : COUNT LPAREN (ASTERISK | expression) RPAREN
    | SUM LPAREN expression RPAREN
    | AVG LPAREN expression RPAREN
    | MIN LPAREN expression RPAREN
    | MAX LPAREN expression RPAREN
    ;

columnIdentifier: (tableIdentifier DOT)? nestedIdentifier;
nestedIdentifier: identifier (DOT identifier)?;


// Constant
constant
    : (DASH|PLUS)? realLiteral
    | (DASH|PLUS)? DECIMAL_LITERAL
    | STRING_LITERAL
    ;

realLiteral
    : DECIMAL_LITERAL DOT (DECIMAL_LITERAL|EXPONENT_NUM_PART)?
    | DOT (DECIMAL_LITERAL|EXPONENT_NUM_PART)
    | EXPONENT_NUM_PART
    ;


// Tables

tableExpr
    : tableIdentifier
    ;

tableIdentifier: (databaseIdentifier DOT)? identifier;

// Databases

databaseIdentifier: identifier;

// //Basics


floatingLiteral
    : FLOATING_LITERAL
    | DOT (DECIMAL_LITERAL | OCTAL_LITERAL)
    | DECIMAL_LITERAL DOT (DECIMAL_LITERAL | OCTAL_LITERAL)?  // can't move this to the lexer or it will break nested tuple access: t.1.2
    ;
numberLiteral: (PLUS | DASH)? (floatingLiteral | OCTAL_LITERAL | DECIMAL_LITERAL | HEXADECIMAL_LITERAL | INF | NAN_SQL);
literal
    : numberLiteral
    | STRING_LITERAL
    | NULL_SQL
    ;
interval: SECOND | MINUTE | HOUR | DAY | WEEK | MONTH | QUARTER | YEAR;
keyword
    :CREATE | DATABASE | DEFAULT| EXISTS | FROM | IF | INSERT | INTO | NOT | NULL_SQL | SELECT | SET | TABLE | VALUE | VALUES
    ;

identifier: IDENTIFIER | interval | keyword;

enumValue: STRING_LITERAL EQ_SINGLE numberLiteral;

ifNotExists
    : IF NOT EXISTS;
