package main

import (
	"fmt"
)

// - Começando com a seguinte slice:
//     - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// - Anexe a ela o valor 52;
// - Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
// - Demonstre a slice;
// - Anexe a ela a seguinte slice:
//     - y := []int{56, 57, 58, 59, 60}
// - Demonstre a slice x.
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x) // Slice inicial

	x = append(x, 52)
	fmt.Println(x) // Add 52

	x = append(x, 53, 54, 55)
	fmt.Println(x) // Add 3 números de uma vez

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x) // Mergeia com outro slice

	fmt.Println("\n----------------------\n")
}
