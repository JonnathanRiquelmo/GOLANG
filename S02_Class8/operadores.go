package main

import "fmt"

func main() {
	//OPERADORES ARITMETICOS
	soma := 1 + 2
	subtracao := 3 - 1
	divisao := 10 / 2
	multiplicacao := 25 * 4
	restoDivisao := 5 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	// tipos diferentes não trabalham, mas é possível fazer Typecasting
	var nr1 int16 = 10
	var nr2 int32 = 50
	result := nr2 + int32(nr1)
	fmt.Println(result)
	result2 := nr1 + int16(nr2)
	fmt.Println(result2)

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)  // maior que
	fmt.Println(1 < 2)  // menor que
	fmt.Println(1 >= 2) // maior-igual
	fmt.Println(1 <= 2) // menor-igual
	fmt.Println(1 == 2) // igual
	fmt.Println(1 != 2) // diferente

	// OPERADORES LÓGICOS
	fmt.Println(true && true) // E
	fmt.Println(false && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	fmt.Println(true || true) // OU
	fmt.Println(false || true)
	fmt.Println(false || false)
	fmt.Println(true || false)

	fmt.Println(true != true) // NEGAÇÃO
	fmt.Println(false != true)

	// OPERADORES UNÁRIOS
	number := 10
	fmt.Println(number)
	number++
	fmt.Println(number)
	number += 9
	fmt.Println(number)
	number *= 12
	fmt.Println(number)
	number /= 5
	fmt.Println(number)
	number %= 2
	fmt.Println(number)

	// OPERADOR TERNÁRIO tipo como Java
	// e.g. var = param > 0 ? 1 : 0     ===== if param > 0 then var is 1 else var is 0
	// SE USA O CLÁSSICO IF-ELSE

	// ISSO É O MAIS PRÓXIMO
	val := 10

	index := val
	if val <= 0 {
		index = -val
	}

	fmt.Println(index)

}
