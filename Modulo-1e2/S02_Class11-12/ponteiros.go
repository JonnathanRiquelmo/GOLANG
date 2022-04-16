package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var v1 int = 10
	var v2 int = v1

	fmt.Println(v1)
	fmt.Println(v2)

	v1++
	// só muda a v1 um, pq o v2 na prática é apenas uma cópia para v1, não um ponteiro.
	fmt.Println(v1)
	fmt.Println(v2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var var3 int
	var pont *int

	fmt.Println(var3, pont)

	var3 = 100
	pont = &var3

	fmt.Println(var3, pont)
	fmt.Println(var3, *pont) // desreferenciação

}
