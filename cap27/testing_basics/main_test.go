package testingbasics

import (
	"fmt"
	"testing"
)

func TestSoma_OK(t *testing.T) {
	x := Soma(1, 2, 3)
	if x != 6 {
		t.Error("Expected:", 6, "Received:", x)
	}
}

func TestSoma_Quebrado(t *testing.T) {
	x := Soma(1, 2, 2)
	if x != 6 {
		t.Error("Expected:", 6, "Received:", x)
	}
}

// Tabela de testes
type testTable struct {
	data   []int
	answer int
}

func TestSoma_Tabela(t *testing.T) {
	tests := []testTable{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{10, 10}, 20},
		{[]int{10, -15}, -5},
		{[]int{1, 2, 3, 4, 5, 0, 85}, 101}, // QUEBRADO
		{[]int{1, 2, 3, 4, 5, 0, 85}, 100},
		{[]int{10, -15}, -7}, // QUEBRADO
	}

	for _, test := range tests {
		result := Soma(test.data...)
		if result != test.answer {
			t.Error("Expected:", test.answer, "Received:", result) // O erro não para a execução, então roda tudo e mostra todos os erros
		}
	}
}

// Exemplos aparecem na documentação (usando godoc, por exemplo), além de executar o próprio teste
// Um exemplo do que rola ao rodar `godoc -http=":3456"` está disponível em `cap27/testing_basics/test_exemples_on_go_doc.mov`
func ExampleSoma_funcionando() { // o nome da função no nome do exemplo precisa ser o mesmo que no código
	fmt.Println(Soma(1, 2, 3))
	// Output: 6
}

func ExampleSoma_varios() { // é possível executar vários casos de uma vez também
	fmt.Println(Soma(1, 2, 3))
	fmt.Println(Soma(1, 2, 4))
	fmt.Println(Soma(2, 3, 4))
	// Output:
	// 6
	// 7
	// 9
}

func ExampleSoma_quebrando() {
	fmt.Println(Soma(1, 2, 3))
	// Output: 7
}

// OUTPUT de ExampleSoma_Quebrando:
// === RUN   ExampleSoma_Quebrando
// --- FAIL: ExampleSoma_Quebrando (0.00s)
// got:
// 6
// want:
// 7
// FAIL
