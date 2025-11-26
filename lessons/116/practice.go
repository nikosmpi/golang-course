package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	// 1
	hobbies := [3]string{"cinema1", "running", "cooking3"}
	fmt.Println(hobbies)
	// 2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])
	// 3
	//selectedHobbies := hobbies[0:2]
	selectedHobbies := hobbies[:2]
	fmt.Println(selectedHobbies)
	// do something for 4
	selectedHobbies = selectedHobbies[1:3]
	fmt.Println(selectedHobbies)
	// 5
	goals := []string{"Learning Go", "Make a CLI"}
	fmt.Println(goals)
	// 6
	goals[1] = "Make an API"
	goals = append(goals, "Make a website")
	fmt.Println(goals)
	// Bonus
	products := []Product{Product{title: "Soublakia", id: 123, price: 13.3}, Product{title: "Koyrkoympinia", id: 124, price: 23.3}}
	products = append(products, Product{title: "Loykanika", id: 125, price: 23.3})
	fmt.Println(products)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
