package main

import "fmt"

func main() {

	// for loop is only contruct in go for looping

	// Implementing "while" using for
	i := 1
	for i < 4 {
		fmt.Println(i)
		i++ // i = i + 1
	}

	fmt.Println("-----------------")

	// we use while for implementing infinite loop with for it can be implemented as
	// infinite loop
	for{
		println("1")
		break
	}
	
	fmt.Println("-----------------")

	// classic for loop
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----------------")

	// break - exits the current loop immediately.
	// continue - skips the rest of the current iteration and moves to the next loop cycle.

	// range -  is a built-in loop construct
	// is used to iterate over elements in arrays, slices, strings, maps, or channels.
	for i := range 11{
		fmt.Println(i)
	} 
}
