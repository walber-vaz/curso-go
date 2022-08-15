package main

import "fmt"

func main() {
	/*
		Nested Loops - Repetição hierárquica
	*/
	// for horas := 0; horas <= 12; horas++ {
	// 	fmt.Println("Hora: ", horas)
	// 	for x := 0; x < 60; x++ {
	// 		fmt.Print(" ", x)
	// 	}
	// 	fmt.Print("\n")
	// }
	for mes := 1; mes <= 12; mes++ {
		fmt.Println("Mes: ", mes)
		for x := 1; x <= 30; x++ {
			fmt.Print("Dia: ", x, "\n")
		}
		fmt.Print("\n")
	}
}
