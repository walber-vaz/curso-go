package main

import "fmt"

func main() {
	/*
		Deletar um item de um map
	*/
	users := map[int]string{
		123: "Muito legal",
		456: "menos",
		677: "Esse é legal",
	}

	fmt.Println(users)

	total := 0

	for k := range users {
		total += k
	}
	fmt.Println(total)

	// Deletando key de um map
	delete(users, 123)
	fmt.Println(users)
}
