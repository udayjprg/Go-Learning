package main

import "fmt"

/*
Fibonacci problem statement:
In mathematics, the Fibonacci numbers, commonly denoted Fₙ, form a sequence, called the Fibonacci sequence,
such that each number is the sum of the two preceding ones, starting from 0 and 1.
That is, and for n > 1.
The sequence starts: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ...
Write a function which takes positive, integer number N as argument and returns N-th element of Fibonacci sequence.
f(0) = 0
f(1) = 1
f(2) = 1
f(3) = 2
f(4) = 3
f(5) = 5
f(6) = 8

*/
func main() {
	var n int
	fmt.Println("Enter a value for n")
	fmt.Scanf("%d", &n)

	trm1, trm2, nxtVal := 0, 1, 0

	for i := 0; i <= n; i++ {
		switch i {
		case 0:
			fmt.Println(trm1)
			continue
		case 1:
			fmt.Println(trm2)
			continue
		default:
			nxtVal = trm1 + trm2
			trm1 = trm2
			trm2 = nxtVal
			fmt.Println(nxtVal)
		}
	}
}
