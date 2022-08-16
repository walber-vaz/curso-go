package main

import "fmt"

func main() {
	/*
		if - else if - else
	*/
	if x := 50; x > 100 {
		fmt.Println("chis é maior que cem")
	} else if x < 10 {
		fmt.Println("chis e manor que dez")
	} else {
		fmt.Println("chis não menor que dez e nem maior que cem")
	}
}
