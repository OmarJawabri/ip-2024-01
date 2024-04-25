package main

import (
	"fmt"
)

func main(){
	var n, k, x, kmaior int
	fmt.Scanln(&n)
	numero := make([]int, n, n)
	armazenamento := make([]int, n, n)
	if n > 0 {
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &numero[i])
		}
		for i := 1; i < n; i++ {
			if numero[i] > numero[i-1] {
				x++
			}
			armazenamento[i] = x 
			if numero[i] < numero[i-1] {
				x = 0
				armazenamento[i] = x
			} 
		}
		for i := 1; i < n; i++ {
			if armazenamento[i] > armazenamento[i-1] {
				k = armazenamento[i]
			} else if armazenamento[i-1] > armazenamento[i]{
				k = armazenamento[i-1]
			}
			if k > kmaior {
				kmaior = k
			}
		}
		fmt.Println("O comprimento do segmento crescente máximo é: ", kmaior)
	}
}