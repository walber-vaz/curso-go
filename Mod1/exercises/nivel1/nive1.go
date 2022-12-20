package main

import "fmt"

func main() {
	// Atribuar valores as variaveis x, y, z com valores "42", "James Bond", "true"
	// E desmotre valores das variaveis em um unico print
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("Valor de x é: %v do tipo %T", x, x)
	fmt.Printf("\nValor de y é: %v do tipo %T", y, y)
	fmt.Printf("\nValor de z é: %v do tipo %T", z, z)
}