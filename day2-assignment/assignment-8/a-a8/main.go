package main

import (
	"fmt"
	"os"
)

func removeCurrentDirectory(f string) string {
	err := os.Remove(f)
	if err != nil {
		return "File is not exists"
	}
	return "File is exists and removed"
}
func main() {
	var f string
	fmt.Println("Enter the file name")
	_, err := fmt.Scanln(&f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(removeCurrentDirectory(f))
}
