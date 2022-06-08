package main

import (
	"fmt"
)
func print_head(n int){
	if(n>0){
		print_head(n-1)
		fmt.Println(n)
		
	}
}
func main(){
	print_head(10)

}