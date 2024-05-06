package main

import (
	"fmt"
	"math"
)

func main(){
	var N int
	var xt, yt, zt float64
	var x[1000]float64
	var y[1000]float64
	var z[1000]float64
	var distancia[1000]float64
	fmt.Scanln(&N)
	if N >= 2 && N <= 1000{
		for i := 0; i < N; i++ {
			fmt.Scanln(&x[i], &y[i], &z[i])
		}
		for k := 1; k <= N-1; k++ {
			if x[k-1] >= -1000 && x[k-1] <= 1000 && y[k-1] >= -1000 && y[k-1] <= 1000 && z[k-1] >= -1000 && z[k-1] <= 1000 {
				if x[k] >= -1000 && x[k] <= 1000 && y[k] >= -1000 && y[k] <= 1000 && z[k] >= -1000 && z[k] <= 1000{
					xt = (x[k-1] - x[k])
					yt = (y[k-1] - y[k])
					zt = (z[k-1] - z[k])
					soma := math.Pow(xt, 2) + math.Pow(yt, 2) + math.Pow(zt, 2)
					distancia[k] = math.Sqrt(soma)
				} else{
					distancia[k] = 0
				}
			} else{
				distancia[k] = 0
			}
		}
		for i := 1; i <= (N-1); i++ {
			if distancia[i] != 0{
				fmt.Printf("%.2f\n", distancia[i])
			} else if distancia[i] == 0{
				fmt.Println("Valores x, y e/ou z excederam o limite")
			}
		}
	}
}