package main

import (
	"log"
)

// - Crie um struct "erroEspecial" que implemente a interface builtin.error.
// - Crie uma função que tenha um valor do tipo error como parâmetro.
// - Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.

type erroEspecial struct {
	customMessage string
}

func (err erroEspecial) Error() string {
	return "iiih, deu mto ruim aqui, em: " + err.customMessage
}

func main() {
	err := erroEspecial{"não era pra ser"}
	dealWith(err)
}

func dealWith(err error) {
	if err != nil {
		log.Panicln(err)
	}
	log.Println("tudo certo por aqui")
}
