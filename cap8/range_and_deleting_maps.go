package main

import (
	"fmt"
)

func main() {
	expectativas := map[int]string{
		1: "muito acima das expectativas",
		2: "acima das expectativas",
		3: "dentro das expectativas",
		4: "abaixo das expectativas",
		5: "muito abaixo das expectativas",
	}
	fmt.Println(expectativas)
	fmt.Println("\n----------------------\n")

	// range de mapas não tem ordem, sempre sai algo aleatório
	for key, value := range expectativas {
		fmt.Println(key, value)
	}
	fmt.Println("\n----------------------\n")

	// delete(map, key)
	delete(expectativas, 1) // não existe ninguém acima da expectativa
	fmt.Println(expectativas)
}
