package main

import (
	"fmt"
	"math"
)

func main() {
	var n, f int
	var Soma, x float64
	f = 1
	x = 1
	fmt.Scan(&n)
	if n >= 1 {
		for n >= f{
			Soma = Soma + 1/x
			x++
			f++
		}
		fmt.Println(math.Round(Soma*1000000)/1000000)
	} else {
		fmt.Println("Número Inválido!")
	}
}
