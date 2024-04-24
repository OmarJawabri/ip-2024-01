package main

import "fmt"

func main(){
	var n, s, y, i float64
	var K, x int
	fmt.Scanln(&n)
	fmt.Scanln(&i)
	fmt.Scanln(&K)
	fmt.Scanln(&s)
	soma := make([]float64, K, K)
	subtracao := make([]float64, K, K)
	multiplicacao := make([]float64, K, K)
	divisao := make([]float64, K, K)
	incremento := make([]float64, K, K)
	if n >= 0 && n <= 9 {
		for x = 0; x < K; x++ {
			incremento[0] = i
			incremento[x] = incremento[0] + s*y
			soma[x] = n + incremento[x]
			subtracao[x] = n - incremento[x]
			multiplicacao[x] = n*incremento[x]
			divisao[x] = n/incremento[x]
			y++
		}
		fmt.Println("Tabuada de soma:")
		for x = 0; x < K; x++ {
			fmt.Printf("%.2f + %.2f = %.2f\n", n, incremento[x], soma[x])
		}
		fmt.Println("Tabuada de subtração:")
		for x = 0; x < K; x++ {
			fmt.Printf("%.2f - %.2f = %.2f\n", n, incremento[x], subtracao[x])
		}
		fmt.Println("Tabuada de multiplicação:")
		for x = 0; x < K; x++ {
			fmt.Printf("%.2f * %.2f = %.2f\n", n, incremento[x], multiplicacao[x])
		}
		fmt.Println("Tabuada de divisão:")
		for x = 0; x < K; x++ {
			fmt.Printf("%.2f / %.2f = %.2f\n", n, incremento[x], divisao[x])
		}
	}
}