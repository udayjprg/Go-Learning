package main

import "fmt"

func main() {
	// arry1 :=[3]string{"BMW","AUDI","FERRARI"}
	// //fmt.Println(arry1)
	// arry2 :=[3]int{10,20,30}
	// fmt.Println(arry1,arry2)
	numb := []int{8, 9}
	fmt.Println("Result:", mulArray(numb...))

}
func mulArray(numbs ...int) int {

	result := 1
	for _, numbs1 := range numbs {

		result *= numbs1
	}
	return result
}
