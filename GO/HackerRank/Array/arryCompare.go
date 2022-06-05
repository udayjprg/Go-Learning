package main

import "fmt"

func main() {

	var a [3]int
	var b [3]int
	fmt.Println("Enter numbers into Arrays")
	fmt.Scan(&a[0], &a[1], &a[2])
	fmt.Scan(&b[0], &b[1], &b[2])
	var c, d int

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			c++
		} else if b[i] > a[i] {
			d++
		}
	}
	fmt.Printf("%d %d \n", c, d)
}
