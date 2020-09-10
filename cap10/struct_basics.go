package main

import (
	"fmt"
)

type cliente struct {
	nome      string
	sobrenome string
	éFumante  bool
}

func main() {
	cliente1 := cliente{
		nome:      "João",
		sobrenome: "Batista",
		éFumante:  false,
	}

	cliente2 := cliente{"Silas", "Soares", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)

	fmt.Println("\n----------------------\n")

	fmt.Println(cliente1.nome)
	fmt.Println(cliente2.sobrenome)

	fmt.Println("\n----------------------\n")
}
