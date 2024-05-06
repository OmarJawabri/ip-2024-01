package main

import "fmt"

func vetorInverso(x int, vetW []int){
	fmt.Println()
	for i := x-1; i >=0; i-- {
		if i != 0{
			fmt.Print(vetW[i], " ")
		} else{
			fmt.Print(vetW[i])
		}
	}
	fmt.Println()
}

func main(){
	var tamanho, maiorV, menorW int
	fmt.Scanln(&tamanho)
	vetorV := make([]int, tamanho)
	menorW = 10000
	if tamanho >= 1 && tamanho <= 1000{
		for i := 0; i < tamanho; i++ {
			fmt.Scan(&vetorV[i])
		}
		for i := 0; i < tamanho; i++ {
			if i != tamanho-1{
				fmt.Print(vetorV[i], " ")
			} else{
				fmt.Print(vetorV[i])
			}
			if vetorV[i] > maiorV{
				maiorV = vetorV[i]
			}
			if vetorV[i] < menorW{
				menorW = vetorV[i]
			}
		}
	}
	if tamanho >= 1 && tamanho <= 1000{
		vetorInverso(tamanho, vetorV)
		fmt.Println(maiorV)
		fmt.Println(menorW)
	}
}