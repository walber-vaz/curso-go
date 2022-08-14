package main

import "fmt"

var x bool

func main() {

	/*
		Type Boolean => Valor true ou false
	*/

	// Valor zero
	fmt.Println(x)
	x = true
	// Atribuição
	fmt.Println(x)
	// Pode fazer assim também => x = (10 < 100)
	x = 10 < 100
	fmt.Println(x)
}
