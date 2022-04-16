package main

import "fmt"

// Retorno variável é a possibilidade de dar nome ao retoro da função
// Normalmente essa função teria apenas int, int no segundo parenteses, indicando o retorno de dois inteiros
// Com retorno variável é possível melhorar a legibilidade da função
// No return não é preciso especificar o retorno, apenas RETURN sozinho
func calculos(n1, n2 int) (soma int, subtracao int) {

	soma = n1 + n2
	subtracao = n1 - n2

	return
}

func main() {

	fmt.Println("Função de Retorno Variável")
	fmt.Println(calculos(5, 2))

	resultSoma, resultSubtracao := calculos(15, 35)

	fmt.Println(resultSoma, resultSubtracao)
}
