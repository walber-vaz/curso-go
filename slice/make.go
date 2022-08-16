package main

import "fmt"

func main() {
	/*
		Slice make -> criar um slice com tanto de elemento com um tanto de capacidade
	*/

	// Ele criou um slice de 5 elemento com capacidade 10
	// se passa a capacidade ele vai criar outra slice copia do slice antigo
	// os elemento colocar no novo e dobra a capacidade. geralmente e dobro
	slice := make([]int, 5, 10)
	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	fmt.Println(slice, len(slice), cap(slice))
}
