package main

import "fmt"

func main() {
	/*
		Doc

		Operador curto de declaração

		Somente declara uma variavel não faz atribuição
		Tem tipagem automatica

		Verb para usar no Printf
		GERAL
			(%v) => the value in a default format. when printing structs, the plus flag (%+v) adds field names
			(%#v) => a Go-syntax representation of the value
			(%T) => a Go-syntax representation of the type of the value
			(%%)  => a literal percent sign; consumes no value
		BOOLEAN
			(%t) => the word true or false
		INTEGER
			(%b) => base 2
			(%c) => the character represented by the corresponding Unicode Codepoint
			(%d) => base 10
			(%o) => base 8
			(%O) => base 8 with 0o prefix
			(%q) => a single-quoted character literal safely escaped with Go syntax.
			(%x) => base 16, with lower-case letters for a-f
			(%X) => base 16, with upper-case letters for A-F
			(%U) => Unicode format: U+1234; same as "U+%04X"
		FLOAT
			(%b) => decimalless scientific notation with exponent a power of two, in the manner of strconv.FormatFloat with the 'b' format, e.g -123456p-78
			(%e) => scientific notation, e.g. -1.234456e+78
			(%E) => scientific notation, e.g. -1.234456E+78
			(%f) => decimal point but no exponent, e.g. 123.456
			(%F) => synonym for %f
			(%g) => %e for large exponents, %f otherwise. Precision is discussed below.
			(%G) => %E for large exponents, %F otherwise
			(%x) => hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
			(%X) => upper-case hexadecimal notation, e.g. -0X1.23ABCP+20
		STRING
			(%s) => the uninterpreted bytes of the string or slice
			(%q) => a double-quoted string safely escaped with Go syntax
			(%x) => base 16, lower-case, two characters per byte
			(%X) => base 16, upper-case, two characters per byte
		SLICE
			(%p) => address of 0th element in base 16 notation, with leading 0x
		POINTER
			(%p) => base 16 notation, with leading 0x
	*/
	
	// Operador curto de declaração
	// So pode atribuir em novas variaveis
	// Operador curto so pode esta dentro do scope
	x := 10
	y := "Olá, Gopher"

	// Printf mostra mostra com % o tipo da variavel e entre outra coisas
	// Ref: https://pkg.go.dev/fmt#Printf
	fmt.Printf("x: %v, %T", x, x)
	fmt.Printf("\ny: %v, %T", y, y)
}