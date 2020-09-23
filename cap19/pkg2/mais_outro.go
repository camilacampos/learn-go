package pkg2

import "fmt"

func Sum(x, y int) int {
	fmt.Println("Dentro de pkg2/mais_outro.go")
	return x + y
}
