package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go sendNumbers(par, impar, quit)
	receiveNumbers(par, impar, quit)
}

func sendNumbers(par, impar chan int, quit chan bool) {
	total := 100
	for i := 0; i < total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	quit <- true
}

func receiveNumbers(par, impar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("O número", v, "é PAR.")
		case v := <-impar:
			fmt.Println("O número", v, "é ÍMPAR.")
		case <-quit:
			fmt.Println("CABÔ!")
			return
		}
	}
}

// Pode acontecer isso aqui, ó:
// OUTPUT:
//
// O número 0 é PAR.
// O número 1 é ÍMPAR.
// O número 2 é PAR.
// ........
// O número 97 é ÍMPAR.
// O número 98 é PAR.
// O número 99 é ÍMPAR.
// O número 0 é ÍMPAR.
// O número 0 é PAR.
// O número 0 é ÍMPAR.
// O número 0 é PAR.
// CABÔ!

// Solução em: cap21/multiple_channels-comma_ok.go
