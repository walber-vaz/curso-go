package main

import "fmt"

type w2k int
var x w2k
var y int

func main() {
	// Utilizando a solução anterior
	// Em package-level scope utilizando a palavra-chave var crie uma variavel
	// com o indentificador y
	// o tipo deve ser subjacente do tipo que voce criou no anterior

	fmt.Printf("Valor de x é: %v", x)
	fmt.Printf("\nTipo de x é: %T", x)
	x = 42
	fmt.Printf("\nValor de x é: %v", x)
	
	y = int(x)
	fmt.Printf("\nValor de y é: %v", y)
	fmt.Printf("\nTipo de y é: %T", y)
}