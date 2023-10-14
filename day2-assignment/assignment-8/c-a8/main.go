package main

import (
	"c-a8/model"
	"c-a8/person"
)

func main() {
	p1 := model.Person{
		Name: "Abhi",
		Age:  23,
	}
	(person.PrintPerson(p1))
}
