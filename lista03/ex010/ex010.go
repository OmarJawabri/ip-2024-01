package main

import "fmt"

func frequencia(x int, y []int)int{
	var freq int
	for _, c:= range y{
		if c == x{
			freq++
		}
	}
	return freq
}

func maiorNota(N int, y []int){
	var maior, indice int
	for i := 0; i < N; i++ {
		if y[i] > maior {
			maior = y[i]
			indice = i
		}
	}
	fmt.Printf("Nota %d, indice %d", maior, indice)
}

func main(){
	var N, x, ultima, y int
	fmt.Scanln(&N)
	notas := make([]int, N)
	if N >= 1 && N <= 10000{
		for i := 0; i < N; i++ {
			fmt.Scanln(&notas[i])
		}
		for i := 0; i < N; i++ {
			if notas[i] >= 0 && notas[i] <= 10{
				x = notas[N-1]
				ultima = frequencia(x, notas)
			} else{
				y++
			}
		}
		if y == 0{
			fmt.Printf("Nota %d, %d vezes\n", x, ultima)
			maiorNota(N, notas)
		} else if y != 0{
			fmt.Println("Notas acima de 10 ou abaixo de 0")
		}
	}
}