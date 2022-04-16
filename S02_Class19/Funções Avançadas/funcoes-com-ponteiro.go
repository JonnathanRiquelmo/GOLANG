package main

import "fmt"

func inverterSinalComCopia(num int) int {
	return num * -1
}

func inverterSinalComPonteiro(num *int) {
	*num = *num * -1

}

func main() {
	numero := 10

	fmt.Println(numero)                        // 10
	fmt.Println(inverterSinalComCopia(numero)) // Retorno da copia com sinal trocado, após o retorno a variável deixa de existir em memória

	fmt.Println(numero)               // A variavel original na realidade segue sendo 10, o sinal não foi trocado no endereço de memória origial
	inverterSinalComPonteiro(&numero) // Aqui se trabalha com o ponteiro, ou seja, parâmetro por nome (referência), diferente da passagem por valor (cópia)
	fmt.Println(numero)
}
