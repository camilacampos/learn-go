package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n---------- COM RECURSÃO ------------\n")
	fmt.Println("Fatorial de", 0, "é", fatorial(0))
	fmt.Println("Fatorial de", 1, "é", fatorial(1))
	fmt.Println("Fatorial de", 2, "é", fatorial(2))
	fmt.Println("Fatorial de", 3, "é", fatorial(3))
	fmt.Println("Fatorial de", 4, "é", fatorial(4))
	fmt.Println("Fatorial de", 5, "é", fatorial(5))

	fmt.Println("\n---------- SEM RECURSÃO ------------\n")
	fmt.Println("Fatorial de", 0, "é", fatorialComLoops(0))
	fmt.Println("Fatorial de", 1, "é", fatorialComLoops(1))
	fmt.Println("Fatorial de", 2, "é", fatorialComLoops(2))
	fmt.Println("Fatorial de", 3, "é", fatorialComLoops(3))
	fmt.Println("Fatorial de", 4, "é", fatorialComLoops(4))
	fmt.Println("Fatorial de", 5, "é", fatorialComLoops(5))
}

func fatorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * fatorial(number-1)
}

func fatorialComLoops(number int) int {
	result := 1
	for currentNumber := number; currentNumber > 0; currentNumber-- {
		result *= currentNumber
	}
	return result
}
