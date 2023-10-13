package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	var sum int
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println("Sum of elements in number", sum)
	avg := sum / len(numbers)
	fmt.Println("average of elements in number", avg)

}
