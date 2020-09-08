package main

import (
	"fmt"
)

var x [5]int // arrays tem tamanhos definidos
var y [6]int // nao são muito utilizados (slices são), apenas quando tem algum uso muito específico

func main() {
	x[0] = 1
	x[1] = 10

	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println("\n----------------------\n")

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	// x = y // PAU: cannot use y (type [6]int) as type [5]int in assignment
	// x[7] = 10 // PAU: invalid array index 7 (out of bounds for 5-element array)
}
