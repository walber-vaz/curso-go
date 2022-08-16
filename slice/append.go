package main

import "fmt"

func main() {
	/*
		Append -> adicionar um elemento ao slice
	*/

	x := []int{1, 2, 3, 4}
	y := []int{5, 6, 7, 8}
	fmt.Printf("%v - %v\n", x, y)

	// y... -> Esta pegando os índice do slice y e colando no final do slice x
	x = append(x, y...)

	fmt.Printf("%v - %v\n", x, y)
}
