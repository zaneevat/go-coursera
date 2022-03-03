package main

import "fmt"

func main() {
	cond := 1
	// множественные if else
	if cond == 1 {
		fmt.Println("cond is 1")
	} else {
		fmt.Println("cond is 2")
	}

	//switch по 1 переменной
	strVal := "name"
	switch strVal {
	case "name":
		fallthrough
	case "test", "lastName":
		//some work
	default:
		// some work
	}

	// switch как замена многим if else
	var val1, val2 = 2, 2
	switch {
	case val1 > 1 || val2 < 11:
		fmt.Println("first block")
	case val2 > 10:
		fmt.Println("second block")
	}
}
