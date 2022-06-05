package main

import "fmt"

func main(){
	n :=[]int{0,10}
	m:= make([]int, 6)
	abc := copy(m,n)
	fmt.Println(n,m, abc)
}