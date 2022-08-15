package main

import "fmt"

func main() {
	/*
		While não tem em golang
	*/
	// x := 0
	// for x < 10 {
	// 	fmt.Printf("x: %v -> x é maior que dez\n", x)
	// 	x++
	// }

	// Usando break para para loop
	x := 0
	for {
		if x < 10 {
			x++
			fmt.Printf("x: %v -> x é menor que dez?\n", x)
		} else {
			fmt.Printf("x: %v -> x é maior que dez. saindo...\n", x)
			break
		}
	}
}
