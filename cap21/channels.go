package main

import (
	"fmt"
)

// Canais são formas de compartilhar informações entre go routines

// Não é possível colocar algo num canal e ler desse canal dentro da mesma go routine.
// Por exemplo, o código abaixo:
// ```
//        canal := make(chan int)
//        canal <- 42
//        fmt.Println(<-canal)
// ```
// Dá o seguinte erro:
//        fatal error: all goroutines are asleep - deadlock!
//        goroutine 1 [chan send]:
//        main.main()
//                /Users/camila.campos/go/src/github.com/camilacampos/learn-go/cap21/channels.go:12 +0x59
//        exit status 2

func main() {
	canal := make(chan int)

	go func() { // ao adicionar algo num canal em uma go routine (função anonima), ela fica disponível em outras (função main)
		canal <- 42 // adicionar informação em um canal
	}()

	fmt.Println(<-canal) // remover informação de um canal

	fmt.Print("\n---------------------------------------------\n\n")

	canal2 := make(chan int, 1) // 1 é o tamanho do buffer
	// com o buffer, não precisa de várias go routines pq só dá o deadlock (erro lá em cima) quando o buffer tá cheio

	canal2 <- 42 // adicionar informação em um canal
	// canal2 <- 43 // PAU: fatal error: all goroutines are asleep - deadlock!

	fmt.Println(<-canal2) // remover informação de um canal

	// Como o buffer tem tamanho 1, eu posso remover e ler valores tranquilamente, sem dar pau, desde que tenha buffer disponível:
	canal2 <- 24
	fmt.Println(<-canal2)

	// NO GERAL, BUFFER NÃO É UTILIZADO
}
