package main

import "fmt"

func main() {
	filaTarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(filaTarefas, resultados)
	go worker(filaTarefas, resultados)
	go worker(filaTarefas, resultados)
	go worker(filaTarefas, resultados)
	go worker(filaTarefas, resultados)
	go worker(filaTarefas, resultados)

	for i := 0; i < 45; i++ {
		filaTarefas <- i
	}
	close(filaTarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
