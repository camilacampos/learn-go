package main

import (
	"fmt"
)

// Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
func main() {
	number := 200

	fmt.Printf("%d\t%b\t%#x\n", number, number, number)
	fmt.Printf("%x\n", number) // print do hexa sem o 0x no começo (pode confundir com decimal)
}
