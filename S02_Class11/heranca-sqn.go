package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    int
}

type estudante struct {
	pessoa
	matricula int
	curso     string
}

func main() {
	fmt.Println("Herança (Só que não)")

	p1 := pessoa{"Jonnathan", "Riquelmo", 32, 175}
	fmt.Println(p1)

	e1 := estudante{p1, 151150, "Engenharia de Software"}
	fmt.Println(e1)

}
