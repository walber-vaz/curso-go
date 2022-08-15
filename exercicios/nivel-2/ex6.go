package main

import "fmt"

const (
	_ = 2022 + iota
	ANO1
	ANO2
	ANO3
	ANO4
)

func main() {
	/*
		Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
	*/
	fmt.Printf("ANO1: %v - ANO2: %v - ANO3: %v - ANO4: %v\n", ANO1, ANO2, ANO3, ANO4)
}
