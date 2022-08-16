package main

import "fmt"

func main() {
	/*
		Mostre a opção else if no ex6
	*/
	estudandoGo := "Não"

	if estudandoGo == "Sim" {
		fmt.Println("Estudando Go")
	} else if estudandoGo == "Não muito" {
		fmt.Println("Não perca o foco.")
	} else {
		fmt.Println("Volte a estudar go")
	}
}
