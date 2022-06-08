package main

import (
	"fmt"
)
func myrec(){
	
	for x:=0; (x<7); x++ {
	myrec()
	fmt.Println("This Is My First Recursive Function....")
}
}
func main(){
  myrec()
}