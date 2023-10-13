package main

import "fmt"

func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	var numbers []int
	numbers = append(numbers, 5)
	fmt.Println(numbers)
	numbers = append(numbers, 10)
	fmt.Println(numbers)
	numbers = append(numbers, 15)
	fmt.Println(numbers)
	numbers = append(numbers, 20)
	fmt.Println(numbers)
	numbers = append(numbers, 25)
	fmt.Println(numbers)
	fmt.Println("Length of numbers", len(numbers))
	fmt.Println("Length of numbers", cap(numbers))

	// var value, j int
	// // i := 0
	// // for i < len(numbers) {
	// // 	if numbers[i] == value {
	// // 		j = i
	// // 		return
	// // 	}
	// // 	i++
	// // }
	var indexvalue int
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == 15 {
			indexvalue = i
		}
	}
	numbers = deleteElement(numbers, indexvalue)
	fmt.Println(numbers)
}
