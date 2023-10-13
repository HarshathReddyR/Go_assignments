package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9}
	for i := 0; i < len(slice2); i++ {
		slice1 = append(slice1, slice2[i])
	}
	fmt.Println(slice1)
}
