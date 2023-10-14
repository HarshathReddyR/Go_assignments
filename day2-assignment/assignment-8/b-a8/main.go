package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	var n int
	fmt.Println("Enetr the number")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Factorial of %v is ", n)
	fmt.Println(factorial(n))
}
