package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Pessoa struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Trabalho"`
	SaldoBancario float64 `json:"SaldoBancario"`
}

func main() {
	jamesBond := Pessoa{
		Nome:          "James",
		Sobrenome:     "Bond",
		Idade:         34,
		Profissao:     "Agente Secreto",
		SaldoBancario: 1020921.30,
	}

	encoder := json.NewEncoder(os.Stdout) // os.Stdout segue a inteface Writer
	// encode escreve direto na interface definida na criação do encoder,
	// sem precisar receber a informação e depois usar ela (como é feito com o marshal)
	encoder.Encode(jamesBond)

	fmt.Println("\n----------------------\n")
}
