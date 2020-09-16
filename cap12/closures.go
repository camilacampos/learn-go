package main

import (
	"fmt"
)

// Closure é cercar ou capturar um scope para que possamos utilizá-lo em outro contexto.

func main() {
	increment := i()         // define o X como 0
	fmt.Println(increment()) // incrementa o X para 1
	fmt.Println(increment()) // incrementa o X para 2
	fmt.Println(increment()) // incrementa o X para 3

	fmt.Println("\n----------------------\n")

	increment2 := i()         // por estar definindo um novo escopo (chamando a função de novo), o X é zerado
	fmt.Println(increment2()) // incrementa o X para 1
	fmt.Println(increment2()) // incrementa o X para 2
	fmt.Println(increment2()) // incrementa o X para 3
}

// usar o X de fora do escopo da função (anon) é a definição de closure
func i() func() int {
	x := 0 // o X está definido no escopo da função

	return func() int { // por estar dentro da função, a função anonima de retorno tem acesso ao X
		x++
		return x
	}
}
