package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) print() {
	fmt.Println(m)
}
func main() {
	userNames := make([]string, 3, 5)
	userNames[0] = "Bob"
	userNames[1] = "Alex"
	userNames = append(userNames, "Jack")

	//fmt.Println(userNames)

	courses := make(floatMap, 3)
	courses["go"] = 7.4
	courses["svelte"] = 4.4
	courses["php"] = 4.5
	//courses.print()

	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for key, value := range courses {
		fmt.Println(key, value)
	}
}
