package main

import "fmt"

func main() {
	/*
		Inteiro ocupa 64bits, 8 bytes na memória
	*/
	i := 1
	//Obs: Go não tem aritmética de ponteiros
	fmt.Println(i)

	var p *int = nil // ponteiro nulo
	p = &i           //pega o endereço

	*p++ //desreferenciando - pegando o valor

	i++
	fmt.Println(p, *p, i, &i)
}
