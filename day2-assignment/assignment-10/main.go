package main

import "fmt"

func Sum(a ...int) int {
	var sum int = 0
	for v, _ := range a {
		sum = sum + v
	}
	return sum
}
func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
}
