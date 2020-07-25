package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	var sub = func(a, b int) int {
		return a - b
	}
	fmt.Println(soma(1, 3))
	fmt.Println(sub(1, 3))
}
