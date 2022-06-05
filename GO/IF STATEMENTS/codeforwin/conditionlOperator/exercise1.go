package main

import "fmt"

func main(){
	var value1, value2, max int

	fmt.Scanf("%d%d",&value1, &value2)

	if (value1>value2){
		max=value1
		fmt.Printf("Maximum value from: %d, %d is: %d\n",value1, value2, max)
	}else{
		max=value2
		fmt.Printf("Maximum value from %d, %d is: %d\n",value1, value2, max)
	}

	
}