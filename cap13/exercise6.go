package main

import (
	"fmt"
)

// - Crie e utilize uma função anônima.

func main() {
	func() { // função anon1, sem argumentos, que é chamada assim que é definida
		sum := func(x, y int) int { // função anon2, com argumentos e retorno, que é guardada em uma variável e depois chamada dentro do println
			return x + y
		}
		fmt.Println(sum(2, 4))
	}()
}
