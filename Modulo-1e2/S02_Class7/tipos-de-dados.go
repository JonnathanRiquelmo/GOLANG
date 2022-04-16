package main

import (
	"errors"
	"fmt"
)

func main() {

	// NUMEROS INTEIROS
	var numero1 int8 = 111
	var numero2 int16 = 22222
	var numero3 int32 = 333333333
	var numero4 int64 = 4444444444444444444
	var numero5 int = 5555555555555555555 //usa a arquiterura do pc (no meu caso 64 bits)
	fmt.Println(numero1, numero2, numero3, numero4, numero5)

	var number uint = 100 //uint significa Unsigned. Significa sem sinal, não assinalado, um int que desconsidera o sinal.
	fmt.Println(number)

	// Alias
	// int32 = RUNE
	var number2 rune = 123456
	fmt.Println(number2)

	// uint8 = BYTE
	var number3 byte = 123
	fmt.Println(number3)

	// NUMEROS REAIS
	var numeroreal1 float32 = 123.45
	fmt.Println(numeroreal1)
	var numeroreal2 float64 = 123.45131123123212312312312312312319
	fmt.Println(numeroreal2)

	// STRINGS
	//Go não tem char (único caractere)

	var str string = "exemplo 1"
	fmt.Println(str)
	str2 := "exemplo 2"
	fmt.Println(str2)

	//Similar a char (ele mapeia o caractere para um número da tabela ASCII)
	char := 'X'
	fmt.Println(char)

	// Conceito de valor 0 em variáveis Go
	var variavel0Texto string
	fmt.Println(variavel0Texto) //retorna espaço em branco
	var variavel0Numero int
	fmt.Println(variavel0Numero) //retorna sempre 0

	// BOOLEAN
	falso := false
	fmt.Println(falso)
	var verdadeiro bool = true
	fmt.Println(verdadeiro)
	var booleanoValor0 bool
	fmt.Println(booleanoValor0) //retorna sempre FALSE

	// ERRO (erro em Go é um tipo de dado)
	var erro error // valor 0 dele é nil ( ou <nil> ), similar ao null do Java por exemplo
	fmt.Println(erro)
	// para criar erros em Go (similar a exceptions do Java) usa-se o pacote errors
	erro = errors.New("Erro Novo (diferente de nil)")
	fmt.Println(erro)
	erro = nil // novamente, nil é similar a null
	fmt.Println(erro)

}
