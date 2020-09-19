package main

import (
	"encoding/json"
	"fmt"
)

// - Partindo do código abaixo, utilize marshal para transformar []user em JSON.
//     - https://play.golang.org/p/U0jea43X55
// - Atenção! Tem pegadinha aqui.

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	serialized, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(serialized))
}
