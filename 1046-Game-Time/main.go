package main

import (
	"fmt"
)

func main() {
	var startTime, endTime int
	fmt.Scan(&startTime, &endTime)

	var duration int

	if startTime == endTime {
		duration = 24
	} else if endTime > startTime {
		duration = endTime - startTime
	} else {
		duration = (24 - startTime) + endTime
	}

	fmt.Printf("O JOGO DUROU %d HORA(S)\n", duration)
}
