package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	soma := a + b

	fmt.Printf("SOMA = %d\n", soma)
}
