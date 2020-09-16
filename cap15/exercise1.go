package main

import (
	"fmt"
)

// - Crie um valor e atribua-o a uma variável.
// - Demonstre o endereço deste valor na memória.

func main() {
	x := 10
	fmt.Println("Valor:", x, "- Endereço:", &x)
}
