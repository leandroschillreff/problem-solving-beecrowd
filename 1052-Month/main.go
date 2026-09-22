package main

import (
	"fmt"
)

func main() {
	var monthNumber int
	fmt.Scan(&monthNumber)

	months := []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	fmt.Println(months[monthNumber-1])
}