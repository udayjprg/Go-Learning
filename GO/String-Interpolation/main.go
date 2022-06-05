package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main(){
	reader=bufio.NewReader((os.Stdin))

	userName:=readString("what is your name ?")
	age:=readInt("What is your Age ?")
	fmt.Printf("Your name is %s. Your Age is %d. \n", userName, age)
}
func prompt(){
	fmt.Println("-->")
}
func readString(s string) string{
		fmt.Println(s)

	for{	
		userInput, _:= reader.ReadString('\n')
		userInput=strings.Replace(userInput, "\n", "",-1)
		if userInput=="" {
			fmt.Println("Enter a value")
		}else{
		 return userInput
		}
}

}

func readInt(s string) int{
	fmt.Println(s)

		for {
			userInput, _:= reader.ReadString('\n')
		userInput=strings.Replace(userInput, "\n", "",-1)

		num, err:=strconv.Atoi(userInput)
		if err!=nil {
			fmt.Println("Enter a Whole number")
		}else{
		return num
		}
	}
}