package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// case2 --> infinite loop
	// 	j:=0
	// 	for {
	// j++
	// 	}
	// 	fmt.Println(j)

	//case 3
	//  j:=10 
	//  for j>=0{
	// 	fmt.Println(j)
	// 	j--
	// }

	// caSE 4
	for k:=0; k <=10; k++{
		if k %2 !=0{
			continue
		}
		fmt.Println(k)
	}

}
