package main

import (
	"fmt"
	"math"
)

func main(){
	var n, i, h int
	var x float64
	n = 99
	p1 := make([]float64, n, n)
	p2 := make([]float64, n, n)
	p3 := make([]float64, n, n)
	p4 := make([]float64, n, n)
	p5 := make([]float64, n, n)
	p6 := make([]float64, n, n)
	p7 := make([]float64, n, n)
	p8 := make([]float64, n, n)
	l1 := make([]float64, n, n)
	l2 := make([]float64, n, n)
	l3 := make([]float64, n, n)
	l4 := make([]float64, n, n)
	l5 := make([]float64, n, n)
	matricula := make([]int, n, n)
	presenca := make([]float64, n, n)
	ppresenca := make([]float64, n, n)
	mp := make([]float64, n, n)
	ml := make([]float64, n, n)
	nf := make([]float64, n, n)
	nt := make([]float64, n, n)
	for i = 0; x != -1; i++ {
		fmt.Scanln(&matricula[i], &p1[i], &p2[i], &p3[i], &p4[i], &p5[i], &p6[i], &p7[i], &p8[i], &l1[i], &l2[i], &l3[i], &l4[i], &l5[i], &nt[i], &presenca[i])
		mp[i] = (p1[i] + p2[i] + p3[i] + p4[i] + p5[i] + p6[i] + p7[i] + p8[i])/8
		ml[i] = (l1[i] + l2[i] + l3[i] + l4[i] + l5[i])/5
		nf[i] = 0.7*mp[i] + 0.15*ml[i] + 0.15*nt[i]
		ppresenca[i] = presenca[i]/120
		x = presenca[i]
		h++
	}
	for i = 0; i < h; i++ {
		if matricula[i] != -1 {
			if nf[i] >= 6 && ppresenca[i] >= 0.75{
				fmt.Println("Matrícula: ", matricula[i], ",Nota final: ", math.Round(nf[i]*100)/100, ",Situação final: APROVADO")
			} else if nf[i] < 6 && ppresenca[i] >= 0.75{
				fmt.Println("Matrícula: ", matricula[i], ",Nota final: ", math.Round(nf[i]*100)/100, ",Situação final: REPROVADO POR NOTA")
			} else if nf[i] >= 6 && ppresenca[i] < 0.75{
				fmt.Println("Matrícula: ", matricula[i], ",Nota final: ", math.Round(nf[i]*100)/100, ",Situação final: REPROVADO POR FREQUÊNCIA")
			} else if nf[i] < 6 && ppresenca[i] < 0.75{
				fmt.Println("Matrícula: ", matricula[i], ",Nota final: ", math.Round(nf[i]*100)/100, ",Situação final: REPROVADO POR NOTA E POR FREQUÊNCIA")
			}
		}
	} 
}