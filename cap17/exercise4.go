package main

import (
	"fmt"
	"sort"
)

// - Partindo do código abaixo, ordene a []int e a []string.
//     - https://play.golang.org/p/H_q75mpmHW

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	fmt.Println("----------------------")

	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("\n----------------------\n")

	fmt.Println(xs)
	fmt.Println("----------------------")

	sort.Strings(xs)
	fmt.Println(xs)
}
