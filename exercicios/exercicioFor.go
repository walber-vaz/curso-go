package main

import "fmt"

func main() {
	/*
		Faça um loop dos números 33 a 122, e utilize format printing para demostrá-los
		como text/string.
	*/
	for i := 33; i <= 122; i++ {
		fmt.Printf("Valor: %d -> Unicode: %v\n ", i, string(rune(i)))
	}
}
