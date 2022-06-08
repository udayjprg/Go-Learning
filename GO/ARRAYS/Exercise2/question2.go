package main

import "fmt"
func main(){
	myArray :=[3]float64{1.2, 5.6}
	myArray[2]=6
	a :=10
	
	// Error --> can't use a type int for float64 
	//     myArray[0]=a
	myArray[0]=float64(a)
	//myArray[0]=a

	// error --> invalid array index 3
	// (out of bound for q 3-element array)

	//myArray[3]=10.10
	fmt.Printf("%#v\n",myArray)
}