package main

import "fmt"

func Divide(a, b int) int {
	if b == 0 {
		panic("Denominator is zero,anything divided by zero is not defined")
	}
	return a / b
}

// func safeDivide(a, b int) (result int) {
// 	defer func() {
// 		if i := recover(); i != nil {
// 			fmt.Println("Recoved")
// 			result = 0
// 		}
// 	}()
// 	result = Divide(a, b)
// 	return result
// }
func main() {
	fmt.Println(Divide(1, 0))
}
