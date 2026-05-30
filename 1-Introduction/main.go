																								/*
Go (Golang) is a compiled programming language developed by Google 
for fast, simple, and scalable systems.
It’s known for built-in concurrency (goroutines/handling multiple tasks/request at same time) 
and easy deployment with single binaries.

Reasons to choose Go:
Fast Build time, Fast startup, compilation and execution.
Built-in support for concurrency with goroutines.
Simple syntax and easy maintenance.
Strong standard library for networking and web development.
Cross-platform and efficient memory management.

Why Golang is so fast
Go compiles directly to a single native binary (no VM, no runtime dependency),
so deployment is fast and simple.
Its compiler is very fast and enforces strict rules,
catching errors early and keeping builds consistent.
																								*/
package main 
																								/*
In Go, every file belongs to a package. 
A package is like a folder of related code 
package main is used to define the package package main 
specifically tells Go that this file
is part of the main executable program (not a reusable library).
main is a special package that tells Go this file is an executable 
program, not a library.
When you run go build or go run, Go looks for the main package and the 
main() function to start execution.
So basically, without package main, you cannot run the program as an executable.				*/

import "fmt" 																					/*

import is used to include a package in your program.
fmt is the standard Go package for formatted I/O (printing to console, reading input). 			*/

func main(){ // func main is the entry point of the program
	
	fmt.Print("Hello World") 
// Calls the Print function from the fmt package to display "Hello World" on the screen.

}																								/*

In Go, semicolons ; are optional at the end of a line 
because the Go compiler automatically inserts them during 
a phase called semicolon insertion.

before running the go program we need to built it
In order to built a program just do:
go build part_of_file
above command gives us and executable file(Machine code,with very small size)
to run the executable file just copy the path of file and paste it over terminal

There is another option to run file directly use command
go run path_of_file																				*/