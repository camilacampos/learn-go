package main

import (
	"fmt"
)

// - Usando uma literal composta:
//     - Crie uma slice de tipo int
//     - Atribua 10 valores a ela
// - Utilize range para demonstrar todos estes valores.
// - E utilize format printing para demonstrar seu tipo.
func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for _, value := range slice {
		fmt.Println(value)
	}

	fmt.Println("\n----------------------\n")

	fmt.Printf("%T", slice)

	fmt.Println("\n----------------------\n")
}
