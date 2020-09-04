package main

import (
	"fmt"
)

// - Crie um loop utilizando a sintaxe: for {}
// - Utilize-o para demonstrar os anos desde que você nasceu.
func main() {
	year := 1992
	for {
		fmt.Printf("%d, ", year)
		year++
		if year > 2020 {
			break
		}
	}

	fmt.Println("\n\n------------------------------------\n")

	year = 1992
	for {
		if year < 2020 {
			fmt.Printf("%d, ", year)
			year++
			continue
		}
		if year == 2020 {
			fmt.Printf("%d\n", year)
			break
		}
	}
}
