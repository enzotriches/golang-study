package main

import (
	"fmt"
	"strconv"
)

func notaParaConceito(n int) rune {
	if n > 1 && n < 3 {
		return 'A'
	} else if n >= 3 {
		return 'B'
	}
	return 'C'

}
func main() {
	fmt.Println("1 á 3 = a \n >= 3 = b \n <= 1 = c")
	fmt.Println("Conceito ", strconv.QuoteRune(notaParaConceito(2)))
}
