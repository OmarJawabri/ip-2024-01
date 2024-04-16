package main

import (
	"fmt"
)

func main() {
	var numerodeCasos int
	var RendaDoJogo, Popular, Geral, Cadeira, Arquibancada float64
	fmt.Scanln(&numerodeCasos)
	Total := make([]float64, numerodeCasos, numerodeCasos)
	PPopular := make([]float64, numerodeCasos, numerodeCasos)
	PGeral := make([]float64, numerodeCasos, numerodeCasos)
	PArquibancada := make([]float64, numerodeCasos, numerodeCasos)
	PCadeira := make([]float64, numerodeCasos, numerodeCasos)
	for i := 0; i < numerodeCasos; i++ {
		fmt.Scanln(&Total[i], &PPopular[i], &PGeral[i], &PArquibancada[i], &PCadeira[i])
	}
	n := 1
	for i := 0; i < numerodeCasos; i++ {
		Popular = (PPopular[i]/100)*Total[i]
		Geral = (PGeral[i]/100)*Total[i]
		Cadeira = (PCadeira[i]/100)*Total[i]
		Arquibancada = (PArquibancada[i]/100)*Total[i]
		RendaDoJogo = Popular + Geral*5 + Arquibancada*10 + Cadeira*20
		fmt.Printf("A RENDA DO JOGO Nº %d É = %.2f\n", n, RendaDoJogo)
		n++
	}
}
