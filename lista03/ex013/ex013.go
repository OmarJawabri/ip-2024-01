package main

import "fmt"

func main(){
	var N, maiorf, valor, y int
	fmt.Scanln(&N)
	inteiros := make([]int, N)
	if N >= 1 && N <= 1000000{
		for i := 0; i < N; i++ {
			fmt.Scanln(&inteiros[i])
		}
		frequencia := make([]int, N)
		for i := 0; i < N; i++ {
			if inteiros[i] >= 0 && inteiros[i] <= 100{
				for j := 0; j < N; j++ {
					if inteiros[i] == inteiros[j]{
						frequencia[i]++
					}
				}
				if frequencia[i] > maiorf{
					maiorf = frequencia[i]
					valor = inteiros[i]
				}
			} else{
				y++
			}
		}
		if y == 0{
			fmt.Println(valor)
			fmt.Println(maiorf)
		} else{
			fmt.Println("Algum dos números digitados excedeu um dos limites")
		}
	}
}