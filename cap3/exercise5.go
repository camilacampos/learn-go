package main

import (
	"fmt"
)

type myType int
var x myType

var y int

func main() {
	fmt.Printf("[1] Variável X - valor: %v | tipo: %T\n", x, x)
	x = 42
	fmt.Printf("[2] Variável X - valor: %v | tipo: %T\n", x, x)

	y = int(x)
	fmt.Printf("[3] Variável Y - valor: %v | tipo: %T\n", y, y)
}
