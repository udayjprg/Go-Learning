package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// a, err := strconv.Atoi("45a")
	// if err != nil {
	// 	fmt.Println(err)
	// } else{
	// 	fmt.Println(a)
	// }

	//     Case2:
	// if i, err :=strconv.Atoi(os.Args[1]); err == nil {
	// 	fmt.Println(i)
	// } else{
	// 	fmt.Println("Error came:",err)
	//}
	// Case3:
	if args := os.Args; len(args) != 2 {
		fmt.Println("One Argument is required")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("The Argument must be an Integer ! Error", err)
	} else {
		fmt.Printf("%d km in miles is: %v\n", km, float64(km)*2.602)
	}

}
