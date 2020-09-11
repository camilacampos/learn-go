package main

import (
	"fmt"
)

func alow() {
	defer fmt.Println("\n----------------------\n")
	defer fmt.Println("teste dentro do defer alow")

	fmt.Println("teste fora do defer alow")
}

func main() {
	defer fmt.Println("\nPortanto, o que for definido no primeiro defer é executado por último")
	defer fmt.Println("Quando tem mais de um, eles são executados de baixo pra cima")
	defer fmt.Println("\nO que for definido em defers é executado após a execução do resto do programa")

	alow()
	olar()
	fmt.Println("O corpo da função (sem defer) é executado primeiro")
	defer fmt.Println("O defer pode ser definido em qualquer lugar da função")
}
