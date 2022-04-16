package main

import (
	"fmt"
)

func diaDaSemana(number int) string {

	switch number {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func modoAlternativo(number int) string {
	switch {
	case number == 1:
		return "Domingo"
	case number == 2:
		return "Segunda-Feira"
	case number == 3:
		return "Terça-Feira"
	case number == 4:
		return "Quarta-Feira"
	case number == 5:
		return "Quinta-Feira"
	case number == 6:
		return "Sexta-Feira"
	case number == 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(number int) string {
	var dia string

	switch {
	case number == 1:
		dia = "Domingo"
	case number == 2:
		dia = "Segunda-Feira"
	case number == 3:
		dia = "Terça-Feira"
	case number == 4:
		dia = "Quarta-Feira"
	case number == 5:
		dia = "Quinta-Feira"
	case number == 6:
		dia = "Sexta-Feira"
	case number == 7:
		dia = "Sábado"
		fallthrough // Função curiosa que faz o switch cair no código da próxima condição, sem avaliar de fato.
		//Ao invés de retornar Sábado, retorna Número Inválido, ou seja, a próxima condição (sem avaliar DE FATO).
	default:
		dia = "Número inválido"
	}
	return dia
}

func main() {
	fmt.Println("Switch")

	fmt.Println(diaDaSemana(7))
	fmt.Println(diaDaSemana(6))
	fmt.Println(diaDaSemana(5))
	fmt.Println(diaDaSemana(4))
	fmt.Println(diaDaSemana(3))
	fmt.Println(diaDaSemana(2))
	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana(99))

	fmt.Println(" ========== ")

	fmt.Println(modoAlternativo(1))
	fmt.Println(modoAlternativo(2))
	fmt.Println(modoAlternativo(3))
	fmt.Println(modoAlternativo(4))
	fmt.Println(modoAlternativo(5))
	fmt.Println(modoAlternativo(6))
	fmt.Println(modoAlternativo(7))
	fmt.Println(modoAlternativo(66))

	fmt.Println(" ========== ")

	fmt.Println(diaDaSemana2(7))
	fmt.Println(diaDaSemana2(6))
	fmt.Println(diaDaSemana2(5))
	fmt.Println(diaDaSemana2(4))
	fmt.Println(diaDaSemana2(3))
	fmt.Println(diaDaSemana2(2))
	fmt.Println(diaDaSemana2(1))
	fmt.Println(diaDaSemana2(99))

}
