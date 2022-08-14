package main

import "fmt"

const x = 10

// var y = 10

// var c int

var d float64

func main() {
	/*
		Contantes -> Valores que você não pode alterar nunca(Valor contante)

		Contantes pode ser tipadas ou não
		Ex:
			const a = "Ola"
			const b string = "Mundo"

			// Outro tipo de criar varias constantes ou variáveis(Somente do mesmo tipo)
			const (
				x = 10
				y = 10
				z = 10
			)

			var (
				x = 10
				y = 10
				z = 10
			)

		TODO Contantes não tipadas so terão un tipo atribuídos a elas quando forem
		usadas.

		TODO Ja tipo de uma variável e definido no momento da atribuição

	*/
	d = x
	fmt.Println(d)
}
