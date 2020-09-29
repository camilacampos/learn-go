package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)

	go func() {
		canal <- 0
		close(canal)
	}()

	v, ok := <-canal

	fmt.Println(v, ok)

	v, ok = <-canal // comma ok (, ok) indica se o valor zerado é zero mesmo ou nem

	fmt.Println(v, ok)
}
