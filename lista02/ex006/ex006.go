package main

import "fmt"

func main(){
	var n, k, x, kmaior int
	x = 1
	fmt.Scanln(&n)
	numero := make([]int, n, n)
	armazenamento := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d",&numero[i])
	}
	if n > 0 {
		for i := 1; i < n; i++ {
			if numero[i] == numero[i-1] {
				x++
			} 
			if numero[i] != numero[i-1]{
				x = 1
			}
			armazenamento[i] = x
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
	}
	fmt.Println("A maior subsequência de números iguais possui ", kmaior, " elementos.")
}