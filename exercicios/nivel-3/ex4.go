package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Faça um loop infinito utilizando o for
		ele demostra desde ano de seu nascimento ate atual.
	*/
	anoNascimento := 1991
	anoAtual := time.Now().Year()

	for {
		if anoNascimento < anoAtual {
			anoNascimento++
			fmt.Println(anoNascimento)
		} else {
			break
		}
	}
}
