package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt" // x no começo do pacote indica que ele está em teste (ainda não foi pra standard library da linguagem)
	// para instalar o pkg acima, foi necessário roda "go get -u golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "07101992"
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(senha)
	fmt.Println(string(hashedPassword))

	valid := isPasswordValid([]byte(senha), hashedPassword)
	fmt.Println("Senha correta:", valid)

	senhaErrada := "07101991"
	valid = isPasswordValid([]byte(senhaErrada), hashedPassword)
	fmt.Println("Senha correta:", valid)

	fmt.Println("\n----------------------\n")
}

func isPasswordValid(password, hashedPassword []byte) bool {
	// func CompareHashAndPassword(hashedPassword, password []byte) error // Returns nil on success, or an error on failure.
	comparison := bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println("Senha:", string(password), " - Resultado do CompareHashAndPassword:", comparison)
	return comparison == nil
}
