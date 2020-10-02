package main

import (
	"fmt"
)

// - Utilizando este código: https://play.golang.org/p/YHOMV9NYKK
// - ...demonstre o comma ok idiom.

func main() {
	c := make(chan int, 1)

	c <- 0

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
