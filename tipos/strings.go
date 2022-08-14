package main

import "fmt"

func main() {
	/*
		Tipo Strings -> sequencias de bytes
	*/
	s := "Olá, Mundo!!"
	// sb := []byte(s)
	// fmt.Printf("%v\n , %T\n", s, s)
	// Convertendo strings para bytes
	// fmt.Printf("%v\n , %T\n", sb, sb)

	// Em Range com string ele vai mostra caractere po r caractere
	for _, v := range s {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}

	// Em for loop string ele vai mostra byte a byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}
}
