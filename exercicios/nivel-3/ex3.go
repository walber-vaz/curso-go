package main

import "fmt"

func main() {
	/*
		Crie um loop utilizando a sintaxe for condition {}
		utilize-o para demostrar os anos desde que você nasceu.
	*/
	x := 1991
	y := 2022
	for x <= y {
		fmt.Printf("Ano: %v\n", x)
		x++
	}
}
