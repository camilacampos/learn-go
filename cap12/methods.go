package main

import (
	"fmt"
)

type Pessoa struct {
	nome            string
	anoDeNascimento int
}

// ESTRUTURA DE UMA FUNÇÃO (pra lembrar)
// func (receiver) identifier(parameters) (returns) { code }

// um método com receiver adiciona algum comportamento a um tipo (receiver)
func (p Pessoa) idade() int {
	return 2020 - p.anoDeNascimento
}

func main() {
	pessoa := Pessoa{
		nome:            "Marquinha",
		anoDeNascimento: 1988,
	}
	fmt.Println(pessoa.idade())
	fmt.Println("\n----------------------\n")
}
