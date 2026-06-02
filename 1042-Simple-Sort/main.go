package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	numbers := []int{a, b, c}

	sort.Ints(numbers)

	for _, num := range numbers {
		fmt.Println(num)
	}

	fmt.Println()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
