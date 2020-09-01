package main

import (
	"fmt"
)

func printLine() {
	fmt.Println("\n\n----------------\n")
}

func main() {
	printLine()
	s := `Hello, 

	 			world`
	fmt.Printf("%v - %T", s, s)

	printLine()
	s = "Hello, world (éòâ中文)" //ASCII é sem os caracteres especiais - UTF8 é com
	fmt.Printf("%v - %T", s, s)

	printLine()
	sb := []byte(s) // sliced bytes - pode dar ruim com caracteres especiais (tem mais de um byte)
	fmt.Printf("%v - %T", sb, sb)

	printLine()
	for _, v := range sb {
		fmt.Printf("%b - %v - %T - %#U - %#x\n", v, v, v, v, v)
	}

	printLine()
	for _, v := range s {
		fmt.Printf("%b - %v - %T - %#U - %#x\n", v, v, v, v, v)
	}

	printLine()
	for i := 0; i < len(s); i++ { // aqui tbm separa em bytes, então não reconhece caracteres especiais
		fmt.Printf("%b - %v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i], s[i])
	}
}
