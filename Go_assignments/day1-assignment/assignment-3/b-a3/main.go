package main

import (
	"b-a3/temperature"
	"fmt"
)

func main() {
	var ft int
	fmt.Println("Enter the temperature")
	_, err := fmt.Scanln(&ft)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(temperature.FahrenheitToCelsius(ft))
}
