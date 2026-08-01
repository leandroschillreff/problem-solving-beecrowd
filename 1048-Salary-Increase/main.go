package main

import (
	"fmt"
)

func main() {
	var currentSalary float64
	fmt.Scan(&currentSalary)

	var percentage int

	switch {
	case currentSalary <= 400.00:
		percentage = 15
	case currentSalary <= 800.00:
		percentage = 12
	case currentSalary <= 1200.00:
		percentage = 10
	case currentSalary <= 2000.00:
		percentage = 7
	default:
		percentage = 4
	}

	moneyEarned := currentSalary * (float64(percentage) / 100.0)
	newSalary := currentSalary + moneyEarned

	fmt.Printf("Novo salario: %.2f\n", newSalary)
	fmt.Printf("Reajuste ganho: %.2f\n", moneyEarned)
	fmt.Printf("Em percentual: %d %%\n", percentage)
}