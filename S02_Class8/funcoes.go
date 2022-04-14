package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

// Go foi construída para suportar múltiplos valores de retorno.
// ps: outra de declarar parametros de funcoes de Go
func calculosMatematicos(n1, n2 int) (int, int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2

	return soma, subtracao, multiplicacao
}

// Funções em Go também são consideradas tipos de dados, ou seja, podem ser atribuídas a variáveis que a executem.
func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função f")
	}

	f()

	var h = func(txt string) {
		fmt.Println(txt)
	}

	h("texto da função h")

	rSoma, rSub, rMult := calculosMatematicos(10, 50)
	fmt.Println("Soma:", rSoma, "| Subtracao:", rSub, "| Multiplicacao:", rMult)

	// uma função com múltiplos retornos pode ser usada mesmo que não se queira todos os valores retornados
	// para isso basta saber a ordem do retorno (o que cada valor é) e omitir aqueles que não se quer com _ (underline)
	// A ORDEM DOS FATORES ALTERA O PRODUTO (MUITO)
	rSoma2, _, rMult2 := calculosMatematicos(10, 15)
	fmt.Println("Soma2:", rSoma2, "| Multiplicacao2:", rMult2)

}
