package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var contador int64 = 0
	totalDeGoRoutines := 30

	var wg sync.WaitGroup
	wg.Add(totalDeGoRoutines)

	for i := 0; i < totalDeGoRoutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1) // no lugar do v++ & contador = v
			runtime.Gosched()             // yield: alternativa time.Sleep(time.Second)
			fmt.Println("Contador atual:", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Valor final:", contador)
}
