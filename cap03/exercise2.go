package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Println("Valores zero de int, string e bool\n")
	fmt.Printf("Valor zero: %v - Tipo: %T\n", x, x)
	fmt.Printf("Valor zero: %v - Tipo: %T\n", y, y)
	fmt.Printf("Valor zero: %v - Tipo: %T\n", z, z)
}
