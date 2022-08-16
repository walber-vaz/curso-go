package main

import "fmt"

var x interface{}

func main() {
	/*
		Como fazer switch pelo tipo da variável
	*/

	x = true
	switch x.(type) {
	case int:
		fmt.Println("É um int")
	case bool:
		fmt.Println("É um bool")
	case string:
		fmt.Println("É uma string")
	case float64:
		fmt.Println("É um float")
	}

}
