package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func Update(a *Student) {
	a.Age++
}
func main() {
	s1 := Student{
		Name: "Abhi",
		Age:  22,
	}
	Update(&s1)
	fmt.Println(s1)
}
