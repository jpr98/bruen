// Proyecto.g4
grammar Proyecto;

// Tokens
INT: [0-9]+;
FLOAT: INT ('.' INT)? ('E' [+-]? INT)?;
CHAR: '"'[A-Za-z]'"';
BOOL: 'true' | 'false';
WS: [ \r\n\t]+ -> skip;
OR: '||';
AND: '&&';
GT: '>';
LT: '<';
EQ: '==';
NEQ: '!=';
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
SEMI: ';';
LPAREN: '(';
RPAREN: ')';

// -- Keywords -- 
// Conditionals
IF: 'if';
ELSE: 'else';
// Loops
WHILE: 'while';
FOR: 'for';
IN: 'in';
// Classes
CLASS: 'class';
ATTRIBUTES: 'attributes';
METHODS: 'methods';
// I/O
WRITE: 'write';
READ: 'read';
// FUNCTIONS
FUNCTION: 'function';
RETURN: 'return';
MAIN: 'main';
// Variables
VAR: 'var';
// Types
INT_TYPE: 'int';
FLOAT_TYPE: 'float';
CHAR_TYPE: 'char';
STRING_TYPE: 'string';
BOOL_TYPE: 'bool';
VOID: 'void';
// Program
PROGRAM: 'program';

ID: [a-zA-Z_][a-zA-Z0-9]*;

// Rules
program: PROGRAM ID SEMI classDef* varsDec* functions* main EOF;

classDef: CLASS ID ('<' ID '>')? classBlock SEMI;
classBlock: '{' ATTRIBUTES varsDec* METHODS functions* '}';

varsDec: VAR ID ':' varsTypeInit SEMI
    | VAR ID '['INT']' ':' varsTypeInit SEMI
    | VAR ID '['INT']['INT']' ':' varsTypeInit SEMI;
varsTypeInit: (typeRule | ID) ('=' exp)?;

vars: ID ('.' ID)? ('[' exp ']')? ('[' exp ']')?;

functions: FUNCTION ID LPAREN parameters? RPAREN (typeRule | VOID) functionBlock;
parameters: parameter (',' parameter)*;
parameter: ID ':' typeRule;
functionBlock: '{' varsDec* statutes* returnRule'}';
returnRule: RETURN exp? SEMI;

block: '{' statutes* '}';

statutes: assignation
        | read
        | write
        | conditional
        | forLoop
        | whileLoop
        | exp;

assignation: ID '=' exp SEMI;

functionCall: ID LPAREN arguments? RPAREN;
arguments: argument (',' argument)*;
argument: (vars | exp);

methodCall: ID '.' ID LPAREN arguments? RPAREN;

call: functionCall | methodCall;

read: READ LPAREN vars (',' vars)* RPAREN SEMI;

write: WRITE LPAREN arguments RPAREN SEMI;

conditional: IF conditional2 conditional3 conditional4?;
conditional2: LPAREN exp RPAREN;
conditional3: block;
conditional4: ELSE block;

forLoop: FOR ID '=' exp IN exp block;

whileLoop: WHILE LPAREN exp RPAREN block;

exp: t_exp exp2*;
exp2: (OR t_exp);

t_exp: g_exp t_exp2*;
t_exp2: (AND g_exp);

g_exp: m_exp (g_exp2)?;
g_exp2: relop m_exp;

m_exp: term m_exp2*;
m_exp2: (ADD | SUB) term;

term: factor term2*;
term2: (MUL | DIV) factor;

factor: LPAREN exp RPAREN
      | varCte
      | vars
      | call;

varCte: cte_i
      | cte_f
      | cte_c
      | cte_s
      | cte_b;

cte_i: INT;
cte_f: FLOAT;
cte_c: CHAR;
cte_b: BOOL;
cte_s: '"' .*? '"'; // checar espacios en strings

main: MAIN LPAREN RPAREN functionBlock;

typeRule: INT_TYPE
    | FLOAT_TYPE
    | CHAR_TYPE
    | STRING_TYPE
    | BOOL_TYPE;

relop: GT
     | LT
     | EQ
     | NEQ;