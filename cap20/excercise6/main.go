package main

import (
	"fmt"
	"runtime"
)

// - Crie um programa que demonstra seu OS e ARCH.
// - Rode-o com os seguintes comandos:
//     - go run
//     - go build
//     - go install

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}

// GO RUN

// $ go run cap20/excercise6/main.go
// => darwin amd64

// ---------------------------------------

// GO BUILD

// $ go build cap20/excercise6/main.go
// $ mv main cap20/excercise6
// $ ./cap20/excercise6/main
// => darwin amd64

// ---------------------------------------

// GO INSTALL

// $ go install cap20/excercise6/main.go
// $ cd ~/go/bin
// $ ./excercise6
// => darwin amd64
