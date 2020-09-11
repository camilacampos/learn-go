package main

import (
	"fmt"
)

// - Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// - Demonstre os valores do map utilizando range.
// - Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.

type Pessoa struct {
	nome             string
	sobrenome        string
	saboresFavoritos []string
}

func printPessoa(pessoa Pessoa) {
	fmt.Println(pessoa.nome, pessoa.sobrenome)

	for _, sabor := range pessoa.saboresFavoritos {
		fmt.Println("\t", sabor)
	}

	fmt.Println("\n----------------------\n")
}

func main() {
	pessoa1 := Pessoa{
		nome:             "Camila",
		sobrenome:        "Campos",
		saboresFavoritos: []string{"limão", "maracujá"},
	}
	pessoa2 := Pessoa{
		nome:             "Daniela",
		sobrenome:        "Jardim",
		saboresFavoritos: []string{"flocos", "FLOCOS", "F L O C O S"},
	}

	fmt.Println("\n----------- EXERCÍCIO ANTERIOR -----------\n")

	printPessoa(pessoa1)
	printPessoa(pessoa2)

	fmt.Println("\n----------- EXERCÍCIO ATUAL -----------\n")

	// pessoas := map[string]Pessoa{
	// 	pessoa1.sobrenome: pessoa1,
	// 	pessoa2.sobrenome: pessoa2,
	// }

	pessoas := make(map[string]Pessoa)
	pessoas[pessoa1.sobrenome] = pessoa1
	pessoas[pessoa2.sobrenome] = pessoa2

	for key, value := range pessoas {
		fmt.Println("KEY:", key)
		fmt.Println("PESSOA:", value.nome, value.sobrenome)

		for _, value := range value.saboresFavoritos { // uma variável (value) pode ser sobrescrita apenas no escopo de um inner for
			fmt.Println("\t", value)
		}
		fmt.Println()
	}
}
