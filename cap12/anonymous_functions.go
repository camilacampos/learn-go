package main

import (
	"fmt"
)

type Alow struct {
	x int
}

func main() {
	x := 492

	// declaração de função anônima, chamando ela ao mesmo tempo
	// por ser anonima, a função não precisa de nome
	// vai ser usada bastante quando aprendermos go routines
	func(x int) {
		fmt.Println(x, "vezes 128 é:")
		fmt.Println(x * 128)
	}(x)

	fmt.Println("\n----------------------\n")

	// podemos usar funções anonimas em expressões também
	// nesse caso, a função não é necessariamente usada apenas uma vez

	mult := func(x int) (int, int) {
		return x * 235, 235
	}
	result, y := mult(x)
	fmt.Println(x, "vezes", y, "é:", result)
	result, y = mult(222)
	fmt.Println(222, "vezes", y, "é:", result)

	fmt.Println("\n----------------------\n")

	// valor da função é o endereço de memória dela
	// tipo é func(tipoParametro) tipoRetorno
	fmt.Printf("%v %T", mult, mult)
}
