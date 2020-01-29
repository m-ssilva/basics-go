package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 1854.94,
			"Guga Pereira":   8456.99,
		},
		"J": {
			"José João": 11452.41,
		},
		"P": {
			"Pedro Junior": 7452.10,
		},
	}

	delete(funcPorLetra, "P")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)
	}

	for _, funcs := range funcPorLetra {
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
