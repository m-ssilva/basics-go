package main

import "fmt"

func notaParaConceito(n float64) string {
	nota := int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(10.0))
	fmt.Println(notaParaConceito(8.5))
	fmt.Println(notaParaConceito(5.5))
	fmt.Println(notaParaConceito(3.3))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(999))
}
