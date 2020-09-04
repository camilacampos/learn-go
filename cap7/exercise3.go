package main

import (
	"fmt"
)

// - Crie um loop utilizando a sintaxe: for condition {}
// - Utilize-o para demonstrar os anos desde que você nasceu.
func main() {
	year := 1992
	for year <= 2020 {
		fmt.Println(year)
		year++
	}
}
