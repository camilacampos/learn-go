package main

import (
	"fmt"
)

func main() {
	// OVERFLOW
	var i uint16
	i = 65535
	// i = 65536 - PAU
	fmt.Println(i)
	i++ // zera (faz o circulo)
	fmt.Println(i)
	i++
	fmt.Println(i)
}
