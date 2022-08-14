package main

import "fmt"

func main() {
	/*
		Escreva um programa que mostre um número em decimal, binário e hexadecimal.
	*/

	a := 10
	fmt.Printf("Valor em binário: %b\n", a)
	fmt.Printf("Valor em decimal: %d\n", a)
	fmt.Printf("Valor em hexadecimal: %#x\n", a)
}
