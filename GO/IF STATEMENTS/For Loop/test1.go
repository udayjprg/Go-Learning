package main

import "fmt"

func main() {
	// using a for loop print out all the numbers between 1 & 500 that divisible by 7

	// for i := 0; i <= 500; i++ {
	// 	if i%7 == 0 && i%5 == 0 {
	// 		fmt.Printf(" %d ", i)
	// 	}

	// }
	//

	// case2 --> using a for loop print out ll the years from your birthday to the current year
	birthyear := 1989
	currentyear := 2022

	for i := birthyear; i <= currentyear; {
		fmt.Printf(" %d ", i)
		i++
	}
	fmt.Println("")

}
