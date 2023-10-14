package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) area() int {
	return r.Height * r.Width
}
func (r Rectangle) perimeter() int {
	return (r.Width + r.Height) * 2
}

// func (r Rectangle) areaAndPerimeter() {
// 	fmt.Println("Area of the ", r, "", r.area())
// 	fmt.Println("Permeter of the ", r, "", r.perimeter())
// }
func main() {
	R1 := Rectangle{
		Width:  20,
		Height: 30,
	}
	R2 := Rectangle{
		Width:  40,
		Height: 60,
	}
	fmt.Println("Area of the R1", R1.area())
	fmt.Println("Perimeter of the R1", R1.perimeter())
	fmt.Println("Area of the R2", R2.area())
	fmt.Println("Perimeter of the R2", R2.perimeter())
	fmt.Println("")
	// R1.areaAndPerimeter()
	// R2.areaAndPerimeter()
}
