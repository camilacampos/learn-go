package main

import (
	"fmt"
)

func main() {
	// make(TIPO DO SLICE, TAMANHO INICIAL, CAPACIDADE DO ARRAY SUBJACENTE)
	slice := make([]int, 5, 10) // cria o slice inicial com os valores zero
	slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4

	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println("\n----------------------\n")

	// Como o tamanho inicial do slice é 5, não é possível referenciar indices a partir do 4
	// slice[5] = 10 // PAU: runtime error: index out of range [5] with length 5
	slice = append(slice, 5)

	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println("\n----------------------\n")

	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9) // Até aqui, é usado o mesmo array subjacente, aumentando o tamanho do slice, mas mantendo a mesma capacidade

	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println("\n----------------------\n")

	slice = append(slice, 10) // Ao adicionar um elemento e o tamanho estourar a capacidade, um novo array subjacente é criado com o dobro da capacidade

	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println("\n----------------------\n")
}
