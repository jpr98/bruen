Bruen programming language
===

Bruen is a small object oriented programming language that takes inspiration from the C-family. 

## Table of Contents

[TOC]

## Beginners Guide

If this is your first time using Bruen, start here!

Installation
---
To be able to compile and execute *.br* files you need to download the Bruen compiler and virtual machine. We created a no hassle download for the most popular operating systems.

[MacOS](https://github.com/jpr98/bruen/raw/main/build/macos/bruen)

[Windows 64-bit](https://github.com/jpr98/bruen/raw/main/build/windows/bruen.exe)

[Linux x86-64](https://github.com/jpr98/bruen/raw/main/build/linux/bruen)

If you are using a different OS you can download Go, the Bruen compiler source code ([link](https://github.com/jpr98/bruen)) and build it for your architecture.


Usage
---
Once you have bruen installed compiling and executing a program is really easy.

Compilation:
```
bruen file-to-path.br b
```
Execution:
```
bruen file-to-path.br c
```
In execution the compiler will take the object file with the same name of the input file generated during compilation. Bruen compiler generates object files as such: *name.obj*.

If you want to compile and execute a program in one step, Bruen's got you covered.

Compilation + Execution:
```
bruen file-to-path.br r
```

**Note:**
It is probable that you will have to use your system's syntax (i.e. `./bruen`) to execute the compiler. If you don't wish to do this, you can install the compiler and be done with it.

If you built the compiler from source (or have Go installed), you can use: `go install github.com/jpr98/bruen`.

Program structure
---
All Bruen programs have a defined structure.

They all begin with a `program` keyword followed by the program name.

Then global variables are declared, followed by classes and then functions.

Finally there's `main` the entry point into every Bruen program.

```=1
program tutorial;

var globalInt: int;

class NumPrinter {
    attributes
    
    init() {
    }
    
    methods
    function print(num: int) void {
        write(num);
    }
};

function getNum() int {
    var num: 22
    return num;
}

main() {
    var printer: NumPrinter = new NumPrinter();
    numPrinter.print(getNum());
}
```


Variables
---
Variable declaration always happens at the top of the current scope. The syntax for declaring variables is the following:
```=
var number: int;
var character: char;
var fNumber: float;
var cool: bool = true;
```

That's right. You can declare a variable and initialize it immediately! Also by know you are probably wondering what are the primitive types supported by Bruen. As of version 1.0 Bruen supports:
- int
- float
- char
- bool

These primitive types are also available in their array forms. Arrays need to have a fixed size at compile time. So you can have this:
```=
var arr: [5]int;
var status: [5]bool;
```
There is no way of initializing all values of an array at once, so you will have to do it manually (although this may become a thing in the future :eyes:)

Functions
---
Once you are done delcaring global variables you are able to start writing functions. In Bruen functions can only take primitive types as parameters. The same applies for returns, except returns can also be void (this needs to be indicated by you).
```=
function add(lhs: int, rhs: int) int {
    var result: int;
    result = lhs + rhs;
    return result;
}

function addToGlobalNumber(a: int) void {
    number + a;
}
```
As you can see, you can declare variables inside of functions. These variables will only be available for the duration of the function. If you decide to name a local variable the same as a global variable, you won't have access to the global variable.

Expressions
---
As every programming lenguage does, Bruen can perform arithmetic, relational and logical operations, as well as assignation. Here are some examples:
```=
a = a + b;
b = (a - c) * b;
c = 10 / a + 1;

greater = a > b;
notEqual = c == b;

oneOrTheOther = a > b || b > c;
oneAndTheOther = a > b && b > c;
```
This operations take into account the types of the variables, any incompatibility issues will be reported during compilation.

Function calls are also an expression.
```=
var a: int = add(b, c);
addToGlobalNomber(10);
```

Conditional Logic
---
Bruen has if statements to handle branching on different code paths. To use an if statement you do the following:
```=
if (a == 1) {
    write(1);
} else {
    write(a);
}
```

Repetition
---
Bruen also has a looping mechanism. First we have the `while` loop. This loop keeps looping until the condition is not met.
```=
var counter: int = 0;
while(counter < 10) {
    write(counter);
    counter = counter + 1;
}
```

For loops are a tad more complex. At the start of a `for` loop you define an int variable and assign it an initial value, then you set the number it needs to reach to exit the loop. Each iteration the variable increases by 1.
```=
for i = 0 in 10 {
    write(i);
}
```

Input and Output
---
All programs need a way to interact with a user. In Bruen you can read and write to standard output (i.e. the console) with two simple built-in functions. No really, they are simple.

```=
var newNumber: int;
read(newNumber);

var boolean: bool = false;
write(10, boolean);
```

Read expects to receive as input a value that is compatible with the type of the variable it is reading into.

Write prints the value of the given variables, with no spaces. After each call to write a new line is printed.

Both functions accept as many variables or values as you want.

Classes
---
Since Bruen is an object-oriented programming language it obviously has support for Classes. Here is an example of how to define a class in Bruen. Remember the program structure, classes go after global variables and before functions.

```=
class Car {
    attributes
    private var tirePressure: [4]int;
    public var serialNumber: int;
    var fast: bool;
    
    init() {
        serialNumber = 2345632;
        fast = true;
    }
    
    methods
    function setTirePressure(tire: int, tirePressure: int) {
        if (tire > 3) {
            return;
        }
        self.tirePressure[tire] = tirePressure;
    }
    
    function getTirePressure(tire: int) int {
        return tirePressure[tire];
    }
};
```

**Sections**

A class has two sections, `attributes` and `methods`. You need to follow that order and write each tag before starting to code the section.

**Initializer**

Every class has a special function called `init`, its declaration goes after the attributes section. The initializer doesn't receive any parameters (for now... :eyes:) and doesn't need the `function` keyword. The init function always returns a new instance of the class.

**Access modifiers**

Everything inside a class is public by default. If you want to specify that it is public you can write `public` before the `var` or `function` keyword. If you want to make something not accessible outside the class you can add the `private`keyword.

**Self**

The keyword `self` is a special variable that exists inside of every class method (and init). It makes reference to the current instance of the class. So for example, if you have a parameter named the same as one of your attributes, you can use `self.attribute` to get access to your attribute. 

So now you have your class defined, how do you use it? To declare an instance of your class you simply set the type of a variable to the name of your class. Once you have your variable declared you can call the initializer by using the `new` keyword. After that you can call methods on the instance.

```=
var myCar: Car;
myCar = new Car();

myCar.setTirePressure(0, 35);
write(myCar.getTirePressure(0));
```