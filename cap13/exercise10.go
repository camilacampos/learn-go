package main

import (
	"fmt"
)

// - Demonstre o funcionamento de um closure.
// - Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.

func main() {
	interact := interaction(0)
	fmt.Println(interact())
	fmt.Println(interact())
	fmt.Println(interact())
	fmt.Println(interact())

	fmt.Println("\n\n")

	interact = interaction(2)
	fmt.Println(interact())
	fmt.Println(interact())
	fmt.Println(interact())
	fmt.Println(interact())
}

func interaction(initialValue int) func() int {
	x := initialValue - 1
	return func() int {
		x++
		return x
	}
}
