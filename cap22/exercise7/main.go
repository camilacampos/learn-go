package main

import (
	"fmt"
	"sync"
)

// - Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
// - Tire estes números do canal e demonstre-os.

func main() {
	canal := make(chan int)

	go send(canal)

	receive(canal)
}

func send(c chan<- int) {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go genGoFunc(c, &wg)()
	}

	wg.Wait()
	close(c)
}

func genGoFunc(c chan<- int, wg *sync.WaitGroup) func() {
	return func() {
		for j := 0; j < 10; j++ {
			c <- j
		}
		wg.Done()
	}
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println("valor:", v)
	}
}
