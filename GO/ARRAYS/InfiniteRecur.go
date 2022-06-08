package main

import "fmt"

func main(){
	show()
}
func show(){
	fmt.Println("Hello")
	show()
}