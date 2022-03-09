# Golang_practice

Go is a popular programming language.<br>
Go is used to create computer programs.<br>

# Go Introduction
## What is Go? <br>
Go is a cross-platform, open source programming language <br>
Go can be used to create high-performance applications <br>
Go is a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language <br>
Go was developed at Google by Robert Griesemer, Rob Pike, and Ken Thompson in 2007 <br>
Go's syntax is similar to C++ <br>

# What is Go Used For?
Web development (server-side)  <br>
Developing network-based programs  <br>
Developing cross-platform enterprise applications <br>

Cloud-native development <br>

# Go Compared to Python and C++
![image](https://user-images.githubusercontent.com/44522472/156924579-edffc322-f670-4e19-b990-2c8b7bef5aa3.png)

# Notes:

## Compilation time refers to translating the code into an executable program
## Concurrency is performing multiple things out-of-order, or at the same time, without affecting the final outcome
## Statically typed means that the variable types are known at compile time

# Go Compact Code
You can write more compact code, like shown below (this is not recommended because it makes the code more difficult to read): <br>
![image](https://user-images.githubusercontent.com/44522472/156932021-b97d408d-1c43-40f4-87fc-7454290ab227.png)


# Go Syntax
## A Go file consists of the following parts:

### Package declaration
### Import packages
### Functions
### Statements and expressions

# Look at the following code, to understand it better:

## Example
package main <br>
import ("fmt") <br>

func main() { <br>
  fmt.Println("Hello World!") <br>
} <br>


# Example explained
### Line 1: In Go, every program is part of a package. We define this using the package keyword. In this example, the program belongs to the main package.

### Line 2: import ("fmt") lets us import files included in the fmt package.

### Line 3: A blank line. Go ignores white space. Having white spaces in code makes it more readable.

### Line 4: func main() {} is a function. Any code inside its curly brackets {} will be executed.

### Line 5: fmt.Println() is a function made available from the fmt package. It is used to output/print text. In our example it will output "Hello World!".


# Note:
 In Go, any executable code belongs to the main package.
 
 # Go Comments
### A comment is a text that is ignored upon execution.
### Comments can be used to explain the code, and to make it more readable.
### Comments can also be used to prevent code execution when testing an alternative code.
### Go supports single-line or multi-line comments.
### Go Single-line Comments
### Single-line comments start with two forward slashes (//).
### Any text between // and the end of the line is ignored by the compiler (will not be executed).

## Go Multi-line Comments
### Multi-line comments start with /* and ends with */.


# Comment to Prevent Code Execution
You can also use comments to prevent the code from being executed. <br>

The commented code can be saved for later reference and troubleshooting. <br>

# Go Variable Types
## In Go, there are different types of variables, for example:

### int- stores integers (whole numbers), such as 123 or -123
### float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
### string - stores text, such as "Hello World". String values are surrounded by double quotes
### bool- stores values with two states: true or false

# Go Variable Types
## In Go, there are different types of variables, for example:

int- stores integers (whole numbers), such as 123 or -123 <br>
float32- stores floating point numbers, with decimals, such as 19.99 or -19.99 <br>
string - stores text, such as "Hello World". String values are surrounded by double quotes <br>
bool- stores values with two states: true or false <br>

# Declaring (Creating) Variables
## In Go, there are two ways to declare a variable:

### 1. With the var keyword:
Use the var keyword, followed by variable name and type:<br>
![image](https://user-images.githubusercontent.com/44522472/156932913-3d03dc55-6521-4cc6-b91c-fec11e8f4b9d.png)

### Note: You always have to specify either type or value (or both).

### 2. With the := sign:
Use the := sign, followed by the variable value: <br>
![image](https://user-images.githubusercontent.com/44522472/156933014-7930bae0-93f8-4c62-ad2d-b6d5bc965b3d.png)

Note: In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).<br>

Note: It is not possible to declare a variable using :=, without assigning a value to it.<br>


# Variable Declaration With Initial Value
If the value of a variable is known from the start, you can declare the variable and assign a value to it on one line:
![image](https://user-images.githubusercontent.com/44522472/156933264-68a7ffd7-632a-4b35-a408-16ae2b4269a4.png)

# Variable Declaration Without Initial Value
In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type: <br>
![image](https://user-images.githubusercontent.com/44522472/156933830-6c088694-9c0c-4f6d-a0b6-8c5e5b259fae.png)

### Example explained
### In this example there are 3 variables:

a <br>
b <br>
c <br>
### These variables are declared but they have not been assigned initial values.

### By running the code, we can see that they already have the default values of their respective types:

a is "" <br>
b is 0 <br>
c is false <br>

# Value Assignment After Declaration
### It is possible to assign a value to a variable after it is declared. This is helpful for cases the value is not initially known.
![image](https://user-images.githubusercontent.com/44522472/156933944-b7577423-af01-4e7f-b845-7e95e8a868e9.png)

Note: It is not possible to declare a variable using ":=" without assigning a value to it. <br>

# Difference Between var and :=
# There are some small differences between the var var := 

![image](https://user-images.githubusercontent.com/44522472/156934019-f3ba87a7-7a41-4e7b-bc36-83a6071bfb22.png)

# Example
This example shows declaring variables outside of a function, with the var keyword: <br>
![image](https://user-images.githubusercontent.com/44522472/156934242-a855bd80-911d-4afa-9600-b8b52b328a07.png)


![image](https://user-images.githubusercontent.com/44522472/156934792-d3c7eb37-2716-435b-90de-7a00468c4181.png)


# Go Multiple Variable Declaration
### In Go, it is possible to declare multiple variables in the same line.

``` golang
package main
import ("fmt")

func main() {
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}
```

### Note: If you use the type keyword, it is only possible to declare one type of variable per line.


# If the type keyword is not specified, you can declare different types of variables in the same line:

``` golang
package main
import ("fmt")

func main() {
  var a, b = 6, "Hello"
  c, d := 7, "World!"

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}
```

# Go Variable Declaration in a Block

## Multiple variable declarations can also be grouped together into a block for greater readability:

``` golang
package main
import ("fmt")

func main() {
   var (
     a int
     b int = 1
     c string = "hello"
   )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```


# Go Variable Naming Rules
## A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

# Go variable naming rules:

A variable name must start with a letter or an underscore character (_) <br>
A variable name cannot start with a digit <br>
A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ ) <br>
Variable names are case-sensitive (age, Age and AGE are three different variables) <br>
There is no limit on the length of the variable name<br>
A variable name cannot contain spaces  <br>
The variable name cannot be any Go keywords <br>


# Multi-Word Variable Names
Variable names with more than one word can be difficult to read. <br>

There are several techniques you can use to make them more readable: <br>

# Camel Case
Each word, except the first, starts with a capital letter: <br>

myVariableName = "John"  <br>

# Pascal Case
Each word starts with a capital letter: <br>

MyVariableName = "John" <br>

# Snake Case
Each word is separated by an underscore character: <br>

my_variable_name = "John" <br>
