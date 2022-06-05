package main

import (
	"fmt"
	//"strconv"
)

func main() {

	// var n int
	// fmt.Scan(&n)

	// for i:=0;i<n;{

	// 	for j:=0;j<n-i-1:j++{
	// 		fmt.Print("")
	// 	}
	// 	for j:=0;j<=i;j++{
	// 		fmt.Print("#")
	// 	}
	// 	fmt.Println()
	// }
	var n int
	fmt.Println("Enter your number Here")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

}
