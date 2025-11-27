package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup("Gopher", numbers...)
	fmt.Println(sum)
}

func sumup(start string, numbers ...int) string {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return fmt.Sprintf("%s %d", start, sum)
}
