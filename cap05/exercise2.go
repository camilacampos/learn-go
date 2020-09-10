package main

import (
	"fmt"
)

// Escreva expressões utilizando os operadores de comparação, e atribua seus valores a variáveis.

const x = 10
const y = 20

func main() {
	igual := (x == y)
	diferente := (x != y)
	maiorOuIgual := (x >= y)
	maior := (x > y)
	menorOuIgual := (x <= y)
	menor := (x < y)

	fmt.Printf("X:\t\t|%v\n", x)
	fmt.Printf("Y:\t\t|%v\n", y)
	fmt.Println("\t\t|")
	fmt.Printf("Igual:\t\t|%v\n", igual)
	fmt.Printf("Diferente:\t|%v\n", diferente)
	fmt.Printf("Maior ou Igual:\t|%v\n", maiorOuIgual)
	fmt.Printf("Maior:\t\t|%v\n", maior)
	fmt.Printf("Menor ou Igual:\t|%v\n", menorOuIgual)
	fmt.Printf("Menor:\t\t|%v\n", menor)
}
