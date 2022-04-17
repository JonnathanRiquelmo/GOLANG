package main

import (
	"app-terminal/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting App...")

	/* Comando para rodar o app, tem como argumento um paramêtro padrão para que os
	*  comandos do sistema operacional sejam reconhecidos pelo terminal
	 */
	aplicacao := app.Gerar()

	/* Ao rodar a aplicação pode retornar um erro, que deve ser tratado, nesse caso
	*  exibe o log
	 */
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
