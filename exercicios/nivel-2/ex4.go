package main

import "fmt"

func main() {
	/*
		Crie um programa que:
		Atribua um valor int a uma variável.
		Demonstre este valor em decimal, binário e hexadecimal.
		Desloque os bits dessa variável 1 para esquerda, e atribua o resultado a outra variável
		Demostre esta outra variável em decimal, binário e hexadecimal.
	*/

	a := 10
	fmt.Printf("%d, %b, %#x\n", a, a, a)
	s := a << 1
	fmt.Printf("%d, %b, %#x\n", s, s, s)
}
