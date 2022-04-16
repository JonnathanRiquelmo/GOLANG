package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "p1"
	array1[2] = "p3"
	array1[4] = "p5"
	fmt.Println(array1)

	array2 := [...]int{1, 2, 3, 4, 5} // Limitação de tamanho
	fmt.Println(array2)

	slice := []int{40, 41, 42, 43, 44, 45} // Sem limitação de tamanho
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array2))
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 100)
	fmt.Println(slice)

	// Primeiro parametro (indice) é inclusivo, o segundo é exclusivo.
	slice = array2[1:4] // Podemos ler: pegue os elementos do indice 1 enquanto indice não for == 4
	fmt.Println(slice)

	slice2 := array1[0:3] // Quando um slice copia parte de um array na verdade não ocorre cópia, apenas referência
	fmt.Println(slice2)
	array1[1] = "|p2|"  // Ou seja, se alterarmos o array, o slice deste array tbm é modificado,
	fmt.Println(slice2) // afinal ele está apenas apontando para o array original

	// ARRAYS INTERNOS
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade

}
