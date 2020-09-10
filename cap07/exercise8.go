package main

import (
	"fmt"
)

// - Crie um programa que utilize a declaração switch, sem switch statement especificado.
var nivelDeCansaço int

func main() {
	fmt.Println("Nível de cansaço:")
	fmt.Scan(&nivelDeCansaço)

	switch {
	case nivelDeCansaço == 0:
		fmt.Println("Tá de boas")
	case nivelDeCansaço == 1:
		fmt.Println("iiih, melhor dormir mais cedo")
	case nivelDeCansaço == 2:
		fmt.Println("mais morta que isso impossível")
	}
}
