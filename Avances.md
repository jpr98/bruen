# Proyecto Compiladores

### Avance 1: Lexico / Sintaxis
(*9 de abril de 2021*)

En este avance definimos la sintaxis inicial de nuestro lenguaje de programación. Creamos los tokens y las reglas de gramática para generar el lexer y el parser correspondiente.

### Avance 2: Directorio de Funciones y Tablas de Variables
(*16 de abril de 2021*)

En este avance creamos una interfaz que escucha los eventos de entrada y salida de las reglas de gramática para ir construyendo el directorio de funciones y sus respectivas tablas de variables. Esto lo hacemos utilizando mapas en Go.

### Avance 3: Generación de código intermedio inicial
(*23 de abril de 2021*)

En este avance creamos los tipos para manejar los cuadruplos en el código y generamos las instrucciones intermedias para las expresiones artiméticas. También creamos el cubo semántico en el código.

### Avance 4: Generación de código de Estatutos Condicionales : Decisiones / Ciclos
(*01 de mayo de 2021*)

En este avance desarrollamos el código que genera los cuádruplos de los estatutos no lineales: if, if-else, while loop y for loop.

### Avance 5: Soporte de constantes en quads y validación de tipo de variables en expresiones.
(*08 de mayo de 2021*)

En este avance aprovechamos para completar algunos TODOs que teníamos pendientes como el manejo de constantes en quads, el fondo falso para las expresiones y la validación de tipos de las variables en ellas.

### Avance 6: Funciones y Ejecución de Expresiones
(*15 de mayo de 2021*)

En este avance se desarrolló la generación de quads para las funciones y llamadas a funciones. Además se realizó el mapa de memoria para la máquina virtual.

### Avance 7: Ejecución de Estatutos Condicionales
(*21 de mayo de 2021*)

En este avance se desarrolló la maquina virtual donde actualmente se ejecutan los quads con operadores aritméticos (ADD,SUB,MUL,DIV), operadores relacionales (GT,LT,EQ,NEQ) y operadores lógicos (AND,OR). Asi mismo se ejecutan ya los estatutos condicionales. También ejecutan todos los quads relacionados a funciones sin embargo, todavía no se maneja recursividad. Finalmente se implemento en la máquina virtual el read y write.

