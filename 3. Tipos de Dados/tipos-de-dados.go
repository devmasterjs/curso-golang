package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero1 int8 = 100
	fmt.Println(numero1)

	var numero2 uint8 = 10
	fmt.Println(numero2)

	// alias
	// int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println((numero4))

	// FIM NUMEROS INTEIROS

	var numero5 float32 = 123.9
	fmt.Println(numero5)

	var numero6 float64 = 166464623.88889
	fmt.Println(numero6)

	// FIM NUMEROS REAIS

	var str1 string = "string 1"
	fmt.Println(str1)

	// FIM STRINGS

	var booleano bool = true
	fmt.Println(booleano)

	// FIM BOOLEANOS

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
