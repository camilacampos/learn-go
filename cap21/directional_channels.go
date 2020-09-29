package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)
	go send(canal)
	receive(canal)

	fmt.Println("\n----------------------\n")
	fmt.Printf("canal\t%T\n", (<-chan int)(canal)) // é possível converter um canal bidirecional para um one-way (específico)

	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send
	fmt.Printf("canal\t%T\n", canal)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	fmt.Printf("canal\t%T\n", (<-chan int)(canal)) // é possível converter um canal bidirecional para um one-way (específico)
	// Não é possível converter um canal específico em outro ou em um geral:
	// fmt.Printf("cr\t%T\n", (chan<- int)(cr)) // PAU: cannot convert cr (type <-chan int) to type chan<- int
	// fmt.Printf("cs\t%T\n", (<-chan int)(cs)) // PAU: cannot convert cs (type chan<- int) to type <-chan int
	// fmt.Printf("cs\t%T\n", (chan int)(cs))   // PAU: cannot convert cs (type chan<- int) to type chan int
}

// se tentar inverter, vai dar merda:   invalid operation: s <- 42 (send to receive-only type <-chan int)
// func send(s <-chan int) {
func send(s chan<- int) { // isso na prática é uma conversão de um canal geral (canal lá em cima) pra um específico
	s <- 42
}

// se tentar inverter, vai dar merda:   invalid operation: <-r (receive from send-only type chan<- int)
// func receive(r chan<- int) {
func receive(r <-chan int) { // isso na prática é uma conversão de um canal geral (canal lá em cima) pra um específico
	fmt.Println("O valor recebido do canal foi:", <-r)
}
