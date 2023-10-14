package main

import (
	"d-a8/model"
	"d-a8/shape"
	"fmt"
)

func main() {
	R1 := model.Circle{
		Radius: 10.1,
	}
	fmt.Println(shape.AreaOfCircle(R1))
}
