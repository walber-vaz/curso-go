package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	/*
		Utilizando a solução do exercício 2:
		para x => 42
		para y => "James Bond"
		para z => atribua true

		Apos use fmt.Sprintf para atribuir todos esses valores a uma variável.
		Faça essa atribuição de tipo string a uma variável de nome "s" utilizando
		operador curto de declaração.
		demostre a variável "s".
	*/
	s := fmt.Sprintf("%T, %v\n, %T, %v\n %T, %v\n", x, x, y, y, z, z)
	fmt.Println(s)
}
