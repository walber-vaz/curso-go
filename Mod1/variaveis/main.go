// Todo programa em pode ser divido em pacotes
package main

// Area onde é importado os pacotes
// Ref: https://pkg.go.dev/fmt
import "fmt"

// Todos programa go tem que ter uma ponto de partida
// tem que ter uma função main
func main() {
	/*
		Doc
		
		Anotação pacote.indetificador
		Por padrão a função println retorna o numero de byte(interger) e o error(string)
		Println e uma função variadica pode recebe any argumentos de qualquer Tipo

	*/

	// Se não queremos um valor do retorno colocamos no lugar do nome da varaiavel um _
	// Ref: https://go.dev/doc/effective_go#blank
	_, error := fmt.Println("Hello, World")
	fmt.Println(error)
}