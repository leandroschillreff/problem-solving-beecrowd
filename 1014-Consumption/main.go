package main

import (
	"fmt"
)

func main() {
	var x int
	var y float64

	fmt.Scan(&x, &y)

	averageConsumption := float64(x) / y
	fmt.Printf("%.3f km/l\n", averageConsumption)
}
