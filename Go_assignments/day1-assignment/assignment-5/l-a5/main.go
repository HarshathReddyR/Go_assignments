package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	var sum int = 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			sum += numbers[i]
		}
	}
	fmt.Println("The sum of the even numbers in the slice is", sum)
}
