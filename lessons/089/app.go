package main

import "fmt"

type str string

func (text str) log() {
	fmt.Print("Hello World,", text)
}

func main() {
	var name str = "Bob"
	name.log()
}
