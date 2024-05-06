package main

import (
	"fmt"
	"math"
)

func comparar(x []int, N int)int{
	var menordif, y float64
	menordif = 10000
	if N >= 2 && N <= 1000{
		for i := 1; i < N; i++ {
			for k := 0; k < N; k++ {
				if math.Abs(float64(x[i] - x[k])) < menordif && math.Abs(float64(x[i] - x[k])) != 0{
					menordif = math.Abs(float64(x[i] - x[k]))
					y++
				} else if y == 0{
					menordif = 0
				}
			}	
		}
	}
	return int(menordif)
}

func progressao(N int)int{
	var soma int
	for i := 0; i < N; i++ {
		soma = soma + i
	}
	return soma
}

func main(){
	var T, N, y, x, l int
	var comparacao[99]int
	var progressaovetor[99]int
	fmt.Scanln(&T)
	if T >= 1 && T <= 10{
		for i := 0; i < T; i++ {
			fmt.Scanln(&N)
			vetor := make([]int, N)
			if N >= 2 && N <= 1000{
				for k := 0; k < N; k++ {
					fmt.Scanln(&vetor[k])
					l++
					if vetor[k] >= -1000 && vetor[k] <= 1000{
						y++
					}
				}
				comparacao[i] = comparar(vetor, N)
				progressaovetor[i] = progressao(N)
				vetor = vetor[:0]
			} else{
				x++
			}
		}
		if y == l && x == 0{
			for i := 0; i < T; i++ {
				fmt.Printf("%d ", comparacao[i])
				fmt.Println(progressaovetor[i])
			}
		}
		if y != l && x == 0{
			fmt.Println("Valor de um ou mais elemento(s) do vetor excedeu o limite")
		} 
		if x != 0 && y == 0{
			fmt.Println("Tamanho do vetor inválido")
		}
	}
}