package main

import (
	"fmt"
)

// Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos e demonstre estes valores.

const (
	_ = iota + 2020
	b
	c
	d
	e
)

func main() {
	fmt.Println(b, c, d, e)
}
