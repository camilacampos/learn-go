package main

import (
	"fmt"
)

// nenhum tipo é definido aqui,
// portando posso usar tanto como int tanto como float
const myConst1 = 10
const (
	myConst2        = 10
	myConst3 string = "turu?"
)

var x int
var y float64

func main() {
	x = myConst1
	y = myConst1
	// x = y - PAU: cannot use y (type float64) as type int in assignment

	fmt.Printf("CONST 1:\t%v %t\n", myConst1, myConst1)
	fmt.Printf("X:\t\t%v %t\n", x, x)
	fmt.Printf("CONST 2:\t%v %t\n", myConst2, myConst2)
	fmt.Printf("Y:\t\t%v %t\n", y, y)
	fmt.Printf("CONST 3:\t%v %t\n", myConst3, myConst3)
}
