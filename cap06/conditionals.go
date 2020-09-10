package main

import (
	"fmt"
)

func main() {
	// ifs simples, podem ser com ou sem parênteses
	x := 1000
	if x < 100 {
		fmt.Println("X é menor que 100")
	}
	if !(x < 100) {
		fmt.Println("X NÃO é menor que 100")
	}

	// ifs podem ter inicialização
	if i := 10; i < 100 {
		fmt.Println("i é menor que 100")
	}

	// podem existir vários "else if"
	x = 200
	if x < 100 {
		fmt.Println("X é menor que 100")
	} else if x > 100 {
		fmt.Println("X NÃO é menor que 100")
	} else {
		fmt.Println("X é igual a 100")
	}

	// operadores:
	// && - e
	// || - ou
	// !  - não
	if y := 6; y%2 == 0 && y%3 == 0 {
		fmt.Println(y, "é multiplo de 2 e 3")
	}
	if y := 3; y%2 == 0 || y%3 == 0 {
		fmt.Println(y, "é multiplo de 2 ou de 3")
	}
	if y := 9; !(y%2 == 0) && y%3 == 0 {
		fmt.Println(y, "é multiplo de 3, mas é não é multiplo de 2")
	}

	// switch case padrão
	x = 10
	switch true { // true é o default
	case x >= 100, x == 10:
		fmt.Println("X é maior ou igual a 100 ou x é igual a 10")
	case x < 5:
		fmt.Println("X é menor que 5")
	case x == 5:
		fmt.Println("X é igual a 5")
	case x > 5:
		fmt.Println("X é maior que 5")
	}

	quemTaNoEscritorioHoje := "joaquina"
	switch quemTaNoEscritorioHoje {
	case "marquinhos":
		fmt.Println("oi")
	case "zezinho":
		fmt.Println("te liguei")
		fallthrough // sempre que cai nesse caso, cai automaticamente no debaixo também
	case "joana":
		fmt.Println("tá com outro")
	case "maria":
		fmt.Println("contatinho")
	case "jorge", "pedrinho":
		fmt.Println("time topper")
	case "mumu", "guigui": // pode dar match em mais de um caso (funciona como ||)
		fmt.Println("time menos topper")
	default:
		fmt.Println("foi ninguém")
	}

	// switch pode ser feito em cima de tipos
	var y interface{}
	y = "1287.9"

	switch y.(type) {
	case int:
		fmt.Println("é um int")
	case bool:
		fmt.Println("é um bool")
	case string:
		fmt.Println("é uma string")
	case float64:
		fmt.Println("é um float")
	}

	// switch pode ter inicialização também
	switch z := 4; {
	case z == 4:
		fmt.Println("é um 4")
	default:
		fmt.Println("não é um 4")
	}
	switch z := 5; z {
	case 4:
		fmt.Println("é um 4")
	default:
		fmt.Println("não é um 4")
	}

}
