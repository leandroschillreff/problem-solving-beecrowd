package main

import (
	"fmt"
)

func main() {
	var value int

	fmt.Scan(&value)

	fmt.Println(value)

	notes := []int{100, 50, 20, 10, 5, 2, 1}

	for _, nota := range notes {
		quantity := value / nota

		fmt.Printf("%d nota(s) de R$ %d,00\n", quantity, nota)

		value %= nota
	}
}
