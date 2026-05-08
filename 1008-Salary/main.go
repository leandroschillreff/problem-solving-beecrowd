package main

import (
	"fmt"
)

func main() {
	var employeeNumber int
	var hoursWorked, receivedPerHour float64

	fmt.Scanln(&employeeNumber)
	fmt.Scanln(&hoursWorked)
	fmt.Scanln(&receivedPerHour)

	salary := hoursWorked * receivedPerHour
	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", employeeNumber, salary)

}
