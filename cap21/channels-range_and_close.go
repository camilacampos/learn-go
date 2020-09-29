package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)

	go meuLoop(10, channel)
	printChannelValues(channel)

	fmt.Println("\n----------------------\n")
}

func meuLoop(quantidade int, s chan<- int) { // canal send
	for i := 0; i < quantidade; i++ {
		s <- i // adiciona 10 números no canal
	}
	close(s) // adicionado para que o canal não dê deadlock quando for rodar o for+range
}

func printChannelValues(r <-chan int) {
	for v := range r { // range para canais não tem índice
		fmt.Println("Recebi do canal o número:", v)
	}
	// No final desse loop, o programa ficaria ouvindo pelo próximo valor do canal.
	// Como não tem nada, então ele dá deadlock: fatal error: all goroutines are asleep - deadlock!
	// Para resolver isso, dá pra fechar o canal (adicionando um close no canal de send (final da função `meuLoop`))

}
