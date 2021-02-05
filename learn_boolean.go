package main

import "fmt"

func main() {
	fmt.Println("Greater than:", 1 > 2)
	fmt.Println("Less than:", 1 < 3)
	fmt.Println("Less than or equal to:", 1 <= 2)
	fmt.Println("Equivalent:", 4.0 == 4)
    fmt.Println("Not equivalent:", 4.1 != 4.1)  // these are equivalent that is why these are FALSE
}
