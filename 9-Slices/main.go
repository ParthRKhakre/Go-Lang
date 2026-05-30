package main 

import "fmt"

func main(){

	// slice are flexible, dynamically sized view of arrays 
	// slices can grow and shrink, they are most used contruct in go
	// also have useful methods

	// uninitialzed slice is nil
	var num []int
	fmt.Println(num)
	fmt.Println(num == nil) // nil is similar to null in other languages
	fmt.Println(len(num))

	// To avoid the creation of nil value slices 
	var x = make([]int,2)
	fmt.Println("made using make()",x)
	fmt.Println("length of slice ",len(x))
	fmt.Println("capacity of slice ",cap(x))
	// make() function is used to create and initialize slices, maps, and channels.
	// make(container,length,capacity)
	// e.g. s := make([]int, 5, 10) // slice of length=5, capacity=10
	
	// Length len() → current number of elements in the slice (how many elements are actually present).
	// Capacity cap() → total space allocated in the underlying array starting from the first element of the slice (how many elements it can hold before needing to allocate more memory).

	// append() is a built-in function used to add elements to a slice.
	x = append(x, 9)
	x = append(x, 10)
	x = append(x, 11)
	x = append(x, 12)
	x = append(x, 13)
	x = append(x, 14)
	x = append(x, 15)
	fmt.Println(x,cap(x))
	// Note: In Go, when a slice exceeds its capacity, Go automatically allocates a new underlying array (usually double the size) and copies the elements.

	// Another way to create slice
	nums := []int{}
	fmt.Println(nums)

	// slices are accessed/updated using indices, just like arrays.
	y := make([]int,2,5)
	y[0] = 10
	fmt.Println(y)

	// copy()
	var a = make([]int,3)
	a = append(a, 10)
	var b = make([]int,len(a))
	fmt.Println(a,b)




}