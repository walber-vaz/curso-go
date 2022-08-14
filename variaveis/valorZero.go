package main

import "fmt"

// Declaração de variavel => go pedir para sistema endereço de memoria.
// var x int
var a int
var b float64
var c string
var d bool

func main() {
	// Declaração de variavel => go pedir para sistema endereço de memoria.
	// Inicialização de variavel => e o primeiro valor colocado na variavel.
	// Atribuição de variavel => e valor que vai se colocado depois.

	// Operador curto de declaração Faz tudo isso ao mesmo tempo (Declara, inicializa e atribui)
	// x := 10

	// Cada tipo da linguagem te mseu valor zero
	/*
		int => 0
		float => 0.0
		boolean => false
		string => ""
		ponteiros, função, interfaces, slices, channels, maps => nil
	*/

	// use := sempre que possivel
	// package level scope use var

	// Inicialização de variavel => e o primeiro valor colocado na variavel.
	// x = 10

	// Atribuição de variavel => e valor que vai se colocado depois.
	// x = 20
	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("c: %v, %T\n", c, c)
	fmt.Printf("d: %v, %T\n", d, d)
}
