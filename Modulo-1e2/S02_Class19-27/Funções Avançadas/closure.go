package main

import "fmt"

// Essa função retorna outra função
func closure() func() {
	texto := "Dentro da função CLOSURE..."

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	fmt.Println("Função CLOSURE")

	//CLOSURE são funções que referência variáveis que estão fora de seu corpo (escopo)
	// Serve como uma "mémoria" que guarda a referência de variáveis mesmo que elas tenham nomes iguais

	texto := "Dentro da função MAIN"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()

}
