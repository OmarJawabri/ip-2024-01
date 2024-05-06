package main

import (
	"fmt"
	"sort"
)

func mediana(x []int, N int){
	var median float64
	for i := 0; i < N; i++ {
		if N % 2 == 0{
			if N/2 == i{
				median = float64(x[i-1] + x[i])/2
			} 
		} else{
			if int(N/2) + 1 == i{
				median = float64(x[i-1])
			}
		}
	}
	fmt.Printf("%.2f", median)
}

func main(){
	var N int
	fmt.Scanln(&N)
	vetor := make([]int, N)
	if N > 0 && N <= 1000000{
		for i := 0; i < N; i++ {
			fmt.Scanln(&vetor[i])
		}
		sort.Ints(vetor)
		mediana(vetor, N)
	}
}
