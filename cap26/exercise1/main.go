package main

import (
	"fmt"

	"github.com/camilacampos/learn-go/cap26/exercise1/cachorro"
)

// - Crie um package "cachorro".
//     - Este package deverá exportar uma função Idade, que toma como parâmetro um número de anos e retorna a idade equivalente em anos caninos. (1 ano humano → 7 anos caninos)
//     - Documente seu código com comentários, e utilize a função Idade na sua função main.
// - Rode seu programa para verificar se ele funciona.
// - Rode um local server com godoc e leia sua documentação.

func main() {
	n := 10
	fmt.Println(n, "anos humanos equivalem a", cachorro.Idade(n), "anos de cachorro.")
}
