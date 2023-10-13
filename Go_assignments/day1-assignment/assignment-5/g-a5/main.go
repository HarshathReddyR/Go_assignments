package main

import "fmt"

func main() {
	numbers := []int{1, 3, 56, 7, 4}
	// for i := len(numbers) - 1; i >= 0; i-- {
	// 	fmt.Print(numbers[i], " ")
	// }
	fmt.Println(numbers)
	start := 0
	end := len(numbers) - 1
	for start < end {
		numbers[start], numbers[end] = numbers[end], numbers[start]
		start++
		end--
	}
	fmt.Println(numbers)
}
