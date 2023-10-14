package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	Age  int
	City string
}

func main() {
	E1 := Employee{
		Id:   1,
		Name: "Abhi",
		Age:  23,
		City: "Bangalore",
	}
	E2 := Employee{
		Id:   2,
		Name: "Giri",
		Age:  22,
		City: "Kolar",
		
	}
	fmt.Println(E1)
	fmt.Println(E2)
}
