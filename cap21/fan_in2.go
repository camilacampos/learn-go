package main

import (
	"fmt"
	"math/rand"
	"time"
)

// - Func trabalho cria um canal, cria uma go func que manda dados pra esse canal, e retorna o canal. Interessante: time.Duration(rand.Intn(1e3))
// - Func converge toma dois canais, cria um canal novo, e cria duas go funcs com for infinito que passa tudo para o canal novo. Retorna o canal novo.
// - Por fim chamamos canal := converge(trabalho(nome1), trabalho(nome2)) e usamos um for para receber dados do canal var.

func main() {
	trabalho1 := trabalho("batata")
	trabalho2 := trabalho("árvore")

	canal := converge(trabalho1, trabalho2)

	for x := 0; x < 16; x++ {
		fmt.Println(<-canal)
	}
}

func trabalho(s string) chan string {
	canal := make(chan string)

	go func() {
		for i := 1; ; i++ {
			canal <- fmt.Sprintf("Canal %v diz: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3))) // demora entre nada e 1 segundo
		}
	}()

	return canal
}

func converge(x, y chan string) chan string {
	novo := make(chan string)
	go func() {
		for {
			novo <- <-x
		}
	}()
	go func() {
		for {
			novo <- <-y
		}
	}()
	return novo
}
