package main

import (
	"fmt"
	"math"
)

func main(){
	var x, N, y, fat, f float64
	fmt.Scanln(&x)
	fmt.Scanln(&N)
	fat = 1
	for i := 0; i <= int(N); i++ {
		f = float64(i)
		for f > 1{
			fat = fat*f
			f = f - 1
		}
		y = y + math.Pow(x, float64(i))/fat
		fat = 1
	}
	fmt.Printf("e^%.2f = %.6f", x, y)
}