package main

import "fmt"

func main() {
	// `` interpletar tudo como esta
	x := "Oi"
	y := "bom dia"

	z := fmt.Sprint(x, " ", y)
	// Println com pular linha no final
	// fmt.Println(x)

	// Print sem pular linha no final
	// fmt.Print(x)

	// Fprint => "Escreve em arquivos"

	fmt.Println(z)
}
