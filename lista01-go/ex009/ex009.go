package main

import (
	"fmt"
	"math"
)

func main(){
	var CoefA, CoefB, CoefC, delta float64
	fmt.Scan(&CoefA)
	fmt.Scan(&CoefB)
	fmt.Scan(&CoefC)
	delta = math.Pow(CoefB, 2) - 4*CoefA*CoefC
	fmt.Printf("O VALOR DE DELTA É = %.2f\n", delta) 
}