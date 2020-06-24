package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// mapas devem ser inicializados

	aprovados := make(map[int]string)
	aprovados[12345678912] = "Jelso"
	aprovados[43214123421] = "Roberto"
	aprovados[51252351235] = "Jurema"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d) \n", nome, cpf)
	}

	fmt.Println(aprovados[12345678912])
	delete(aprovados, 12345678912)
}
