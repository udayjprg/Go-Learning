package main

import (
	"fmt"
	//"log"
	"math"
)
func main(){
	result, err:=sqRoot(-6)
	if (err!=nil){
		fmt.Println(err)
	}else{
		fmt.Printf("Squaeroot is %.2f \n",result)
	}
}
func sqRoot(num float64)(float64,error){
	if num <0 {
		return 0, fmt.Errorf("%.2f value is invalid",num)
	}
	return math.Sqrt(num), nil // here "nil" indicates no error
}