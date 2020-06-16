package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	//fmt.Println(x / float64(y)) ERROR
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("Teste " + string(97)) // a

	// int para string
	fmt.Println("Inteiro para String = " + strconv.Itoa(123))

	// string para int === (_ = ignore)
	num, _ := strconv.Atoi("123") // retorna dois valores => (num, error)
	fmt.Println(num - 122)
	// string para boolean
	bo, _ := strconv.ParseBool("true") // retorna dois valores => (num, error)
	if bo {
		fmt.Println(bo)
	}
}
