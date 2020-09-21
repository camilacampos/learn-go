package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// pode ser que o programa execute antes de terminar as threads das funções
// para resolver isso, existe o conceito de wait groups:
// Ele define:
//   - quantas threads vão ter no programa principal
//   - quando alguma thread terminou de executar
//   - quando o programa principal deve esperar as demais threads finalizarem
func main() {
	fmt.Println("# Processors", runtime.NumCPU())
	fmt.Println("# GO Routines (initial)", runtime.NumGoroutine())

	// WaitGroup: add(total de go routines)
	wg.Add(2)

	fmt.Println("# GO Routines (after waitGroup.add)", runtime.NumGoroutine())

	go func1()
	go func2()
	func3()

	fmt.Println("# GO Routines (after function calls)", runtime.NumGoroutine())

	// WaitGroup: espera terminar tudo
	wg.Wait()

	fmt.Println("# GO Routines (after waitGroup.wait)", runtime.NumGoroutine())
}

// Exemplo de output:
// func3: 0
// func3: 1
// func3: 2
// func3: 3
// func3: 4
// func2: 0
// func2: 1
// func2: 2
// func2: 3
// func2: 4
// func2: 5
// func2: 6
// func1: 0
// func1: 1
// func1: 2
// func1: 3
// func1: 4
// func1: 5
// func1: 6
// func1: 7
// func1: 8
// func1: 9
// func3: 5
// func3: 6
// func3: 7
// func3: 8
// func3: 9
// func2: 7
// func2: 8
// func2: 9

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
	}
	// WaitGroup: Terminei a execução
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2:", i)
	}
	// WaitGroup: Terminei a execução
	wg.Done()
}

func func3() {
	for i := 0; i < 10; i++ {
		fmt.Println("func3:", i)
	}
}
