package main

import "fmt"

type user struct {
	nome     string
	idade    int
	endereco address
}

type address struct {
	rua    string
	numero int
}

func main() {
	fmt.Println("Arquivos de Structs")

	var u1 user
	u1.nome = "Fulano"
	u1.idade = 69
	fmt.Println(u1)

	// dá pra aninhar structs dentro de structs (nesse caso, ADDRESS dentro de USER)
	e1 := address{"rua da saudade", 124}
	// u2 := user{"Cliclado", 32}
	u2 := user{"Cliclado", 32, e1}
	fmt.Println(u2)

	// pode-se omitir valores, nesse caso o struct é criado apenas com nome (a idade fica com valor 0)
	u3 := user{nome: "Beltrano"}
	fmt.Println(u3)

}
