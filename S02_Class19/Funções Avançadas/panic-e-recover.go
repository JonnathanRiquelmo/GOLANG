package main

import "fmt"

func recuperarExecucao() {
	// Se a função recover retornar algo diferente de nulo o sistema é retomado e evita o panic
	fmt.Println("Porgrama entrou em PANIC, tentando recuperar a execução...")
	if r := recover(); r != nil {
		fmt.Println("Programa recuperado!!!")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	// O DEFER sempre será chamado antes de uma função dar um return
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// A função PANIC basicamente para todo o fluxo de execução do programa
	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println("Funções PANIC e RECOVER")

	alunoAprovado(6, 6)
	fmt.Println("Pós execução")

}
