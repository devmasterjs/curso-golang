package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var usuario1 usuario
	usuario1.nome = "John Doe"
	usuario1.idade = 33
	fmt.Println(usuario1)

	enderecoComercial := endereco{"Rua dos Jaguaras", 171}

	usuario2 := usuario{"Billy Jean", 21, enderecoComercial}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 22}
	fmt.Println(usuario3)
}
