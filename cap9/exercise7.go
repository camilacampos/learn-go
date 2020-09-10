package main

import (
	"fmt"
)

// - Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
//     - "Nome", "Sobrenome", "Hobby favorito"
// - Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
// desafio camila: usar um map para mostrar o tipo do atributo durante o for

func main() {
	dictionary := map[int]string{
		0: "Nome",
		1: "Sobrenome",
		2: "Hobby",
	}

	slice := [][]string{
		[]string{"Camila", "Campos", "dançar"},
		[]string{"Alimac", "Sopman", "cantar"},
		[]string{"Leonam", "Manoel", "bailar"},
	}

	for _, person := range slice {
		for key, attribute := range person {
			fmt.Println(dictionary[key], "-", attribute)
		}
		fmt.Println("----------------------")
	}
}
