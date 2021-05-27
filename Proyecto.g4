// Proyecto.g4
grammar Proyecto;

// Tokens
INT: '-'? [0-9]+;
FLOAT: INT ('.' [0-9]+)?;
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
ASSIGN: '=';
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
BOOL_TYPE: 'bool';
VOID: 'void';
// Program
PROGRAM: 'program';

ID: [a-zA-Z_][a-zA-Z0-9]*;

// Rules
program: PROGRAM ID SEMI classDef* varsDec* program2;
program2: functions* main EOF;

classDef: CLASS ID ('<' ID '>')? classBlock SEMI;
classBlock: '{' ATTRIBUTES varsDec* METHODS functions* '}';

varsDec: VAR ID ':' varsTypeInit SEMI
    | VAR ID '['INT']' ':' varsTypeInit SEMI
    | VAR ID '['INT']['INT']' ':' varsTypeInit SEMI;
varsTypeInit: (typeRule | ID) varsTypeInit2?;
varsTypeInit2: (ASSIGN exp);

vars: ID ('.' ID)? ('[' exp ']')? ('[' exp ']')?;

functions: FUNCTION ID LPAREN parameters? RPAREN (typeRule | VOID) functionBlock;
parameters: parameter (',' parameter)*;
parameter: ID ':' typeRule;
functionBlock: '{' varsDec* statutes* '}';
returnRule: RETURN exp? SEMI;

block: '{' statutes* '}';

statutes: assignation
        | read
        | write
        | conditional
        | forLoop
        | whileLoop
        | expression
        | returnRule;

assignation: ID ASSIGN exp SEMI;

functionCall: ID LPAREN arguments? RPAREN;
arguments: exp arguments2;
arguments2: ',' arguments | ;

methodCall: ID '.' ID LPAREN arguments? RPAREN;

call: functionCall | methodCall;

read: READ LPAREN read2 RPAREN SEMI;
read2: vars read3;
read3: ',' read2 | ;

write: WRITE LPAREN arguments RPAREN SEMI;

conditional: IF conditional2 conditional3 conditional4?;
conditional2: LPAREN exp RPAREN;
conditional3: block;
conditional4: ELSE block;

forLoop: FOR forLoop2 IN forLoop3 block;
forLoop2: ID ASSIGN exp;
forLoop3: exp;

whileLoop: WHILE whileLoop2 block;
whileLoop2: LPAREN exp RPAREN;

expression: exp SEMI;
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

factor: factor2
      | varCte
      | vars
      | call;
      
factor2: LPAREN exp RPAREN;

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

main: MAIN LPAREN RPAREN mainBlock;
mainBlock: '{' varsDec* statutes* '}';

typeRule: INT_TYPE
    | FLOAT_TYPE
    | CHAR_TYPE
    | BOOL_TYPE;

relop: GT
     | LT
     | EQ
     | NEQ;