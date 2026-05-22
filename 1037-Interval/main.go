package main

import (
	"fmt"
)

func main() {
	var number float64
	fmt.Scan(&number)

	if number < 0 || number > 100 {
		fmt.Println("Fora de intervalo")
	} else if number <= 25 {
		fmt.Println("Intervalo [0,25]")
	} else if number <= 50 {
		fmt.Println("Intervalo (25,50]")
	} else if number <= 75 {
		fmt.Println("Intervalo (50,75]")
	} else {
		fmt.Println("Intervalo (75,100]")
	}
}
