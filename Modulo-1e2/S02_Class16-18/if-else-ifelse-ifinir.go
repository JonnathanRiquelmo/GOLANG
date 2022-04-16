package main

import "fmt"

func main() {
	fmt.Println("IF, ELSE, IF ELSE, IF INIT")

	number := 15

	if number > 15 {
		fmt.Println("É maior que 15")
	} else if number < 15 {
		fmt.Println("Menor que 15")
	} else {
		fmt.Println("É igual a 15")
	}

	// Em GO é possível inicializar variáveis em uma condição IF
	// Essas variáveis são limitas ao escopo do IF, fora dele ela não existe (impossível acessar após o teste)
	if otherNumber := 10; otherNumber > 0 {
		fmt.Println("Número maior que 0")
	}

	// É possível inicializar inclusive múltilpas variáveis
	if x, y := 5, 38; x == 5 {
		fmt.Printf("Whee! %d\n", y)
	}

}
