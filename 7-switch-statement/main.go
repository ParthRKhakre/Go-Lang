package main

import ("fmt"
		"time"
	)

func main() {
	// switch condition is a control statement used to compare a value against multiple cases and execute the matching block.
	// It’s a cleaner alternative to multiple if–else.
	// Go’s switch breaks automatically (no break needed).

	// Go offers multiple type of switch cases :

	// 1.Traditional Switch
	i := 3
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")

	// default case is option in Go, default is recommended for safety, but not mandatory.	
	default:
		fmt.Println("Default Case executed")	
	}

	// 2. Multiple condition switch
	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
			fmt.Println("it's Weekend")
	default:
		fmt.Println("it's Workday")	
	}

	// 3.type switch

	whoAmI := func(i interface{}){
		switch t := i.(type){
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its an string")
		case bool:
			fmt.Println("its an bool")	
		default:
			fmt.Println("other",t)		
		}
	}

	whoAmI("golang")
	whoAmI(98)
	whoAmI(true)
	whoAmI(9.23)

	// whoAmI := func(i interface{}) { ... } → defines an anonymous function that accepts any type (interface{}).
    // switch t := i.(type) → checks the actual data type stored in i.


}