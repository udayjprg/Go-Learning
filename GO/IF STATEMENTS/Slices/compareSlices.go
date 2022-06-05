package main

import "fmt"

func main() {
	// we cant compare two sloces.
	
	m, n := []int{1, 2, 3}, []int{1, 5, 3}
	var eq bool = true
	m =nil
	for i, valueM := range m {
		if valueM != n[i] {
			eq = false
			break
		}
	}
		if eq {
			fmt.Println("m & n are equal slices")
		} else {
			fmt.Println(" m & n are not equal slices")
		}

	

}
