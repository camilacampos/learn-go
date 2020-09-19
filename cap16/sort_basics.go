package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"camila", "natalia", "daniela", "ana", "jordana"}

	fmt.Println("Sorted:", sort.StringsAreSorted(names), "-", names)
	sort.Strings(names)
	fmt.Println("Sorted:", sort.StringsAreSorted(names), "-", names)

	numbers := []int{127, 23498, 1298, 29, 1, 394, 934, 4389}

	fmt.Println("\n\n--------------------------")
	fmt.Println()

	fmt.Println("Sorted:", sort.IntsAreSorted(numbers), "-", numbers)
	sort.Ints(numbers)
	fmt.Println("Sorted:", sort.IntsAreSorted(numbers), "-", numbers)
	numbers = append(numbers, 2)
	fmt.Println("Sorted:", sort.IntsAreSorted(numbers), "-", numbers)
	sort.Ints(numbers)
	numbers = append(numbers, 99999)
	fmt.Println("Sorted:", sort.IntsAreSorted(numbers), "-", numbers)
}
