package main

import (
	"fmt"
)

func main() {
	for letter := 33; letter <= 122; letter++ {
		fmt.Printf("%v ", string(letter))
	}
	fmt.Println()
}
