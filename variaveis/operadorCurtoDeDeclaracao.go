package main

import "fmt"

func main() {
	// Operador curto de declaração => :=
	// Operador curto serve para declara variavel, pode ser usando apenas em
	// variaveis novas Ex: x := 10

	// Tipagem automatica no operador curto de declaração
	x := 10
	y := "Olá, Mundo!"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	// Operador curto de declaração serve para declara apenas uma variavel
	// Se não for usar para declara uma variavel não pode ser usado
	// O operador curto so pode ser usando dentro de contexto do bloco. não
	// pode ser usado no contexto global.
	x, z := 20, 30
	fmt.Printf("z: %v, %T\n", z, z)
	fmt.Printf("x: %v, %T\n", x, x)
}
