package main

import (
	"fmt"
)

// - Usando uma literal composta:
//      - Crie um array que suporte 5 valores to tipo int
//      - Atribua valores aos seus índices
// - Utilize range e demonstre os valores do array.
// - Utilizando format printing, demonstre o tipo do array.
func main() {
	var array [5]int
	array[0] = 10
	array[1] = 20
	array[2] = 30
	array[3] = 40
	array[4] = 50

	for _, value := range array {
		fmt.Println(value)
	}

	fmt.Println("\n----------------------\n")

	fmt.Printf("%T", array)

	fmt.Println("\n----------------------\n")

	// ALTERNATIVA
	array2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(array2)
}
