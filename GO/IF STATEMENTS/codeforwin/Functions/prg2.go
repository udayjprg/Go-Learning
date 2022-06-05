package main

import "fmt"

func main(){
	var num1, num2  int
	fmt.Println("Input two numbers: ")
	fmt.Scanf("%d %d", &num1, &num2)
	status(num1, num2)
	fmt.Printf("Maximum number: %d\n",num1 )
	fmt.Printf("Maximum number: %d\n",num2 )
}

func status(max, min int)int {
	var num1, num2 int
	if num1 > num2 {
		//max=num1
		fmt.Printf("%d\n", max)
	}else if num2>num1 {
		max=num2
		//fmt.Printf("%d\n", max)
	}else{
		fmt.Println("Enter two valid numbers")
	}
	
	return num1, num2
}