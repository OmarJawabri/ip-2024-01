package main

import (
	"fmt"
	"math"
)

func main() {
	var Horas, Valor float64
	fmt.Scan(&Horas)
	Valor = ((Horas - math.Mod(Horas, 3))/3)*10 + math.Mod(Horas, 3)*5
	fmt.Printf("O VALOR A PAGAR É = %.2f", Valor)
}
