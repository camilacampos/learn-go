package testingbasics

func Soma(valores ...int) int {
	total := 0
	for _, v := range valores {
		total += v
	}
	return total
}
