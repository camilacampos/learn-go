package main

import (
	"fmt"
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
	maximoGoFuncs := 5

	go manda(100, canal1)
	go recebeAi(maximoGoFuncs, canal1, canal2)

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

func recebeAi(maximoGoFuncs int, canal1, canal2 chan int) {
	var wg sync.WaitGroup
	// Com throttling! Ou seja, com um número máximo de go funcs
	// estamos executando apenas maximoGoFuncs (5) go routines de cada vez

	for i := 0; i < maximoGoFuncs; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalhoDoido(v) // expandiu um trabalho de um canal pra várias go routines (divergência) e depois colocou o resultado em um canal (convergência)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)
}

func trabalhoDoido(n int) int {
	time.Sleep(time.Millisecond * 1000)
	return n * 10
}
