package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	p = &i // pega endereço da variável
	*p++   // o asterisco serve para pegar o valor da variável associada ao ponteiro
	i++

	fmt.Println(p, *p, i, &i)
}
