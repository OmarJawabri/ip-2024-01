package main

import (
	"fmt"
	"math"
)

func main(){
	var x, N float64
	fmt.Scan(&N)
	x = 2
	if N > 5 {
		if N < 2000 {
			if math.Mod(N, 2) == 0 {
				for N >= x{
					s := x*x
					fmt.Printf("%.0f^2 = %.0f\n",x, s)
					x = x + 2
				}
			}else{
				for (N-1) >= x{
					s:= x*x
					fmt.Printf("%.0f^2 = %.0f\n", x, s)
					x = x + 2					
				}
			} 
		}else{
			fmt.Println("Número Inválido!")
		}
	}else{
		fmt.Println("Número Inválido!")
	}
}

