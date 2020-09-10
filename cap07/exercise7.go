package main

import (
	"fmt"
)

// - Crie um programa que demonstre o funcionamento da declaração if.
// - Utilizando a solução anterior, adicione as opções else if e else.

var number int

func main() {
	fmt.Println("Digita ae um número:")
	fmt.Scan(&number)

	if number > 10 {
		fmt.Println("O número digitado é maior que 10")
	} else if number == 10 {
		fmt.Println("O número digitado é 10 :O")
	} else {
		fmt.Println("O número digitado é menor que 10")
	}

}
