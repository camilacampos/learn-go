package main

import (
	"fmt"
)

func main() {
	sabores := []string{"4 queijos", "mussarela", "pepperoni", "abobrinha", "escarola"}

	// slice[FROM_INCLUSIVE:TO_EXCLUSIVE] (não inclui o slice no índice TO no slice do slice)
	// O último item na hora de cortar é identificado pelo tamanho do slice

	// fatia := sabores[2:5]
	// fatia := sabores[2:len(sabores)]
	// fatia := sabores[2:] // se omitir o TO, vai até o final
	fatia := sabores[:3] // se omitir o FROM, vai do começo até o índice do TO

	fmt.Println(fatia)
	fmt.Println("\n----------------------\n")

	fatia2 := sabores[:] // Todos os items do slice
	fmt.Println(fatia2)
	fmt.Println("\n----------------------\n")

	// PERCORRENDO SLICES COM FOR
	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}
	fmt.Println("\n----------------------\n")

	for _, sabor := range sabores {
		fmt.Println(sabor)
	}
	fmt.Println("\n----------------------\n")

	// PARA DELETAR UM VALOR DE UM SLICE:
	// 1. PEGAR OS VALORES ATÉ O VALOR E OS VALORES APÓS O VALOR
	// 2. JUNTAR OS DOIS "SUBSLICES" COM O APPEND
	// EXEMPLO: remover "pepperoni"

	fatia3 := append(sabores[:2], sabores[3:]...)
	fmt.Println(fatia3)
	fmt.Println("\n----------------------\n")

	// APPEND DE SLICES
	sabores = append(sabores, "portuguesa", "margeritha") // append aceita um primeiro slice (que será appendado) e mais 1-n argumentos do tipo que compoe o primeiro slice
	fmt.Println(sabores)
	novos_sabores := []string{"vegetariana", "cogumelos"}
	sabores = append(sabores, novos_sabores...) // append tbm pode aceitar com 2o argumento um outro slice, mas daí precisa colocar "..." depois.
	//Isso pode ser chamado de "unfurl" ou "enumerate", e é basicamente transformar um slice em uma série de argumentos
	// Equivalente a "arguments = [1, 2, 3]; my_method(*arguments)" (splat) em ruby (https://ruby-doc.org/core-2.0.0/doc/syntax/calling_methods_rdoc.html#label-Array+to+Arguments+Conversion)

	fmt.Println(sabores)
}
