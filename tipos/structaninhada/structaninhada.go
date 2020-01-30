package main

import "fmt"

type item struct {
	produtoID int
	qtd       int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, qtd: 2, preco: 12.10},
			item{produtoID: 2, qtd: 1, preco: 23.49},
			item{produtoID: 11, qtd: 100, preco: 3.13},
		},
	}

	fmt.Printf("O valor total do pedido é R$ %.2f", pedido.valorTotal())
}
