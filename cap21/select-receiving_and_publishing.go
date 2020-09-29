package main

import (
	"fmt"
)

var quit = make(chan int)

func main() {
	canal := make(chan int)

	go recebeQuit(canal)
	enviaPraCanal(canal) // para evitar o uso de wait groups e simplificar o exemplo, uma das funções não será go func

	fmt.Println("\n----------------------\n")
}

func recebeQuit(c chan int) {
	for i := 0; i < 20; i++ {
		fmt.Println("Recebido:", <-c)
	}
	quit <- 0
}

func enviaPraCanal(c chan int) {
	qualquercoisa := 9
	for {
		select {
		case c <- qualquercoisa: // manda coisa pra um canal
			qualquercoisa++
		case <-quit: // recebe coisa de um canal
			return
		}
	}
}
