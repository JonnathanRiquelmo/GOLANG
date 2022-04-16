package main

import "fmt"

var n int

// É uma função que é chamada antes mesmo da MAIN executar
// É possível ter uma função INIT por ARQUIVO
// Sidenote: me lembra a função before do Xtend
func init() {
	fmt.Println("0 - Executando a função INIT, antes da função MAIN...")
	n = 10
	fmt.Println(n)
}

func main() {
	fmt.Println("1 - Após a INIT, a função MAIN é executada...")
	fmt.Println(n * 10)
}
