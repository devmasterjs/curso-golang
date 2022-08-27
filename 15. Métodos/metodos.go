package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário '%s' no banco de dados", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func main() {
	fmt.Println("Métodos")

	usuario1 := usuario{"John Doe", 21}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Willy Wonka", 55}
	fmt.Println(usuario2)
	usuario2.salvar()

}
