package main

import "fmt"

func main() {
	/*
		slice Multidimecional -> são slice dentro de slice
	*/

	ss := [][]int{
		// 0  1  2 // Indice
		{7, 8, 9}, // 0
		{1, 2, 3}, // 1
		{4, 5, 6}, //2s
	}

	// Pegando indice 1 do slice do slice 1
	fmt.Println(ss[1][1])
}
