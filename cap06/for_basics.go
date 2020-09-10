package main

import (
	"fmt"
)

func main() {
	// inicialização; condição; pós
	for x := 0; x < 3; x++ {
		fmt.Println(x)
	}

	// nested loops
	fmt.Println("\n------------------------------------\n")
	for hours := 0; hours < 12; hours++ {
		fmt.Println("Hora: ", hours)
		for minutes := 0; minutes < 60; minutes++ {
			fmt.Print("  ", minutes)
		}
		fmt.Println()
	}

	// go não tem while, mas dá pra fazer algo parecido com o for
	fmt.Println("\n------------------------------------\n")
	fmt.Println("Sem inicialização")
	i := 0
	for ; i < 2; i++ {
		fmt.Print("  ", i)
	}

	fmt.Println("\nSem inicialização e iteração")
	i = 0
	for i < 4 {
		fmt.Print("  ", i)
		i++
	}

	// for pode ser sem condição (loop infinito). mas pode ser parado com break
	fmt.Println("\n------------------------------------\n")

	i = 0
	for {
		if i < 10 {
			i++
			fmt.Print("  ", i)
		} else {
			break
		}
	}
}
