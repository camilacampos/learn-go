package main

import (
	"fmt"
)

// iota = consts com valores sucessivos inteiros não tipados
// usada quando não importa o valor da const, apenas que ela seja diferente das demais
// const (
// 	a = iota + 10 // faz a conta apenas nesse
// 	b = iota
// 	_ = iota
// 	d = iota
// 	_ = iota
// 	f = iota
// )

// é possível omitir o " = iota" das consts subsequentes
const (
	a = iota * 10 // aplica a conta para todos os itens abaixo
	b
	_
	d
	_
	f
)

func main() {
	fmt.Printf("A:\t%v %t\n", a, a)
	fmt.Printf("B:\t%v %t\n", b, b)
	fmt.Printf("D:\t%v %t\n", d, d)
	fmt.Printf("F:\t%v %t\n", f, f)
}
