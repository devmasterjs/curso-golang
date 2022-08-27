package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Log do arquivo main")
	auxiliar.Escrever()

	erro:= checkmail.ValidateFormat("email-invalido")
	fmt.Println(erro)
}