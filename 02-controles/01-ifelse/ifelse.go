package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota < 7 {
		fmt.Println("Reprovado com nota ", nota)
	} else {
		fmt.Println("Aprovado com nota ", nota)
	}
}
func main() {
	imprimirResultado(5.2)
	imprimirResultado(8.6)
}
