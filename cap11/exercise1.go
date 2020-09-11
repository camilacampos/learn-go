package main

import (
	"fmt"
)

// - Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
//     - Nome
//     - Sobrenome
//     - Sabores favoritos de sorvete
// - Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete

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

	printPessoa(pessoa1)
	printPessoa(pessoa2)
}
