package main

import "fmt"

func troca(vetor []int, i, j int){
	vetor[i], vetor[j] = vetor[j], vetor[i]
}

func trocaOpostosSeMenor(vetor []int, tamanho int){
	for i := 0; i < tamanho; i++ {
		for j := 0; j < tamanho; j++ {
			oposto := tamanho - (i+1)
			if vetor[j] > vetor[i] && j > i && j == oposto{
				troca(vetor, i, j)
			}
		}
	}
}

func main(){
	var numCasos, N int
	fmt.Scanln(&numCasos)
	for i := 0; i < numCasos; i++ {
		fmt.Scanln(&N)
		if N <= 500{
			vetor := make([]int, N)
			for i := 0; i < N; i++ {
				fmt.Scan(&vetor[i])
			}
			trocaOpostosSeMenor(vetor, N)
			fmt.Println(vetor)
		}
	}
}