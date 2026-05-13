package main

import (
	"fmt"
)

func main() {
	var amount float64
	fmt.Scan(&amount)

	totalCents := int(amount*100 + 0.5)

	fmt.Println("NOTAS:")
	banknotes := []int{10000, 5000, 2000, 1000, 500, 200}
	for _, note := range banknotes {
		quantity := totalCents / note
		fmt.Printf("%d nota(s) de R$ %.2f\n", quantity, float64(note)/100.0)
		totalCents %= note
	}

	fmt.Println("MOEDAS:")
	coins := []int{100, 50, 25, 10, 5, 1}
	for _, coin := range coins {
		quantity := totalCents / coin
		fmt.Printf("%d moeda(s) de R$ %.2f\n", quantity, float64(coin)/100.0)
		totalCents %= coin
	}
}
