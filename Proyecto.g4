// Proyecto.g4
grammar Proyecto;

// Tokens
INT: [0-9]+;
FLOAT: INT ('.' INT)? ('E' [+-]? INT)?;
CHAR: '"'[A-Za-z]'"';
BOOL: 'true' | 'false';
WS: [ \r\n\t]+ -> skip;
GT: '>';
LT: '<';
EQ: '==';
NEQ: '!=';
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
SEMI: ';';

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

functions: FUNCTION ID '(' parameters? ')' (typeRule | VOID) functionBlock;
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

functionCall: ID '(' arguments? ')';
arguments: argument (',' argument)*;
argument: (vars | exp);

methodCall: ID '.' ID '(' arguments? ')';

call: functionCall | methodCall;

read: READ '(' vars (',' vars)* ')' SEMI;

write: WRITE '(' arguments ')' SEMI;

conditional: IF '(' exp ')' block (ELSE block)?;

forLoop: FOR ID '=' exp IN exp block;

whileLoop: WHILE '(' exp ')' block;

exp: t_exp ('||' t_exp)*;
t_exp: g_exp ('&&' g_exp)*;
g_exp: m_exp (relop m_exp)?;
m_exp: term ((ADD | SUB) term)*;
term: factor ((MUL | DIV) factor)*;
factor: '(' exp ')'
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

main: MAIN '('')' functionBlock;

typeRule: INT_TYPE
    | FLOAT_TYPE
    | CHAR_TYPE
    | STRING_TYPE
    | BOOL_TYPE;

relop: GT
     | LT
     | EQ
     | NEQ;