package main

import "fmt"

type Shape interface {
	Area() float64
}
type Circle struct {
	Radius float64
}
type Rectangle struct {
	Length float64
	Width  float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}
func main() {
	C1 := Circle{
		Radius: 3,
	}
	R1 := Rectangle{
		Length: 20,
		Width:  30,
	}
	fmt.Println(C1.Area())
	fmt.Println(R1.Area())
}
