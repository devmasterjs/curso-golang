package main

import (
	"fmt"
)

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Tipos Genéricos")
	generica("String")
	generica(1)
	generica(true)
}
