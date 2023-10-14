package main

import "fmt"

func main() {
	var r float32
	fmt.Println("Enter the radius:")
	_, err := fmt.Scanln(&r)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(shape.AreaOfCircle(r))
}
