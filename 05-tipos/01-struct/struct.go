package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//Método: receiver
func (p produto) precoComDesconto() float64 {
	return p.preco * (p.desconto / 100)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Caderno",
		preco:    100,
		desconto: 10,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Caixa de som", 50.24, 10}
	fmt.Println(produto2,produto2.precoComDesconto())
}
