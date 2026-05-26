package main

import (
	"fmt"
	"math"
)

func main() {
	var n1, n2, n3, n4 float64

	fmt.Scan(&n1, &n2, &n3, &n4)

	average := (n1*2 + n2*3 + n3*4 + n4*1) / 10.0

	average = math.Floor(average*10) / 10.0

	fmt.Printf("Media: %.1f\n", average)

	if average >= 7.0 {
		fmt.Println("Aluno aprovado.")
	} else if average < 5.0 {
		fmt.Println("Aluno reprovado.")
	} else {
		fmt.Println("Aluno em exame.")

		var examGrade float64
		fmt.Scan(&examGrade)
		fmt.Printf("Nota do exame: %.1f\n", examGrade)

		finalAverage := (average + examGrade) / 2.0

		if finalAverage >= 5.0 {
			fmt.Println("Aluno aprovado.")
		} else {
			fmt.Println("Aluno reprovado.")
		}

		fmt.Printf("Media final: %.1f\n", finalAverage)
	}
}
