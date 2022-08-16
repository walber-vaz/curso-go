package main

import "fmt"

func main() {
	/*
		Escreva n tela o unicode point de todas as letras Maiúsculas do alfabeto tres vez cada.
	*/
	for i := 65; i <= 90; i++ {
		fmt.Printf("Unicode: %#U\n", i)
	}
}
