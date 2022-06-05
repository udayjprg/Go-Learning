package main

import (
	//"bufio"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	//	"os"
)

// "bufio"
// "fmt"
// "os"
//	"time"
//"strings"

func main() {
	// var now = time.Now()
	// var year int = now.Year()
	//fmt.Print("Enter a value here:")
	//reader := bufio.NewReader(os.Stdin)
	// var inp int
	// fmt.Scanf("%d", &inp)
	//input := reader.ReadString('\n')
	//fmt.Println(inp)
	//**********************************************************
	// user can enter a value here
	fmt.Print("Enter  Grade here:")
	// created a "reader" variable to store user input value
	reader := bufio.NewReader(os.Stdin)
	// ReadString() has always return two values so the secondone is error
	// to store the error we created "err" variable.and passing "error" into it.
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		//fmt.Println("Passing")
		status = "passing"
	} else {
		//fmt.Println("Failed")
		status = "Failed"
	}

	fmt.Println("A grade of", grade, "is", status) // printing the input
}
