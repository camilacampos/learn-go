package main

import (
	"fmt"
)

// - Crie um tipo struct "pessoa" que contenha os campos:
//     - nome
//     - sobrenome
//     - idade
// - Crie um método para "pessoa" que demonstre o nome completo e a idade;
// - Crie um valor de tipo "pessoa";
// - Utilize o método criado para demonstrar esse valor.

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p Pessoa) toString() string {
	return fmt.Sprint("O nome da pessoa é: ", p.nome, " ", p.sobrenome, " e ela tem ", p.idade, " anos.")
}

func main() {
	pessoa := Pessoa{
		"Camila", "Campos", 27,
	}
	fmt.Println(pessoa.toString())
	fmt.Println("\n----------------------\n")
}
