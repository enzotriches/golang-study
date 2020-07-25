package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}
func exec(fn func(int, int) int, x, y int) int {
	return fn(x, y)
}
func main() {
	result := exec(multiplicacao, 3, 2)
	fmt.Println(result)
}
