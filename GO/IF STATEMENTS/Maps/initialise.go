package main

import "fmt"

func main() {
	// emp := map[int]int{
	// 	1: 20,
	// 	2: 30, 3: 30,
	// 	//4: 40
	// }
	// _=emp
	// balances := map[string]int{
	// 	"uday":10,
	// 	"madhav":20,
	// 	"abc":30,
	// 	}
	// 	_=balances

	abc := map[int]int{1: 10, 2: 20}
	_ = abc

	fmt.Println(abc)

	abc[3] = 0
	v, ok := abc[3]
	if ok {
		fmt.Println("abc[3] exists", v)
	} else {
		fmt.Println("It doesnt exists")
	}
	for k, v := range abc {
		fmt.Printf("Key is: %#v, Value is: %3v\n", k, v)
	}
	delete(abc, 2)
	fmt.Println(abc)
}
