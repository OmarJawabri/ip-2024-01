package main

import (
	"fmt"
	"math"
)

func main(){
	var x, N, y, taylorsen float64
	var f, fat, k int
	fmt.Scanln(&x)
	fmt.Scanln(&N)
	fat = 1
	if N <= 9 {
		for i := 0; i <= int(N); i++ {
			k = 2*i + 1
			f = 2*i + 1
			for f > 1{
				fat = fat*f
				f = f - 1
			}
			taylorsen = math.Pow(x, float64(k))/float64(fat)
			if i % 2 != 0{
				y = y - taylorsen
			} else if i % 2 == 0 || i == 0{ 
				y = y + taylorsen
			}
			fat = 1
		}
		fmt.Printf("seno(%.2f) = %.6f", x, y)
	}
}