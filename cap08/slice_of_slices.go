package main

import (
	"fmt"
)

func main() {
	ss := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	ss2 := [][][]int{
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		},
		[][]int{
			[]int{13, 14, 15},
			[]int{10, 11, 12},
		},
	}
	fmt.Println(ss[0])
	fmt.Println(ss[1][1])
	fmt.Println(ss2[1])
	fmt.Println(ss2[0][2])
	fmt.Println(ss2[0][1][2])
	fmt.Println("\n----------------------\n")
}
