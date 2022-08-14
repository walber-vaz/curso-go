package main

import "fmt"

type w2k int

var x w2k

func main() {
	/*
		Crie um tipo. O tipo subjacente deve ser int.
		Crie uma variável para este tipo, com o identificador "x", utilizando a
		palavra-chave var.

		Apos demostre o valor da variável "x"
		Demostre o tipo da variável "x"
		Atribua 42 á variável "x" utilizando o operador "="
		Demostre o valor da variável "x"
	*/

	fmt.Printf("Valor de x: %v\n", x)

	fmt.Printf("Tipo de x: %T\n", x)

	x = 42

	fmt.Printf("Valor atribuindo a x: %v\n", x)

}
