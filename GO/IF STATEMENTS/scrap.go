package main

import "fmt"

func main() {
	// assigning the returned func to a variable
	var myFloatPointer = floatFunc()
	fmt.Println(myFloatPointer)

	// assigning the returned pointer to a variable
	var flaoNum *float64 = floatPoint2()
	fmt.Println(*flaoNum)
}

// creating a func with return type "float64"
func floatFunc() float64 {
	//  assigning value to a varibale
	myFloat := 66.9
	// returning variable
	return myFloat
}

// declaring that the func returns a float64 pointer
func floatPoint2() *float64 {
	floatNumb := 55.2
	// returning a pointer of the specified type
	return &floatNumb
}
