package main

import "fmt"

var first string = "Hello"

func main() {
	second := "tiny hare"
	printer(first, second)
	looper()
	init := closure(15)
	fmt.Println(init())
}

func printer(a string, b string) {
	fmt.Printf("%v, %v! \n", a, b)
}

func looper() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func closure(number int) func() int {
	add := 3
	return func() int {
		return number + add
	}
}
