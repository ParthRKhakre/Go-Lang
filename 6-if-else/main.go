package main

import "fmt"


func main() {

	age := 13

	// control statement : if-else
	if age > 18 {
		fmt.Print("person is an adult")
	}else if age >= 12 {
		fmt.Println("person is teenager")
	}else{
		fmt.Println("person is not adult")
	}

// 	In an if–else if–else control statement, once a condition evaluates to true, 
// its block executes and all remaining conditions are automatically skipped.

	fmt.Println("---------------------------------")

	var role = "admin"
	var hasPermission = false

	// OR operator
	if role == "admin" || hasPermission{
		fmt.Println("yes")
	}

	// AND operator
	if role == "admin" || hasPermission{
		fmt.Println("no")
	}

	// we can declare variable inside if construct
	if age := 15; age > 18{
		fmt.Println("person is an adult ",age)
	} else if age >= 12{
		fmt.Println("person is an adult ",age)
	}
	// age := 15 is declared and initialized only for this if–else block.
	// age exists only inside the if, else if, and else.

	// Note GO doesnt have a ternary operator -- it must be implemented using a if-else statement
}