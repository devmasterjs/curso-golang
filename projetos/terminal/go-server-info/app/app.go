package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Aplicação de linha de comando
func GoServerInfo() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidores na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: getIps,
		},
		{
			Name:   "server",
			Usage:  "Busa o nome dos servidores na internet",
			Flags:  flags,
			Action: getServerNames,
		},
	}

	return app
}

// Busca IPs na internet
func getIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println("IP:", ip)
	}
}

// Busca Nomes dos Servidoes na Internet
func getServerNames(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println("Server:", server.Host)
	}
}
