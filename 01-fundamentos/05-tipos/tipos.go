package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(3200))

	//sem sinal (só positivos) - uint8 uint16 uint32 uint64
	var b byte = 255 // byte == uint8
	fmt.Println("Byte é", reflect.TypeOf(b))

	//com sinal - int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("Valor máximo do int é", i1)
	fmt.Println("O tipo do il é", reflect.TypeOf(i1))

	//representar um mapeamento da tabela Unicode (int32)
	var i2 rune = 'a'
	fmt.Println("Valor ", i2)
	fmt.Println("O tipo do i2 é", reflect.TypeOf(i2))

	//decimais float32 float64
	var x float32 = 9.34
	//var x = float32(9.34)
	fmt.Println("Valor ", x)
	fmt.Println("Tipo padrão de float ", reflect.TypeOf(9.34)) // float64 por padrão
	fmt.Println("O tipo do x é", reflect.TypeOf(x))            //float32

	//booleano
	bo := false
	//var x = float32(9.34)
	fmt.Println("Valor ", bo)
	fmt.Println("Tipo padrão de boolean é ", reflect.TypeOf(true)) // bool
	fmt.Println("O tipo do bo é", reflect.TypeOf(bo))              //bool

	//string
	s1 := "Minha frase"
	fmt.Println(s1 + "!!! Legal..")
	fmt.Println("Tipo padrão de caracteres ", reflect.TypeOf(s1))
	//multiplas linhas
	s2 := `
	Similar 
	a
	Javascript, hu
	`
	fmt.Println("Tipo padrão de caracteres de varias linhas", reflect.TypeOf(s2))
	fmt.Println("Valor ", s2)
	fmt.Println("Tamanho ", len(s2))
	//char = rune
}
