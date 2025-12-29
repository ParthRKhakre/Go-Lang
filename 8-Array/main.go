package main

import "fmt"

func main() {

// 	arrays are fixed-size sequences of elements of the same type.
	
// Ways to declare array
	var a[5] int                  // default values = 0
	b := [3] int{1, 2, 3}          // with values
	c := [...] int{4, 5, 6, 7}     // size inferred means the compiler automatically determines the array’s size from the number of elements you provide.
	
	fmt.Println(a,b,c)

// len() - is a built-in function that returns the number of elements in a collection.
	var nums [4]int
	fmt.Println("array length",len(nums)) 

//  Can access or update array elements using their index.
	nums[0] = 9 
	fmt.Println(nums)

//  Normal Iteration
	for i := 0;i<len(nums);i++{
		fmt.Print(i," ")
	}

//  Iteration using range
	fmt.Println("nums"+" "+"value")
	for index,value := range nums{
		fmt.Println(index,"  ",value)
	}

//  _ use in iteration
	for _,i := range nums{  // _ can be used to ignore index or value:
		fmt.Println(" ",i)
	}

	fmt.Println("-------------------------")

//  This loop only prints the index value not the element value
	for i := range nums{ 
		fmt.Println(i) 
	}

// Default value of array when not initializes are Zero or Falsy
// int -> 0
// float -> 0.0
// string -> " "(empty string)
// bool -> false	

	x := [3]int{1,2,3}
	fmt.Println(x)

	// 2d arrays
	y := [5][5]int{{1,2,3,4,5},{6,7,8,9,10},{11,12,13,14,15},{16,17,18,19,20},{21,22,23,24,25}}
	fmt.Println(y)

	for i := range y{
		for j:= range y{
		fmt.Print(y[i][j]," ")
		}
	}
	fmt.Println()

	for i := 0;i<len(y);i++{
		for j := 0;j<len(y);j++{
			fmt.Print(y[i][j]," ")
		}
	}
	// arrays are not used in project they are fixed size 
	// used only when size is known and also arrry provide memory optimization
}