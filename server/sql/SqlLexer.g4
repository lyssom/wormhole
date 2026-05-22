
lexer grammar SqlLexer;

/**
 * Whitespace
 */

// Instead of discarding whitespace completely, send them to a channel invisable to the parser, so
// that the lexer could still produce WS tokens for the CLI's highlighter.
WS
    :
    [ \u000B\t\r\n]+ -> channel(HIDDEN)
    ;


/**
 * Keywords
 */

// Common Keywords

AND: A N D;
ASC: A S C;
AVG: A V G;
BY: B Y;
COUNT: C O U N T;
CREATE: C R E A T E;
DATABASE: D A T A B A S E;
DEFAULT: D E F A U L T;
DELETE: D E L E T E;
DESC: D E S C;
EXISTS: E X I S T S;
FROM: F R O M;
GROUP: G R O U P;
HAVING: H A V I N G;
IF: I F;
INF: I N F | I N F I N I T Y;
INNER: I N N E R;
INSERT: I N S E R T;
INTO: I N T O;
JOIN: J O I N;
LEFT: L E F T;
LIMIT: L I M I T;
MAX: M A X;
MIN: M I N;
NAN_SQL: N A N;
NOT: N O T;
NULL_SQL: N U L L;
OFFSET: O F F S E T;
ON: O N;
OR: O R;
ORDER: O R D E R;
SELECT: S E L E C T;
SET: S E T;
SUM: S U M;
TABLE: T A B L E;
UPDATE: U P D A T E;
USE: U S E;
VALUE: V A L U E;
VALUES: V A L U E S;
WHERE: W H E R E;


// interval keywords

SECOND: S E C O N D;
MINUTE: M I N U T E;
HOUR: H O U R;
DAY: D A Y;
WEEK: W E E K;
MONTH: M O N T H;
QUARTER: Q U A R T E R;
YEAR: Y E A R;




/**
 * Tokens
 */

IDENTIFIER
    : (LETTER | UNDERSCORE) (LETTER | UNDERSCORE | DEC_DIGIT)*
    | BACKQUOTE ( ~([\\`]) | (BACKSLASH .) | (BACKQUOTE BACKQUOTE) )* BACKQUOTE
    | QUOTE_DOUBLE ( ~([\\"]) | (BACKSLASH .) | (QUOTE_DOUBLE QUOTE_DOUBLE) )* QUOTE_DOUBLE
    ;
FLOATING_LITERAL
    : HEXADECIMAL_LITERAL DOT HEX_DIGIT* (P | E) (PLUS | DASH)? DEC_DIGIT+
    | HEXADECIMAL_LITERAL (P | E) (PLUS | DASH)? DEC_DIGIT+
    | DECIMAL_LITERAL DOT DEC_DIGIT* E (PLUS | DASH)? DEC_DIGIT+
    | DOT DECIMAL_LITERAL E (PLUS | DASH)? DEC_DIGIT+
    | DECIMAL_LITERAL E (PLUS | DASH)? DEC_DIGIT+
    ;
OCTAL_LITERAL: '0' OCT_DIGIT+;
DECIMAL_LITERAL: DEC_DIGIT+;
HEXADECIMAL_LITERAL: '0' X HEX_DIGIT+;
EXPONENT_NUM_PART: DEC_DIGIT+ ('e'|'E') ('+'|'-')? DEC_DIGIT+;

// It's important that quote-symbol is a single character.
STRING_LITERAL: QUOTE_SINGLE ( ~([\\']) | (BACKSLASH .) | (QUOTE_SINGLE QUOTE_SINGLE) )* QUOTE_SINGLE;

// Alphabet and allowed symbols

fragment A: [aA];
fragment B: [bB];
fragment C: [cC];
fragment D: [dD];
fragment E: [eE];
fragment F: [fF];
fragment G: [gG];
fragment H: [hH];
fragment I: [iI];
fragment J: [jJ];
fragment K: [kK];
fragment L: [lL];
fragment M: [mM];
fragment N: [nN];
fragment O: [oO];
fragment P: [pP];
fragment Q: [qQ];
fragment R: [rR];
fragment S: [sS];
fragment T: [tT];
fragment U: [uU];
fragment V: [vV];
fragment W: [wW];
fragment X: [xX];
fragment Y: [yY];
fragment Z: [zZ];

fragment LETTER: [a-zA-Z];
fragment OCT_DIGIT: [0-7];
fragment DEC_DIGIT: [0-9];
fragment HEX_DIGIT: [0-9a-fA-F];

ARROW: '->';
ASTERISK: '*';
BACKQUOTE: '`';
BACKSLASH: '\\';
COLON: ':';
COMMA: ',';
CONCAT: '||';
DASH: '-';
DOT: '.';
EQ_DOUBLE: '==';
EQ_SINGLE: '=';
GE: '>=';
GT: '>';
LBRACE: '{';
LBRACKET: '[';
LE: '<=';
LPAREN: '(';
LT: '<';
NOT_EQ: '!=' | '<>';
PERCENT: '%';
PLUS: '+';
QUERY: '?';
QUOTE_DOUBLE: '"';
QUOTE_SINGLE: '\'';
RBRACE: '}';
RBRACKET: ']';
RPAREN: ')';
SEMICOLON: ';';
SLASH: '/';
UNDERSCORE: '_';

// Comments and whitespace

MULTI_LINE_COMMENT: '/*' .*? '*/' -> skip;
SINGLE_LINE_COMMENT: '--' ~('\n'|'\r')* ('\n' | '\r' | EOF) -> skip;
WHITESPACE: [ \u000B\u000C\t\r\n] -> skip;  // '\n' can be part of multiline single query
