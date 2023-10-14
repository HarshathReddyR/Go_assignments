package shape

import "d-a8/model"

var pi float64 = 3.14

func AreaOfCircle(c model.Circle) float64 {
	return pi * c.Radius * c.Radius
}
