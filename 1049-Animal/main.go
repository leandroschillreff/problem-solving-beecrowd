package main

import (
	"fmt"
)

func main() {
	var phylum, class, diet string
	fmt.Scan(&phylum, &class, &diet)

	if phylum == "vertebrado" {
		if class == "ave" {
			if diet == "carnivoro" {
				fmt.Println("aguia")
			} else {
				fmt.Println("pomba")
			}
		} else {
			if diet == "onivoro" {
				fmt.Println("homem")
			} else {
				fmt.Println("vaca")
			}
		}
	} else {
		if class == "inseto" {
			if diet == "hematofago" {
				fmt.Println("pulga")
			} else {
				fmt.Println("lagarta")
			}
		} else {
			if diet == "hematofago" {
				fmt.Println("sanguessuga")
			} else {
				fmt.Println("minhoca")
			}
		}
	}
}
