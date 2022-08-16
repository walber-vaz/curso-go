package main

import (
	"fmt"
)

func main() {
	/*
		Slice e tipo composto de dados
		Tipo de dados composto e qualquer tipo de dados
		que podem ser construindo em um programa utilizando tipos primitivos
	*/

	// outra forma de usar arrays -> declara o array e ja colocando valores no indices
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	// Slice -> não tem numero finito de elementos
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// Adicionando elementos no slice
	slice2 := append(slice, 6)
	fmt.Println(slice2)

	// Trocando valor dos elementos do slice
	slice[3] = 43
	fmt.Println(slice)

}
