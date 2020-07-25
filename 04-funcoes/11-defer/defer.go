package main

import "fmt"

func obterValorAprovado(n int) int {
	defer fmt.Println("fim")
	if n >= 5000 {
		fmt.Println("Alto")
		return 5000
	}
	fmt.Println("baixo")
	return n

}

func main() {
	fmt.Println(obterValorAprovado(6000))
	obterValorAprovado(100)
}
