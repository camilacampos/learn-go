package main

import (
	"fmt"
)

func main() {
	slice := []string{"banana", "maça", "jaca"}
	for indice, valor := range slice {
		fmt.Println("No índice", indice, "tem o valor", valor)
	}

	fmt.Println("\n----------------------\n")

	// slice[3] = "melancia" // PAU: runtime error: index out of range [3] with length 3
	slice = append(slice, "melancia")
	for indice, valor := range slice {
		fmt.Println("No índice", indice, "tem o valor", valor, "\b.") // \b dá um backspace
	}

}
