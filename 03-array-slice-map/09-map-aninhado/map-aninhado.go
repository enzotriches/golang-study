package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"E": {
			"Enzo Trichês": 10.00,
		},
		"B": {
			"Bruno Espíndola": 1800.00,
		},
		"A": {
			"Ariel Schmmit": 1800.00,
		},
		"J": {
			"João Borges": 50000.00,
		},
	}

	delete(funcPorLetra, "J")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)
		for funcionario, salario := range funcs {
			fmt.Printf("\t %s (R$%.2f)\n ", funcionario, salario)
		}
	}
}
