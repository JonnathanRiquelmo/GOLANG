package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1...")
}

func funcao2() {
	fmt.Println("Executando a função 2...")
}

func funcao3() {
	fmt.Println("Executando a função 3...")
}

func alunoAprovado(n1, n2 float64) bool {

	fmt.Println("Esperando a função verificar se o aluno está aprovado...")

	defer fmt.Println("Média calculada! O resultado será retornado...")

	media := (n1 + n2) / 2

	if media >= 6 {
		// fmt.Println("Média calculada! O resultado será retornado...")
		return true
	}

	// fmt.Println("Média calculada! O resultado será retornado...")

	return false
}

func main() {

	funcao1()
	funcao2()

	fmt.Println(" --- ")
	// Defer em inglês quer dizer Adiar, isso significa que a função será adiada até o último momento possível
	// É um recurso útil para, por exemplo, usar em funções que lidam com bancos de dados
	defer funcao1()
	funcao2()
	funcao3()

	fmt.Println(alunoAprovado(6, 6))
}
