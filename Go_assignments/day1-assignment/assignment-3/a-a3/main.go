package main

import (
	"a-a3/calculator"
	"fmt"
)

func main() {
	fmt.Println(calculator.Square(3))
	fmt.Println(calculator.Sum(13, 5))
	fmt.Println(calculator.Difference(20, 2))
	fmt.Println(calculator.Product(6, 3))
	fmt.Println(calculator.QuotientRemainde(10, 4))
}
