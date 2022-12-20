package main

import "fmt"

type w2k int
var x w2k

func main() {
	// Crie um tipo subjacente deve ser int
	// Crie uma variável para este tipo, com o indentificador "x", utilizando a palavra-chave
	// var.

	// Demostre o valor da variável x
	// Demostre o tipo da variável x
	// Atribua 42 à variável x utilizando o operador =
	// demostre o valor da variável
	fmt.Printf("Valor de x é: %v do tipo %T", x, x)
	x = 42
	fmt.Printf("\nValor de x é: %v do tipo %T", x, x)
}