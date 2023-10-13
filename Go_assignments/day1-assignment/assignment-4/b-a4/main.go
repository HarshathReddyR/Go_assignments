package main

import (
	"fmt"
	"math/rand"
)

//	func game(i int) string {
//		r := rand.Intn(10)
//		if r < i {
//			return "too low"
//		}
//		if r > i {
//			return "too high"
//		}
//		return "you guessed it right"
//	}
func main() {
	var a int
	for {
		fmt.Println("Enter the Number")
		_, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Println(err)
		}
		r := rand.Intn(10)
		if r > a {
			fmt.Println("too low")
		} else if r < a {
			fmt.Println("too high")
		} else {
			fmt.Println("you guessed it right")
			return
		}

	}
}
