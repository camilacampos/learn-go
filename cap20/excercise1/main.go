package main

import (
	"fmt"
	"sync"
)

// - Alem da goroutine principal, crie duas outras goroutines.
// - Cada goroutine adicional devem fazer um print separado.
// - Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go print1()
	go print2()

	wg.Wait()
	fmt.Println("\n FIM DO PROGRAMA!")
}

func print1() {
	for i := 0; i < 20; i++ {
		fmt.Println("[1]", i)
	}
	wg.Done()
}

func print2() {
	for i := 20; i > 0; i-- {
		fmt.Println("[2]", i)
	}
	wg.Done()
}
