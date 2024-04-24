package main

import (
	"fmt"
	"math"
)

func main(){
	var mp, mi, x, par, impar, y float64
	slice := 100
	x, y = 1, 1	
	n := make([]float64, slice)
	for i := 0; x != 0; i++ {
		fmt.Scanf("%f",&n[i])
		x = n[i]
	}		
	for i := 0; y != 0; i++ {
		if math.Mod(n[i], 2) == 0 && n[i] != 0 {
			par++
			mp = mp + n[i]
		} else if math.Mod(n[i], 2) != 0 {
			impar++
			mi = mi + n[i]
		}
		y = n[i]
	}
	mp = mp/par
	mi = mi/impar
	if math.IsNaN(mp){
		fmt.Print("MÉDIA PAR: 0.00\n")
	} else{
		fmt.Printf("MÉDIA PAR: %.2f\n", mp)
	}
	if math.IsNaN(mi) {
		fmt.Printf("MÉDIA ÍMPAR: 0.00")
	} else{
		fmt.Printf("MÉDIA ÍMPAR: %.2f", mi)
	}
}