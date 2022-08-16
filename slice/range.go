package main

import "fmt"

func main() {
	/*
		Função range em slice
	*/

	slice := []string{"banana", "maça", "uva"}

	for indice, valor := range slice {
		fmt.Println("No indice", indice, "temos o valor: ", valor)
	}

	slice = append(slice, "goiaba")

	for _, valor := range slice {
		fmt.Printf("Valor %v\n", valor)
	}
}
