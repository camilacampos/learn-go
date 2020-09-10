package main

import (
	"fmt"
)

// - Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
func printAmigues(amigues map[string][]string) {
	for name, hobbies := range amigues {
		fmt.Println(name, "-", hobbies)
		for index, hobby := range hobbies {
			fmt.Println("\t", index, "-", hobby)
		}
	}
}

func main() {
	fmt.Println("\n----------- EXERCÍCIO 8 (ANTERIOR DO ANTERIOR) -----------\n")
	amigues := map[string][]string{
		"campos_camila": []string{"cantar", "dançar"},
		"korbes_ellen":  []string{"ensinar", "contar piada"},
	}
	printAmigues(amigues)

	fmt.Println("\n----------- EXERCÍCIO 9 (ANTERIOR) -----------\n")

	amigues["jardim_daniela"] = []string{"estudar", "encher o saco", "dançar"}

	printAmigues(amigues)

	fmt.Println("\n----------- EXERCÍCIO 10 (ATUAL) -----------\n")

	delete(amigues, "korbes_ellen")

	printAmigues(amigues)
}
