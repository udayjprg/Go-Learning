package main

import (
	"fmt"
	"strings"
)

func main() {
	result   := "Uday is Hello, Uday has is a handsome guy, Uday is an Indian"
	//fmt.Println("Enter String here:")
	//fmt.Scanf("%s", &result)
	for index, element := range repitition(result) {
		fmt.Println(index, "=", element)
	}

}

func repitition(st string) map[string]int {
	input := strings.Fields(st) // using strings.Field() function
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		 if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
		//fmt.Println(word)
	}
	return wc
}
