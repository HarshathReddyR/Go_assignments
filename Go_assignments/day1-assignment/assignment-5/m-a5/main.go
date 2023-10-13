package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 5}
	// var c int
	if len(slice1) != len(slice2) {
		fmt.Println("Slices are not equal")
		return
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			fmt.Println("Slices are not equal")
			return
		}
	}
	fmt.Println("Slices are equal")

}
