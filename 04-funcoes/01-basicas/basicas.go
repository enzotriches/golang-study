package main

import "fmt"

/**
*Função pura não altera nada externamente
 */
func f1() {
	fmt.Println("Sem parametro nem retorno")
}
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}
func f3() string {
	return "F3"
}
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}
func f5() (string, string) {
	return "Retorno 1", "Outro retorno"
}
func main() {
	f1()
	f2("Texto 1", "Text two")
	r3, r4 := f3(), f4("P1", "P2")
	fmt.Printf("R3: %s \n R4: %s \n", r3, r4)
	r51, r52 := f5()
	fmt.Println("F5 1: ", r51)
	fmt.Println("F5 2: ", r52)
}
