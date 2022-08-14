package main

import (
	"fmt"
	"runtime"
)

var x = 10
var y = 10.0

func main() {
	/*
		Tipos numéricos => int e float
		TODO -> https://go.dev/ref/spec#Numeric_types

		uint8       the set of all unsigned  8-bit integers (0 to 255)
		uint16      the set of all unsigned 16-bit integers (0 to 65535)
		uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
		uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

		int8        the set of all signed  8-bit integers (-128 to 127)
		int16       the set of all signed 16-bit integers (-32768 to 32767)
		int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
		int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

		float32     the set of all IEEE-754 32-bit floating-point numbers
		float64     the set of all IEEE-754 64-bit floating-point numbers

		complex64   the set of all complex numbers with float32 real and imaginary parts
		complex128  the set of all complex numbers with float64 real and imaginary parts

		byte        alias for uint8
		rune        alias for int32
	*/

	// Go tipos são únicos
	// Go linguagem estática

	// Regra geral quando for usar int => use somente int e não int32 ou int64
	// Deixe o computador decide.

	// Regra geral quando for usa float => use somente float64

	a := "e"
	b := "é"
	c := "f"

	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("%v, %v, %v\n", d, e, f)

	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)

	fmt.Println(runtime.GOOS)   // Saber qual OS esta sendo executado
	fmt.Println(runtime.GOARCH) // Sabe qual Arch do seu OS
}
