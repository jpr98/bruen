// Proyecto.g4
grammar Proyecto;

// Tokens
INT: [0-9]+;
FLOAT: INT ('.' INT)? ('E' [+-]? INT)?;
CHAR: '"'[A-Za-z]'"';
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
VOID: 'void';
// Program
PROGRAM: 'program';

ID: [a-zA-Z_][a-zA-Z0-9]*;

// Rules
program: PROGRAM ID SEMI classDef* vars* functions* main EOF;

classDef: CLASS ID ('<' ID '>')? classBlock SEMI;
classBlock: '{' ATTRIBUTES vars* METHODS functions* '}';

vars: VAR ID ':' varsTypeInit SEMI
    | VAR ID '['INT']' ':' varsTypeInit SEMI
    | VAR ID '['INT']['INT']' ':' varsTypeInit SEMI;
varsTypeInit: typeRule ('=' expression)?;

functions: FUNCTION ID '(' parameters? ')' typeRule functionBlock;
parameters: parameter (',' parameter)*;
parameter: ID ':' typeRule;
functionBlock: '{' vars* statutes* returnRule'}';
returnRule: RETURN expression? SEMI;

block: '{' statutes* '}';

statutes: assignation
        | functionCall
        | methodCall
        | read
        | write
        | conditional
        | forLoop
        | whileLoop
        | expression;

assignation: ID '=' (expression SEMI| functionCall | methodCall);

functionCall: ID '(' arguments? ')' SEMI;
arguments: argument (',' argument)*;
argument: (ID | expression | functionCall | methodCall);

methodCall: ID '.' ID '(' arguments? ')' SEMI;

read: READ '(' ids ')' SEMI;

ids: id (',' id)*;
id: (ID | ID '.' ID);

write: WRITE '(' arguments ')' SEMI;

conditional: IF '(' expression ')' block (ELSE block)?;

forLoop: FOR ID '=' exp IN exp block;

whileLoop: WHILE '(' expression ')' block;

expression: exp (relop exp)*;

exp: term ((ADD | SUB) term)*;

term: factor ((MUL | DIV) factor)*;

factor: '(' expression ')'
      | (ADD | SUB)? varCte;

varCte: ID
      | cte_i
      | cte_f
      | cte_c
      | cte_s;

cte_i: INT;
cte_f: FLOAT;
cte_c: CHAR;
cte_s: '"' . '"'; // checar espacios en strings

main: MAIN '('')' functionBlock;

typeRule: INT_TYPE
    | FLOAT_TYPE
    | CHAR_TYPE
    | STRING_TYPE
    | VOID;

relop: GT
     | LT
     | EQ
     | NEQ;