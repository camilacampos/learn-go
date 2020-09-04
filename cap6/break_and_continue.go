package main

import (
	"fmt"
)

func main() {
	x := 0
	for x <= 20 {
		x++
		if x == 7 {
			fmt.Println("eae, parei tudo irmão hihihi")
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}
