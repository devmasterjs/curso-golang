package app

import "github.com/urfave/cli"

// Aplicação de linha de comando
func GoServerInfo() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidores na Internet"
	return app
}
