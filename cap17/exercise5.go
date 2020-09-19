package main

import (
	"fmt"
	"sort"
)

// - Partindo do código abaixo, ordene os []user por idade e sobrenome.
//     - https://play.golang.org/p/BVRZTdlUZ_
// - Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (u user) toString() string {
	text := fmt.Sprintf("%v, %v - Age: %v\nSayings:\n", u.Last, u.First, u.Age)

	for _, saying := range u.Sayings {
		text += fmt.Sprintf("\t- %v\n", saying)
	}
	text += fmt.Sprintf("\n")
	return text
}

func (u user) withSortedSayings() user {
	sort.Strings(u.Sayings)
	return u
}

type orderUserByLastAge []user

func (users orderUserByLastAge) Len() int           { return len(users) }
func (users orderUserByLastAge) Less(i, j int) bool { return users[i].Age < users[j].Age }
func (users orderUserByLastAge) Swap(i, j int)      { users[i], users[j] = users[j], users[i] }

func main() {
	u1 := createUser1().withSortedSayings()
	u2 := createUser2().withSortedSayings()
	u3 := createUser3().withSortedSayings()
	users := []user{u1, u2, u3}

	printUsers("Sem ordem:", users)
	fmt.Println("\n----------------------\n")

	sort.Sort(orderUserByLastAge(users))
	printUsers("Ordenado por idade:", users)
	fmt.Println("\n----------------------\n")

	sort.Slice(users, func(i, j int) bool {
		return users[i].Last < users[j].Last
	})
	printUsers("Ordenado por sobrenome:", users)
	fmt.Println("\n----------------------\n")
}

func printUsers(desc string, users []user) {
	fmt.Println("**", desc, "**")
	for _, user := range users {
		fmt.Println(user.toString())
	}
}

func createUser1() user {
	return user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}
}

func createUser2() user {
	return user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}
}

func createUser3() user {
	return user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}
}
