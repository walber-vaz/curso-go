package main

import "fmt"

func main() {
	/*
		Escreva n tela o unicode point de todas as letras Maiúsculas do alfabeto tres vez cada.
	*/
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for in := 1; in <= 3; in++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
