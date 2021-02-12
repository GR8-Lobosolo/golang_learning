package main

import (
	"fmt"
)

func main() {
	ages := map[string]int{}
	ages["Keith"] = 28
	ages["Kevin"] = 71
	ages["Kayla"] = 43
	ages["Corey"] = 10

	// IF STATEMENT
	// This this the example of a "If statement"
	/*
		if ages["Corey"] < 18 {
			// if statement for conditional
			fmt.Println("Corey can't vote")
		} else if ages["Corey"] < 67 {
			// else if to utilize multiple conditionals
			fmt.Println("Corey is not of retirement age")
		} else {
			// else closes the if statement and always returns something
			fmt.Println("Corey is of retirement age")
		}

		fmt.Println(ages)
		fmt.Println(ages["Keith"])
	*/

	// SWITCH STATEMENT
	// This is an example of the same above using switch
	// comment out the If statment above to run this part of the code.
	/*
		switch {
		// evaluates from top to bottom with default executing if the other cases are false
		case ages["Kevin"] < 18:
			fmt.Println("Kevin can't vote yet")
		case ages["Kevin"] < 67:
			fmt.Println("Kevin is not of retirement age just yet")
		default:
			fmt.Println("Enjoy retirement Kevin!")
		}
	*/

	// OR this version of switch
	switch ages["Kevin"] {
	case 1, 2, 3, 5, 7, 11, 13, 17, 19: // allows the comparison between several cases
		fmt.Println("Kevin's age is a prime number!")
	case 16:
		fmt.Println("Kevin can drive now")
	case 18:
		fmt.Println("Kevin can vote now!")
	case 67:
		fmt.Println("Kevin can retire now!")
	default:
		fmt.Println("There's nothing special about Kevin's current age.")
	}

}
