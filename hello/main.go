package main

import "fmt"

var first string = "Hello"
const ex rune = 2305888

func main() {
	second := "tiny hare"
	printer(first, second)
	looper()
	addThreeToFifteen := closure(15)
	fmt.Println(addThreeToFifteen(3))
	fmt.Println(closure(10)(4))
	looper()
	Breaker()
}

func printer(a string, b string) {
	fmt.Printf("%v, %v%v \n", a, b, ex)
}

func looper() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func closure(number int) func(add int) int {
	return func(add int) int {
		return number + add
	}
}
