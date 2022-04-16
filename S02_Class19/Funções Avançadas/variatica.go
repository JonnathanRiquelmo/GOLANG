package main

import (
	"fmt"
)

// Uma função variática possibilita que uma função receba N paramêtros, inclusive 0 parametros
func soma(numbers ...int) int {

	total := 0

	for i, n := range numbers {
		fmt.Println(i)
		total += n
	}

	return total
}

// Existe a limitação de uma função variática somente poder ter um paramêtro variático por vez, e ele OBRIGATÓRIAMENTE DEVE SER O ÚLTIMO
func escrever(palavra string, numero ...int) {
	println()
	for _, letra := range palavra {
		fmt.Println(string(letra))
	}
	println()
	for _, n := range numero {
		fmt.Println(n)
	}

}

func main() {
	fmt.Println("Função Variática")

	// fmt.Println(soma(5, 10, 15, 20))

	escrever("Ola mundo!", 1, 2, 3, 4, 5, 6)
}
