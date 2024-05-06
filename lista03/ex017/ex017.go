package main

import "fmt"

func unico(x []int, N int)int{
	var diferente, unico int
	for i := 0; i < N; i++ {
		for k := 0; k < N; k++ {
			if x[i] != x[k] {
				diferente++
			}
		}
		if diferente == N-1{
			unico++
		}
		diferente = 0
	}
	return unico
}

func main(){
	var N int
	fmt.Scanln(&N)
	inteiros := make([]int, N)
	if N < 5000{
		for i := 0; i < N; i++ {
			fmt.Scan(&inteiros[i])
		}
		valorunico := unico(inteiros, N)
		if valorunico != 0{
			fmt.Println(valorunico)
		} else{
			fmt.Println("Não há valor único nesse vetor")
		}
	}
}