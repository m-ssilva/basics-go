package main

import "fmt"

func main() {
	aprovados := make(map[int]string)
	aprovados[123456789] = "Maria"
	aprovados[123654789] = "Pedro"
	aprovados[987456321] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	delete(aprovados, 987456321) // <-- deleta valor do map
	fmt.Println(aprovados)
}
