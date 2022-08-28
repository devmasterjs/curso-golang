package main

import (
	"fmt"
	"time"
)

func main() {
	// CANAIS == OPERAÇÃO BLOQUEANTE

	fmt.Println("Canais")
	canal := make(chan string)

	go escrever("Olá Mundo!", canal)

	// for {
	// 	mensagem, open := <-canal
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
