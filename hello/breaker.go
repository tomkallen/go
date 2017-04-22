package main

import "fmt"

func Breaker() {
	condition := true
	x := 0
	for condition {
		x++
		if x == 10 {
			fmt.Println("broke out with value ", x)
			break
		}
	}
}
