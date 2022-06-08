package main

import "fmt"

func main() {
	firstMap := map[string]int{"Uday": 10, "Abc": 20, "Xyz": 30}
	//secondMap := firstMap
	firstMap["Uday"] = 40

	//fmt.Println("SecondMap is:", secondMap)

	thirdMap := map[string]int{}

	for k, v := range firstMap {
		thirdMap[k] = v
	}
	//fmt.Println("ThirdMap is:", thirdMap)
	firstMap["Madhav"] = 70
	firstMap["Madhav"]=45
	firstMap["Uday"]=20
	fmt.Printf("3rd Map:%#v --------------- 1stMap:%#v\n", thirdMap, firstMap)
}
