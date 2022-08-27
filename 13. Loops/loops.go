package main

import (
	"fmt"
	"time"
)

func sleep(timeSleep bool) {
	if timeSleep {
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("Loops")

	i := 0
	timeSleep := false

	for i < 10 {
		fmt.Println("Incrementando i:", i)
		i++
		sleep(timeSleep)
	}
	fmt.Println("Valor final de i:", i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j:", j)
		sleep(timeSleep)
	}

	nomes := [3]string{"John", "Mike", "Dave"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "John",
		"sobrenome": "Doe",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Executando infinitamente...", "[Ctrl+C] para encerrar")
		sleep(true)
	}
}
