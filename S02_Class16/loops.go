package main

import (
	"fmt"
	"time"
)

func main() {

	// Diferente de outras linguagens, onde estruturas de repetição podem ter diversas formas (FOR, WHILE, DO..WHILE, FOR..EACH, FOR..IN)
	// Em GO existe apenas o FOR
	fmt.Println("Estruturas de Repetição")

	i := 0

	for i < 5 {
		// time.Sleep(time.Second)
		fmt.Println("Incrementando i:", i)
		i++
	}

	for j := 0; j < 5; j++ {
		// time.Sleep(time.Second)
		fmt.Println("Incrementando j:", j)
	}

	nomes := [3]string{"Fulano", "Ciclano", "Beltrano"}

	// For para iterar, estilo for each do Java
	for indice, valor := range nomes {
		fmt.Println("Posição:", indice, "| Nome: ", valor)
	}

	// Para omitir um valor, quando se quer apenas um indice, é preciso usar underline ( _ ) para omitir.
	for _, valor := range nomes {
		fmt.Println("Nome: ", valor)
	}

	// Se iterar sobre Strings ele retorna o código equivalente da letra na tabela ASCII
	// Para pegar o valor é preciso castar usando string(letra)
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
		fmt.Println(indice, string(letra))
	}

	var1 := map[string]string{
		"nome":      "Jonnathan",
		"sobrenome": "Riquelmo",
	}

	// Range em cima de MAPs é possível, mas em Structs não
	for chave, valor := range var1 {
		fmt.Println(chave, valor)
	}

	for {
		time.Sleep(time.Second)
		fmt.Println("Executando indefinidademente...")
	}

}
