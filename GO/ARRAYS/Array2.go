package main

import (
	 "fmt"
	 "strings"
)

func main(){
	// Multi Dimensional Array
	// we declared 2 Rows & 3 Columns
	// array3 := [2][3] int{
	// 	{1,2,3},
	// 	{3,2,1},
	// 	// [3]int{3,2,1}, --> we can declare like this as well
	// }
	// fmt.Println(array3)
	// fmt.Println(strings.Repeat("-", 20))

	// -----> Ellipsis Operator
	// abc:= [...]string{"","","","madhav"}
	// fmt.Println(abc,"length of Array is:",len(abc))
	// fmt.Println(cap(abc))
	//fmt.Println(strings.Repeat("-", 20))
	aba :=[...]string{
		8:   "uday",
		"madhav",
		3:   "m",
	}
	fmt.Println(aba, len(aba))
	//aba[0]="maddy"
	fmt.Println(aba, len(aba))
	fmt.Println(strings.Repeat("-", 20))
	fmt.Printf("%#v\n" ,aba)
	fmt.Println(strings.Repeat("-", 20))

	asd:= [5]bool{1: true, 4:true}	
	fmt.Printf("%#v\n" ,asd)
}