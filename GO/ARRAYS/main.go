package main

import (
	 "fmt"
	"strings"
)

func main(){

	//W3 Schools

	// arr1 :=[5]int{1,2,3}
	// fmt.Println(arr1)

	// arr2 :=[2]string{"uday", "madhav"}
	// fmt.Println(len(arr2))
	// fmt.Println(arr2)

	// fmt.Println(strings.Repeat("-", 20))

	// //Access elements of Array
	// arr3 :=[]int{1,2,3}
	// fmt.Println(arr3[2])
	// fmt.Println("Array3")
	//fmt.Println(strings.Repeat("-", 20))
	//Change elements of Array
	// arr4 :=[]int{1,2,3}
	// arr4[2]=22
	// fmt.Println(arr4)
	// fmt.Println("Array4")
	// fmt.Println(strings.Repeat("-", 20))
	//Initialize Only Specific Elements

	//arr5 :=[6]int{1:20,3:50}
	//arr4[2]=22
	// fmt.Println(arr5)
	// fmt.Println("Array5")

	//fmt.Println(strings.Repeat("-", 20))
	//Find the Length of an Array
	// arr6 :=[]string{"uday", "madhav"}
	// fmt.Println(len(arr6))
	// fmt.Println(cap(arr6))

	//fmt.Println(strings.Repeat("-", 20))
	//  PArt-1 How to do Iterations in Array
	// array7 :=[3]int {1,13,12}
	// for i,v :=range array7{
	// 	fmt.Println("index is:", i, "Value is:",v)
	// }

	// // Part-2 How to do Iterations in Array
	// fmt.Println(strings.Repeat("-",10))
	// for i:=0; i<len(array7); i++{
	// 	fmt.Println("index is:", i, "Value is:",array7[i])
	// }
	//fmt.Println(strings.Repeat("-", 20))
	// 2 Dimensional Array( Array inside an Array)
	// declaring 2D array with 3 lines & 2 Columns		
	array8 :=[3][2]int{
				{1,2},
				{3,4},
			}
			
			fmt.Println(array8)
			//o/p will be --> [1,2] [3,4] [0,0]
			fmt.Println(strings.Repeat("-", 20))

			// Arrays Equality --> 
// when 2 Arrays have same length and values then they are equal
array9 := [3]int {1,2,3}	
n :=array9
	fmt.Println("n is equals to array9 :", n == array9)
	// Hee e have changed the array9[0] vaue 
	array9[1] = 10
	fmt.Println("n is equals to array9 :", n == array9)
	fmt.Println(strings.Repeat("-", 20))
	var arry[]int
	//fmt.Println(arry ==nil)
	fmt.Printf("arry %#v\n",arry)
	fmt.Println(strings.Repeat("-", 20))
	nums :=make([] int, 2)
	fmt.Printf("%#v\n",nums)
} 