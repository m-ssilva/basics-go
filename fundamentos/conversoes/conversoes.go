package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("Teste " + string(97))       // converte 97 utilizando a tabela ASCII, ou seja "a"
	fmt.Println("Teste " + strconv.Itoa(97)) // int para string
	num, _ := strconv.Atoi("123")            // string para int o método Atoi retorna a converrsão, porém caso aconteça erro é retornado um erro na variável _ (underscore)

	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") // a conversão acontece apenas com "true" ou "1"

	if b {
		fmt.Println("Verdadeiro")
	}
}
