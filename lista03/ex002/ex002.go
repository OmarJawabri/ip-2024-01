package main

import "fmt"

func main(){
	var N, K, x int
	fmt.Scanln(&N)
	for N < 1 || N > 1000{
		fmt.Println("Numero inválido, digite outro número: ")
		fmt.Scanln(&N)
	}
	V := make([]int, 10, 10)
	for i := 0; i < N; i++ {
		fmt.Scanln(&V[i]) //SCANF ERROR
	}

	fmt.Scanln(&K)

	for i := 0; i < N; i++ {
		if K <= V[i]{
			x++
		}
	}
	fmt.Println(x)
}