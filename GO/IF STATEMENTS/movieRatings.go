package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main(){

// 	if (len(os.Args)!=2) {
// 		fmt.Println("Enter the Age")
// 		return
// 	}
// 	age, err:=strconv.Atoi(os.Args[1])
// 	if err!= nil || age<0{
// 		fmt.Printf("Wrong age: %q"+"\n", os.Args[1])
// 	}else if age>17 {
// 		fmt.Printf("R-Rated \n")
// 	}else if age>=13 && age<=17 {
// 		fmt.Printf("PG-13 \n")
// 	}else if age<13 {
// 		fmt.Printf("PG-Rated\n")
// 	}
// }
