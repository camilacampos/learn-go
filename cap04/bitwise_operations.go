package main

import (
	"fmt"
)

func main() {
	x := 24
	y := x << 2
	z := x >> 2

	fmt.Printf("x:\t%b - %v %T\n", x, x, x)
	fmt.Printf("y:\t%b - %v %T\n", y, y, y)
	fmt.Printf("z:\t%b - %v %T\n", z, z, z)
}
