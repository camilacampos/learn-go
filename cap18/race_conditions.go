package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Podemos usar a flag -race no `go run` para identificar race conditions

func main() {
	contador := 0
	totalDeGoRoutines := 10

	var wg sync.WaitGroup
	wg.Add(totalDeGoRoutines)

	for i := 0; i < totalDeGoRoutines; i++ {
		go func() {
			fmt.Println(i, contador, "inicial")
			v := contador
			runtime.Gosched() // yield: alternativa time.Sleep(time.Second)
			v++
			contador = v
			fmt.Println(i, contador, "final")
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(contador)
}

// SAMPLE OUTPUT:
// 3 0 inicial
// 10 0 inicial
// 10 0 inicial
// 10 0 inicial
// 10 0 inicial
// 10 0 inicial
// 10 0 inicial
// 10 1 final
// 10 1 final
// 10 1 final
// 10 1 final
// 10 2 final
// 10 1 inicial
// 10 2 inicial
// 10 2 final
// 10 1 final
// 10 2 inicial
// 10 2 final
// 10 2 final
// 10 3 final
// 3
