package main

import "fmt"

const (
	x = iota
	y = iota
	z = iota
)

func main() {
	/*
		Iota -> São números inteiros que ainda não foram tipados. E so vai recebe
		o seu tipo assim que for usando(pode fazer em float).
	*/

	// Ira colocar valores sucessíveis inteiros não tipados.
	fmt.Println(x, y, z)
}
