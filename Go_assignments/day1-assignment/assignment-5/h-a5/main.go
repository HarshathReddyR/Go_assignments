package main

import "fmt"

func main() {
	numbers := []int{1, 3, 56, 7, 3, 4, 6, 5}
	var numood, numeven int
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			numeven++
		} else {
			numood++
		}
	}
	fmt.Println("numbers of even elements is ", numeven)
	fmt.Println("numbers of odd elements is ", numood)
}
