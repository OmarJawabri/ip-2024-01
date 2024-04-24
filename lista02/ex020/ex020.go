package main

import "fmt"

func main(){
	const numValor = 999999999
	var n, soma, maior int
	var d[numValor]int
	fmt.Scanln(&n)
	for i := 1; i < n; i++ {
		if n % i == 0{
			d[i] = i
			soma = soma + i
			if d[i] > maior {
				maior = d[i]
			}
		}
	}
	fmt.Printf("%d = ", n)
	for i := 1; i < n; i++ {
		if n % i == 0 {
			fmt.Printf("%d", d[i])
			if d[i] != maior {
				fmt.Printf(" + ")
			}
		}
	}
	if soma == n {
		fmt.Printf(" = %d (NÚMERO PERFEITO)", soma)
	} else{
		fmt.Printf(" = %d (NÚMERO NÃO É PERFEITO)", soma)
	}
}