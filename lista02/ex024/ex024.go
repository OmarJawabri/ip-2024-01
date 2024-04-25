package main

import (
	"fmt"
	"math"
)

func main(){
	var x, N, fat, y, i float64
	fmt.Scanln(&x)
	fmt.Scanln(&N)
	fat = 1
	if N <= 9 {
		for k := 0;k <= int(N); k++ {
			fat = 1
			i = float64(k)*2
			for i > 1{
				fat = fat*i
				i = i - 1
			}
			taylorcos := math.Pow(x, 2*float64(k))/fat
			if k % 2 == 0 {
				y = y + taylorcos
			} else{
				y = y - taylorcos
			}
		}
		fmt.Printf("cos(%.2f) = %.6f", x, y)
	}
}