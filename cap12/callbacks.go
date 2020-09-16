package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(soma(slice...)) // chamando uma função simples

	fmt.Println("\n----------------------\n")

	fmt.Println(selecionaImparesE(soma, slice...)) // chamando uma função simples

	fmt.Println("\n----------------------\n")
}

func soma(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

// callback: passar uma função como argumento de outra
func selecionaImparesE(operacao func(...int) int, numbers ...int) int {
	var slice []int
	for _, number := range numbers {
		if number%2 != 0 {
			slice = append(slice, number)
		}
	}
	return operacao(slice...)
}
