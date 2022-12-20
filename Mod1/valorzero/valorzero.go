package main

import "fmt"

var zeroint int64 // valor zero 0
var zerofloat float64 // valor zero 0.0
var zerobool bool // valor zero bool false
var zerostr string // valor zero string ""
// pointers, functions, interfaces, slices, channels, maps => valor zero nil

func main() {
	/*
		Valor Zero
		
		São valores default na declaração da variavel sem inicialização
	*/

	fmt.Printf("Valor zero do tipo: %T, %v", zeroint, zeroint)
	fmt.Printf("\nValor zero do tipo: %T, %v", zerobool, zerobool)
	fmt.Printf("\nValor zero do tipo: %T, %v", zerofloat, zerofloat)
	fmt.Printf("\nValor zero do tipo: %T, %v", zerostr, zerostr)
}