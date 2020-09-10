package main

import (
	"fmt"
)

// - Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto (65 a 90), três vezes cada.
// - Por exemplo:
//     65
//         U+0041 'A'
//         U+0041 'A'
//         U+0041 'A'
//     66
//         U+0042 'B'
//         U+0042 'B'
//         U+0042 'B'
//     ...e por aí vai.
func main() {
	for upperCaseLetter := 65; upperCaseLetter <= 90; upperCaseLetter++ {
		fmt.Printf("%v - %v - ", upperCaseLetter, string(upperCaseLetter))
		for i := 0; i < 3; i++ {
			fmt.Printf("%#U ", upperCaseLetter)
		}
		fmt.Println()
	}
}
