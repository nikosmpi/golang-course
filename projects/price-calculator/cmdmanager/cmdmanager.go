package cmdmanager

import "fmt"

type CMDManager struct {
}

func New() CMDManager {
	return CMDManager{}
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please Enter your prices")
	var prices []string
	for {
		fmt.Print("Price: ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			return nil, err
		}
		if input == "0" {
			break
		}
		prices = append(prices, input)
	}
	return prices, nil
}

func (cmd CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}
