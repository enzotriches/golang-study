package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		fatorialAnterior := fatorial(n - 1)
		return n * fatorialAnterior
	}
}

func main() {
	fmt.Println(fatorial(3))
}
