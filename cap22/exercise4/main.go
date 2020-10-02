package main

import (
	"fmt"
)

// - Utilizando este código: https://play.golang.org/p/MvL6uamrJP
// - ...use um select statement para demonstrar os valores do canal.

func main() {
	quit := make(chan int)
	c := gen(quit)

	receive(c, quit)

	fmt.Println("About to exit")
}

func gen(quit chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		quit <- 1
		close(quit)
	}()

	return c
}

func receive(c, quit <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("Valor:", v)
		case <-quit:
			fmt.Println("\nExiting select")
			return
		}
	}
}
