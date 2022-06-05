package main

import "fmt"

func main() {
	// var week int
	// fmt.Println("Input week number (1-7): ")
	// fmt.Scanf("%d", &week)
	// switch week {
	// case 1:
	// 	{
	// 		fmt.Println("Monday")
	// 	}
	// case 2:
	// 	{
	// 		fmt.Println("Tuesday")
	// 	}
	// case 3:
	// 	{
	// 		fmt.Println("Wednesday")
	// 	}
	// case 4:
	// 	{
	// 		fmt.Println("Thursday")
	// 	}
	// case 5:
	// 	{
	// 		fmt.Println("Friday")
	// 	}
	// case 6:
	// 	{
	// 		fmt.Println("Saturday")
	// 	}
	// case 7:
	// 	{
	// 		fmt.Println("Sunday")
	// 	}
	// default:
	// 	{
	// 		fmt.Println("Invalid number enter valid number")
	// 	}
	// }
	//fmt.Println(week)

	//***************************************************
	// var month int
	// fmt.Println("Input month number (1-12): ")
	// fmt.Scanf("%d", &month)
	// switch month {
	// case 1:
	// 	{
	// 		fmt.Println("Total number of days = 31")
	// 	}
	// case 2:
	// 	{
	// 		fmt.Println("Total number of days = 29")
	// 	}
	// case 3:
	// 	{
	// 		fmt.Println("Total number of days = 31")
	// 	}
	// case 4:
	// 	{
	// 		fmt.Println("Total number of days = 30")
	// 	}
	// case 5:
	// 	{
	// 		fmt.Println("Total number of days = 31")
	// 	}
	// case 6:
	// 	{
	// 		fmt.Println("Total number of days = 30")
	// 	}
	// case 7:
	// 	{
	// 		fmt.Println("Total number of days = 31")
	// 	}
	// default:
	// 	{
	// 		fmt.Println("Invalid number enter valid number")
	// 	}

	// }
	//***************************************************
	var num int
	fmt.Println("Input  number : ")
	fmt.Scanf("%d", &num)
	switch {
	case (num > 0):
		fmt.Printf("%d is Positive\n", num)
	case (num < 0):
		fmt.Printf("%d is Negative\n", num)
	case (num == 0):
		fmt.Printf("%d is Zero\n", num)

	}
}
