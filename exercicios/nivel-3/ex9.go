package main

import "fmt"

func main() {
	/*
		Crie programa usando switch statement seja uma variavel do tipo string e identificador "esporteFavorito"
	*/
	esporteFavorito := "Futebol"

	switch esporteFavorito {
	case "Futebol":
		fmt.Printf("%v é Legal\n", esporteFavorito)
	case "Video Game":
		fmt.Printf("%v é Moderno\n", esporteFavorito)
	case "Natação":
		fmt.Printf("%v Hum\n", esporteFavorito)
	}
}
