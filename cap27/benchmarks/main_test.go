package benchmarks

import "testing"

type testTable struct {
	data []int
}

var tests = []testTable{
	{[]int{1, 2, 3, 4}},
	{[]int{10, 10}},
	{[]int{10, -15}},
	{[]int{1, 2, 3, 4, 5, 0, 85}},
}

func BenchmarkSomaComRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			SomaComRange(test.data...)
		}
	}
}
func BenchmarkSomaSemRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			SomaSemRange(test.data...)
		}
	}
}

// OUTPUT: $ go test -bench Soma

// goos: darwin
// goarch: amd64
// pkg: github.com/camilacampos/learn-go/cap27/benchmarks
// BenchmarkSomaComRange-8         90757116                13.3 ns/op
// BenchmarkSomaSemRange-8         68434290                15.9 ns/op
// PASS
// ok      github.com/camilacampos/learn-go/cap27/benchmarks       3.405s
