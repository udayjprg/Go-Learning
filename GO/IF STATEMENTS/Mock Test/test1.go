package main

import "fmt"

func main() {
	var date int
	week := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(week)
	fmt.Println("Enter date here: ")
	fmt.Scanf("%d", &date)
	for _, v := range week {
		if date!=32{

		fmt.Println("Date is: ", date, v)
		break	
		}
	}
}

//fmt.Scanf("%d", &date)
// switch (date) {
// case 1:
// 	{
// 		fmt.Println("Monday")
// 	}
// case 2:
// 	{
// 		fmt.Println("Tuesday")
// 	}
// 	default	:{
// 		fmt.Println("Please Enter valid date")
// 	}

//}
