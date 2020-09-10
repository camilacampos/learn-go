package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Println("Package level vars inicializadas\n")
	fmt.Printf("Valor: %v - Tipo: %T\n", x, x)
	fmt.Printf("Valor: %v - Tipo: %T\n", y, y)
	fmt.Printf("Valor: %v - Tipo: %T\n", z, z)

	s := fmt.Sprintf("%v\n%v\n%v", x, y, z)
	fmt.Printf("\n\nVariável escrita com Sprint:\n%v\n", s)
}
