package main

import (
	"fmt"
)

// ESTRUTURA DAS FUNÇÕES
// func (receiver) identifier(parameters) (returns) { code }

func printLine() { // FUNCÃO BÁSICA
	fmt.Println("\n--------------------------------------------\n")
}

func printLineWith(text string) { // FUNCÃO BÁSICA COM ARGUMENTO
	fmt.Println("\n-----------", text, "-----------\n")
}

//        x int, y int (ALTERNATIVA)
func soma(x, y int) int { // FUNÇÃO COM RETORNO
	return x + y
}

// "...int" significa que recebe vários ints (função com parâmetro variádico), dentro da função é usado como slice
func somaTurbo(x ...int) (int, int) { // FUNÇÃO COM MÚLTIPLOS RETORNOS
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)
}

// parametros variádicos sempre precisam ser o último parâmetro
func calcula(operacao string, valores ...int) (string, int) {
	switch operacao {
	case "+":
		resposta := 0
		for _, valor := range valores {
			resposta += valor
		}
		return "Resultado da SOMA", resposta
	case "-":
		resposta := 0
		for _, valor := range valores {
			resposta -= valor
		}
		return "Resultado da SUBTRAÇÃO", resposta
	case "*":
		resposta := 1
		for _, valor := range valores {
			resposta *= valor
		}
		return "Resultado da MULTIPLICAÇÃO", resposta
	case "/":
		resposta := valores[0]
		for _, valor := range valores[1:] {
			resposta *= valor
		}
		return "Resultado da DIVISÃO", resposta
	default:
		return "Not implemented yet", 0
	}
}

func main() {
	printLineWith("usando a função soma")
	valor := soma(1, 2)

	fmt.Println(valor) // go sempre usa "pass by value", é como se a chamada fosse Println(3)

	printLineWith("usando a somaTurbo")

	novoValor, quantidadeDeItems := somaTurbo(1, 2, 3, 4)

	fmt.Println(novoValor, quantidadeDeItems)

	printLineWith("usando a CALCULADORA TURBO")

	text, valor := calcula("+", 1, 2, 3, 4, 5)
	fmt.Println(text, "-", valor)
	printLine()

	text, valor = calcula("-", 1, 2, 3, 4, 5)
	fmt.Println(text, "-", valor)
	printLine()

	text, valor = calcula("*", 1, 2, 3, 4, 5)
	fmt.Println(text, "-", valor)
	printLine()

	text, valor = calcula("/", 1, 2, 3, 4, 5)
	fmt.Println(text, "-", valor)
	printLine()

	text, valor = calcula("%", 1, 2, 3, 4, 5)
	fmt.Println(text, "-", valor)
	printLine()
}
