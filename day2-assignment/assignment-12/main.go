package main

import "fmt"

type Animal interface {
	MakeSound() string
}
type Dog struct {
	Sound string
}
type Cat struct {
	Sound string
}

func (d Dog) MakeSound() string {
	return "Dog makes sound as " + d.Sound
}
func (d Cat) MakeSound() string {
	return "Cat makes sound as " + d.Sound
}
func main() {
	C1 := Cat{
		Sound: "miyauu",
	}
	D1 := Dog{
		Sound: "Bouuu",
	}
	fmt.Println(C1.MakeSound())
	fmt.Println(D1.MakeSound())
}
