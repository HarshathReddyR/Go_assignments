package main

import "fmt"

func main() {
	fruits := make(map[string]int)
	fruits["Apple"] = 2
	fruits["Banana"] = 3
	fruits["Grapes"] = 5
	fruits["Mango"] = 4
	for k, v := range fruits {
		fmt.Println("Fruits = ", k, " & Quantities", v)
	}
}
