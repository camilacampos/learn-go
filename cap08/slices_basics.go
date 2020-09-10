package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5} // fixo, impossível de mudar o tamanho
	fmt.Println(array)
	fmt.Println("\n----------------------\n")

	slice := []int{1, 2, 3, 4, 5} // tamanho indefinido
	fmt.Println(slice)
	fmt.Println("\n----------------------\n")

	// array2 := append(array, 6) // PAU: first argument to append must be slice; have [5]int
	slice2 := append(slice, 6)
	fmt.Println("Slice novo:", slice2)
	fmt.Println("\n----------------------\n")

	fmt.Println(slice[3])
	slice[3] = 1289128
	fmt.Println(slice[3])
	fmt.Println("\n----------------------\n")

	// fmt.Println(slice[29]) // PAU: runtime error: index out of range [29] with length 5
}
