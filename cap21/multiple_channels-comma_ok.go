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
		case v, ok := <-par:
			if !ok {
				fmt.Println("Deu ruim no par, segue em frente aí:", v, ok)
				continue
			}
			fmt.Println("O número", v, "é PAR.")
		case v, ok := <-impar:
			if !ok {
				fmt.Println("Deu ruim no ímpar, segue em frente aí:", v, ok)
				continue
			}
			fmt.Println("O número", v, "é ÍMPAR.")
		case v, ok := <-quit:
			if !ok {
				fmt.Println(v, ok)
				fmt.Println("CAIU NO NOK DO QUIT!")
			} else {
				fmt.Println("CAIU NO OK DO QUIT:", v)
			}
			return
		}
	}
}

// Com o comma-ok, podemos resolver o problema de um canal estar recebendo além do que tem coisas dentro dele.
// OUTPUT:
//
// O número 0 é PAR.
// O número 1 é ÍMPAR.
// O número 2 é PAR.
// ........
// O número 96 é PAR.
// O número 97 é ÍMPAR.
// O número 98 é PAR.
// O número 99 é ÍMPAR.
// Deu ruim no par, segue em frente aí: 0 false
// Deu ruim no par, segue em frente aí: 0 false
// Deu ruim no par, segue em frente aí: 0 false
// Deu ruim no ímpar, segue em frente aí: 0 false
// CAIU NO OK DO QUIT: true
