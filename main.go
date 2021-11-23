package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Start  here --> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "quit" {
			fmt.Println("You Entered correct word")
			break
		} else {
			fmt.Println("Hello Mastaru, You Entered wrong word, Please try again")
		}
	}
}
