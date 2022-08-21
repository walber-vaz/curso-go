package main

import "fmt"

func main() {
	/*
		Slice Array Subjacente -> Cuidado
	*/
	primeiroSlice := []int{1, 2, 3, 4, 5}
	segundoSlice := append(primeiroSlice[:2], primeiroSlice[4:]...)
	fmt.Println(segundoSlice)
}
