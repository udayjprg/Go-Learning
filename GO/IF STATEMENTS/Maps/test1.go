package main

import "fmt"

func main() {
	// without initialization
	var emp map[int]string
	//fmt.Printf("%#v\n", emp)
	//fmt.Printf("%d\n", len(emp))
	//_=emp
	fmt.Println(emp)
	//emp[22]="programmer"
	//*******************&&&&&&&^^^^%%%%$$$$$$$$$$$$$$$$$$
	// this is first way to initialize in maps
	map1 := map[int]int{}
	map1[1] = 22
	//*******************&&&&&&&^^^^%%%%$$$$$$$$$$$$$$$$$$
	// this is second way to initialize in maps using make function
	map2 := make(map[int]int)
	map2[1] = 22
	fmt.Println(map2)

}
