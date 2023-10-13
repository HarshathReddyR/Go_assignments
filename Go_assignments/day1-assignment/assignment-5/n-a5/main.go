package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	var searchnumber int = 3
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == searchnumber {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("Element is not found")
	return
}
