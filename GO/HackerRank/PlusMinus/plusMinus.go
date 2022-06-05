package main

import "fmt"

func main() {
	// create variables to store length of array and array
	// to store postive, negative and zero values
	var n, arry, p, negve, z int
	// storing user length of array in n
	fmt.Scan(&n)
	// looping through array
	for i := 0; i < n; i++ {
		// storing user array in arry
		fmt.Scan(&arry)
		// checking the arry values positive or not
		if arry > 0 {
			// incrementing positive values count
			p++
		} else if arry < 0 {
			// incrementing negative values count
			negve++
		} else {
			// incrementing zero values count
			z++
		}
	}
	// storing ratios between postive and length of arry
	pf := float64(p) / float64(n)
	negvef := float64(negve) / float64(n)
	zf := float64(z) / float64(n)

	// printing float values of ratios
	fmt.Printf("%.6f \n", pf)
	fmt.Printf("%.6f \n", negvef)
	fmt.Printf("%.6f \n", zf)
}
