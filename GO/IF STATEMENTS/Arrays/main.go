package main

import "fmt"

func main() {
	var cities [2]string
	fmt.Printf("%v\n",cities)
	grades :=[3] float64{1.6, 3.4,}
	fmt.Printf("%v\n",grades)
	taskDone :=[...]bool{true, false, false}
	fmt.Println(taskDone)
	for i:=0; i<len(cities); i++{
		fmt.Println(cities[i], " from cities")
	}
	for index,value := range grades{
		fmt.Println("at index:",index,"value is:", value)
	}
}
