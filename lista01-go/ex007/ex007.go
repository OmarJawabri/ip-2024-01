package main

import (
	"fmt"
)

func main(){
	var Fahrenheit, QtdChuva, Celsius, Polegadas float64
	fmt.Scan(&Fahrenheit)
	fmt.Scan(&Polegadas)
	Celsius = (5*Fahrenheit - 160)/9
	QtdChuva = Polegadas*25.4
	fmt.Printf("O VALOR EM CELSIUS É = %.2f\n", Celsius)
	fmt.Printf("A QUANTIDADE DE CHUVA É = %.2f", QtdChuva)
}