package main

import "fmt"

func main() {
	/*
		Map -> Eles não são ordenados
	*/

	contatos := map[string]int{
		"Pedro": 55125678,
		"Ana":   45349889,
	}

	// Mostrando todos valores do map
	fmt.Println(contatos)
	// Mostrando um valor especifico
	fmt.Println(contatos["Ana"])

	// Adicionando um valo ao map
	contatos["João"] = 45454545

	fmt.Println(contatos)

	// Verificando se existe em key no map
	if isTrue, ok := contatos["Matheus"]; !ok {
		fmt.Println("não tem")
	} else {
		fmt.Println(isTrue, ok)
	}

}
