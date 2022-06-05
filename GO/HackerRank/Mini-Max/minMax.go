package main

import (
	"fmt"
	"sort"
)

func main() {
	// declaring sum of Min and max
	sumMin, sumMax := 0, 0
	//sumMin:=0
	// creating an array of arry with make function of length 5
	arry := make([]int, 5)
	//allowing user to give values to arry
	fmt.Scan(&arry[0], &arry[1], &arry[2], &arry[3], &arry[4])
	// using sort package to get Ints() function
	sort.Ints(arry)
	// concatinating value from index 0 to 4 sending hem to sumMin
	sumMin = arry[0] + arry[1] + arry[2] + arry[3]
	sumMax = arry[1] + arry[2] + arry[3] + arry[4]

	// printing sunMin & sumMax
	fmt.Printf("Minimum Sum Value is: %d \n", sumMin)
	fmt.Printf("Maximum Sum Value is: %d\n", sumMax)
}

//***********************************************
// for i:=0; i<=len(arry)-2;i++{
//  	sumMin+=arry[i]
// }
// for j:=1; j<len(arry);j++{
// 	sumMax+=arry[j]
// }
//*****************************************************

//fmt.Printf("",sumMax)
//fmt.Printf("%d \n",sumMin)
//fmt.Println(len(arry))
// for i,_:=range arry{
// 	sumMin+=arry[i]
// }
// fmt.Println(sugvhhhhh
