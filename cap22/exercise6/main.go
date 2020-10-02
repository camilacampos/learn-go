package main

import (
	"fmt"
)

// - Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.

func main() {
	canal := make(chan int)

	go send(canal)
	receive(canal)
}

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
