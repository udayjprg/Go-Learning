package main

import "fmt"

func main() {
	var  result, num1 int

	fmt.Println("Input any number: ")
	fmt.Scanf("%d", &num1)
	result =calc(num1)
	fmt.Printf("Cube of %d = %d\n", num1, result)

}
func calc(num int) int {
	//var num int
	return	 (num * num * num)
//return result	
}
