package main

import "fmt"

func fibonnaci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonnaci(posicao-2) + fibonnaci(posicao-1)

}

func main() {
	fmt.Println("Função Recursiva")

	//Sequência de fibonnaci
	//0 1 1 2 3 5 8 13 ...
	fmt.Println("Posição[3] :", fibonnaci(3)) // retorna 2
	fmt.Println("Posição[5] :", fibonnaci(5)) // retorna 5
	fmt.Println("Posição[7] :", fibonnaci(7)) // retorna 13

	posicao := 13

	for i := int(0); i < posicao; i++ {
		fmt.Println(fibonnaci(i))
	}
}
