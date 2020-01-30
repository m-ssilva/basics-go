package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func main() {
	n := 1
	inc1(n) // passagem por valor
	fmt.Println(n)
	inc2(&n) // passagem por referência
	fmt.Println(n)
}
