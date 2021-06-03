// Bruen.g4
grammar Bruen;

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
INIT: 'init';
PRIVATE: 'private';
PUBLIC: 'public';
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
program: PROGRAM ID SEMI variableDeclaration* program2;
program2: classDef* functions* main EOF;

classDef: CLASS ID classBlock SEMI;
classBlock: '{' classAttributes classInit METHODS classMethod* '}';
classAttributes: ATTRIBUTES attributesDeclaration*;
classMethod: (PRIVATE | PUBLIC?) functions;
classInit: INIT LPAREN RPAREN '{' variableDeclaration* statutesNoReturn* '}';

variableDeclaration: varsDec | varsDecArray | varsDecMat;
varsDec: VAR ID ':' varsTypeInit SEMI;
varsDecArray: VAR ID ':' '['INT']'typeRule SEMI;
varsDecMat: VAR ID ':' '['INT']['INT']'typeRule SEMI;

attributesDeclaration: (PRIVATE | PUBLIC?) (attributesDec | varsDecArray | varsDecMat);
attributesDec: VAR ID ':' typeRule SEMI;

varsTypeInit: (typeRule | ID) varsTypeInit2?;
varsTypeInit2: (ASSIGN exp);

vars: ID ('.' ID)?;
varArray: ID ('.' ID)? '[' exp ']';
varMat: ID ('.' ID)? ('[' exp ']')? ('[' exp ']')?;

functions: FUNCTION ID LPAREN parameters? RPAREN (typeRule | VOID) functionBlock;
parameters: parameter (',' parameter)*;
parameter: ID ':' typeRule;
functionBlock: '{' variableDeclaration* statutes* '}';
returnRule: RETURN exp? SEMI;

block: '{' statutes* '}';

statutesNoReturn: assignation
        | read
        | write
        | conditional
        | forLoop
        | whileLoop
        | expression;

statutes: assignation
        | read
        | write
        | conditional
        | forLoop
        | whileLoop
        | expression
        | returnRule;

assignation: (vars | varArray | varMat) ASSIGN exp SEMI;

functionCall: ID LPAREN arguments? RPAREN;
arguments: exp arguments2;
arguments2: ',' arguments | ;

methodCall: ID '.' ID LPAREN arguments? RPAREN;

initCall: 'new' ID LPAREN RPAREN;

call: functionCall | methodCall | initCall;

read: READ LPAREN read2 RPAREN SEMI;
read2: (vars | varArray | varMat) read3;
read3: ',' read2 | ;

write: WRITE LPAREN w_arguments RPAREN SEMI;
w_arguments: exp w_arguments2;
w_arguments2: ',' w_arguments | ;

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
      | (vars | varArray | varMat)
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
mainBlock: '{' variableDeclaration* statutes* '}';

typeRule: INT_TYPE
    | FLOAT_TYPE
    | CHAR_TYPE
    | BOOL_TYPE;

relop: GT
     | LT
     | EQ
     | NEQ;