package main

import "fmt"

func main() {
	/*
		Escreva expressões utilizando os seguintes operadores, e atribua seus valores
		a variável.

		1 ==
		2 !=
		3 <=
		4 <
		5 >=
		6 >
	*/

	a := 1 == 1 // true
	b := 2 != 1 // true
	c := 3 <= 3 // true
	d := 4 < 3  // false
	e := 5 >= 6 // false
	f := 6 > 8  // false

	fmt.Printf("%v - %v - %v - %v - %v - %v\n", a, b, c, d, e, f)
}
