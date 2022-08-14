package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB
	GB
	TB
)

func main() {
	/*
		Deslocamento de bits é deslocamos dígitos binários para esquerda ou direita

		>> -> Deslocar para direita
		<< -> Deslocar para esquerda
	*/

	// x := 1
	// y := x << 1
	// z := x >> 1

	// fmt.Printf("%b\n", x)
	// fmt.Printf("%b\n", y)
	// fmt.Printf("%b\n", z)

	fmt.Println("binary\t\t\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t\t\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}
