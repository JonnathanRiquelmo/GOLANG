package main

import "fmt"

func main() {

	fmt.Println("Maps")

	// Parece com Struct, mas é mais rígido.
	// Conjunto de pares chave/valor, onde deve ser definido o tipo.
	// O tipo deve ser respeitado a cada inserção de chave/valor.
	user1 := map[string]string{
		"nome":      "Jonnathan",
		"sobrenome": "Riquelmo",
	}

	fmt.Println(user1)
	fmt.Println(user1["nome"])

	var2 := map[int]string{
		1: "Jonnathan",
		2: "Riquelmo",
	}

	fmt.Println(var2)
	fmt.Println(var2[1])

	var3 := map[string]int{
		"idade": 32,
		"resta": 3,
	}

	fmt.Println(var3)
	fmt.Println(var3["resta"])

	// Pode-se criar maps aninhados, como nos structs (ma a leitura fica mais confusa)
	var4 := map[int]map[int]string{
		1: {
			1: "primeiro item da primeira chave do map aninhado",
			2: "segundo item da primeira chave do map aninhado",
		},
		2: {
			1: "primeiro item da segunda chave do map aninhado",
		},
	}

	fmt.Println(var4)
	fmt.Println(var4[1])
	fmt.Println(var4[1][2])
	fmt.Println(var4[2])
	fmt.Println(var4[2][1])

	// Para deletar a chave de um map pode-se usar a função nativa <delete>
	delete(var4, 2)
	fmt.Println(var4[2])
	fmt.Print(var4)

	// Para adicionar novos registros no map deve-se respeitar a restrição de tipos (chaves) previamente estabelecido
	var4[3] = map[int]string{
		1: "tem algo aqui nessa nova chave no map aninhado",
	}

	fmt.Print(var4)

}
