package main

import (
	"fmt"
)

// Crie um programa que:
// - Atribua um valor int a uma variável
// - Demonstre este valor em decimal, binário e hexadecimal
// - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
// - Demonstre esta outra variável em decimal, binário e hexadecimal
func printNumber(text string, number int) {
	fmt.Printf("%v:\n%d\t%b\t%#x\n\n", text, number, number, number)
}

func main() {
	number := 10
	printNumber("Valor inicial", number)

	left := number << 1
	printNumber("1 bit deslocado a esquerda", left)

	right := number >> 1
	printNumber("1 bit deslocado a direita", right)
}
