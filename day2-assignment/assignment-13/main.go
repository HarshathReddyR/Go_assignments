package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func CountWords(c string) int {
	// ioutil.ReadAll(c)
	var count int = 0
	l := len(c)
	if l != 0 {
		s := strings.Fields(c)
		//  var count int = 0
		count = len(s)
	}
	return count
}
func main() {
	v, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := string(v)
	//l := len(s)
	fmt.Println(CountWords(s))
	// fmt.Println(l)
}
