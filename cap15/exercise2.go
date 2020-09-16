package main

import (
	"fmt"
)

// - Crie um struct "pessoa"
// - Crie uma função chamada mudeMe que tenha *pessoa como parâmetro. Essa função deve mudar um valor armazenado no endereço *pessoa.
//     - Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
//     - Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors
//     - "As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method),
//     - → x.f is shorthand for (*x).f." ←
//     - Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
// - Crie um valor do tipo pessoa;
// - Use a função mudeMe e demonstre o resultado.

type Pessoa struct {
	nome string
}

func main() {
	pessoa := Pessoa{
		nome: "Um nome qualquer",
	}
	fmt.Println(pessoa.nome)
	mudeMe(&pessoa)
	fmt.Println(pessoa.nome)
}

func mudeMe(pessoa *Pessoa) {
	// pessoa.nome = "MUDEI" // EXCEÇÃO QUE FUNCIONA
	(*pessoa).nome = "MUDEI" // MANEIRE MAIS COMUM/CORRETA
}
