package main

import (
	"fmt"
)

func main(){
	var ValorA, ValorB, ValorC, ValorD, determinante float64
	fmt.Scan(&ValorA)
	fmt.Scan(&ValorB)
	fmt.Scan(&ValorC)
	fmt.Scan(&ValorD)
	determinante = ValorA*ValorD - ValorB*ValorC
	fmt.Printf("O VALOR DO DETERMINANTE É = %.2f", determinante)
}