package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada no main.go
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Terminal Application"
	app.Usage = "Returns IPs and SERVER NAMES on the internet"

	flags := []cli.Flag{
		//é a definição da flag em um slice, a flag é o que se passa como parãmetro
		cli.StringFlag{
			//nome da flag
			Name: "host",
			//valor padrão executado na flag caso nada seja passado no paramêtro
			Value: "lesse.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			//nome do comando
			Name: "ip",
			//descrição de uso do comando definição dos parâmetros que podem ser passados ao comando
			Usage: "Search ip addresses on the internet",
			Flags: flags,
			//execução do comando, define uma ação válida do app
			Action: searchIps,
		},
		// Segundo comando
		{
			Name:   "server",
			Usage:  "Search the server(s) name on the internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	//para buscar é necessário salvar o host que está sendo passado
	host := c.String("host")

	//pacote NET para armazenar os ips em um slice de ips (pq um host pode ter mais de um ip público)
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Address(es) found for the host", host)
	fmt.Println("-------------------------------")
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host) // name servers
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Server(s) found for the host", host)
	fmt.Println("-------------------------------")
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
