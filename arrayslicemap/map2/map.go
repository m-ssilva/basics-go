package main

import "fmt"

func main() {
	funcsSalarios := map[string]float64{
		"José":  1200.43,
		"Ana":   1804.78,
		"Pedro": 2541.78,
	}

	funcsSalarios["José"] = 1350.90
	delete(funcsSalarios, "Unknown") // <-- tentar elemento inexistente não retorna erro

	for nome, salario := range funcsSalarios {
		fmt.Println(nome, salario)
	}
}
