package main

import "fmt"

func main() {
	// Ref: https://pkg.go.dev/fmt#pkg-index

	// interpreted string literals
	x := "oi bom dia\ncomo vai?\tespero \"que\" esteja tudo bem."
	fmt.Println(x)

	// interpreted raw string
	y := `oi bom dia\ncomo vai?\tespero \"que\" esteja tudo bem.`
	fmt.Println(y)

	// Print => não quebra linha
	fmt.Print("Olá")
	// Printf => Print formatado
	fmt.Printf("Olá")
	// Println => Print com quebra de linha
	fmt.Println("Olá")

	// Sprint não mostra na tela ele retorna uma string
	a := "oi"
	b := " bom, dia"
	c := fmt.Sprint(a, b)

	fmt.Println(c)
}