package main

import "fmt"

type Address struct {
	Street  string
	City    string
	ZipCode int
}
type Person struct {
	Name string
	Address
}

func main() {
	P1 := Person{
		Name: "Abhi",
		Address: Address{
			Street:  "RG Layout",
			City:    "Kolar",
			ZipCode: 563101,
		},
	}
	fmt.Println(P1)
}
