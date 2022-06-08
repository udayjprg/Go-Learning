package main

import "fmt"

func main() {
	// here we initialised the array so its not nil anymore
	m := [2]int{}

	// here we have not initialised so the value will be nil
	var n []int
	fmt.Println(n == nil, m)
}
