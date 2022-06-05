package main

import "fmt"

func main() {
slice1 := []int{1,2,3}

//slice2 :=make([]int, len(slice1))
//slice1 =append(slice1, 20,30,40)

n :=[]int{2,3,4}
slice1 =append(slice1, n...)
m :=[]int{}
m =append(m,slice1...)
fmt.Println(slice1, m)
}
