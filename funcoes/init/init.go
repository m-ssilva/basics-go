package main

import "fmt"

/* a função init() é executada antes do main() */
func init() {
	fmt.Println("inicializando...")
}

func main() {
	fmt.Println("main...")
}
