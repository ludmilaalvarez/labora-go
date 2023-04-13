package main

import (
	"fmt"
	"strings"
)

var (
	monedas  = 50
	usuarios = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(usuarios))
)

func main() {
	distribucion := distribuirBTC(usuarios, monedas, distribution)
	fmt.Println(distribucion)
	fmt.Println("Coins left:", monedas)
}

func contarVocales(usuario string) int {
	vocales := 0

	for _, v := range strings.ToLower(usuario) {

		if string(v) == "a" {
			vocales++
		}
		if string(v) == "e" {
			vocales++
		}
		if string(v) == "i" {
			vocales += 2
		}
		if string(v) == "o" {
			vocales += 3
		}
		if string(v) == "u" {
			vocales += 4
		}
	}
	return vocales
}

func distribuirBTC(usuarios []string, monedas int, distribution map[string]int) map[string]int {
	for _, nombre := range usuarios {
		vocales := contarVocales(nombre)
		if vocales > 10 {
			vocales = 10
		}
		distribution[nombre] = vocales
	}
	return distribution
}
