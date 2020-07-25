package main

import "fmt"

func inc1(n int) {
	n++
}
func inc2(n *int) {
	//desreferenciar um ponteiro
	//valor do ponteiro

	*n++
}
func main() {
	n := 1
	inc1(n)
	fmt.Println(n)
	inc2(&n) //por referencia
	fmt.Println(n)
}

//evitar usar esse  metodo
