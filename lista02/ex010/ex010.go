package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	const numMatricula = 999
	const numHorasTrabalhadas = 999
	const valorHora = 999
	const salarioT = 999
	var matricula[numMatricula]int
	var horasT[numHorasTrabalhadas]float64
	var valor[valorHora]float64
	var salario[salarioT]float64
	x := 1
	y := 0
	for i := 0; x != 0; i++ {
		fmt.Scanf("%d %f %f", &matricula[i], &horasT[i], &valor[i])
		salario[i] = valor[i]*horasT[i]
		bufio.NewReader(os.Stdin).ReadByte()
		x = matricula[i]
		y++
	}
	for i := 0; i < y-1; i++ {
		fmt.Printf("%d %.2f\n", matricula[i], salario[i])
	}
}