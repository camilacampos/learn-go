package main

import (
	"fmt"
	"runtime"
	"sync"
)

// usando o sync.Mutex, paramos de ter problemas de race condition e alteramos a variável conforme deveria

func main() {
	contador := 0
	totalDeGoRoutines := 10

	var wg sync.WaitGroup
	wg.Add(totalDeGoRoutines)

	var mu sync.Mutex

	for i := 0; i < totalDeGoRoutines; i++ {
		go func() {
			mu.Lock()

			// fmt.Println(i, contador, "inicial")
			v := contador
			runtime.Gosched() // yield: alternativa time.Sleep(time.Second)
			v++
			contador = v
			// fmt.Println(i, contador, "final")

			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(contador)
}
