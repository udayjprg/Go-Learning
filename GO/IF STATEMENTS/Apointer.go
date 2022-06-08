package main

import "fmt"

func main() {
	// // declaring a vairable "number"
	// var number int
	// // you can get address of a variable with &number
	// result :=&number
	// //printing the result
	// fmt.Println("Value of result is: ",result)
	// // you can print value of number to "result"
	// //by using pointer "*result"
	// fmt.Println("Value of *result-->pointer is: ",*result) // o/p is "0"

	//******************************************

	// Assigning a value to "number1"
	number1 := 2
	// declare a variable that holds a pointer to an int
	var result2 *int
	// Assign a pointer to the variable
	result2 = &number1
	// assigning "number1" value to "result2"
	*result2 = number1
	// printing "*result2" pointer
	fmt.Println("Value of result2 is: ", *result2)
	// updatting "*result2" value with 8
	*result2 = 8
	// printing "*result2" pointer
	fmt.Println("After Updating Value of *result2 is: ", *result2)
	//result2 = &number1
	//fmt.Println("Value of result is: ", *result2)
	//fmt.Println("Value of result is: ", *result2)
}
