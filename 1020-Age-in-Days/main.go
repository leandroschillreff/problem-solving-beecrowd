package main

import (
	"fmt"
)

func main() {
	var ageInDays int

	fmt.Scan(&ageInDays)

	years := ageInDays / 365
	remainder := ageInDays % 365
	months := remainder / 30
	days := remainder % 30

	fmt.Printf("%d ano(s)\n", years)
	fmt.Printf("%d mes(es)\n", months)
	fmt.Printf("%d dia(s)\n", days)

}
