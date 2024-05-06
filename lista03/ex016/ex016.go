package main

import "fmt"

func presencas(x []int)int{
	var presencatempo int
	for _, c := range x{
		if c <= 0{
			presencatempo++
		}
	}
	return presencatempo
}

func imprimir(x []int, N int){
	for i := N-1; i >= 0; i-- {
		if x[i] <= 0{
			fmt.Println(i+1)
		}
	}
}

func main(){
	var N, K, tempocerto int
	fmt.Scanln(&N, &K)
	presenca := make([]int, N)
	if N >= 0 && N <= 1000 && K >= 0 && K <= 1000{
		for i := 0; i < N; i++ {
			fmt.Scan(&presenca[i])
		}
		tempocerto = presencas(presenca)
		if tempocerto >= K{
			fmt.Println("NÃO")
			imprimir(presenca, N)
		} else{
			fmt.Println("SIM")
		}
	}
}