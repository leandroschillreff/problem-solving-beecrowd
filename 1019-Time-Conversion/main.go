package main

import (
	"fmt"
)

func main() {
	var duration int

	fmt.Scan(&duration)

	hours := duration / 3600
	remainder := duration % 3600
	minutes := remainder / 60
	seconds := remainder % 60

	fmt.Printf("%d:%d:%d\n", hours, minutes, seconds)
}
