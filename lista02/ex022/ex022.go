package main

import (
	"fmt"
	"math"
)

func main(){
	var N, i, num, den, k float64
	fmt.Scanln(&N)
	x := 10000.0
	num = N*x
	den = 10000
	for i = 1; i < num; i++{
		for k = 0; k < num; k++{
			if math.Mod(num, i) == 0 && math.Mod(den, i) == 0{
				num = num/i
				den = den/i
			} else{
				k++
			}
		}
	}
	if N > 0 {
		fmt.Printf("%.0f/%.0f", num, den)
	} else if N == 0{
		fmt.Printf("%.0f/0", N)
	}
}