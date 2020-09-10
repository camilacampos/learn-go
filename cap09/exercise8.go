package main

import (
	"fmt"
)

// - Crie um map com key tipo string e value tipo []string.
//     - Key deve conter nomes no formato sobrenome_nome
//     - Value deve conter os hobbies favoritos da pessoa
// - Demonstre todos esses valores e seus índices.

func main() {
	amigues := map[string][]string{
		"campos_camila": []string{"cantar", "dançar"},
		"korbes_ellen":  []string{"ensinar", "contar piada"},
	}

	for name, hobbies := range amigues {
		fmt.Println(name, "-", hobbies)
		for index, hobby := range hobbies {
			fmt.Println("\t", index, "-", hobby)
		}
	}
}
