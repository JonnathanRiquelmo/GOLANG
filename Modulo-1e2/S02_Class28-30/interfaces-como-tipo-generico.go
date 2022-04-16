package main

import "fmt"

// É uma solução muito específica (usada por exemplo no método fmt.Println),
// mas melhor procurar não abusar (gambiarra)
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("INTERFACES COMO TIPOS GENÉRICOS")

	generica("Olá mundo")
	generica(666)
	generica(false)
	generica(55 + 45)

	// Esse seria um exemplo de gambiarra:
	mapa := map[interface{}]interface{}{
		1:      "palavra",
		true:   float64(69),
		"algo": 666,
	}

	fmt.Println(mapa)
}
