package main

import (
	"fmt"
)

// - Crie um programa que demonstre o funcionamento da declaração if.

var number int

func main() {
	fmt.Println("Digita ae um número:")
	fmt.Scan(&number)

	if number > 10 {
		fmt.Println("O número digitado é maior que 10")
	}
}
