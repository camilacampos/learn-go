package main

import (
	"fmt"
)

// - Crie uma função que retorna uma função.
// - Atribua a função retornada a uma variável.
// - Chame a função retornada.

func main() {
	sum := calculator("+")
	fmt.Println(sum(1, 2, 3))

	invalidOperation := calculator("|")
	fmt.Println(invalidOperation())
}

func calculator(operacao string) func(...int) int {
	switch operacao {
	case "+":
		return func(numbers ...int) int {
			total := 0
			for _, number := range numbers {
				total += number
			}
			return total
		}
	default:
		return func(...int) int {
			fmt.Println("Operação não definida")
			return 0
		}
	}

}
