package main

import "fmt"

func calculoArray(x []float64, n int)float64{
	if n < 0{
		return 0
	}
	return x[n] + calculoArray(x, n-1)
}

func main(){
	var tamanho int
	fmt.Println("Digite o tamanho da array que deseja:")
	fmt.Scanln(&tamanho)
	array := make([]float64, tamanho)
	fmt.Println("Digite o número de cada elemento da array:")
	for i := 0; i < len(array); i++{
		fmt.Scanln(&array[i])
	}
	fmt.Printf("%.2f\n", calculoArray(array, tamanho-1))
}
