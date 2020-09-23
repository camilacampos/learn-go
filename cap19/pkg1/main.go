package main

import (
	"fmt"

	"github.com/camilacampos/learn-go/cap19/pkg2"
)

// Para usar funções do mesmo pacote, não precisa de importe, basta usar diretamente.
// Para outros pacotes, é necessário fazer o import do caminho completo a partir do GOPATH
// Para que as coisas de outros pacotes sejam visíveis, é necessário que elas estejam com letra maiúscula

// Quando tem mais de um pacote/arquivo, devemos rodar `go run *.go` para executar o rolê todo.

func main() {
	fmt.Println("Dentro de pkg1/main.go")
	printLine()
	result := pkg2.Sum(3, 6)
	fmt.Println("Resultado da função de outro pacote:", result)
}
