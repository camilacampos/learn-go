package main

import (
	"fmt"
)

// - Utilizando este código: https://play.golang.org/p/sfyu4Is3FG
// - ...use um for range loop para demonstrar os valores do canal.

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
