package main

import "fmt"

func main(){
	var N, M, l int
	fmt.Scanln(&N)
	V := make([]int, N, N)
	if N >= 1 && N <= 100000{
		for i := 0; i < N; i++ {
			fmt.Scanln(&V[i]) //SCANF ERROR
		}
		fmt.Scanln(&M)
		Numero := make([]int, M)
		if M >= 1 && M <= 1000{
			for i := 0; i < M; i++ {
				fmt.Scanln(&Numero[i])
			}
			for i := 0; i < M; i++ {
				for k := 0; k < N; k++ {
					if Numero[i] == V[k]{
						fmt.Println("ACHEI")
						l++
					}
				}
				if l == 0{
					fmt.Println("NÃO ACHEI")
				}
				l = 0
			}
		}
	}
}