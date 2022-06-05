package main

import "fmt"

func main(){
	Array1 := [4] string {"uday", "Madhav", "Harsha", "A"}
	Array2 := [2]string {"B", "A"}

	outer:
	for index, name := range Array1{
		for _, friend := range Array2{
			if name ==friend{
				fmt.Println("Friend", friend, "at index" , index)
				break outer
			}
		}
		}
		fmt.Println("Message After Outer loop ")
}