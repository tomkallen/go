package main

import "fmt"

var first string = "Hello"

func main() {
	second := "tiny hare"
	Printer(first, second)
}

func Printer(a string, b string) {
	fmt.Println(a, b)
}
