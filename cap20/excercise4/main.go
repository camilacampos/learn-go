package main

import (
	"fmt"
	"runtime"
	"sync"
)

// - Utilize mutex para consertar a condição de corrida do exercício anterior.

var wg sync.WaitGroup
var mutex sync.Mutex

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
	mutex.Lock()
	fmt.Println("[START]\t", index, "- Contador:", *number)
	x := *number
	runtime.Gosched()
	x++
	*number = x
	fmt.Println("[END]\t", index, "- Contador:", *number)
	mutex.Unlock()
	wg.Done()
}

// OUTPUT RODANDO COM -race

// [START]  0 - Contador: 0
// [END]    0 - Contador: 1
// [START]  1 - Contador: 1
// [END]    1 - Contador: 2
// [START]  2 - Contador: 2
// [END]    2 - Contador: 3
// [START]  3 - Contador: 3
// [END]    3 - Contador: 4
// [START]  4 - Contador: 4
// [END]    4 - Contador: 5
// [START]  5 - Contador: 5
// [END]    5 - Contador: 6
// [START]  6 - Contador: 6
// [END]    6 - Contador: 7
// [START]  7 - Contador: 7
// [END]    7 - Contador: 8
// [START]  8 - Contador: 8
// [END]    8 - Contador: 9
// [START]  9 - Contador: 9
// [END]    9 - Contador: 10

// Contador final: 10
