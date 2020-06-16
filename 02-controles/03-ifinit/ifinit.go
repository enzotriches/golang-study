package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}
func main() {
	if i := numeroAleatorio(); i > 3 { //suportado com switch
		fmt.Println(i, "maior que 3")
	} else {
		fmt.Println(i, "menor que 3")
	}
}
