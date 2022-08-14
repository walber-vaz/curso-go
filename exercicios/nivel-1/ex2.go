package main

import "fmt"

var x int
var y string
var z bool

func main() {
	/*
		Use var para declarar três varieaveis, Elas devem ter package-level scope
		Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos
		para estas variáveis:

		x => tipo int
		y => tipo string
		z => tipo bool

		Ops demostre os valores da cada edentificador
		O compilador atribuiu valores essa variáveis. como esses valores se chaman?
	*/
	fmt.Printf("%v\n%v\n%v\n", x, y, z)

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n", z, z)
}
