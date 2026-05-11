package main

import (
	"fmt"
)

func main() {
	var spentTime, averageSpeed int

	fmt.Scan(&spentTime, &averageSpeed)

	liters := (float64(spentTime) * float64(averageSpeed)) / 12.0

	fmt.Printf("%.3f\n", liters)
}
