package functionsarevalues

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformerNumbers := transformNumbers(&numbers, transformerFn1)
	fmt.Println(transformerNumbers)
	transformerNumbers2 := transformNumbers(&moreNumbers, transformerFn2)
	fmt.Println(transformerNumbers2)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNambers := []int{}
	for _, val := range *numbers {
		dNambers = append(dNambers, transform(val))
		//(*numbers)[i] = val * 2
	}
	return dNambers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
