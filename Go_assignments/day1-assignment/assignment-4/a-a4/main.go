package main

import "fmt"

func main() {
	var i int
	fmt.Println("Enter the Number")
	_, err := fmt.Scanln(&i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(checkEvenOdd(i))

}
func checkEvenOdd(a int) string {
	if a%2 == 0 {
		return "Number is even"
	}
	return "Number is odd"
}
