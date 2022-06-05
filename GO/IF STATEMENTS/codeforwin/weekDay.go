package main

import "fmt"

func main() {
	// week of the day.
	// i/P ENTER NUMBER --> 1 O/P: Its MONDAY
	// var num1 int
	// fmt.Println("Enter Number here")
	// fmt.Scanf("%d", &num1)
	// if num1 == 1 {
	// 	fmt.Println("Monday")
	// } else if num1 == 2 {
	// 	fmt.Println("Tuesday")
	// } else if num1 == 3 {
	// 	fmt.Println("Wednesday")
	// } else if num1 == 4 {
	// 	fmt.Println("Thurssday")
	// } else if num1 == 5 {
	// 	fmt.Println("Friday")
	// } else if num1 == 6 {
	// 	fmt.Println("Saturday")
	// } else if num1 == 7 {
	// 	fmt.Println("Sunday")
	// } else {
	// 	fmt.Printf("Enter Number between: 1-7\n")
	// }

	//*************************************************************
	// program to calculate profit or loss.
	// var costPrice, sellingPrice, profit, loss int
	// fmt.Println("Enter Price here:")
	// fmt.Scanf("%d%d", &costPrice, &sellingPrice)
	// if costPrice < sellingPrice {
	// 	profit= sellingPrice - costPrice
	// 	fmt.Printf("Profit:  %d \n", profit)
	// }else if (costPrice > sellingPrice){
	// 	loss= costPrice - sellingPrice
	// 	fmt.Printf("Loss:  %d \n", loss)
	// }else{
	// 	fmt.Println("Enter numbers only")
	// }

	//*************************************************************
	// program to enter student marks & find % & grade
	var Physics, Chemistry, Biology, Mathematics, Computer int
	//Percentage :=100
	fmt.Println("Input marks of five subjects: ")
	fmt.Scanf("%d %d %d %d %d", &Physics, &Chemistry, &Biology, &Mathematics, &Computer)
	Percentage := ((Physics + Chemistry + Biology + Mathematics + Computer) / 5)
	//fmt.Println(Percentage)
	fmt.Printf("Percentage = %.2d \n", Percentage)
	if Percentage >= 90 {
		fmt.Println(" Grade A")
	} else if Percentage >= 80 {
		fmt.Println("Grade B")
	} else if Percentage%5 >= 70 {
		fmt.Println("Grade C")
	} else if Percentage >= 60 {
		fmt.Println("Grade D")
	} else if Percentage >= 50 {
		fmt.Println("Grade E")
	} else if Percentage >= 40 {
		fmt.Println("Grade F")
	} else {
		fmt.Println("Enter five subjects marks only")
	}
}
