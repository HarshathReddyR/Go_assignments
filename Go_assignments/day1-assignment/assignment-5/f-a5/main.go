package main

import "fmt"

func main() {
	numbers := []int{1, 3, 56, 7, 3}
	var j int
	fmt.Println("Enter the serach element")
	_, err := fmt.Scanln(&j)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(numbers); i++ {
		if j == numbers[i] {
			j = numbers[i]
			fmt.Println("Element is found")
			return
		}
	}
	fmt.Println("element in not found")
}
