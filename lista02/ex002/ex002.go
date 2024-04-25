package main

import "fmt"

func main(){
	var paisA, paisB float64
	Ano := 0
	fmt.Scanln(&paisA)
	fmt.Scanln(&paisB)
	if paisB > paisA {
		for paisA < paisB {
			paisA = paisA*1.03
			paisB = paisB*1.015
			Ano++
		}
		fmt.Println("ANOS = ", Ano)
	} else{
		fmt.Println("Valor de moradores do país A maior que moradores do país B é inválido!")
	}
}