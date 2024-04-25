package main

import "fmt"

func main(){
	var n, valorTotal, valorPrimeiraLinha int
	fmt.Scanln(&n)
	valorPrimeiraLinha = n*5 + n*2*3
	valorTotal = valorPrimeiraLinha + (n*4 + n*2*4)*7
	fmt.Println(valorTotal)
}