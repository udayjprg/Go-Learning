package main

import (
	"fmt"
	"strings"
)

func main(){
	arry1:=[3]int{1,2,3}
	//fmt.Printf("%v \n", arry1)
	fmt.Println(strings.Repeat("-",24))
	fmt.Printf("%#v \n", (arry1))
	fmt.Printf("Capacity of array is: %v \n", cap(arry1))
	//fmt.Printf("length of array is:%d \n", (arry1))
	fmt.Println(strings.Repeat("-",14))

	// how to iterate over an array
	
	// for i,v:= range arry1{
	// 	fmt.Println("index is:",i, " & value is:",v)
	// }
	fmt.Println(strings.Repeat("-", 25))
	for i:=0; i<len(arry1);i++ {
		fmt.Println("index is:",i, " & value is:",arry1[i])
		fmt.Printf("length of array is: %v \n" ,len(arry1))
		fmt.Printf("Capacity of array is: %v \n", cap(arry1))
	}
}