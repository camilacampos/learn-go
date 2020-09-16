package main

import (
	"fmt"
)

// - Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
// - Passe um valor do tipo slice of int como argumento para a função;
// - Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
// - Passe um valor do tipo slice of int como argumento para a função.

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// podemos chamar funções com parametros variádicos usando slice... ou passando diretamente os parâmetros
	fmt.Println(sum1(slice...))
	fmt.Println(sum1(1, 2, 3, 4, 5, 6, 7, 8, 9))

	fmt.Println("\n----------------------\n")

	fmt.Println(sum2(slice))
}

func sum1(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func sum2(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}
