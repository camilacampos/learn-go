package main

import (
	"fmt"
)

type myType int

var x myType

func main() {
	fmt.Printf("[1] Variável X - valor: %v | tipo: %T\n", x, x)
	x = 42
	fmt.Printf("[2] Variável X - valor: %v | tipo: %T\n", x, x)
}
