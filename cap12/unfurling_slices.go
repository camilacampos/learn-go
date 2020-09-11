package main

import (
	"fmt"
)

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func main() {
	valor := soma(1, 2, 3, 4)
	fmt.Println(valor)
	fmt.Println("\n----------------------\n")

	slice := []int{1, 2, 3, 4}
	valor = soma(slice...)
	fmt.Println(valor)
	fmt.Println("\n----------------------\n")

	// PARAMETROS VARIÁDICOS RECEBEM DE 0 A N ARGUMENTOS
	valor = soma()
	fmt.Println(valor)
	fmt.Println("\n----------------------\n")
}
