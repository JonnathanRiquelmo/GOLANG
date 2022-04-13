package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
	// se usa "go mod tidy" para remover eventuais dependencias não usada do .mod do projeto
)

func main() {
	fmt.Println("1 - Escrevendo do arquivo main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("jonnathan.riquelmo@gmail.com")
	fmt.Println(erro)
	erro = checkmail.ValidateFormat("123456")
	fmt.Println(erro)
}
