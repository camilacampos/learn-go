package main

import (
	"fmt"
)

// Uma interface é um conjunto de métodos
// Todo tipo que criarmos que tiver os métodos da interface, vai implementar a interface (não tem declaração explícita)

type Pessoa struct {
	nome            string
	anoDeNascimento int
}

type Dentista struct {
	Pessoa
	especialização string
	salário        float64
}

type Engenheiro struct {
	Pessoa
	tipo                  string
	trabalhaComEngenharia bool
}

func (p Dentista) oiBomDia() {
	fmt.Println("Meu nome é", p.nome, ". Eu ganho R$", p.salário, "por dia e estou te desejando bom dia!")
}

func (p Engenheiro) oiBomDia() {
	msgTrabalhoNaArea := "trabalho com engenharia"

	if !p.trabalhaComEngenharia {
		msgTrabalhoNaArea = "não " + msgTrabalhoNaArea
	}
	fmt.Println("Meu nome é", p.nome, ". Eu", msgTrabalhoNaArea, "e estou te desejando bom dia!")
}

// Não é possível criar um método NA interface (aqui vai apenas as definições)
type Profissão interface {
	oiBomDia()
}

func serHumano(p Profissão) {
	p.oiBomDia()

	switch p.(type) {
	case Dentista:
		fmt.Println("Minha especialização é:", p.(Dentista).especialização)
	case Engenheiro:
		fmt.Println("Minha engenharia é:", p.(Engenheiro).tipo)
	}
}

func main() {
	thais := Dentista{
		Pessoa: Pessoa{
			nome:            "Thais",
			anoDeNascimento: 1992,
		},
		especialização: "ortodontia",
		salário:        1000.4,
	}

	caue := Engenheiro{
		Pessoa: Pessoa{
			nome:            "Cauê",
			anoDeNascimento: 1994,
		},
		tipo:                  "Produção",
		trabalhaComEngenharia: false,
	}

	thais.oiBomDia()
	caue.oiBomDia()

	fmt.Println("\n------------------\n\n")

	serHumano(thais)
	fmt.Println()
	serHumano(caue)

	fmt.Println("\n------------------\n\n")
}
