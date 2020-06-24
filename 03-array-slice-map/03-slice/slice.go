package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	s2 := a2[1:3]       // 2,3 índice 1 ao 3, mas não incluindo o 3
	fmt.Println(a2, s2) //não gerou um array diferente, mas sim um ponteiro para um array

	s3 := s1[:2] // 1,2 início até o 2 não incluindo o 2
	fmt.Println(a2, s3)
	/*
		Slice não é um arrau, e sim é um PEDAÇO de um array.
		mais leve por ele é uma referencia de um array
	*/
}
