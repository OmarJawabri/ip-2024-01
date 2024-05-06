package main

import "fmt"

func main(){
	var n, soma int
	fmt.Scanln(&n)
	vetor := make([]int, n, n)
	if n < 5000 && n > 0{
		for i := 0; i < n; i++ {
			fmt.Scan(&vetor[i])
			soma = soma + vetor[i]
		}
		fmt.Println(soma)
	}
}
