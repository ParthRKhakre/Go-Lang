package main 

import "fmt"

func main(){

// constants are fixed values that cannot change once declared.	
// Declared using the const keyword.Must be assigned a value at declaration.

	const name string = "Larry"
	fmt.Println(name)

	const PI = "3.14"
	fmt.Println(PI)

// const does NOT use :=.
// := is only for short variable declaration, not constants.

// Constant grouping
const (
	port = 5000
	host = "localhost"
)

	fmt.Println(port,host)
}