package main

import (
	"fmt"
	"os"
)

func main(){
	const num = 1000
	var maiorindice[num]int
	var maiornumero[num]int
	var N, x, y int
	x = 1
	N = 2
	if N > 1 && N <= 10000{
		for i := 0; x != 0; i++ {
			fmt.Scanln(&N)
			x = N
			vetor := make([]int, 1000, 1000)
			if N > 1 && N <= 10000{
				for k := 0; k < N; k++ {
					fmt.Scanln(&vetor[k])
					if vetor[k] < 0 || vetor[k] > 1000 {
						os.Exit(0)
					}
				}
				for j := 0; j < N; j++ {
					if vetor[j] > maiornumero[i]{
						maiornumero[i] = vetor[j]
						maiorindice[i] = j
					}
				}
				y++ 
				vetor = nil
			} else if N < 1 && N != 0 || N > 10000{
				break
			}
		}
		for i := 0; i < y; i++ {
			fmt.Printf("%d %d\n", maiorindice[i], maiornumero[i])
		}
	} 
}