package main

import (
	"fmt"
)

// - Atribua uma função a uma variável.
// - Chame essa função.

func main() {
	// função anonima, com argumentos e retorno, que é guardada em uma variável e depois chamada dentro do println
	sum := func(x, y int) int {
		return x + y
	}
	fmt.Println(sum(2, 4))
}
