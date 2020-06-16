package main

import "fmt"

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	//não existe operador ternário em Go
	//alternativa
	fmt.Print(obterResultado(7))
}
