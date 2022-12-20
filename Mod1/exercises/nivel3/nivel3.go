package main

import "fmt"

var (
	x int = 42
	y string = "James Bond"
	z bool = true
)

func main() {
	// Utilizando a solução do exercicio anterior
	// Atribua os sequintes valores no package-level scope
	// 42, James Bond, true

	s := fmt.Sprintln(x, y, z)
	fmt.Println(s)
}