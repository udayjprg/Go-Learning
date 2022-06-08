package main

import (
	"fmt"
)

func main() {
	// today := time.Now().Hour()

	// switch {
	// case today >8:
	// 	fmt.Println("Good Evening")
	// case today > 19:
	// 	fmt.Println("Good Afternoon")
	// 	default :
	// 	fmt.Println(today)
	// }

	// case 2
	num := 25

	switch {
	case num > 25:
		fmt.Println("This is 1st case")
	case num <= 25:
		fmt.Println("This is 3rd case")
	case num > -1:
		fmt.Println("This is 2nd case")

	}
}
