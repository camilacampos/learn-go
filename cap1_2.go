package main

import (
	"fmt"
)

type hotdog int
var flubs hotdog
var flabs int

func main() {
	fmt.Println(`Digita ae um         número: \n`)
	var number1, number2 int
	fmt.Scan(&number1)
	// fmt.Println(`Digita    \t ae oto        "número":  `)
	fmt.Scan(&number2)
	fmt.Println("Soma:", number1+number2)
	_, err := fmt.Println("Subtração:", number1-number2)
	fmt.Println("Errors:", err)
	flubs = 123

	// flabs = flubs => PAU - tipos diferentes
	// flabs - flubs => PAU - tipos diferentes
	// flabs == flubs => PAU - tipos diferentes

	fmt.Println(flubs)
	fmt.Println(flubs == hotdog(flabs))
}
