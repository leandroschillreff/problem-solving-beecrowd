package main

import (
	"fmt"
	"sort"
)

func main() {
	var n1, n2, n3 float64
	fmt.Scan(&n1, &n2, &n3)

	sides := []float64{n1, n2, n3}
	sort.Float64s(sides)

	a := sides[2]
	b := sides[1]
	c := sides[0]

	if a >= b+c {
		fmt.Println("NAO FORMA TRIANGULO")
		return
	}

	if a*a == (b*b + c*c) {
		fmt.Println("TRIANGULO RETANGULO")
	} else if a*a > (b*b + c*c) {
		fmt.Println("TRIANGULO OBTUSANGULO")
	} else if a*a < (b*b + c*c) {
		fmt.Println("TRIANGULO ACUTANGULO")
	}

	if a == b && b == c {
		fmt.Println("TRIANGULO EQUILATERO")
	} else if a == b || b == c || a == c {
		fmt.Println("TRIANGULO ISOSCELES")
	}
}
