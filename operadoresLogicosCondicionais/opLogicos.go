package main

import "fmt"

func main() {
	/*
		Operadores Lógicos Condicionais
		&& -> E lógicos -> As 2 expressões tem que ser true para retorna true
		|| -> OU lógicos -> So e true se apenas uma das expressões e true
		!  ->
	*/
	x := 2

	if x == 2 || x == 3 {
		fmt.Println("Chis é 2 ou 3")
	}
}
