package main

import (
	"encoding/json"
	"fmt"
)

// Em go, tudo que tem a letra maíscula é visível para outros pacotes que importarem o seu. Tudo que é minúsculo não pode ser importado.
// Portanto, para o json conseguir entender que campos devem ser serializados, os atributos precisam começar com letra maíuscula.
type Pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	SaldoBancario float64
}

func main() {
	jamesBond := Pessoa{
		Nome:          "James",
		Sobrenome:     "Bond",
		Idade:         34,
		Profissao:     "Agente Secreto",
		SaldoBancario: 1020921.30,
	}

	darthVader := Pessoa{
		"Anakin", "Skywalker", 50, "Vilão Intergalático", 501293042.30,
	}

	if j, err := json.Marshal(jamesBond); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(j))
	}

	if d, err := json.Marshal(darthVader); err != nil {
		fmt.Println(err)
	} else {
		// fmt.Println(d)         // d é []byte
		fmt.Println(string(d)) // para ver o json, é necessário converter em string
	}
	fmt.Println("\n----------------------\n")
}
