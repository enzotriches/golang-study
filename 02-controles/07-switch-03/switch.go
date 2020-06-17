package main

//refatorando o elseif.go
import (
	"fmt"
	"strconv"
)

func notaParaConceito(n int) rune {
	switch {
	case n >= 1 && n <= 3:
		return 'A'
	case n >= 3:
		return 'B'
	default:
		return 'C'
	}
}
func main() {
	fmt.Println("1 á 3 = a \n >= 3 = b \n <= 1 = c")
	fmt.Println("Conceito ", strconv.QuoteRune(notaParaConceito(2)))
}
