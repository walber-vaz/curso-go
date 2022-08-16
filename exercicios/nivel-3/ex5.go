package main

import "fmt"

func main() {
	/*
		Mostre o resto da divisão por 4 de todos os números entre 10 e 100
	*/

	x := 10
	y := 100
	for {
		if x < y {
			soma := x % 4
			fmt.Printf("Valor %v -> modulo: %v\n", x, soma)
			x++
		} else {
			break
		}
	}
}
