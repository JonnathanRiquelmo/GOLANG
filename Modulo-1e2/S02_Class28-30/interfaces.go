package main

import (
	"fmt"
	"math"
)

// Extremamente similar as interfaces de Java, a diferença apenas é que a verificação do Implements
// se dá de forma implícita, ou seja,  como n usa classes, e o próprio Struct não tem algo similar,
// a verificação se dá na hora que o GO verifica a combinação da variavel (parametro) com o
// método invocado (deve ter a mesma assinatura que a assitura definida na interface)

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f \n", f.area())
}

func main() {
	fmt.Println("INTERFACES")

	r := retangulo{10, 15}
	escreverArea(r)
	c := circulo{10}
	escreverArea(c)
}
