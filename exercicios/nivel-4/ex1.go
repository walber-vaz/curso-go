package main

import "fmt"

func main() {
	/*
		Usando uma literal composta
		- Crie um array que suporte 5 valores do tipo int
		- Atribua valores aos seus indices
		- utilize range e demostre os valores do array.
		- utilizando format printing, demostre o tipo do array
	*/
	array := [5]int{1, 2, 4, 5}

	for i, v := range array {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", array)
}
