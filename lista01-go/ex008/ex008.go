package main

import (
	"fmt"
	"math"
)

func main() {
	var RaioLata, Altura, Ac, Al, At, Custo, pi float64
	fmt.Scan(&RaioLata)
	fmt.Scan(&Altura)
	pi = 3.14159
	Ac = pi * math.Pow(RaioLata, 2)
	Al = 2 * pi * RaioLata * Altura
	At = 2*Ac + Al
	Custo = 100 * At
	fmt.Printf("O VALOR DE CUSTO É = %.2f", Custo)
}
