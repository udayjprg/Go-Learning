package main

import (
	"fmt"
	"math"
)

func main() {
	// create an integer n forlength of array
	var n int
	fmt.Scan(&n)
	// create a 2d array with a using make func and
	// declare rows with length n
	a := make([][]int, n)
	// create leftdiagonal and right diagonal variables
	ld, rd := 0, 0
	// we are looping over number of rows to allocate memory for columns
	for i := 0; i < n; i++ {
		// we are using make to allocate memory and specifiying
		//  a[i] will be an aray of int type of length n
		a[i] = make([]int, n)
		//loop through j which is columns and
		for j := 0; j < n; j++ {
			//fmt.Println("Engter 2d array here")
			// now we can take inputs for rows and cloumns
			fmt.Scan(&a[i][j])
			// now add values to leftdiagonal and right diagonal
			if i == j {
				ld += a[i][j]
			}
			if i+j == n-1 {
				rd += a[i][j]
			}
		}
	}
			// now we have to write difference of ld and rd
			// we are using math package and using Abs func with float64
			diff := math.Abs(float64(ld - rd))
			fmt.Println(diff)
		
	}

