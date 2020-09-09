package main

import (
	"fmt"
)

func main() {
	primeiroSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(primeiroSlice)
	fmt.Println("\n----------------------\n")

	segundoSlice := append(primeiroSlice[:2], primeiroSlice[4:]...)
	fmt.Println(segundoSlice)
	// como tinha espaço para mudar o primeiro slice (tava dentro da capacidade), então o append modifica o array subjacente do primeiro, alocando o segundo.
	// isso faz com que o segundo slice tenha os items 1, 2 e 5, mas como o tamanho é menor que 5, então o array subjacente é alterado; e, portanto,
	// o primeiro slice passa a começar com o segundo slice (ficando [1, 2, *5*, 4, 5], com o 5 no lugar do 3)
	fmt.Println(primeiroSlice)
	fmt.Println("\n----------------------\n")

	// RECOMENDAÇÃO PARA RESOLVER ESSE PROBLEMA:
	// - Usar um for loop para copiar os itens do primeiro slice
	// - Não gerar novas variáveis ao mudar o tamanho de um slice (fazer primeiroSlice = append(...) ao invés de segundoSlice := append(...))
}
