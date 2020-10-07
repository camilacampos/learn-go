package main

import (
	"fmt"
)

// OUTPUT:
// Calling g()
// Printing in g 0
// Printing in g 1
// Printing in g 2
// Printing in g 3
// Panicking!
// Defer in g 3
// Defer in g 2
// Defer in g 1
// Defer in g 0
//
// Error when reaching: 4
//
// Recovered in f
// Returned normally from f

func main() {
	f()
	fmt.Println("Returned normally from f")
}

func f() {
	defer func() {
		if r := recover(); r != nil { // se não teve nenhum panic, retorna nulo
			fmt.Println(r) // r é preechido com o que foi usado como parâmetro do panic(...)
			fmt.Println("Recovered in f")
		}
	}()

	fmt.Println("Calling g()")
	g(0)

	fmt.Println("Returned normally from g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("\nError when reaching: %v\n", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
	fmt.Println("This is never printed because all g() calls panicked", i)
}
