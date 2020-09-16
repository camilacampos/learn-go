package main

import (
	"fmt"
)

// - Callback: passe uma função como argumento a outra função.

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	calculate(sum, numbers...)
}

func calculate(operation func(...int) int, numbers ...int) {
	fmt.Println("O resultado da operação é:")
	fmt.Println(operation(numbers...))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
