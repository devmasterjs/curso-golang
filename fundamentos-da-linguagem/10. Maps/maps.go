package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "John",
		"sobrenome": "Doe",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"], usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "John",
			"ultimo":   "Doe",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["primeiro"], usuario2["nome"]["ultimo"])

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Leão",
	}
	fmt.Println(usuario2)

}
