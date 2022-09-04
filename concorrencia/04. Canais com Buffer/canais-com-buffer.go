package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"
	// canal <- "Se rodar essa linha ocorrerá deadlock!!!"

	mensagem := <-canal
	fmt.Println(mensagem)

	mensagem2 := <-canal
	fmt.Println(mensagem2)

}
