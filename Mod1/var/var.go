package main

import "fmt"

// Variavel "Global"
var y = 10

func main() {
	z := 20
	// Chamando a função para imprimir valor y
	printvar(z)
}

// Parâmetro da função recebe um inteiro
func printvar (x int) {
	fmt.Println(y)
	fmt.Println(x)
}