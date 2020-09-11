package main

import (
	"fmt"
)

// - Crie um novo tipo: veículo
//     - O tipo subjacente deve ser struct
//     - Deve conter os campos: portas, cor
// - Crie dois novos tipos: caminhonete e sedan
//     - Os tipos subjacentes devem ser struct
//     - Ambos devem conter "veículo" como struct embutido
//     - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
//     - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
// - Usando os structs veículo, caminhonete e sedan:
//     - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
//     - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
// - Demonstre estes valores.
// - Demonstre um único campo de cada um dos dois.
type Veiculo struct {
	portas int
	cor    string
}

type Caminhonete struct {
	Veiculo
	traçãoNasQuatro bool
}

type Sedan struct {
	veiculo    Veiculo
	modeloLuxo bool
}

func main() {
	veiculo := Veiculo{
		portas: 4, cor: "azul",
	}
	caminhonete := Caminhonete{
		Veiculo: Veiculo{
			portas: 2, cor: "preta",
		},
		traçãoNasQuatro: false,
	}
	sedan := Sedan{
		Veiculo{4, "vermelha"},
		true,
	}

	fmt.Println("Veiculo:", veiculo)
	fmt.Println("Caminhonete:", caminhonete)
	fmt.Println("Sedan:", sedan)

	fmt.Println("\n----------------------\n")

	fmt.Println("Cor do Veículo:", veiculo.cor)
	fmt.Println("Portas da Caminhonete (v1):", caminhonete.Veiculo.portas)
	fmt.Println("Portas da Caminhonete (v2):", caminhonete.portas)
	fmt.Println("Sedan é modelo de luxo:", sedan.modeloLuxo)

	fmt.Println("\n----------------------\n")
}
