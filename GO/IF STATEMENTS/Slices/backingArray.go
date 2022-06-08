package main

import "fmt"

func main() {
	//m := []int{1, 2, 3, 4, 5}
	//****************###@!@@@&&&*****************
	// here we have created a slice --> n & adding slice m to it
	// now if we change values in slice --.m tehn it will
	// effect slcie --> n as well
	//n := m[0:2]
	//newM := []int{}
	// here if we change values in slice (m) then still it
	// wont effect slice (newM) as becoz we are appending the values
	// newM = append(newM, m...)
	// m[0] = 22
	// fmt.Println(m, newM, n)

	//arrry1 := [5]int{1, 2, 3, 4, 5}
	//slice1, slice2 := arrry1[0:3], arrry1[2:4]
	// elements of Slice stored in Backing array (arrry1)
	// Changing backing array values will effect slices as well.
	//arrry1[2] = 40
	//fmt.Println(arrry1, slice1, slice2)
	//************%%%%%%%%%%%$$$$$$$########@@@@@@@@******
	arry2 := [6]int{1, 2, 3, 4, 5, 6}
	slice11 := []int{}
	slice11 = append(slice11, arry2[0:4]...)
	fmt.Println(arry2, slice11)
	slice11 = append(slice11, 10, 20, 30)
	fmt.Println(slice11)
	arry2[2] = 66
	fmt.Println("arry2 values are:", arry2, "Slice11 values are:", slice11)
	//slice12:=make([]int, 0)
	slice12 := arry2[0:6]
	fmt.Println("Slice12 values are:", slice12)
	arry2[2] = 42
	fmt.Println("arry2 values are:", arry2, "slice12 values are:", slice12)
}
