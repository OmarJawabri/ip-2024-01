package main

import (
	"fmt"
)

func main(){
	var Nota float64
	var Conceito string
	fmt.Scan(&Nota)
	if Nota <= 10 && Nota >=9{
		Conceito = "A"
	} else if Nota >= 7.5 && Nota < 9 {
		Conceito = "B"	
	} else if Nota >= 6 && Nota < 7.5 {
		Conceito = "C"	
	} else if Nota >= 0 && Nota < 6{
		Conceito = "D"
	}
	fmt.Printf("NOTA = %.1f CONCEITO = %s", Nota, Conceito)
}