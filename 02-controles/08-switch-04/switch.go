package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Desconhecido"
	}
}
func main() {
	fmt.Println("Tipo: ", tipo(64.3))
	fmt.Println("Tipo: ", tipo("ola"))
	fmt.Println("Tipo: ", tipo(2))
	fmt.Println("Tipo: ", tipo(func() {}))
	fmt.Println("Tipo: ", tipo(time.Now()))
}
