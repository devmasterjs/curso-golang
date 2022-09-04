package main

import (
	"fmt"
	"time"
)

// função que encapsula uma chamada para uma 'go routine' e devolve um canal

func main() {
	canal := escrever("Olá Mundo!")
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
