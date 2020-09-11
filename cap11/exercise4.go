package main

import (
	"fmt"
)

// - Crie e use um struct anônimo.
// - Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.

func main() {
	agenda := struct {
		numeros map[string]int
		nomes   []string
	}{
		numeros: map[string]int{
			"home": 12345678,
			"cel":  987654321,
		},
		nomes: []string{"Camila", "Camis", "Cá"},
	}

	fmt.Println(agenda)
}
