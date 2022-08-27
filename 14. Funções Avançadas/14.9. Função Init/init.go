package main

import (
	"fmt"
)

var n int

func init() {
	fmt.Println("Função Init")
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Valor de [n] inicializado pela função init:", n)
}
