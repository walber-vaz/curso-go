package main

import (
	"fmt"
)

func main() {
	/*
		Crie um programa que utilize a declaração switch,
	*/

	horaDoAlmoco := 13

	switch {
	case horaDoAlmoco == 12:
		fmt.Printf("Hora do almoço: %v\tPode almoçar\n", horaDoAlmoco)
	case horaDoAlmoco < 12:
		fmt.Printf("Hora do almoço: %v\tAinda não e hora\n", horaDoAlmoco)
	case horaDoAlmoco > 12:
		fmt.Printf("Hora do almoço: %v\tPassou da hora\n", horaDoAlmoco)

	}
}
