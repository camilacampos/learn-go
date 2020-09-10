package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	// pessoinha  pessoa // => assim é possivel deixar com nome diferente
	pessoa  // => isso define um atributo pessoa dentro de profissional, que é do tipo pessoa definido acima
	titulo  string
	salario float64
}

func main() {
	pessoa1 := pessoa{nome: "Joana", idade: 30}
	pessoa2 := profissional{
		pessoa:  pessoa{nome: "Marcos", idade: 40},
		titulo:  "dev",
		salario: 124.5,
	}

	pessoa3 := profissional{pessoa{"Valdisney", 22}, "vereador", 126389127639816.99}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	fmt.Println(pessoa3)

	fmt.Println("\n----------------------\n")

	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.titulo)
	fmt.Println(pessoa3.pessoa.nome)
	fmt.Println(pessoa3.idade) // pode omitir o `.pessoa` caso não tenha conflito nos nomes dos atributos de cada struct - promoted fields

	fmt.Println("\n----------------------\n")
}
