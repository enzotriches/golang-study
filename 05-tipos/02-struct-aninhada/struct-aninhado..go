package main

import "fmt"

type item struct {
	id    int
	qnt   int
	valor float64
}
type pedido struct {
	idUser int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.valor * float64(item.qnt)
	}
	return total
}

func main() {
	pedido := pedido{
		idUser: 1,
		itens: []item{
			item{id: 1, qnt: 2, valor: 50.33},
			item{id: 2, qnt: 1, valor: 150.33},
		},
	}
	fmt.Printf("Valor total %.2f", pedido.valorTotal())
}

//estou apaixonado por go
