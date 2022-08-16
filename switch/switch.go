package main

import "fmt"

func main() {
	/*
		Switch - ter vários casos diferentes
		Ele avaliar as expressões e escolhe qual executar com esta correta

		fallthrough -> ele executar sempre o proximo case.
		case default -> se não entrou em nenhum ele executar o default
	*/
	// x := 10

	// switch {
	// case x < 5:
	// 	fmt.Println("chis é menor que cinco")
	// case x == 5:
	// 	fmt.Println("chis é igual a cinco")
	// case x > 5:
	// 	fmt.Println("chis é maior que cinco")
	// }

	// Swicth default statement
	// switch true {
	// case x == 5:
	// 	fmt.Println("chis é igual a cinco")
	// case x == 5:
	// 	fmt.Println("chis é igual a cinco")
	// case x > 5:
	// 	fmt.Println("chis é maior que cinco")
	// }

	// responsavelDaAguaHJ := "w2k"

	// switch responsavelDaAguaHJ {
	// case "Pedro":
	// 	fmt.Println("Pedro é responsável pela água.")
	// case "Walber":
	// 	fmt.Println("Walber é responsável pela água.")
	// case "w2k":
	// 	fmt.Println("w2k é responsável pela água.")
	// }

	// Exemplo com fallthrough
	// responsavelDaAguaHJ := "Walber"

	// switch responsavelDaAguaHJ {
	// case "Pedro":
	// 	fmt.Println("Pedro é responsável pela água.")
	// case "Walber":
	// 	fmt.Println("Walber é responsável pela água.")
	// 	fallthrough
	// case "w2k":
	// 	fmt.Println("Sempre w2k ajuda ele.")
	// }

	// Default
	responsavelDaAguaHJ := "Fechado"

	switch responsavelDaAguaHJ {
	case "Pedro":
		fmt.Println("Pedro é responsável pela água.")
	case "Walber":
		fmt.Println("Walber é responsável pela água.")
		fallthrough
	case "w2k":
		fmt.Println("Sempre w2k ajuda ele.")
	default:
		fmt.Println("Hoje esta fechado.")
	}
}
