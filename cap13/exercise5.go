package main

import (
	"fmt"
	"math"
)

// - Crie um tipo "quadrado"
// - Crie um tipo "círculo"
// - Crie um método "área" para cada tipo que calcule e retorne a área da figura
//     - Área do círculo: 2 * π * raio
//     - Área do quadrado: lado * lado
// - Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
// - Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
// - Crie um valor de tipo "quadrado"
// - Crie um valor de tipo "círculo"
// - Use a função "info" para demonstrar a área do "quadrado"
// - Use a função "info" para demonstrar a área do "círculo"

type Figura interface {
	area() float64
}

type Quadrado struct {
	lado float64
}

func (q Quadrado) area() float64 {
	return q.lado * q.lado
}

type Círculo struct {
	raio float64
}

func (c Círculo) area() float64 {
	return 2 * math.Pi * c.raio
}

func info(f Figura) {
	fmt.Println("A área é:", f.area())
}

func main() {
	quadrado := Quadrado{2}
	círculo := Círculo{2}

	fmt.Println("\n----------- USANDO CADA TIPO -----------\n")
	fmt.Println("A área do quadrado é:", quadrado.area())
	fmt.Println("A área do círculo é:", círculo.area())

	fmt.Println("\n----------- USANDO O INFO -----------\n")
	info(quadrado)
	info(círculo)
}
