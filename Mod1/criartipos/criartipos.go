package main

import "fmt"

// Criando proprios tipos
type hotdog int
var b hotdog = 10

func main() {
	fmt.Printf("%T", b)
}