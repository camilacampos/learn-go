package main

import (
	"fmt"
)

// - Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
//     - Usando buffer

func main() {
	c := make(chan int, 1) // 1 é o tamanho do buffer

	c <- 42

	fmt.Println(<-c)
}
