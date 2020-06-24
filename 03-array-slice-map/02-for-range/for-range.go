package main

import "fmt"

func main() {
	numeros := [...]int{5, 4, 3, 2, 1} //compilador conta
	// index , value
	for i, numero := range numeros {
		fmt.Printf("%d) %d \n", i+1, numero)
	}

	for _, num := range numeros {
		fmt.Printf(" %d \t", num)
	}
}
