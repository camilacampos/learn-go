package main

import (
	"fmt"
)

// - Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
//     - Crie um tipo para um struct chamado "pessoa"
//     - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
//     - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
//     - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
//     - Demonstre no seu código:
//         - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
//         - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
// - Se precisar de dicas, veja: https://gobyexample.com/interfaces

type humanos interface {
	falar()
}

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println("Oi, meu nome é", p.nome, "e eu tenho", p.idade, "anos.")
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	p := pessoa{"Daniela", 32}

	p.falar()    // existe um atalho de p.m() para (&p).m() pra que não precise fazer da forma descrita abaixo (exceção)
	(&p).falar() // é a mais correta de ser usada

	dizerAlgumaCoisa(&p)
	// dizerAlgumaCoisa(p) PAU: cannot use p (type pessoa) as type humanos in argument to dizerAlgumaCoisa:
	//                          pessoa does not implement humanos (falar method has pointer receiver)
}
