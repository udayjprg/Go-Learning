package main

import "fmt"

func main() {

	var myBool bool = true
	floatPointer3(&myBool)
}

// you can pass pointers to funcs as args.
// just specify that the typeof one or more params should be a pointer
func floatPointer3(numb1 *bool) {
	fmt.Println(*numb1)
}

// 	// assigning func to varibale
// 	var result1 = floatPointer1()
// 	fmt.Println(result1)

// }

// func floatPointer1() float64 {
// 	myfloat1 := 10.2
// 	return myfloat1
// }

//****************************************** page 335
// returning pointers from functions...........declare func return type as poimter type
// assigning pointer to retrun float64
// func floatPointer2() *float64 {
// 	myfloat2 := 33.33
// 	// returning pointer of the specified type
// 	return &myfloat2
// }

//****************************************** page 336

// we can also pass pointer as arguments to funcs
