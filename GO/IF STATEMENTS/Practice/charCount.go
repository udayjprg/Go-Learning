package main

import (
	"fmt"
	"strings"
)

func main() {
	// Initializing the Strings
	// x := "Golang Programming Language"
	// y := "Language"
	// // Display the Strings
	// fmt.Println("First String:", x)
	// fmt.Println("Second String:", y)
	// // Using Count Function
	// test1 := strings.Count(x, "g")
	// test2 := strings.Count(y, "b")
	// // Diplay the Count Output
	// fmt.Println("Count of 'g' in the First String:", test1)
	// fmt.Println("Count of 'b' in the Second String:", test2)

	fruit := "BATMAAN, SUPERMAN, HEMAN, AQUAMAN"
	name := strings.Count(fruit, "MAN")
	fmt.Println("Count of 'MAN' in the First String:", name)

}
