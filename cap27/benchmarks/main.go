package benchmarks

func SomaComRange(valores ...int) int {
	total := 0
	for _, v := range valores {
		total += v
	}
	return total
}

func SomaSemRange(valores ...int) int {
	total := 0
	for i := 0; i < len(valores); i++ {
		total += valores[i]
	}
	return total
}
