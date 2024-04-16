package main

import "fmt"

func main(){
	var Horas, Minutos, Segundos, TempoTotal int
	fmt.Scan(&Horas)
	fmt.Scan(&Minutos)
	fmt.Scan(&Segundos)
	TempoTotal = Horas*3600 + Minutos*60 + Segundos
	fmt.Println("O TEMPO EM SEGUNDOS É = ", TempoTotal)
}