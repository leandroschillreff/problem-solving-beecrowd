package main

import (
	"fmt"
)

func main() {
	var code, quantity int

	fmt.Scan(&code, &quantity)

	var price float64

	switch code {
	case 1:
		price = 4.00
	case 2:
		price = 4.50
	case 3:
		price = 5.00
	case 4:
		price = 2.00
	case 5:
		price = 1.50
	}

	total := float64(quantity) * price

	fmt.Printf("Total: R$ %.2f\n", total)
}
