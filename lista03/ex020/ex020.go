package main

import (
	"fmt"
	"math"
)

func maior(xt, yt, zt float64){
	var maior float64
	if xt >= yt && xt >= zt{
		maior = xt
	} else if yt >= xt && yt >= zt{
		maior = yt
	} else if zt >= xt && zt >= yt{
		maior = zt
	}
	fmt.Printf("%.2f\n", maior)
}

func main(){
	var N int
	fmt.Scanln(&N)
	x := make([]float64, N)
	y := make([]float64, N)
	z := make([]float64, N)
	h := make([]int, 1000)
	if N >= 2 && N <= 1000{
		for i := 0; i < N; i++ {
			fmt.Scanln(&x[i], &y[i], &z[i])
			if x[i] >= -1000 && x[i] <= 1000 && y[i] >= -1000 && y[i] <= 1000 && z[i] >= -1000 && z[i] <= 1000{
				continue
			} else{
				h[i]++
			}
		}
		for i := 1; i < N; i++ {
			if h[i] == 0 && h[i-1] == 0{
				xt := math.Abs(x[i] - x[i-1])
				yt := math.Abs(y[i] - y[i-1])
				zt := math.Abs(z[i] - z[i-1])
				maior(xt, yt, zt)
			} else{
				fmt.Println("Numero do vetor excedeu o limite")
			}
		}
	}
}