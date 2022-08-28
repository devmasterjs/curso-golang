package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println("[numeros] é um slice, e pode ser iterado:", numeros)
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// Parâmetros varáticos devem ser os últimos parâmetros da função
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3)
	fmt.Println("Total da soma:", totalDaSoma)

	escrever("Olá Mundo", 1, 2, 3, 4, 5, 6, 7, 8, 9)
}
