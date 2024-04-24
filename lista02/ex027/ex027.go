package main

import "fmt"

func main(){
	var naoPrimo string
	var N int
	const numVezes = 999
	fmt.Scanln(&N)
	numero := N
	var nprimos[numVezes]int
	x := 0
	for i := 0; i < 5; i++ {
		if N < 1{
			fmt.Printf("Fatoração não é possivel para o número %d!\n", N)
			fmt.Scanln(&N)
			numero = N
		}
	}
	if N > 1{
		for i := 2; i <= numero; i++ {
			for k := 2; k <= numero; k++ {
				if i % k != 0 && N % i == 0{
					N = N/i
					nprimos[x] = i
					x++
				}
			}
		}
	}
	for i := 2; i < numero; i++ {
		if numero % i == 0 {
			naoPrimo = "não primo"
		}	
	}
	if naoPrimo == "não primo" {
		fmt.Printf("%d = ", numero)
		for i := 0; i < x; i++ {
			if nprimos[i+1] != 0 {
				fmt.Printf("%d x ", nprimos[i])
			} else{
				fmt.Printf("%d ", nprimos[i])
			}
		}
	} else{
		fmt.Printf("%d = %d", numero, numero)
	}
}
