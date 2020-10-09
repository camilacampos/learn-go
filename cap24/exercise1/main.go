package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// - Utilizando este código: https://play.golang.org/p/3W69TH4nON
// - ...remova o underscore e verifique e lide com o erro de maneira apropriada.

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println("An error occurred when marshalling person:", err)
	}
	fmt.Println(string(bs))

}
