package main

import "fmt"

func main() {

	studentData := make(map[string]map[any]any)
	// studentData["Abhi"] = map[any]any{
	// 	"Age":   23,
	// 	"Grade": "A",
	// 	"City":  "Bangalor",
	// }
	// studentData["Madhan"] = map[any]any{
	// 	"Age":   22,
	// 	"Grade": "C",
	// 	"City":  "Mysore",
	// }
	// studentData["Giri"] = map[any]any{
	// 	"Age":   24,
	// 	"Grade": "B",
	// 	"City":  "Kolar",
	// }
	// fmt.Println(studentData)
	studentData = map[string]map[any]any{
		"Abhi": {
			"Age":   23,
			"Grade": "A",
			"City":  "Bangalor",
		},
		"Giri": {
			"Age":   24,
			"Grade": "B",
			"City":  "Kolar",
		},
		"Madhan": {
			"Age":   22,
			"Grade": "C",
			"City":  "Mysore",
		},
	}
	fmt.Println(studentData)
}
