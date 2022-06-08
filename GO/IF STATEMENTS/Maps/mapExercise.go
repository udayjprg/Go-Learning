package main

import "fmt"

//import "fmt"

func main(){
	// var m10 map[int]string

	// fmt.Printf("M1 Map Type is: %T \n M1 Map value is: %#v\n", m10, m10)
	// m9 :=map[int]string{ 1:"Uday", 2:"Abc"}
	// m9[10]="Abba"
	// fmt.Println(m9[1])
	// fmt.Println(m9[3])


	// Exercise2:

//var m1 map[int]bool --> panic error can't assign to entry in nil map.
	// m1:=map[int]bool{}
	// m1[5] = true
 
	// m2 := map[int]int{3: 10, 4: 40}
	// m3 := map[int]int{3: 10, 4: 40}
 
	// //fmt.Println(m2, m3)
	// _,_,_=m1,m2,m3

		//Exercise 3 --> 

		m :=map[int]bool{1: true, 2: false, 3: false} // type int but assgined as string
		delete(m,1)
		for k,v := range m{
			fmt.Printf("Key is:%d, Value is:%t\n",k,v)
		}
		fmt.Println(m)

}