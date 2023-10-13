package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	var doubleNumbers []int
	for _, v := range numbers {
		doubleNumbers = append(doubleNumbers, 2*v)
	}
	fmt.Println(doubleNumbers)
}
