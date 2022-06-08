package main

import "fmt"

func main() {
	a := map[string]string{"A": "B"}
	b := map[string]string{"A": "C"}

	// c := map[int]int{2: 4}
	// d := map[int]int{2: 3}

	// s3 := fmt.Sprintf("%s", c)
	// s4 := fmt.Sprintf("%s", d)

	// if s3 == s4 {
	// 	fmt.Println("Mpas are equal")
	// } else {
	// 	fmt.Println("Mpas are equal")
	// }

	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)
	if s1 == s2 {
		fmt.Println("Mpas are equal")
	} else {
		fmt.Println(("Maps are not equal"))
	}
}
