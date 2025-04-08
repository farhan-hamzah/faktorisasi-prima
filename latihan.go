package main

import (
	"fmt"
)

func primeFactors(n int) {
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			fmt.Printf("%d ", i)
			n /= i
		}
	}
	if n > 1 {
		fmt.Printf("%d", n)
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)
	fmt.Print("Output: ")
	primeFactors(n)
}
