package main

import (
	"fmt"
)

// - Utilizando como base o slice do exercício anterior, utilize slicing para demonstrar os valores:
//     - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
//     - Do quinto ao último item do slice (incluindo o último item!)
//     - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
//     - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
//     - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println("1o até o 3o", slice[:3])
	fmt.Println("5o até último", slice[4:])
	fmt.Println("2o até o 7o", slice[1:7])
	fmt.Println("3o até o 9o (v1)", slice[2:9])
	fmt.Println("3o até o 9o (v2)", slice[2:len(slice)-1])
}
