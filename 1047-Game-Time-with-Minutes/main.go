package main

import (
	"fmt"
)

func main() {
	var startHour, startMinute, endHour, endMinute int
	fmt.Scan(&startHour, &startMinute, &endHour, &endMinute)

	startTotalMinutes := startHour*60 + startMinute
	endTotalMinutes := endHour*60 + endMinute

	if endTotalMinutes <= startTotalMinutes {
		endTotalMinutes += 24 * 60
	}

	durationInMinutes := endTotalMinutes - startTotalMinutes

	hours := durationInMinutes / 60
	minutes := durationInMinutes % 60

	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", hours, minutes)
}