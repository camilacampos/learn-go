package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"alfredo": 12983923,
		"joana":   18292183,
		"denis":   0,
	}
	fmt.Println(amigos)
	fmt.Println(amigos["joana"])

	amigos["gopher"] = 18299128
	fmt.Println(amigos)

	fmt.Println("\n----------------------\n")

	fantasma := amigos["fantasma"] // quando não acha uma chave no map, atribui o valor zero quando ela é buscada
	fmt.Println(fantasma)

	fantasma, okFantasma := amigos["fantasma"] // para diferenciar zero-value de não-existente, é possível usar
	denis, okDenis := amigos["denis"]          // uma segunda variável, que indica se a chave existe ou não
	fmt.Println("fantasma", fantasma, okFantasma)
	fmt.Println("denis", denis, okDenis)

	keys := []string{"someKey", "alfredo", "denis"}
	for _, key := range keys {
		if someValue, ok := amigos[key]; !ok {
			fmt.Println("não existe a chave:", key)
		} else {
			fmt.Println("existe a chave", key, "e o valor é:", someValue)
		}
	}

	// o ", ok" é conhecido como COMMA OK IDIOM ou apenas COMMA OK e pode ser usado pra mais coisas

	fmt.Println("\n----------------------\n")
}
