package main

import "fmt"

type conv int
var b conv = 10

func main() {
	x := 10
	// Converter tipo para int
	x = int(b)
	fmt.Println(x)
}