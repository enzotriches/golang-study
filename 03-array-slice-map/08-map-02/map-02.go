package main

import "fmt"

func main() {
	funcEsalario := map[string]float64{
		"Enzo":             15323.33,
		"Bill Gates":       1300.42,
		"Bezerra da Silva": 50000.55,
	}

	funcEsalario["Roberto Carlos"] = 500.55
	delete(funcEsalario, "Inexistente") //não gera erro, apenas não acontece nada
	fmt.Println(funcEsalario)

	for funcionario, salario := range funcEsalario {
		fmt.Printf("Funcionario %s tem o salário de %.2f \n", funcionario, salario)
	}
}
