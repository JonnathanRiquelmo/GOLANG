package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "(1) uma variável"
		variavel4 string = "em outra (2)"
	)

	fmt.Println(variavel3 + " concatenando " + variavel4)

	variavel5, variavel6 := "e + outra variável", "com another variável"
	fmt.Println(variavel5, variavel6)

	var pi string = "O número Pi (π) é um número irracional cujo valor é uma constante:"
	const constante1 float64 = 3.14159265358979323846
	fmt.Println(pi, constante1)

	// para inverter valores entre variáveis em GO não é necessário variáveis auxiliares
	variavel3, variavel4 = variavel4, variavel3
	fmt.Println(variavel3, variavel4)
}
