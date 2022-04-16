package main

import "fmt"

type user struct {
	nome  string
	idade int
}

// GO não tem suporte a OO, mas oferece suporte a métodos para Structs
// De certa forma isso simula o comportamento que existe por ex no Java
// A declaração se dá como a seguir
func (u user) metodoDaStructUser() {
	fmt.Printf("Testando o método da struct com o nome %v e com idade %v\n", u.nome, u.idade)
}

// Passando o endereço de memória é possível alterar o valor do atributo da Struct
func (u *user) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("MÉTODOS")

	u1 := user{"Jonnathan", 32}

	fmt.Println(u1)
	u1.metodoDaStructUser()
	u1.fazerAniversario() // chama método que usa a referecia ao endereço de memória
	fmt.Println(u1)

}
