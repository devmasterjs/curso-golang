package main

import (
	"fmt"
	"go-server-info/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Go Server Info")

	application := app.GoServerInfo()
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
