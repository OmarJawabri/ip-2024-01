package main

import (
	"fmt"
	"strconv"
	"strings"
)
func maior(x []int)int{
	var maior int
	for _, c:= range x{
		if c > maior{
			maior = c
		}
	}
	return maior
}

func qtdMenores(x []int, maior int)[]int{
	qtd := make([]int, maior+1)
	for i := 0; i < maior+1; i++ {
		for k := 0; k < len(x); k++ {
			if i >= x[k]{
				qtd[i]++
			}
		}
	}
	return qtd
} 

func converterString(x []int)string{
	convertido := ""
	for i := 0; i < len(x); i++ {
		convertido += strconv.Itoa(x[i])
		convertido += " "
	}
	return convertido
}

func main(){
	var N, y int
	N = 1
	var convertido string
	maiores := []int{}
	if N >= 1 && N <= 10000{
		for i := 0;N != 0;i++{
			fmt.Scanln(&N)
			inteiro := make([]int, N)
			for i := 0; i < N; i++ {
				fmt.Scanln(&inteiro[i])
			}
			maiorN := maior(inteiro)
			maiores = append(maiores, maiorN)
			qtd := qtdMenores(inteiro, maiorN)
			convertido += converterString(qtd)
			inteiro = nil
		}
		armazenar := strings.Split(convertido, " ")
		for i := 0; i < len(maiores); i++ {
			for k := 0; k < maiores[i]+1; k++ {
				if y != len(armazenar)-2{
					fmt.Printf("(%d) %s\n", k, armazenar[y])
					y++
				}
			}
		}
	}
}
