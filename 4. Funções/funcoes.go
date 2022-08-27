package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subracao := n1 - n2
	return soma, subracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return "retorno = " + txt
	}

	resultado := f("Texto de f(txt string)")
	fmt.Println(resultado)

	resultadoSoma1, resultadoSubtracao1 := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma1, resultadoSubtracao1)

	resultadoSoma2, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma2)
}
