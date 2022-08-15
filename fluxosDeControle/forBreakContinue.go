package main

import "fmt"

func main() {
	/*
		Loop for -> Break, continue
	*/
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			// entra aqui quando numero for não é impar
			continue
		}
		fmt.Println(i)
	}
}
