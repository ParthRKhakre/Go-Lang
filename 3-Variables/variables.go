package main

import "fmt"

func main() {
// If you declare a variable but don’t use it, the compiler will throw an error.
	
// Ways to declare a variable

// 1.Traditional useful only when you want to declare a variable and dont want to assign value. the value will be assigned further
	var name string = "golang"
	fmt.Println(name)
// 2.infer - type inference means the compiler can automatically determine the type of a variable based on the value you assign.
	var language = true
	fmt.Println(language)
// 3.short-hand syntax
	age := 20
	fmt.Println(age)

	var price float32 
	price = 99.23
	fmt.Println(price)

// := can be used inside the main function using it outside is not valid
// In Go, if you declare a variable outside of any function, it becomes package-level (global) variable.

}