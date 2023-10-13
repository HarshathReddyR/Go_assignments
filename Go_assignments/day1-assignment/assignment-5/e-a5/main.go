package main

import "fmt"

func main() {
	numbers := []int{1, 3, 56, 7, 3}
	var j int = 0
	for i := 0; i < len(numbers); i++ {
		if j < numbers[i] {
			j = numbers[i]
		}
	}
	fmt.Println("Maximun value in slices is", j)
}
