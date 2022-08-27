package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá Mundo")
	}()

	func(texto string) {
		fmt.Println(fmt.Sprintf("Recebido ==> %s", texto))
	}("Olá Mundo")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido ==> %s", texto)
	}("Olá Mundo")

	fmt.Println(retorno)

}
