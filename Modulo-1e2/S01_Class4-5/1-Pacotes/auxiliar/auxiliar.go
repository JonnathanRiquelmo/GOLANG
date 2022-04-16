package auxiliar

import "fmt"

// função que começa com letra maiuscula é PUBLIC
// função que começa com letra minuscula é PRIVATE
func Escrever() {
	fmt.Println("2 - Escrevendo do arquivo do pacote auxiliar")
	EscrevendoMais()
	escrever2()
}

func EscrevendoMais() {
	fmt.Println("3 - E aqui se escreve de um método privado do arquivo do pacote auxiliar")
}
