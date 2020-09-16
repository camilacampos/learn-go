package main

import (
	"fmt"
)

func main() {
	mult := buildMultiplyFuncion(20)

	result, y := mult(123)
	fmt.Println(123, "vezes", y, "é:", result)

	fmt.Println("\n----------------------\n")

	result, y = mult(222)
	fmt.Println(222, "vezes", y, "é:", result)

	fmt.Println("\n----------------------\n")

	fmt.Println(buildMultiplyFuncion(30)(999))
}

func buildMultiplyFuncion(y int) func(int) (int, int) {
	return func(x int) (int, int) {
		return x * y, y
	}
}
