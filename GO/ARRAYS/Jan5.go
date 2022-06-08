package main

import (
	"fmt"
	"strings"
)

func main(){

	//Arrays exercis1:
	// var cities=[2]string{"Sanjose", "PaloAlto"}
	// grades:=[3]float64{10.1, 1.2}
	// fmt.Printf("%v \n",cities)
	// taskDone:=[...]bool{true,false,true}
	// fmt.Println(taskDone)
	// for i:=0; i<len(cities);i++{
	// 	fmt.Println("index is:",cities[i])
	// }

	fmt.Println(strings.Repeat("#", 25))
	// for i,v:= range grades{
	// 	fmt.Println("index is: ",i,"& value is:",v)
	// }
	
	//Arrays exercise2:
	// nums:=[...]int{30,-1,-6,90,-6}
	// var count=0
	// for _,v:=range nums{
	// 	if v%2==0 && v>0{
	// 		count++
	// 	}
	// }
	// 	fmt.Println(count)

	//Arrays exercis3:
	myArray:=[3]float64{1.2,5.6}
		myArray[2]=6
		a:=10
		myArray[0]=float64(a)
		myArray[2]=10.10
fmt.Println(myArray)
}