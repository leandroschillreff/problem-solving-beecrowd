package main

import (
	"fmt"
)

const commissionRate = 0.15

func main() {
	var name string
	var baseSalary, totalSales float64

	fmt.Scan(&name, &baseSalary, &totalSales)

	totalSalary := baseSalary + (totalSales * commissionRate)
	fmt.Printf("TOTAL = R$ %.2f\n", totalSalary)
}
