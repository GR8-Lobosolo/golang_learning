package main

import (
	"fmt"
)

func main() {
	birthdays := map[string]string{ //specified the type of map
		"Keith": "02/06/1990",
		"Kevin": "01/01/1957",
		"Kayla": "06/24/1975",
	}

	delete(birthdays, "Keith") // deleteing or removing an item from the map
	fmt.Println(birthdays)

	ages := map[string]int{}
	ages["Keith"] = 28
	ages["Kevin"] = 61
	ages["Kayla"] = 43

	fmt.Println(ages)
	fmt.Println(ages["Keith"]) // printing a specific key's value
}

// map - Simply means we're creating a map
// [string] - We're specifiying that the keys wil lbe strings
// string - The values will also be strings
