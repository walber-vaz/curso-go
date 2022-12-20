package main

import "fmt"

var (
	x int
	y string
	z bool
)

func main() {
	// Usar var para declarar três variáveis, Elas devem ter package-level scope
	// Não atribua valores as mesma utilizer os seguites identificadores

	// x => int
	// y => string
	// z => bool

	// demostre as valores de cada identificador
	fmt.Printf("Valor default de x é: %v tipo %T", x, x)
	fmt.Printf("\nValor default de y é: %v tipo %T", y, y)
	fmt.Printf("\nValor default de z é: %v tipo %T", z, z)
}