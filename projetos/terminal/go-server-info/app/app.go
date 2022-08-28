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

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "github.com",
				},
			},
			Action: getIps,
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
		fmt.Println(ip)
	}
}
