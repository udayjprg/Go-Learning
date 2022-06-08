package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	success :=false
	seconds := time.Now().Unix()// Unix method will convert time into Integer and will pass into seconds
	rand.Seed(seconds) // we give seconds to "Seed function" to generate random numb
	target := rand.Intn(100) + 1
	fmt.Println(target)
	fmt.Println("I have a Guess number between 1-100")
	fmt.Println("Can you Guess ? ")

	for guessess:=0; guessess <3; guessess ++{
		fmt.Println("You have", 3-guessess, "Guess Left")
		fmt.Println("Guess your number here: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input) // it will remove the new line
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Your Guess is :", guess)
	//	fmt.Println(guess)
	if guess < target {
		fmt.Println("Oops. Your Guess was LOW")
	} else if guess > target {
		fmt.Println("Oops your Guess was High")
	}else{
		fmt.Println("Your Guess is Correct")
		break
	}
}
if !success{
		fmt.Println("Sorry you didn't Guess my number. It was :",target)
	}
}



