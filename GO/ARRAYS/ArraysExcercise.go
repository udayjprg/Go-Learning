package main

import (
	"fmt"
	"strings"
) 

func main(){
	var cities=[2]string{}
	for i:=0; i<len(cities); i++ {
	fmt.Println(cities[i])
	}
	fmt.Println(strings.Repeat("-",20))
	// q2
	grades:=[3]float64{1.2, 3.4}
	for i,v:= range grades{
		fmt.Println("For index",i, "value is:", v)
	}
	

	//q3
	var taskDone=[...]bool{true, false}
	fmt.Println(taskDone)
}