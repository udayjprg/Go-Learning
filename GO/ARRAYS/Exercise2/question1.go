package main

import "fmt"
func main(){

	//Q1 --> write a GO prg that counts the number 
	// of positive even numbers in the array
	nums :=[...]int {30, -1, -6, 90, -6}
	var count int
	for _,v:= range nums{
		if v%2==0 && v>0{
			count++
		}
	}
	fmt.Println("NUmber of positve even numbers are:", count)
}