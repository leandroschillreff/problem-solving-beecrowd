package main

import (
	"fmt"
)

func main() {
	var code1, quantity1, code2, quantity2 int
	var price1, price2 float64
	fmt.Scan(&code1, &quantity1, &price1, &code2, &quantity2, &price2)

	valueToPay := (float64(quantity1) * price1) + (float64(quantity2) * price2)
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", valueToPay)
}
