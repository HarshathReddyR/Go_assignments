package main

import "fmt"

func main() {
	// var studentGrades map[string]string
	// studentGrades["abhi"] = "shek"
	studentGrades := make(map[string]string)
	studentGrades["Abhi"] = "A"
	studentGrades["Giri"] = "B"
	studentGrades["Madhan"] = "C"
	for k, v := range studentGrades {
		fmt.Println("names ", k, " grades", v)
	}
	delete(studentGrades, "Madhan")
	fmt.Println(studentGrades)
}
