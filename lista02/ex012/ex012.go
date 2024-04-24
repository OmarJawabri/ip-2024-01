package main

import "fmt"

func main(){
	const numVezes = 9999
	var ValorIngresso, ValorInicial, ValorFinal, ValorMaior, LucroMaior, IngressoMaior float64
	fmt.Scanln(&ValorIngresso)
	fmt.Scanln(&ValorInicial)
	fmt.Scanln(&ValorFinal)
	var Ingresso[numVezes]float64
	var Lucro[numVezes]float64
	y := ValorInicial
	if ValorInicial >= ValorFinal{
		fmt.Println("INTERVALO INVÁLIDO.")
	} else{
		for i := 0; y <= ValorFinal; i++ {
			if  y < ValorIngresso{
				x := (ValorIngresso - y)/0.5
				Ingresso[i] = 120 + 25*x
			} else if y > ValorIngresso{
				x := (y - ValorIngresso)/0.5
				Ingresso[i] = 120 - 30*x
			} else if y == ValorIngresso{
				Ingresso[i] = 120
			}
			Despesas := 200 + Ingresso[i]*0.05
			Lucro[i] = y*Ingresso[i] - Despesas
			if LucroMaior < Lucro[i]{
				LucroMaior = Lucro[i]
				ValorMaior = y
				IngressoMaior = Ingresso[i]
			} else if LucroMaior < 0{
				LucroMaior = 0
				ValorMaior = 0
				IngressoMaior = 0
			}
			y++
		}
		h := ValorInicial
		for i := 0; h <= ValorFinal; i++ {
			fmt.Printf("V: %.2f, N: %.0f, L: %.2f\n", h, Ingresso[i], Lucro[i])
			h++
		}
		if ValorInicial < ValorFinal{
		fmt.Printf("Melhor valor final: %.2f\n", ValorMaior)
		fmt.Printf("Lucro: %.2f\n", LucroMaior)
		fmt.Printf("Número de ingressos: %.0f", IngressoMaior)
		}
	}
}