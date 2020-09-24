package main

import (
	"fmt"
	"runtime"
	"sync"
)

// - Utilizando goroutines, crie um programa incrementador:
//     - Tenha uma variável com o valor da contagem
//     - Crie um monte de goroutines, onde cada uma deve:
//         - Ler o valor do contador
//         - Salvar este valor
//         - Fazer yield da thread com runtime.Gosched()
//         - Incrementar o valor salvo
//         - Copiar o novo valor para a variável inicial
//     - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
//     - Demonstre que há uma condição de corrida utilizando a flag -race

var wg sync.WaitGroup

func main() {
	interactions := 10

	wg.Add(interactions)
	number := 0

	for i := 0; i < interactions; i++ {
		go increment(i, &number)
	}

	wg.Wait()
	fmt.Println("\nContador final:", number)
}

func increment(index int, number *int) {
	fmt.Println("[START]\t", index, "- Contador:", *number)
	x := *number
	runtime.Gosched()
	x++
	*number = x
	fmt.Println("[END]\t", index, "- Contador:", *number)
	wg.Done()
}

// PARTE DO OUTPUT RODANDO COM -race

// ==================
// WARNING: DATA RACE
// Read at 0x00c0000be008 by goroutine 10:
//   main.increment()
//       /Users/camila.campos/go/src/github.com/camilacampos/learn-go/cap20/excercise3/main.go:37 +0x72

// Previous write at 0x00c0000be008 by goroutine 7:
//   main.increment()
//       /Users/camila.campos/go/src/github.com/camilacampos/learn-go/cap20/excercise3/main.go:41 +0x19b

// Goroutine 10 (running) created at:
//   main.main()
//       /Users/camila.campos/go/src/github.com/camilacampos/learn-go/cap20/excercise3/main.go:29 +0xa4

// Goroutine 7 (finished) created at:
//   main.main()
//       /Users/camila.campos/go/src/github.com/camilacampos/learn-go/cap20/excercise3/main.go:29 +0xa4
// ==================

// Found 3 data race(s)
// exit status 66
