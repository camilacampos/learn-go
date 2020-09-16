package main

import (
	"fmt"
)

// - Crie uma função que retorne um int
// - Crie outra função que retorne um int e uma string
// - Chame as duas funções
// - Demonstre seus resultados

func main() {
	result := soma(2, 5)
	text, _ := generateText(result)
	fmt.Println(text)
}

func soma(x, y int) int {
	return x + y
}

func generateText(result int) (string, int) {
	return fmt.Sprintln("O resultado da soma é", result), 10
}
