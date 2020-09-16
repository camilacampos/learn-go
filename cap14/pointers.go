package main

import (
	"fmt"
)

func main() {
	x := 10
	y := &x
	fmt.Println("X:", x, " - Endereço:", &x, "- Valor (de-reference): ", *&x)
	fmt.Println("Y:", y, "- Valor (de-reference): ", *y)

	fmt.Println("\n----------------------\n")

	fmt.Printf("X: %T - %T\n", x, &x)
	fmt.Printf("Y: %T - %T\n", y, *y)

	fmt.Println("\n----------------------\n")

	*y++
	fmt.Println("X:", x, " - Endereço:", &x, "- Valor (de-reference): ", *&x)
	fmt.Println("Y:", y, "- Valor (de-reference): ", *y)

	fmt.Println("\n----------------------\n")

	passingByValue(x)
	fmt.Println("[VALOR][FORA DA FUNÇÃO] Valor de X: ", x)

	passingByReference(&x)
	fmt.Println("[PONTEIRO][FORA DA FUNÇÃO] Valor de X: ", x)
}

func passingByValue(x int) {
	x++ // aqui é uma cópia do valor do que foi passado como argumento
	fmt.Println("[VALOR][DENTRO DA FUNÇÃO] Valor de X: ", x)
}

func passingByReference(x *int) {
	*x++ // aqui é o próprio endereço da variável que foi passada como argumento
	fmt.Println("[PONTEIRO][DENTRO DA FUNÇÃO] Valor de X: ", *x)
}
