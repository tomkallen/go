package main

import "fmt"

func Breaker() {
	x := 0
	for {
		x++
		if x == 10 {
			fmt.Println("broke out with value ", x)
			break
		}
	}
}
