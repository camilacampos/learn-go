package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// - Canais par, ímpar, e converge.
// - Func send manda pares pra um, ímpares pro outro, depois fecha.
// - Func receive cria duas go funcs, cada uma com um for range, enviando dados dos canais par e ímpar pro canal converge. Não esquecer de WGs!
// - Por fim um range retira todas as informações do canal converge.

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	go manda(20, canal1)
	go recebeAi(canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}
	fmt.Println("\n----------------------\n")
}

func manda(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func recebeAi(canal1, canal2 chan int) {
	var wg sync.WaitGroup

	for v := range canal1 {
		wg.Add(1)
		go func(x int) {
			canal2 <- trabalhoDoido(x) // expandiu um trabalho de um canal pra várias go routines (divergência) e depois colocou o resultado em um canal (convergência)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(canal2)
}

func trabalhoDoido(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n
}
