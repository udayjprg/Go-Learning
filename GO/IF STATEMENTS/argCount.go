package main

import (
	"fmt"
	"os"
)

func main(){
	abc:=(os.Args)
	l:=len(abc)-1
	if l==0 {
		fmt.Println("Give me Args")
	}else if l==1 {
		// %q will print the string with double quotes
		fmt.Printf("There is one argument: %q \n", abc[1])
	}else if l==2 {
		// %s will print the string withut double quotes
		fmt.Printf("There are two arguments: %s %s \n", abc[1], abc[2])
	}else{
	fmt.Printf("There are %d arguments \n",l)
	}
	fmt.Println(abc)
}