package main

import "fmt"

// Quando faz a declaração da variavel sem atribuir valor
// a tribuição do valor do pode ser atribuida dentro de code block
//var x int
var x int = 10

func main() {
	/*
		DOC
		
		Tipos em Go são estaticos
		Go e de tipagem estaticas
	*/

	// Se variavel foi declarada com interge não pode recebe outro valores além do
	// interge
	x = 20
	fmt.Println(x)
}