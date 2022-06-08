package main

import (
	//"crypto/rand"
	"bufio"
	"fmt"
	"log"
	"math/rand" // import math/rand package
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// it will get time now in seconds
	seconds := time.Now().Unix()
	// it will seed the random number generator with a new value
	rand.Seed(seconds)
	// generating a random number into "target"
	target := rand.Intn(100) +1// "rand.Intn" will genrate random numb
	// printing target
	fmt.Println(target)
	fmt.Println("I have a Guess Number here between 1-100")
	
	fmt.Println("Can you Guess it? ")
	//var guessNumber int
	//
	reader := bufio.NewReader(os.Stdin)

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "Guesses Left")
		fmt.Println("Make a Guess ? ")
		input, err := reader.ReadString('\n')
		

		//
		
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess > target {
			fmt.Println("Oops your Guess is too High")
		} else if guess < target {
			fmt.Println("Oops your Guess is too Low")
		}else {
			fmt.Println("Your Guess is Correct")
			break
		}

	}

	//fmt.Println(input)
}
