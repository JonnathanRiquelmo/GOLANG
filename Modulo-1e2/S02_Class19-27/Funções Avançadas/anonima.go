package main

import (
	"fmt"
)

func main() {
	fmt.Println("Função Anônima")

	// De forma resumida, uma função anônima é uma função que NÃO TEM NOME. Ela é declarada e imediatamente após é executada.
	func() {
		fmt.Println("Olá mundo!")
	}()

	// É possível passar parâmetros também.
	func(n1 int, n2 int, texto string) {
		fmt.Println(n1 + n2)
		fmt.Println(texto)
	}(5, 10, "Olha um parametro aqui!")

	// Funções anônimas ainda podem ter retorno definido (return)
	retorno := func(texto string) string {
		return fmt.Sprintf("RECEBIDO: %s %d", texto, 100)
	}("Olá mundo!")

	fmt.Println(retorno)

}
