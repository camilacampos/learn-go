package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("Printf único:\n%v\n%v\n%v\n\n", x, y, z)

	fmt.Println("Println único:", x, y, z, "\n")

	fmt.Println("Multiplos println")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
