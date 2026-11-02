package main

import (
	"fmt"
)

func main() {
	var positiveCount int

	for i := 0; i < 6; i++ {
		var number float64
		fmt.Scan(&number)

		if number > 0 {
			positiveCount++
		}
	}

	fmt.Printf("%d valores positivos\n", positiveCount)
}
