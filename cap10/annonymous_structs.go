package main

import (
	"fmt"
)

func main() {
	// structs anonimos, descartados logo após o primeiro (e único) uso.
	x := struct {
		nome  string
		idade int
	}{
		nome:  "mario",
		idade: 1294,
	}

	fmt.Println(x)
}
