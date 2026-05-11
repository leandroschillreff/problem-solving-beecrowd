package main

import (
	"fmt"
)

func main() {
	var distance int

	fmt.Scan(&distance)

	minutes := distance * 2

	fmt.Printf("%d minutos\n", minutes)
}
