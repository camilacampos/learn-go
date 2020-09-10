package main

import (
	"fmt"
	"runtime"
)

func main() {
	var unsignedInt uint8 //sem sinal (+ / -)
	var signedInt int8    //com sinal (+ / -)

	fmt.Println(unsignedInt, signedInt)

	// BYTE = UINT8
	// RUNE = INT32 (um caracter)

	a := "e"
	b := "é"
	c := "文"
	fmt.Printf("%v, %v, %v\n", a, b, c)

	e := []byte(a) // 1 byte
	f := []byte(b) // 2 bytes
	g := []byte(c) // 3 bytes

	fmt.Printf("%v, %v, %v\n", e, f, g)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
