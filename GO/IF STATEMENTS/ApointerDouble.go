package main

import "fmt"

func main() {
	numb2 := 10
	double(&numb2)
	fmt.Println(numb2)
}
func double(numb1 *int) {
	*numb1 *= 2
}
