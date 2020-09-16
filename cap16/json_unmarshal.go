package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Trabalho"`
	SaldoBancario float64 `json:"SaldoBancario"`
}

func (p Pessoa) toString() string {
	return fmt.Sprintf(`{ nome: %v, sobrenome: %v, idade: %v, profissao: %v, saldoBancario: %v }`,
		p.Nome, p.Sobrenome, p.Idade, p.Profissao, p.SaldoBancario)
}

func main() {
	jsonDePessoas := []byte(`[
		{"Nome":"James","Sobrenome":"Bond","Idade":34,"Trabalho":"Agente Secreto","SaldoBancario":1020921.3},
		{"Nome":"Anakin","Sobrenome":"Skywalker","Idade":50,"Trabalho":"Vilão Intergalático","SaldoBancario":501293042.3}
	]`)

	var pessoas []Pessoa
	err := json.Unmarshal(jsonDePessoas, &pessoas)
	if err != nil {
		fmt.Println(err)
	}
	printPessoas(pessoas...)

	fmt.Println("\n----------------------\n")

	jsonDePessoa := []byte(`{"Nome":"Camila","Sobrenome":"Campos","Idade":25,"Trabalho":"Dev","SaldoBancario":3.0}`)
	var pessoa Pessoa

	err = json.Unmarshal(jsonDePessoa, &pessoa)
	if err != nil {
		fmt.Println(err)
	}
	printPessoas(pessoa)
}

func printPessoas(pessoas ...Pessoa) {
	for _, pessoa := range pessoas {
		fmt.Printf("%+v\n", pessoa.toString())
	}
}
