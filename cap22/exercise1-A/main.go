package main

import (
	"fmt"
)

// - Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
//     - Usando uma função anônima auto-executável

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
