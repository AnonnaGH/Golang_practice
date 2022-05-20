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


# Go Constants
If a variable should have a fixed value that cannot be changed, you can use the const keyword. <br>

The const keyword declares the variable as "constant", which means that it is unchangeable and read-only. <br>
![image](https://user-images.githubusercontent.com/44522472/157428124-f2b20e62-e3cb-45e9-acbc-67495213d49e.png)

### Note: The value of a constant must be assigned when you declare it.

# Declaring a Constant
### Here is an example of declaring a constant in Go:

``` golang
package main
import ("fmt")

const PI = 3.14

func main() {
  fmt.Println(PI)
}
```
# Constant Rules
Constant names follow the same naming rules as variables <br>
Constant names are usually written in uppercase letters (for easy identification and differentiation from variables) <br>
Constants can be declared both inside and outside of a function <br>


# Constant Types
## There are two types of constants:

### Typed constants
### Untyped constants


# Typed Constants
## Typed constants are declared with a defined type:

``` golang
package main
import ("fmt")

const A int = 1

func main() {
  fmt.Println(A)
}
```

# Untyped Constants
## Untyped constants are declared without a type:

``` golang 
package main
import ("fmt")

const A = 1

func main() {
  fmt.Println(A)
}
```
### Note: In this case, the type of the constant is inferred from the value (means the compiler decides the type of the constant, based on the value).

# Constants: Unchangeable and Read-only
## When a constant is declared, it is not possible to change the value later:
![image](https://user-images.githubusercontent.com/44522472/157436096-3672e95d-9ccc-4ca2-a2e7-176007763e1c.png)

# Multiple Constants Declaration
### Multiple constants can be grouped together into a block for readability:

``` golang
package main
import ("fmt")

const (
  A int = 1
  B = 3.14
  C = "Hi!"
)

func main() {
  fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)
}
```

# Go Output Functions
Go has three functions to output text: <br>
![image](https://user-images.githubusercontent.com/44522472/157436805-39199c5c-ada1-473c-a942-9705a97aeda1.png)

@ The Print() Function
The Print() function prints its arguments with their default format. <br>

### Print the values of i and j:
![image](https://user-images.githubusercontent.com/44522472/157437589-382b5b1b-8c27-41aa-94b5-d4c0b692e737.png)

# If we want to print the arguments in new lines, we need to use \n.

``` golang
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, "\n")
  fmt.Print(j, "\n")
}
```
![image](https://user-images.githubusercontent.com/44522472/157437832-14a89e57-b03e-4612-95e5-ac255e68c248.png)

# It is also possible to only use one Print() for printing multiple variables.

``` golang
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, "\n",j)
}
```
![image](https://user-images.githubusercontent.com/44522472/157438011-c688ac46-4115-41fc-8a45-5202a629a0a4.png)

If we want to add a space between string arguments, we need to use " ":

 we want to add a space between string arguments, we need to use " ":

``` golang
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, " ", j)
}
```
![image](https://user-images.githubusercontent.com/44522472/157438175-e7a10720-1931-4167-b748-bcdfa7336692.png)

# Print() inserts a space between the arguments if neither are strings:
![image](https://user-images.githubusercontent.com/44522472/157438329-26983b83-9e65-43ae-ae04-491f93c63d05.png)

# The Println() Function
## The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end:

``` golang
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Println(i,j)
}
```
![image](https://user-images.githubusercontent.com/44522472/157438480-86cf538e-9bea-46db-a336-b2d702ed0c3b.png)

# The Printf() Function
### The Printf() function first formats its argument based on the given formatting verb and then prints them.

### Here we will use two formatting verbs:

### %v is used to print the value of the arguments
### %T is used to print the type of the arguments

``` golang
package main
import ("fmt")

func main() {
  var i string = "Hello"
  var j int = 15

  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
}
```
![image](https://user-images.githubusercontent.com/44522472/157438781-d06992da-7481-4f79-9097-a859decead0c.png)

# Go Formatting Verbs

## Formatting Verbs for Printf()
Go offers several formatting verbs that can be used with the Printf() function. <br>


# General Formatting Verbs
The following verbs can be used with all data types: <br>
![image](https://user-images.githubusercontent.com/44522472/157441625-a31b9b7e-e32a-4676-9f8c-d173e569123a.png)

``` golang 
package main
import ("fmt")

func main() {
  var i = 15.5
  var txt = "Hello World!"

  fmt.Printf("%v\n", i)
  fmt.Printf("%#v\n", i)
  fmt.Printf("%v%%\n", i)
  fmt.Printf("%T\n", i)

  fmt.Printf("%v\n", txt)
  fmt.Printf("%#v\n", txt)
  fmt.Printf("%T\n", txt)
}
```
![image](https://user-images.githubusercontent.com/44522472/157441714-91acb6ec-768c-44ab-bb57-fb88684d70b7.png)

# Integer Formatting Verbs
### The following verbs can be used with the integer data type:

![image](https://user-images.githubusercontent.com/44522472/157442074-46c11860-d99b-4e9c-b922-b020cc59c973.png)

``` golang 
package main
import ("fmt")

func main() {
  var i = 15
 
  fmt.Printf("%b\n", i)
  fmt.Printf("%d\n", i)
  fmt.Printf("%+d\n", i)
  fmt.Printf("%o\n", i)
  fmt.Printf("%O\n", i)
  fmt.Printf("%x\n", i)
  fmt.Printf("%X\n", i)
  fmt.Printf("%#x\n", i)
  fmt.Printf("%4d\n", i)
  fmt.Printf("%-4d\n", i)
  fmt.Printf("%04d\n", i)
}
```
![image](https://user-images.githubusercontent.com/44522472/157442140-88c3e548-1ea9-41b8-bd08-fb53716d915d.png)


# String Formatting Verbs
### The following verbs can be used with the string data type:
![image](https://user-images.githubusercontent.com/44522472/157442267-e7b6d0d2-a4fe-4ef4-8c7a-3b11a7e6f021.png)

`` golang
package main
import ("fmt")

func main() {
  var txt = "Hello"
 
  fmt.Printf("%s\n", txt)
  fmt.Printf("%q\n", txt)
  fmt.Printf("%8s\n", txt)
  fmt.Printf("%-8s\n", txt)
  fmt.Printf("%x\n", txt)
  fmt.Printf("% x\n", txt)
}
```
![image](https://user-images.githubusercontent.com/44522472/157442364-ee1be2cc-b3c1-431d-aa34-730dd4eaf56e.png)

# Boolean Formatting Verbs
### The following verb can be used with the boolean data type:
![image](https://user-images.githubusercontent.com/44522472/157442435-d0dcf9e4-3a01-456b-908b-298ef100eda5.png)

``` golang 
package main
import ("fmt")

func main() {
  var i = true
  var j = false

  fmt.Printf("%t\n", i)
  fmt.Printf("%t\n", j)
}
```
![image](https://user-images.githubusercontent.com/44522472/157442476-d120d4f8-ce52-431f-bac1-2e590a3874ec.png)

# Float Formatting Verbs
### The following verbs can be used with the float data type:
![image](https://user-images.githubusercontent.com/44522472/157443828-5b9a2cc3-29cc-4fd7-a1a6-58dda688aae1.png)


# Go Data Types
Data type is an important concept in programming. Data type specifies the size and type of variable values. <br>

Go is statically typed, meaning that once a variable type is defined, it can only store data of that type. <br>


## Go has three basic data types:

## bool: represents a boolean value and is either true or false
## Numeric: represents integer types, floating point values, and complex types
## string: represents a string value

``` golang
package main
import ("fmt")

func main() {
  var a bool = true     // Boolean
  var b int = 5         // Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hi!"  // String

  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float:   ", c)
  fmt.Println("String:  ", d)
}
```

# Boolean Data Type
### A boolean data type is declared with the bool keyword and can only take the values true or false.

### The default value of a boolean data type is false.

```
package main
import ("fmt")

func main() {
  var b1 bool = true // typed declaration with initial value
  var b2 = true // untyped declaration with initial value
  var b3 bool // typed declaration without initial value
  b4 := true // untyped declaration with initial value

  fmt.Println(b1) // Returns true
  fmt.Println(b2) // Returns true
  fmt.Println(b3) // Returns false
  fmt.Println(b4) // Returns true
}
```
### Note: Boolean values are mostly used for conditional testing

# Go Integer Data Types
## Integer data types are used to store a whole number without decimals, like 35, -50, or 1345000.

### The integer data type has two categories:

<b>Signed integers: </b>  can store both positive and negative values <br>
<b>Unsigned integers: </b> - can only store non-negative values <br>

## Signed Integers
### Signed integers, declared with one of the int keywords, can store both positive and negative values:

## Go has five keywords/types of signed integers:
![image](https://user-images.githubusercontent.com/44522472/157640359-ae71eefc-93da-49e9-bff1-6c0e04d34edc.png)

# Unsigned Integers
## Unsigned integers, declared with one of the uint keywords, can only store non-negative values:

``` golang
package main
import ("fmt")

func main() {
  var x uint = 500
  var y uint = 4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}
```

# Go Arrays
### Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

# Declare an Array
### In Go, there are two ways to declare an array:

## 1. With the var keyword:
![image](https://user-images.githubusercontent.com/44522472/157693321-e13bcf84-3c20-4138-b2a2-6dd2507e3d36.png)

## 2. With the := sign:
![image](https://user-images.githubusercontent.com/44522472/157693568-bc5d2efe-94c1-4b90-bf38-2e5a6a66837d.png)

# Note: The length specifies the number of elements to store in the array. In Go, arrays have a fixed length. The length of the array is either defined by a number or is inferred (means that the compiler decides the length of the array, based on the number of values).

# Array Examples
# This example declares two arrays (arr1 and arr2) with defined lengths:

``` golang
package main
import ("fmt")

func main() {
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}
```
![image](https://user-images.githubusercontent.com/44522472/157693985-abf75548-074b-436d-9a7b-6f24291583e4.png)

``` golang 
package main
import ("fmt")

func main() {
  var arr1 = [...]int{1,2,3}
  arr2 := [...]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}
```
![image](https://user-images.githubusercontent.com/44522472/157696234-006d3bce-f324-4a04-a361-74da765f4c40.png)

``` golang
package main
import ("fmt")

func main() {
  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  fmt.Print(cars)
}
```
![image](https://user-images.githubusercontent.com/44522472/157696341-a420da17-8e36-4d45-95d0-7521d7ff7524.png)

# Access Elements of an Array
You can access a specific array element by referring to the index number. <br>

In Go, array indexes start at 0. That means that [0] is the first element, [1] is the second element, etc. <br>

This example shows how to access the first and third elements in the prices array: <br>

``` golang
package main
import ("fmt")

func main() {
  prices := [3]int{10,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[2])
}
```

![image](https://user-images.githubusercontent.com/44522472/157709176-84fecba4-ad2b-4a07-9dfd-735921fb59f7.png)

# Change Elements of an Array
## You can also change the value of a specific array element by referring to the index number.
``` golang
package main
import ("fmt")

func main() {
  prices := [3]int{10,20,30}

  prices[2] = 50
  fmt.Println(prices)
}
```
![image](https://user-images.githubusercontent.com/44522472/157709870-3c8a43c9-1fe5-438b-90b2-79b4444a7906.png)

# Change Elements of an Array
### You can also change the value of a specific array element by referring to the index number.
``` golang 
package main
import ("fmt")

func main() {
  prices := [3]int{10,20,30}

  prices[2] = 50
  fmt.Println(prices)
}
```
#Array Initialization
### If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.

### Tip: The default value for int is 0, and the default value for string is "".

``` golang
package main
import ("fmt")

func main() {
  arr1 := [5]int{} //not initialized
  arr2 := [5]int{1,2} //partially initialized
  arr3 := [5]int{1,2,3,4,5} //fully initialized

  fmt.Println(arr1)
  fmt.Println(arr2)
  fmt.Println(arr3)
}
```
# Initialize Only Specific Elements
## It is possible to initialize only specific elements in an array.

``` golang
package main
import ("fmt")

func main() {
  arr1 := [5]int{1:10,2:40}

  fmt.Println(arr1)
}
```
## Example Explained 
 The array above has 5 elements. <br>

1:10 means: assign 10 to array index 1 (second element). <br>
2:40 means: assign 40 to array index 2 (third element). <br>

# Find the Length of an Array
## The len() function is used to find the length of an array:

``` golang 
package main
import ("fmt")

func main() {
  arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  arr2 := [...]int{1,2,3,4,5,6}

  fmt.Println(len(arr1))
  fmt.Println(len(arr2))
}
```
![image](https://user-images.githubusercontent.com/44522472/157811215-57feea89-d78e-44a1-a36f-e4160db77f70.png)



# Go Slices
Slices are similar to arrays, but are more powerful and flexible. <br>

Like arrays, slices are also used to store multiple values of the same type in a single variable. <br>

However, unlike arrays, the length of a slice can grow and shrink as you see fit. <br>

In Go, there are several ways to create a slice: <br>

Using the []datatype{values} format <br>
Create a slice from an array <br>
Using the make() function <br>


# Create a Slice With []datatype{values}

![image](https://user-images.githubusercontent.com/44522472/157810510-8d4b3c5e-f7f0-4bce-ba55-50ff160f0f0b.png)

A common way of declaring a slice is like this: <br>
![image](https://user-images.githubusercontent.com/44522472/157810542-89ac964a-8ad0-4c70-a67b-89dfb97ba593.png)

The code above declares an empty slice of 0 length and 0 capacity. <br>

To initialize the slice during declaration, use this: <br>
![image](https://user-images.githubusercontent.com/44522472/157810652-4221f2f1-2731-4230-a0e3-1429a9bec813.png)
The code above declares a slice of integers of length 3 and also the capacity of 3.

### In Go, there are two functions that can be used to return the length and capacity of a slice:

<b> len() function - </b> returns the length of the slice (the number of elements in the slice)
<b> cap() function - </b> returns the capacity of the slice (the number of elements the slice can grow or shrink to)
``` golang
package main
import ("fmt")

func main() {
  myslice1 := []int{}
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)
}
```
![image](https://user-images.githubusercontent.com/44522472/157811092-61e216ef-e4d7-432b-9280-9fe88b5656c4.png)

# Create a Slice From an Array
### You can create a slice by slicing an array:

``` golang
var myarray = [length]datatype{values} // An array
myslice := myarray[start:end] // A slice made from the array
```
## This example shows how to create a slice from an array:
``` golang
package main
import ("fmt")

func main() {
  arr1 := [6]int{10, 11, 12, 13, 14,15}
  myslice := arr1[2:4]

  fmt.Printf("myslice = %v\n", myslice)
  fmt.Printf("length = %d\n", len(myslice))
  fmt.Printf("capacity = %d\n", cap(myslice))
}
```
![image](https://user-images.githubusercontent.com/44522472/158009548-acde0a38-fc6f-4093-b86e-5805bab7feff.png)


```
In the example above myslice is a slice with length 2. It is made from arr1 which is an array with length 6.

The slice starts from the second element of the array which has value 12. The slice can grow to the end of the array. This means that the capacity of the slice is 4.

If myslice started from element 0, the slice capacity would be 6.
```
# Create a Slice With The make() Function
### The make() function can also be used to create a slice.
```
slice_name := make([]type, length, capacity)
```
### Note: If the capacity parameter is not defined, it will be equal to length.

``` golang 
package main
import ("fmt")

func main() {
  myslice1 := make([]int, 5, 10)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  // with omitted capacity
  myslice2 := make([]int, 5)
  fmt.Printf("myslice2 = %v\n", myslice2)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n", cap(myslice2))
}
```
![image](https://user-images.githubusercontent.com/44522472/158010715-68c5ae00-b3eb-41b1-bed4-6d6bf4a0650f.png)

# Go Access, Change, Append and Copy Slices
Access Elements of a Slice <br>
You can access a specific slice element by referring to the index number.<br>

In Go, indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.<br>

``` golang
package main
import ("fmt")

func main() {
  prices := []int{10,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[2])
}
```
## Output
![image](https://user-images.githubusercontent.com/44522472/159422725-6454e0a2-a9d3-4010-a537-7671f04d4c8c.png)

# Change Elements of a Slice
You can also change a specific slice element by referring to the index number. <br>
``` golang
package main
import ("fmt")

func main() {
  prices := []int{10,20,30}
  prices[2] = 50
  fmt.Println(prices[0])
  fmt.Println(prices[2])
}
```
### output
![image](https://user-images.githubusercontent.com/44522472/159425201-ce95d12e-bc54-4650-9a74-f40dee5817dc.png)

# Append Elements To a Slice
You can append elements to the end of a slice using the append()function: <br>
## Syntax
![image](https://user-images.githubusercontent.com/44522472/159426129-bcf402a7-ebab-4153-a6b4-36b1755fe7b0.png)

``` golang
package main

import (
	"fmt"
)

func main() {

	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("s = %v\n", s)
	fmt.Printf("length = %d\n", len(s))
	fmt.Printf("capacity = %d\n", cap(s))

	s = append(s, 20, 21, 22, 23, 24, 25)
	fmt.Printf("s = %v\n", s)
	fmt.Printf("length = %d\n", len(s))
	fmt.Printf("capacity =%d\n", cap(s))

	s = append(s, 30, 36)
	fmt.Printf("s = %v\n", s)
	fmt.Printf("length = %d\n", len(s))
	fmt.Printf("capacity =%d\n", cap(s))

}
```
### output
![image](https://user-images.githubusercontent.com/44522472/159431217-d8c74d73-0060-4939-8dd8-d71d33005650.png)

# Go Operators

Conditional statements are used to perform different actions based on different conditions. <br>

Go Conditions <br>
A condition can be either true or false. <br>

Go supports the usual comparison operators from mathematics: <br>

Less than < <br>
Less than or equal <= <br>
Greater than > <br>
Greater than or equal >= <br>
Equal to == <br>
Not equal to != <br>
Additionally, Go supports the usual logical operators: <br>

Logical AND && <br>
Logical OR || <br>
Logical NOT ! <br>
You can use these operators or their combinations to create conditions for different decisions. <br>

![image](https://user-images.githubusercontent.com/44522472/159440942-bf6f6804-aaba-40c9-bb5b-7f08a6365a4f.png)

``` golang
package main
import ("fmt")

func main() {
  x := 10
  y := 5
  fmt.Println(x > y)
}
```
![image](https://user-images.githubusercontent.com/44522472/159441076-7b114cf8-a4f5-43da-8c37-9fbac86e4bc9.png)

``` golang
package main
import ("fmt")

func main() {
  x := 10
  y := 5
  fmt.Println(x != y)
}
```
![image](https://user-images.githubusercontent.com/44522472/159441791-bd83b04a-cded-4777-8cd9-0409051e45dc.png)

``` golang
package main
import ("fmt")

func main() {
  x := 10
  y := 5
  z := 2
  fmt.Println((x > y) && (y > z))
}
```
![image](https://user-images.githubusercontent.com/44522472/159441992-a48bd845-c158-451b-81f7-b05d1bd5059d.png)

``` golang
package main
import ("fmt")

func main() {
  x := 10
  y := 5
  z := false
  fmt.Println((x == y) || z)
}
```
![image](https://user-images.githubusercontent.com/44522472/159442099-1424a628-4452-4a79-9c8f-d94834d4cee4.png)

### Go has the following conditional statements:

Use if to specify a block of code to be executed, if a specified condition is true <br>
Use else to specify a block of code to be executed, if the same condition is false <br>
Use else if to specify a new condition to test, if the first condition is false <br>
Use switch to specify many alternative blocks of code to be executed <br>

# Go if statement
Use the if statement to specify a block of Go code to be executed if a condition is true. <br>
### Syntax
![image](https://user-images.githubusercontent.com/44522472/159612294-caa311cf-760b-4066-bd8f-fd19bf9ad356.png)

``` golang
package main
import ("fmt")

func main() {
  if 20 > 18 {
    fmt.Println("20 is greater than 18")
  }
}
```
![image](https://user-images.githubusercontent.com/44522472/159612387-d61fc960-d0ad-40d2-9d70-367fde93d755.png)

``` golang
package main
import ("fmt")

func main() {
  x := 20
  y := 18
  if x>y {
    fmt.Println("x is greater than y")
  }
}
```
![image](https://user-images.githubusercontent.com/44522472/159612528-18070a3f-8824-481e-8479-6464ab91b6c2.png)

## Example explained <br>
In the example above we use two variables, x and y, to test whether x is greater than y (using the > operator). As x is 20, and y is 18, and we know that 20 is <br> greater than 18, we print to the screen that "x is greater than y". <br>


# Go if else Statement
The else Statement <br>
Use the else statement to specify a block of code to be executed if the condition is false. <br>

### Syntax
![image](https://user-images.githubusercontent.com/44522472/159613292-81216980-5eae-4129-afe4-b39bf2ace694.png)

#### Using The if else Statement

## In this example, time (20) is greater than 18, so the if condition is false. Because of this, we move on to the else condition and print to the screen "Good evening". If the time was less than 18, the program would print "Good day":
``` golang
package main
import ("fmt")

func main() {
  time := 20
  if (time < 18) {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
}
```
![image](https://user-images.githubusercontent.com/44522472/159613503-209b2553-2ced-411d-862d-13eff9d95e6a.png)

## In this example, the temperature is 14 so the condition for if is false so the code line inside the else statement is executed:
``` golang
package main
import ("fmt")

func main() {
  temperature := 14
  if (temperature > 15) {
    fmt.Println("It is warm out there")
  } else {
    fmt.Println("It is cold out there")
  }
}
```
### The brackets in the else statement should be like } else {:

![image](https://user-images.githubusercontent.com/44522472/159613821-db763a11-f77b-4b26-be3b-869ce2958e00.png)

# Go else if Statement
Use the else if statement to specify a new condition if the first condition is false. <br>
### Syntax
![image](https://user-images.githubusercontent.com/44522472/159614422-f7580f1c-aa26-48f8-aa82-7aa2eecfeb9d.png)

# Using The else if Statement
``` golang
package main
import ("fmt")

func main() {
  time := 22
  if time < 10 {
    fmt.Println("Good morning.")
  } else if time < 20 {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
}
```
![image](https://user-images.githubusercontent.com/44522472/159614590-ec38579e-8caa-419b-9fcd-68a22290ed45.png)

## Example explained
In the example above, time (22) is greater than 10, so the first condition is false. The next condition, in the else if statement, is also false, so we move on to <br> else condition since condition1 and condition2 are both false - and print to the screen "Good evening". <br>

However, if the time was 14, our program would print "Good day." <br>
`` golang
package main
import ("fmt")

func main() {
  a := 14
  b := 14
  if a < b {
    fmt.Println("a is less than b.")
  } else if a > b {
    fmt.Println("a is more than b.")
  } else {
    fmt.Println("a and b are equal.")
  }
}
```
![image](https://user-images.githubusercontent.com/44522472/159615073-27ecab29-842d-4180-b106-9ea47bf70cc8.png)

## If conditions1 and condition2 are both correct, only the code for condition1 are executed:
``` golang
package main
import ("fmt")

func main() {
  x := 30
  if x >= 10 {
    fmt.Println("x is larger than or equal to 10.")
  } else if x > 20
    fmt.Println("x is larger than 20.")
  } else {
    fmt.Println("x is less than 10.")
  }
}
```
![image](https://user-images.githubusercontent.com/44522472/159615272-6b37ee01-35a8-41fb-a1e1-c9d79716d5d3.png)

# Go Nested if Statement
You can have if statements inside if statements, this is called a nested if. <br>

## Syntax
![image](https://user-images.githubusercontent.com/44522472/159616278-0f067a9c-ad0f-4874-a8fe-f3215a3d6e24.png)

``` golang
package main
import ("fmt")

func main() {
  num := 20
  if num >= 10 {
    fmt.Println("Num is more than 10.")
    if num > 15 {
      fmt.Println("Num is also more than 15.")
     }
  } else {
    fmt.Println("Num is less than 10.")
  }
}
```
![image](https://user-images.githubusercontent.com/44522472/159616348-bf1f0d9c-8ab7-4a6f-911c-efe139020ee9.png)

# Go switch Statement
Use the switch statement to select one of many code blocks to be executed. <br>

The switch statement in Go is similar to the ones in C, C++, Java, JavaScript, and PHP. The difference is that it only runs the matched case so it does not need a  <br> break statement. <br>
### Single-Case switch Syntax
![image](https://user-images.githubusercontent.com/44522472/159616661-f82aac37-5cf7-4225-93a8-611a00e303e4.png)

## This is how it works:

The expression is evaluated once <br>
The value of the switch expression is compared with the values of each case <br>
If there is a match, the associated block of code is executed <br>
The default keyword is optional. It specifies some code to run if there is no case match <br>

# Single-Case switch Example
### The example below uses a weekday number to calculate the weekday name:

``` golang
package main
import ("fmt")

func main() {
  day := 4

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  }
}
```
![image](https://user-images.githubusercontent.com/44522472/159616984-5b9f78c9-2917-4595-b684-6eefd0a771c3.png)

# The default Keyword
The default keyword specifies some code to run if there is no case match: <br>

``` golang
package main
import ("fmt")

func main() {
  day := 8

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  default:
    fmt.Println("Not a weekday")
  }
}
```
![image](https://user-images.githubusercontent.com/44522472/159617126-a0c80c5b-f7af-45c2-a252-2944b060b307.png)

### All the case values should have the same type as the switch expression. Otherwise, the compiler will raise an error:
![image](https://user-images.githubusercontent.com/44522472/159617189-e79ab0db-8c91-448e-8e5d-864e34d58475.png)


# The Multi-case switch Statement
It is possible to have multiple values for each case in the switch statement: <br>

## Syntax
![image](https://user-images.githubusercontent.com/44522472/159617345-1f5ffd13-fca6-422c-b27a-8d9e68a715e6.png)
#### Multi-case switch Example
The example below uses the weekday number to return different text: <br>

``` golang

package main
import ("fmt")

func main() {
   day := 5

   switch day {
   case 1,3,5:
    fmt.Println("Odd weekday")
   case 2,4:
     fmt.Println("Even weekday")
   case 6,7:
    fmt.Println("Weekend")
  default:
    fmt.Println("Invalid day of day number")
  }
}
```
# Go For Loops
The for loop loops through a block of code a specified number of times. <br>

The for loop is the only loop available in Go. <br>

# Go for Loop
Loops are handy if you want to run the same code over and over again, each time with a different value. <br>

Each execution of a loop is called an iteration. <br>

The for loop can take up to three statements:  <br>
### Syntax
![image](https://user-images.githubusercontent.com/44522472/159859039-7aa44603-2034-4c4f-acf4-4911cdf4b883.png)
#### statement1 Initializes the loop counter value.

#### statement2 Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends.

#### statement3 Increases the loop counter value.

#### Note: These statements don't need to be present as loops arguments. However, they need to be present in the code in some form.

# for Loop Examples
``` golang
package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}

}
```
![image](https://user-images.githubusercontent.com/44522472/159862612-a118a640-18bb-414c-a138-a2f99502e55b.png)

# Example 1 explained
i:=0; - Initialize the loop counter (i), and set the start value to 0 <br>
i < 5; - Continue the loop as long as i is less than 5 <br>
i++ - Increase the loop counter value by 1 for each iteration <br>

``` golang
package main
import ("fmt")

func main() {
  for i:=0; i <= 100; i+=10 {
    fmt.Println(i)
  }
}
```
![image](https://user-images.githubusercontent.com/44522472/159862873-dd8c1297-f05e-490f-b342-200a65870924.png)
# Example 2 explained
i:=0; - Initialize the loop counter (i), and set the start value to 0 <br>
i <= 100; - Continue the loop as long as i is less than or equal to 100 <br>
i+=10 - Increase the loop counter value by 10 for each iteration <br>

# The continue Statement
The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop. <br>

## Example
``` golang
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      continue
    }
    fmt.Println(i)
  }
}
```
![image](https://user-images.githubusercontent.com/44522472/159863307-5b68f548-f418-47e3-a6e9-0c9be9b1fc6a.png)

# The break Statement
The break statement is used to break/terminate the loop execution. <br>
This example breaks out of the loop when i is equal to 3: <br>
``` golang
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      break
    }
   fmt.Println(i)
  }
}
```
# Nested Loops
It is possible to place a loop inside another loop. <br>

Here, the "inner loop" will be executed one time for each iteration of the "outer loop":  <br>
``` golang
package main
import ("fmt")

func main() {
  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }
}
```

![image](https://user-images.githubusercontent.com/44522472/159902903-05d005f6-b701-4fa2-a772-4692f16fa66e.png)

# The Range Keyword
The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value. <br>

The range keyword is used like this: <br>

## Syntax
![image](https://user-images.githubusercontent.com/44522472/159903052-2b0584cf-cb74-45ca-b02b-27703e68c854.png)


### Example
This example uses range to iterate over an array and print both the indexes and the values at each (idx stores the index, val stores the value):  <br>

``` golang 
package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
     fmt.Printf("%v\t%v\n", idx, val)
  }
}
```
![image](https://user-images.githubusercontent.com/44522472/159903573-d6461d1c-30a5-478b-b988-7b42a2e4606b.png)

## To only show the value or the index, you can omit the other output using an underscore (_).
### Here, we want to omit the indexes (idx stores the index, val stores the value):
``` golang
package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for _, val := range fruits {
     fmt.Printf("%v\n", val)
  }
}
```
![image](https://user-images.githubusercontent.com/44522472/159903782-b454087a-5d6d-476d-9c8b-f6368d6ec80b.png)

Here, we want to omit the values (idx stores the index, val stores the value): <br>

``` golang
package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}

  for idx, _ := range fruits {
     fmt.Printf("%v\n", idx)
  }
```
![image](https://user-images.githubusercontent.com/44522472/159903928-f6721126-d5e6-4c1e-b4f4-be7e831dfd74.png)

# Go Functions

A function is a block of statements that can be used repeatedly in a program. <br>

A function will not execute automatically when a page loads. <br>

A function will be executed by a call to the function.  <br>


# Create a Function
To create (often referred to as declare) a function, do the following:  <br>

## Use the func keyword.

Specify a name for the function, followed by parentheses ().  <br>
Finally, add code that defines what the function should do, inside curly braces {}.  <br>
![image](https://user-images.githubusercontent.com/44522472/169294459-86a8d57a-bfaf-4236-b384-d10d47e3c10e.png)  <br>

# Call a Function
Functions are not executed immediately. They are "saved for later use", and will be executed when they are called. <br>

In the example below, we create a function named "myMessage()". The opening curly brace ( { ) indicates the beginning of the function code, and the closing curly brace ( } ) indicates the end of the function. The function outputs "I just got executed!". To call the function, just write its name followed by two parentheses ():
 <br>
 ## Example
 
 ``` golang
 package main
import ("fmt")

func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage() // call the function
}
```

# Result:
![image](https://user-images.githubusercontent.com/44522472/169295103-ad0a4774-5722-434c-a596-d19aeb44b2ec.png)

## A function can be called multiple times.

``` golang
package main
import ("fmt")

func FunctionName() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage()
  myMessage()
  myMessage()
}
```

## Result:
![image](https://user-images.githubusercontent.com/44522472/169295391-8ea2c859-a448-46fe-9390-48ba89dc1f42.png)

# Go Function Parameters and Arguments
## Parameters and Arguments
Information can be passed to functions as a parameter. Parameters act as variables inside the function.<br>

Parameters and their types are specified after the function name, inside the parentheses. You can add as many parameters as you want, just separate them with a comma: <br>

### Syntax
![image](https://user-images.githubusercontent.com/44522472/169484734-b30c2dff-2acf-47bf-b65e-bc7b0998a67a.png)

# Function With Parameter Example
The following example has a function with one parameter (fname) of type string. When the familyName() function is called, we also pass along a name (e.g. Liam), and the name is used inside the function, which outputs several different first names, but an equal last name: <br>

## Example
``` golang
package main
import ("fmt")

func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

func main() {
  familyName("Liam")
  familyName("Jenny")
  familyName("Anja")
}
```
# result
![image](https://user-images.githubusercontent.com/44522472/169485431-da0db327-9813-4bce-848b-16bb307cbbf1.png)

# Multiple Parameters
Inside the function, you can add as many parameters as you want: <br>

### Example
``` golang
package main
import ("fmt")

func familyName(fname string, age int) {
  fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func main() {
  familyName("Liam", 3)
  familyName("Jenny", 14)
  familyName("Anja", 30)
}
```
## Result:
![image](https://user-images.githubusercontent.com/44522472/169486397-d1d377a1-83c2-4e17-83d0-3f8bc6076b4f.png)
### Note: When you are working with multiple parameters, the function call must have the same number of arguments as there are parameters, and the arguments must be passed in the same order.

# Go Function Returns
## Return Values
If you want the function to return a value, you need to define the data type of the return value (such as int, string, etc), and also use the return keyword inside the function: <br>
### Syntax
![image](https://user-images.githubusercontent.com/44522472/169487702-d93d3d39-55e8-4bca-949f-30f70001ff1e.png)

# Function Return Example
Here, myFunction() receives two integers (x and y) and returns their addition (x + y) as integer (int): <br>

``` golang
package main
import ("fmt")

func myFunction(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(myFunction(1, 2))
}
```
### Result:
![image](https://user-images.githubusercontent.com/44522472/169488037-70eca03b-9762-46ce-9f47-f3c55d128197.png)

# Named Return Values
In Go, you can name the return values of a function. <br>

## Example
Here, we name the return value as result (of type int), and return the value with a naked return (means that we use the return statement without specifying the variable name): <br>
``` golang
package main
import ("fmt")

func myFunction(x int, y int) (result int) {
  result = x + y
  return
}

func main() {
  fmt.Println(myFunction(1, 2))
}
```
# Result:
![image](https://user-images.githubusercontent.com/44522472/169488815-c4404ceb-ea11-4ccb-8ae4-12868425660b.png)

The example above can also be written like this. Here, the return statement specifies the variable name: <br>

``` golang
package main
import ("fmt")

func myFunction(x int, y int) (result int) {
  result = x + y
  return result
}

func main() {
  fmt.Println(myFunction(1, 2))
}
```
# Store the Return Value in a Variable
You can also store the return value in a variable, like this: <br>
### Example
Here, we store the return value in a variable called total:  <br>

``` golang 
package main
import ("fmt")

func myFunction(x int, y int) (result int) {
  result = x + y
  return
}

func main() {
  total := myFunction(1, 2)
  fmt.Println(total)
}
```
# Multiple Return Values
Go functions can also return multiple values. <br>
# Example
Here, myFunction() returns one integer (result) and one string (txt1): <br>

``` golang
package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
  fmt.Println(myFunction(5, "Hello"))
}
```
## Result:
![image](https://user-images.githubusercontent.com/44522472/169495135-db2e5c29-995a-442a-a83c-478ae7781260.png)

## Example
Here, we store the two return values into two variables (a and b):
``` golang
package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
  a, b := myFunction(5, "Hello")
  fmt.Println(a, b)
}
```
## Result:
![image](https://user-images.githubusercontent.com/44522472/169495357-60bce139-b656-42c9-81aa-c94db6d31d0d.png)

If we (for some reason) do not want to use some of the returned values, we can add an underscore (_), to omit this value. <br>
## Example
Here, we want to omit the first returned value (result - which is stored in variable a): <br>

``` golang
package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
   _, b := myFunction(5, "Hello")
  fmt.Println(b)
}
```
## Result:
![image](https://user-images.githubusercontent.com/44522472/169495736-529bb0e1-178a-496a-a323-c8acc2325f8d.png)

## Example
Here, we want to omit the second returned value (txt1 - which is stored in variable b):  <br>
``` golang
package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
   a, _ := myFunction(5, "Hello")
  fmt.Println(a)
}
```
## Result:
![image](https://user-images.githubusercontent.com/44522472/169495904-cb12a921-2cd6-4825-89da-bbf8f7b3952a.png)

# Go Recursion Functions

