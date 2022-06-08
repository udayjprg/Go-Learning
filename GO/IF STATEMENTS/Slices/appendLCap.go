package main

import "fmt"

func main(){
	var nums []int
	fmt.Printf("%#v\n", nums)
	fmt.Printf("Length of slcie is : %d, Capacity of slice is: %d\n", len(nums), cap(nums))
	nums= append(nums, 2,3,4,5,6,7,8,9,10,10,11)
	fmt.Printf("Length of slcie is : %d, Capacity of slice is: %d\n", len(nums), cap(nums))
	//  abc :=[]int{}
	// fmt.Printf("%#v, %d, %d\n",abc, len(abc), cap(abc))

}