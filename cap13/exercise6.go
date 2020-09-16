package main

import (
	"fmt"
)

// - Crie e utilize uma função anônima.

func main() {
	func() { // função anon, sem argumentos, que é chamada assim que é definida
		fmt.Println("Oi, eu estou dentro de uma função anônima :D")
	}()
}
