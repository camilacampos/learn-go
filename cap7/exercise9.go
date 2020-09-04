package main

import (
	"fmt"
)

// - Crie um programa que utilize a declaração switch, sem switch statement especificado.
var esporteFavorito string

func main() {
	fmt.Println("Qual o seu esporte favorito?")
	fmt.Scan(&esporteFavorito)

	switch esporteFavorito {
	case "basquete", "volei":
		fmt.Println("Gosta de uma bola aérea, eem")
	case "futebol":
		fmt.Println("Padrão demais")
	case "handball", "rugby":
		fmt.Println("Esse é pra quem curte umas pancadas, em hihihi")
	case "curling":
		fmt.Println("ah lá o cult")
	default:
		fmt.Println("non conosco")
	}
}
