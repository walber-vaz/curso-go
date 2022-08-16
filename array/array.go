package main

import "fmt"

var x [5]int

func main() {
	/*
		Array -> tipo de estrutura de dados
		valores vão ter um tipo
	*/
	x[0] = 1
	x[1] = 10

	fmt.Println(x[0], x[1])
	fmt.Println(x)
	// len() -> mostra quantidade de itens do array
	fmt.Println(len(x))
}
