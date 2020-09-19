package main

import (
	"fmt"
	"sort"
)

// Para fazer um sort customizado, é necessário usar a func sort.Sort(data Interface)
// O "data" recebido no método pode ser qualquer coisa que implemente a interface sort.Interface
// Para isso, né necessário implementar:
//		 Len is the number of elements in the collection.
//				Len() int
//		 Less reports whether the element with
//		 index i should sort before the element with index j.
//		 		Less(i, j int) bool
//		 Swap swaps the elements with indexes i and j.
//		 		Swap(i, j int)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

// tipo ordenarPorPotencia implementa a interface sort.Interface
type ordenarPorPotencia []carro

func (a ordenarPorPotencia) Len() int           { return len(a) }
func (a ordenarPorPotencia) Less(i, j int) bool { return a[i].potencia < a[j].potencia }
func (a ordenarPorPotencia) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	carros := []carro{
		{"Vectra", 10, 30},
		{"Astra", 40, 40},
		{"Gol", 30, 10},
		{"Fusca", 20, 20},
	}

	fmt.Println(carros)
	// ordenarPorPotencia(carros): isso é uma conversão de []carros em []ordenarPorPotencia
	// mesma coisa quando fazemos string(12) para transformar 12 (int) em "12" (string)
	sort.Sort(ordenarPorPotencia(carros))
	fmt.Println(carros)

	fmt.Println("\n----------------------\n")

	// Ao invés de definir um tipo que implemente a interface interface,
	// é possível passar uma função (less) para o sort.Slice.
	sort.Slice(carros, func(i, j int) bool {
		return carros[i].consumo < carros[j].consumo
	})
	fmt.Println(carros)

}
